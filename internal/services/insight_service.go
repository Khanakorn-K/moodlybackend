package services

import (
	"errors"
	"moodly/internal/repositoriesImpl"
	"moodly/utils"
	"strings"
	"time"
)

var ErrInvalidDateFormat = errors.New("invalid date format")

type InsightService struct {
	repo *repositoriesImpl.InsightRepository
}

func NewInsightService(repo *repositoriesImpl.InsightRepository) *InsightService {
	return &InsightService{repo: repo}
}

type InsightResult struct {
	TotalLogs        int                       `json:"totalLogs"`
	AverageMood      float64                   `json:"averageMood"`
	MoodDistribution map[string]int            `json:"moodDistribution"`
	CausesAnalysis   map[string]map[string]int `json:"causesAnalysis"`
}

func (s *InsightService) GetInsights(userID uint, selectedDate string) (*InsightResult, error) {
	if userID == 0 {
		return nil, errors.New("user id is required")
	}

	selectedDate = strings.TrimSpace(selectedDate)

	var selectedDateFilter *time.Time

	if selectedDate != "" {
		parsedDate, err := time.Parse("2006-01-02", selectedDate)
		if err != nil {
			return nil, ErrInvalidDateFormat
		}

		selectedDateFilter = &parsedDate
	}

	logs, err := s.repo.FindInsightLogs(userID, selectedDateFilter)
	if err != nil {
		return nil, err
	}

	return &InsightResult{
		TotalLogs:        len(logs),
		AverageMood:      utils.CalculateAverageMood(logs),
		MoodDistribution: utils.CalculateMoodDistributionRecord(logs),
		CausesAnalysis:   utils.CalculateCauseAnalysis(logs),
	}, nil
}
