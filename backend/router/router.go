package router

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/haron1996/inboxzen/api/account"
	"github.com/haron1996/inboxzen/logger"
	"github.com/haron1996/inboxzen/mw"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	styles := log.DefaultStyles()

	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERROR").
		Foreground(lipgloss.Color("#FF0060"))

	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("ERROR").
		Foreground(lipgloss.Color("#22A699"))

	l := log.New(os.Stderr)

	l.SetStyles(styles)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "https://bookmarkmonster.xyz", "https://*.bookmarkmonster.xyz"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH", "HEAD"},
		AllowedHeaders:   []string{"User-Agent", "Content-Type", "Accept", "Accept-Encoding", "Accept-Language", "Cache-Control", "Connection", "DNT", "Host", "Origin", "Pragma", "Referer", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
		Debug:            false,
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.AllowContentEncoding("application/json", "application/x-www-form-urlencoded"))
	r.Use(middleware.CleanPath)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.RealIP)

	r.Route("/account", func(r chi.Router) {
		r.Get("/getauthurl", logger.MakeHandler(account.GenerateGoogleAuthURL, l))
		r.Post("/comletegoogleauth", logger.MakeHandler(account.CompleteGoogleAuth, l))
	})

	r.Route("/private", func(r chi.Router) {
		r.Use(mw.AuthenticateRequest())
		r.Get("/checkuserloginstatus", logger.MakeHandler(account.CheckUserLoginStatus, l))
	})

	return r
}
