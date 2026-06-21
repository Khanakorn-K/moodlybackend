package customercontroller

import (
	"moodly/internal/domain/entities"
	"moodly/internal/services"
	"moodly/pkg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomCauseController struct {
	service *services.CustomCauseService
}

func NewCustomCauseController(service *services.CustomCauseService) *CustomCauseController {
	return &CustomCauseController{service: service}
}

// CreateCause godoc
//
// @Summary Create Custom Cause
// @Description Create a new custom cause for the authenticated user
// @Tags Custom Cause
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateCustomCausesRequest true "Create Custom Cause Request"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /custom-causes/create-custom-cause [post]
func (cc *CustomCauseController) CreateCause(c *gin.Context) {
	userID, ok := pkg.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	var req CreateCustomCausesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	customcauses := entities.CustomCauseEntity{
		UserID: userID,
		Name:   req.Name,
	}

	if err := cc.service.CreateCause(&customcauses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "cause created",
		"cause":   customcauses.Name,
	})
}

// GetCauses godoc
//
// @Summary Get Custom Causes
// @Description Get all custom causes for the authenticated user
// @Tags Custom Cause
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /custom-causes/get-custom-causes [get]
func (cc *CustomCauseController) GetCauses(c *gin.Context) {
	userID, ok := pkg.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	causes, err := cc.service.GetCauses(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"causes": causes,
	})
}

// UpdateCause godoc
//
// @Summary Update Custom Cause
// @Description Update a custom cause by ID
// @Tags Custom Cause
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Cause ID"
// @Param request body CreateCustomCausesRequest true "Update Custom Cause Request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /custom-causes/update-custom-cause/{id} [patch]
func (cc *CustomCauseController) UpdateCause(c *gin.Context) {
	userID, ok := pkg.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid cause id"})
		return
	}

	var req CreateCustomCausesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	cause := entities.CustomCauseEntity{
		ID:     uint(id),
		UserID: userID,
		Name:   req.Name,
	}

	if err := cc.service.UpdateCause(&cause); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cause updated",
		"cause":   cause,
	})
}

// DeleteCause godoc
//
// @Summary Delete Custom Cause
// @Description Delete a custom cause by ID
// @Tags Custom Cause
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Cause ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /custom-causes/delete-custom-cause/{id} [delete]
func (cc *CustomCauseController) DeleteCause(c *gin.Context) {
	userID, ok := pkg.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid cause id"})
		return
	}

	if err := cc.service.DeleteCause(uint(id), userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cause deleted",
	})
}
