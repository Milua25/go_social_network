package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/Milua25/go_social/internal/mailer"
	"github.com/Milua25/go_social/internal/store"
	"github.com/google/uuid"
)

type RegisterPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

type userWithToken struct {
	*store.User
	Token string `json:"token"`
}

// registerUserHandler godoc
//
//	@Summary		Register a user
//	@Description	Creates a user account and invitation token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			payload	body	RegisterPayload	true	"User registration payload"
//	@Success		201
//	@Failure		400	{object}	map[string]string
//	@Failure		409	{object}	map[string]string
//	@Router			/authentication/user [post]
func (app *application) registerUserHandler(w http.ResponseWriter, req *http.Request) {
	var payload RegisterPayload

	if err := readJSON(w, req, &payload); err != nil {
		app.badRequestError(w, req, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestError(w, req, err)
		return
	}
	// hash the user password
	user := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
	}

	// hash the password
	if err := user.Password.Set(payload.Password); err != nil {
		app.internalServerError(w, req, err)
		return
	}

	ctx := req.Context()

	plainToken := uuid.New()

	// store token on DB
	hashedTokenByte := sha256.Sum256([]byte(plainToken.String()))

	hashedToken := hex.EncodeToString(hashedTokenByte[:])

	// store the user
	err := app.store.Users.CreateAndInvite(ctx, user, app.config.mail.exp, hashedToken)
	if err != nil {
		switch err {
		case store.ErrDuplicateEmail:
			app.badRequestError(w, req, err)
		case store.ErrDuplicateUsername:
			app.badRequestError(w, req, err)
		default:
			app.internalServerError(w, req, err)
			return
		}
		return
	}

	userWithToken := userWithToken{
		User:  user,
		Token: plainToken.String(),
	}

	activationURL := fmt.Sprintf("%s/confirm/%s", app.config.frontendURL, plainToken)

	isProdEnv := app.config.env == "production"
	vars := struct {
		Username      string
		ActivationURL string
	}{
		Username:      user.Username,
		ActivationURL: activationURL,
	}

	// send mail
	_, err = app.mailer.Send(mailer.UserWelcometemplate, user.Username, user.Email, vars, !isProdEnv)
	if err != nil {
		app.config.logger.Errorw("error sending welcome email", "error", err)

		// rollback user creating if email fails (SAGA pattern)
		if err := app.store.Users.Delete(ctx, user.ID); err != nil {
			app.config.logger.Errorw("error deleting user", "error", err)
		}
	}

	if err := writeJSON(w, http.StatusCreated, userWithToken); err != nil {
		app.internalServerError(w, req, err)
	}
}
