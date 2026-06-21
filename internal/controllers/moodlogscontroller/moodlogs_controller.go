package moodlogscontroller

import (
	"moodly/internal/domain/entities"
	"moodly/internal/services"
	"moodly/pkg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MoodLogsController struct {
	service *services.MoodLogsService
}

func NewMoodLogsController(service *services.MoodLogsService) *MoodLogsController {
	return &MoodLogsController{service: service}
}

// HandleCreateMoodLog godoc
//
// @Summary Create Mood Log
// @Description Create Mood Log
// @Tags Moodlog
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body MoodlogsCreateRequest true "MoodlogsCreate Request"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /mood-logs/create-mood-log [post]
func (mc *MoodLogsController) CreateMoodLog(c *gin.Context) {
	userID, ok := pkg.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	var req MoodlogsCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	moodLog := entities.MoodLogEntity{
		UserID: userID,
		Mood:   req.Mood,
		Note:   req.Note,
		Causes: req.Causes,
	}

	if err := mc.service.CreateMoodLog(&moodLog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "mood log created",
		"mood_log": moodLog,
	})
}

// GetMoodLogsByDate godoc
//
// @Summary Get Mood Logs By Date
// @Description Get user's mood logs by date
// @Tags Moodlog
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param date query string true "Date (YYYY-MM-DD)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /mood-logs/get-mood-logs [get]
func (mc *MoodLogsController) GetMoodLogsByDate(c *gin.Context) {
	userID, ok := pkg.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	date := c.Query("date")

	moodLogs, err := mc.service.GetMoodLogsByDate(userID, date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mood_logs": moodLogs,
	})
}

// UpdateMoodLog godoc
//
// @Summary Update Mood Log
// @Description Update an existing mood log
// @Tags Moodlog
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Mood Log ID"
// @Param request body MoodlogsCreateRequest true "Update Mood Log Request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /mood-logs/update-mood-log/{id} [patch]
func (mc *MoodLogsController) UpdateMoodLog(c *gin.Context) {
	userID, ok := pkg.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid mood log id"})
		return
	}

	var req MoodlogsCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	moodLog := entities.MoodLogEntity{
		ID:     uint(id),
		UserID: uint(userID),
		Mood:   req.Mood,
		Note:   req.Note,
		Causes: req.Causes,
	}

	if err := mc.service.UpdateMoodLog(&moodLog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "mood log updated",
		"mood_log": &moodLog,
	})
}

// DeleteMoodLog godoc
//
// @Summary Delete Mood Log
// @Description Delete a mood log
// @Tags Moodlog
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Mood Log ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /mood-logs/delete-mood-log/{id} [delete]
func (mc *MoodLogsController) DeleteMoodLog(c *gin.Context) {
	userID, ok := pkg.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid mood log id"})
		return
	}

	if err := mc.service.DeleteMoodLog(uint(id), userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "mood log deleted",
	})
}
