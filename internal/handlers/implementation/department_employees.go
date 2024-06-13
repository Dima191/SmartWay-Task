package handlerimpl

import (
	"github.com/Dima191/SmartWay-Task/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

func (h *handler) DepartmentEmployees() http.HandlerFunc {
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

		employees, err := h.service.DepartmentEmployees(r.Context(), companyID, departmentID)
		if err != nil {
			http.Error(w, "invalid department id", http.StatusBadRequest)
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, employees)
	}
}
