package handlerimpl

import (
	"github.com/Dima191/SmartWay-Task/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

func (h *handler) Departments() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		companyIDStr := chi.URLParam(r, handlers.CompanyIDKey)
		companyID, err := strconv.Atoi(companyIDStr)
		if err != nil {
			http.Error(w, "invalid company id", http.StatusBadRequest)
			return
		}

		departments, err := h.service.Departments(r.Context(), companyID)
		if err != nil {
			http.Error(w, "failed to get list of departments", http.StatusInternalServerError)
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, departments)
	}
}
