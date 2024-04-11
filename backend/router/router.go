package router

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/haron1996/inboxzen/api/handlers"
	"github.com/haron1996/inboxzen/logger"
	"github.com/haron1996/inboxzen/mw"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	l := log.New(os.Stderr)

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

	r.Route("/user", func(r chi.Router) {
		r.Get("/getauthurl", logger.MakeHandler(handlers.GenerateGoogleAuthURL, l))
		r.Post("/comletegoogleauth", logger.MakeHandler(handlers.CompleteGoogleAuth, l))
	})

	r.Route("/private", func(r chi.Router) {
		r.Use(mw.AuthenticateRequest())
		r.Get("/checkloginstatus", logger.MakeHandler(handlers.CheckUserLoginStatus, l))
		r.Get("/getuseraccount", logger.MakeHandler(handlers.GetUserAccount, l))
		r.Get("/logout", logger.MakeHandler(handlers.LogOut, l))
		r.Get("/switchaccount/{email}", logger.MakeHandler(handlers.SwitchAccount, l))
		r.Get("/getdomains", logger.MakeHandler(handlers.GetDomains, l))
		r.Get("/getemails", logger.MakeHandler(handlers.GetVipEmails, l))
		r.Get("/getkeywords", logger.MakeHandler(handlers.GetVipKeywords, l))
		r.Post("/updatedomains", logger.MakeHandler(handlers.AddDomains, l))
		r.Post("/updateemails", logger.MakeHandler(handlers.UpdateVipEmails, l))
		r.Post("/updatekeywords", logger.MakeHandler(handlers.UpdateVipKeywords, l))
	})

	return r
}
