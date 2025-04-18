package main

import "net/http"

func (app *App) authTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := app.models.Token.AuthenticateToken(r)
		if err != nil {
			payload := jsonResponse{
				Error:   true,
				Message: "Invalid auth credentials",
			}
			_ = app.writeJSON(w, http.StatusUnauthorized, payload)
			return
		}
		next.ServeHTTP(w, r)
	})
}
