package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func (app *application) BasicAuthMiddleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// read the auth header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				app.unAuthorizedResponseError(w, r, fmt.Errorf("authorization header missing"))
				return
			}

			// parse it -> get the base64
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Basic" {
				app.unAuthorizedResponseError(w, r, fmt.Errorf("authorization header is malformed"))
				return
			}

			// decode it
			decoded, err := base64.StdEncoding.DecodeString(parts[1])
			if err != nil {
				app.unAuthorizedResponseError(w, r, err)
				return
			}

			// check the credentials
			username := app.config.auth.basic.user
			pass := app.config.auth.basic.pass
			creds := strings.SplitN(string(decoded), ":", 2)
			if len(creds) != 2 || creds[0] != username || creds[1] != pass {
				app.unAuthorizedResponseError(w, r, fmt.Errorf("invalid credentials"))
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}

func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// read the auth header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			app.unAuthorizedResponseError(w, r, fmt.Errorf("authorization header missing"))
			return
		}

		// parse it -> get the base64
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			app.unAuthorizedResponseError(w, r, fmt.Errorf("authorization header is malformed"))
			return
		}

		token := parts[1]

		jwtToken, err := app.authenticatior.ValidateToken(token)
		if err != nil {
			app.unAuthorizedResponseError(w, r, err)
			return
		}

		claims := jwtToken.Claims.(jwt.MapClaims)

		userID, err := strconv.ParseInt(fmt.Sprint("%.f", claims["sub"]), 10, 64)

		if err != nil {
			app.unAuthorizedResponseError(w, r, err)
			return
		}
		ctx := r.Context()
		user, err := app.store.Users.GetUserByID(r.Context(), int(userID))
		if err != nil {
			app.unAuthorizedResponseError(w, r, err)
			return
		}

		ctx = context.WithValue(ctx, userKeyCtx, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
