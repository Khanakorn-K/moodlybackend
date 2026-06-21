package insightcontroller

import (
	"errors"
	"moodly/internal/services"
	"moodly/pkg"
	"moodly/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InsightController struct {
	service *services.InsightService
}

func NewInsightController(service *services.InsightService) *InsightController {
	return &InsightController{service: service}
}

// GetInsights godoc
//
// @Summary Get Insights
// @Description Get mood insights for the selected date
// @Tags Insight
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param selectedDate query string true "Selected Date (YYYY-MM-DD)"
// @Success 200 {object} InsightResultResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /insights/get-insights [get]
func (ic *InsightController) GetInsights(c *gin.Context) {
	userID, ok := pkg.GetUserIDFromContext(c)
	if !ok {
		pkg.CreateAPIErrorResponse(
			c,
			http.StatusUnauthorized,
			"UNAUTHORIZED",
			"unauthorized",
		)
		return
	}

	selectedDate := c.Query("selectedDate")

	result, err := ic.service.GetInsights(userID, selectedDate)
	if err != nil {
		if errors.Is(err, services.ErrInvalidDateFormat) {
			pkg.CreateAPIErrorResponse(
				c,
				http.StatusBadRequest,
				"INVALID_DATE_FORMAT",
				err.Error(),
			)
			return
		}

		pkg.CreateAPIErrorResponse(
			c,
			http.StatusInternalServerError,
			"INTERNAL_SERVER_ERROR",
			"internal server error",
		)
		return
	}
	mapperResult := &InsightResultResponse{
		TotalLogs:        len(*result),
		AverageMood:      utils.CalculateAverageMood(*result),
		MoodDistribution: utils.CalculateMoodDistributionRecord(*result),
		CausesAnalysis:   utils.CalculateCauseAnalysis(*result),
	}
	pkg.CreateAPIResponse(c, http.StatusOK, mapperResult)
}
