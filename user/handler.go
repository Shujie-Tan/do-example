package user

import (
	"example/domain"
	"net/http"

	"github.com/samber/do"
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

func NewHandler(i *do.Injector) (domain.UserHandler, error) {
	svc := do.MustInvoke[domain.UserService](i)
	return &handler{svc: svc}, nil
}
