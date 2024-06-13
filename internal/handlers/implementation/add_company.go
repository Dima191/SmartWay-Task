package handlerimpl

import (
	"errors"
	"github.com/Dima191/SmartWay-Task/internal/models"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func (h *handler) AddCompany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		company := models.Company{}
		if err := render.DecodeJSON(r.Body, &company); err != nil {
			h.logger.Error("failed to decode body", slog.String("error", err.Error()))
			http.Error(w, "failed to decode body", http.StatusBadRequest)
			return
		}

		companyID, err := h.service.AddCompany(r.Context(), company)
		if err != nil {
			if errors.Is(err, models.ErrAlreadyExists) {
				http.Error(w, "failed to add company. company already exists", http.StatusBadRequest)
				return
			}

			http.Error(w, "failed to add company", http.StatusBadRequest)
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, struct {
			CompanyID int `json:"company_id"`
		}{companyID})
	}
}
