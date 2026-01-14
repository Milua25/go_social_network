package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Milua25/go_social/docs"
	"github.com/Milua25/go_social/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
	"go.uber.org/zap"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr   string
	db     dbConfig
	env    string
	apiURL string
	logger *zap.SugaredLogger
	mail   mailConfig
}

type mailConfig struct {
	exp time.Duration
}

type dbConfig struct {
	addr         string
	port         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

// func (app *application) userHandler(w http.ResponseWriter, req *http.Request) {
// 	if req.URL.Path != "/v1/users" {
// 		http.Error(w, "404 not Found", 404)
// 		return
// 	}
// 	if req.Method != "GET" {
// 		http.Error(w, "method not allowed", http.StatusNotFound)
// 		return
// 	}
// 	w.Write([]byte("users available"))
// }

// mount configures the HTTP router and middleware stack.
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/v1", func(r chi.Router) {
		// health
		r.HandleFunc("/health", app.healthCheckHandler)

		// config
		docsURL := fmt.Sprintf("%s/swagger/doc.json", app.config.addr)
		r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(docsURL)))

		// posts
		r.Route("/posts", func(r chi.Router) {
			r.Post("/create", app.createPostHandler)

			r.Route("/{postID}", func(r chi.Router) {
				r.Use(app.postsContextMiddleware)
				r.Get("/", app.getPostHandler)
				r.Patch("/", app.patchPostHandler)
				r.Delete("/", app.deletePostHandler)
			})
		})

		// users
		r.Route("/users", func(r chi.Router) {
			r.Put("/activate/{token}", app.activateUserHandler)

			r.Get("/", app.getAllUsersHandler)
			r.Group(func(r chi.Router) {
				r.Get("/feed", app.getUserFeedHandler)
			})
			r.Route("/{userID}", func(r chi.Router) {
				r.Use(app.getUserContextIdMiddleware)
				r.Get("/", app.getUserByIdHandler)
				r.Put("/follow", app.followUserHandler)
				r.Put("/unfollow", app.unfollowUserHandler)
			})

		})

		// public routes
		r.Route("/authentication", func(r chi.Router) {
			r.Post("/user", app.registerUserHandler)
		})

	})

	// auth

	return r
}

// run starts the HTTP server with timeouts applied.
func (app *application) run() error {
	// Docs
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = app.config.apiURL
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Title = "Social Network API"

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      app.mount(),
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	app.config.logger.Infow("Server running on", "addr", srv.Addr, "env", app.config.env)
	return srv.ListenAndServe()
}
