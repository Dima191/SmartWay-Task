package handlerimpl

import (
	"errors"
	"github.com/Dima191/SmartWay-Task/internal/models"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func (h *handler) AddEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employee := models.EmployeeBase{}
		if err := render.DecodeJSON(r.Body, &employee); err != nil {
			h.logger.Error("failed to decode body", slog.String("error", err.Error()))
			http.Error(w, "failed to decode body", http.StatusBadRequest)
			return
		}

		employeeID, err := h.service.AddEmployee(r.Context(), employee)
		if err != nil {
			switch {
			case errors.Is(err, models.ErrPassportAlreadyExists):
				http.Error(w, "failed to add passport. passport already exists", http.StatusBadRequest)
				return
			case errors.Is(err, models.ErrEmployeeAlreadyExists):
				http.Error(w, "failed to add employee. employee already exists", http.StatusBadRequest)
				return
			}

			http.Error(w, "failed to add employee", http.StatusInternalServerError)
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, struct {
			ID int `json:"employee_id"`
		}{employeeID})
	}
}
