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

func (h *handler) AddDepartment() http.HandlerFunc {
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

		department := models.Department{}
		if err = render.DecodeJSON(r.Body, &department); err != nil {
			h.logger.Error("failed to decode body", slog.String("error", err.Error()))
			http.Error(w, "failed to decode body", http.StatusBadRequest)
			return
		}

		department.CompanyID = companyID

		departmentID, err := h.service.AddDepartment(r.Context(), department)
		if err != nil {
			switch {
			case errors.Is(err, models.ErrAlreadyExists):
				http.Error(w, "failed to add department. department already exists", http.StatusBadRequest)
				return
			case errors.Is(err, models.ErrUnknownReference):
				http.Error(w, "failed to add department to a non-existent company", http.StatusBadRequest)
				return
			}
			http.Error(w, "failed to add department", http.StatusInternalServerError)
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, struct {
			DepartmentID int `json:"department_id"`
		}{departmentID})
	}
}
