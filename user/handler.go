package user

import (
	"example/domain"
	"net/http"
)

type handler struct {
	svc domain.UserService
}

func (h *handler) FetchByUsername() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		ctx := r.Context()
		user, err := h.svc.FetchByUsername(ctx, username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(user.Username))
	}
}
