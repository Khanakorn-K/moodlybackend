package overviewcontroller

import (
	"moodly/internal/services"
	"moodly/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OverviewController struct {
	service *services.OverviewService
}

func NewOverviewController(svc *services.OverviewService) *OverviewController {
	return &OverviewController{service: svc}
}

// GetMonthlyAverageMood godoc
//
// @Summary Get Monthly Average Mood
// @Description Get average mood statistics for a specific month
// @Tags Overview
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param month query string true "Month (YYYY-MM)"
// @Success 200 {object} MonthlyAverageMoodResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /overview/get-monthly-average-mood [get]
func (oc *OverviewController) GetMonthlyAverageMood(c *gin.Context) {
	userID, ok := pkg.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	result, err := oc.service.GetMonthlyAverageMood(userID, c.Query("month"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toMonthlyAverageMoodResponse(result))
}

// GetOverview godoc
//
// @Summary Get Overview
// @Description Get mood overview within a date range
// @Tags Overview
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param startDate query string true "Start Date (YYYY-MM-DD)"
// @Param endDate query string true "End Date (YYYY-MM-DD)"
// @Success 200 {object} OverviewResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /overview/get-overview [get]
func (oc *OverviewController) GetOverview(c *gin.Context) {
	userID, ok := pkg.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	result, err := oc.service.GetOverview(userID, c.Query("startDate"), c.Query("endDate"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toOverviewResponse(result))
}
