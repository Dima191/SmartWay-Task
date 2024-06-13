package handlerimpl

import (
	"errors"
	"github.com/Dima191/SmartWay-Task/internal/handlers"
	"github.com/Dima191/SmartWay-Task/internal/models"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *handler) FireEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employeeIDStr := chi.URLParam(r, handlers.EmployeeIDKey)
		if employeeIDStr == "" {
			http.Error(w, "invalid employee id", http.StatusBadRequest)
			return
		}

		employeeID, err := strconv.Atoi(employeeIDStr)
		if err != nil {
			http.Error(w, "invalid employee id", http.StatusBadRequest)
			return
		}

		if err = h.service.FireEmployee(r.Context(), employeeID); err != nil {
			if errors.Is(err, models.ErrNoData) {
				http.Error(w, "can not to fire a non-existent employee", http.StatusBadRequest)
				return
			}
			http.Error(w, "failed to fire employee", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
