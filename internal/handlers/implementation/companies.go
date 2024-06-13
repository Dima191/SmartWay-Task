package handlerimpl

import (
	"github.com/go-chi/render"
	"net/http"
)

func (h *handler) Companies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		companies, err := h.service.Companies(r.Context())
		if err != nil {
			http.Error(w, "failed to get list of companies", http.StatusInternalServerError)
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, companies)
	}
}
