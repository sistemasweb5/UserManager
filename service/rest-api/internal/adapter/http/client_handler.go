package adapter

import (
	"log"
	"net/http"
	"service/rest-api/internal/core/domain"
	"service/rest-api/internal/port/in"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ClientHandler struct {
	applicantService in.ApplicantService
	workerService    in.WorkerService
	categoryService  in.CategoryService
	clientService    in.ClientService
}

func NewClientHandler(
	applicant in.ApplicantService,
	worker in.WorkerService,
	category in.CategoryService,
	client in.ClientService,
) ClientHandler {
	return ClientHandler{
		applicantService: applicant,
		workerService:    worker,
		categoryService:  category,
		clientService:    client,
	}
}

func (handler *ClientHandler) GetAllCategory(context echo.Context) error {
	items, err := handler.categoryService.GetAllCategory(context.Request().Context())
	if err != nil {
		log.Printf("Error: %v", err)

		return context.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Unable to retrieve category",
		})
	}
	return context.JSON(http.StatusOK, items)
}

// Deprecated: This endpoint could return null WorkSchedule.
// Use GetApplicantById or GetWorkerById instead.
func (handler *ClientHandler) GetClientById(context echo.Context) error {
	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		log.Printf("Could not parse id: %v", err)
		return err
	}

	client, err := handler.clientService.GetById(context.Request().Context(), &id)
	if err != nil {
		log.Printf("Could not get client: %v", err)
		return err
	}

	return context.JSON(http.StatusOK, client)
}

func (handler *ClientHandler) GetAllApplicant(context echo.Context) error {
	items, err := handler.applicantService.GetAllApplicant(context.Request().Context())
	if err != nil {
		log.Printf("Error: %v", err)

		return context.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Unable to retrieve clients",
		})
	}
	return context.JSON(http.StatusOK, items)
}

func (handler *ClientHandler) GetAllWorker(context echo.Context) error {
	items, err := handler.workerService.GetAllWorker(context.Request().Context())
	if err != nil {
		log.Printf("Error: %v", err)
		return context.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Unable to retrieve clients",
		})
	}
	return context.JSON(http.StatusOK, items)
}

func (handler *ClientHandler) GetWorkerById(context echo.Context) error {
	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		log.Printf("Could not parse id: %v", err)
		return err
	}

	client, err := handler.workerService.GetWorkerById(context.Request().Context(), &id)
	if err != nil {
		log.Printf("Could not get client: %v", err)
		return err
	}

	return context.JSON(http.StatusOK, client)
}

func (handler *ClientHandler) GetApplicantById(context echo.Context) error {
	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		log.Printf("Could not parse id: %v", err)
		return err
	}

	client, err := handler.applicantService.GetApplicantById(context.Request().Context(), &id)
	if err != nil {
		log.Printf("Could not get client: %v", err)
		return err
	}

	return context.JSON(http.StatusOK, client)
}

func (handler *ClientHandler) CreateApplicant(c echo.Context) error {
	var user domain.ApplicantRequest
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	client, err := handler.applicantService.CreateApplicant(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, client)
}

func (handler *ClientHandler) CreateWorker(c echo.Context) error {
	var user domain.WorkerRequest

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	client, err := handler.workerService.CreateWorker(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, client)
}
