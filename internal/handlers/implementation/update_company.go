package handlerimpl

import (
	"errors"
	"github.com/Dima191/SmartWay-Task/internal/handlers"
	"github.com/Dima191/SmartWay-Task/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *handler) UpdateCompany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		companyIDStr := chi.URLParam(r, handlers.CompanyIDKey)
		if companyIDStr == "" {
			http.Error(w, "invalid company id", http.StatusBadRequest)
			return
		}

		companyID, err := strconv.Atoi(companyIDStr)
		if err != nil {
			http.Error(w, "invalid company id", http.StatusBadRequest)
			return
		}

		company := models.Company{}

		if err = render.DecodeJSON(r.Body, &company); err != nil {
			h.logger.Error("failed to decode body", slog.String("error", err.Error()))
			http.Error(w, "failed to decode body", http.StatusBadRequest)
			return
		}

		company.ID = companyID

		if err = h.service.UpdateCompany(r.Context(), company); err != nil {
			if errors.Is(err, models.ErrAlreadyExists) {
				http.Error(w, "failed to update company. company should be unique", http.StatusBadRequest)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
	}
}
