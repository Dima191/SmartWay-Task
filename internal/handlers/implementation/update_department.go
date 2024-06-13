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

func (h *handler) UpdateDepartment() http.HandlerFunc {
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

		departmentIDStr := chi.URLParam(r, handlers.DepartmentIDKey)
		if departmentIDStr == "" {
			http.Error(w, "invalid department id", http.StatusBadRequest)
			return
		}
		departmentID, err := strconv.Atoi(departmentIDStr)
		if err != nil {
			http.Error(w, "invalid department id", http.StatusBadRequest)
			return
		}

		department := models.Department{}

		if err = render.DecodeJSON(r.Body, &department); err != nil {
			h.logger.Error("failed to decode body", slog.String("error", err.Error()))
			http.Error(w, "failed to decode body", http.StatusBadRequest)
			return
		}

		department.CompanyID = companyID
		department.ID = departmentID

		if err = h.service.UpdateDepartment(r.Context(), department); err != nil {
			switch {
			case errors.Is(err, models.ErrAlreadyExists):
				http.Error(w, "failed to update department. department should be unique", http.StatusBadRequest)
				return
			case errors.Is(err, models.ErrUnknownCompanyReference):
				http.Error(w, "failed to add department to a non-existent company", http.StatusBadRequest)
				return
			}
			http.Error(w, "failed to update department", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
	}
}
