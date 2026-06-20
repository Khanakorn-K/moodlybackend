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
