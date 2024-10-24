package adapter

import (
	"log"
	"net/http"
	"service/rest-api/internal/port/in"

	"github.com/labstack/echo/v4"
)

type ApplicantHandler struct {
	service in.ApplicantService
}

func NewApplicantHandler(s in.ApplicantService) *ApplicantHandler {
	return &ApplicantHandler{service: s}
}

func (handler *ApplicantHandler) GetAllApplicants(context echo.Context) error {
	items, err := handler.service.GetAll(context.Request().Context())
	if err != nil {
		log.Printf("Error: %v", err)

		return context.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Unable to retrieve clients",
		})
	}
	return context.JSON(http.StatusOK, items)
}
