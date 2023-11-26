package chi

import (
	"Rebrain/module09/internal/constants"
	"Rebrain/module09/internal/handlers/user"
	"github.com/go-chi/chi"
	"net/http"
)

func MiddlewareAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name, pass, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "failed authorization", http.StatusUnauthorized)
			return
		}

		if name != constants.AdminName || pass != constants.AdminPass {
			http.Error(w, "wrong login or password", http.StatusUnauthorized)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func Start() {
	r := chi.NewRouter()
	r.Route("/user", func(r chi.Router) {
		r.With(MiddlewareAuth).Post("/", user.CreateUserHandler)
		r.Get("/", user.GetUserListHandler)
		r.Get("/{id:[0-9]+}", user.GetUserHandler)
		r.With(MiddlewareAuth).Delete("/{id:[0-9]+}", user.DeleteUserHandler)
	})

	http.ListenAndServe("127.0.0.1:8080", r)
}
