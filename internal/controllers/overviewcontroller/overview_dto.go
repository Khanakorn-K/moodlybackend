package overviewcontroller

import "moodly/internal/services"

type DailyCauseDistributionResponse struct {
	Cause string `json:"cause"`
	Count int    `json:"count"`
}

type DailyMoodAverageResponse struct {
	Date              string                           `json:"date"`
	AverageMood       float64                          `json:"averageMood"`
	TotalLogs         int                              `json:"totalLogs"`
	CauseDistribution []DailyCauseDistributionResponse `json:"causeDistribution"`
}

type MoodDistributionResponse struct {
	Mood  int `json:"mood"`
	Count int `json:"count"`
}

type MoodNoteResponse struct {
	ID        uint     `json:"id"`
	Date      string   `json:"date"`
	Mood      int      `json:"mood"`
	Note      string   `json:"note"`
	Causes    []string `json:"causes"`
	CreatedAt string   `json:"createdAt"`
}

type CauseSummaryResponse struct {
	Cause         string                     `json:"cause"`
	TotalCount    int                        `json:"totalCount"`
	MoodBreakdown []MoodDistributionResponse `json:"moodBreakdown"`
}

type MonthlyAverageMoodResponse struct {
	Month             string                     `json:"month"`
	StartDate         string                     `json:"startDate"`
	EndDate           string                     `json:"endDate"`
	TotalLogs         int                        `json:"totalLogs"`
	AverageMood       float64                    `json:"averageMood"`
	DailyMoodAverages []DailyMoodAverageResponse `json:"dailyMoodAverages"`
}

type OverviewResponse struct {
	StartDate         string                     `json:"startDate"`
	EndDate           string                     `json:"endDate"`
	TotalLogs         int                        `json:"totalLogs"`
	AverageMood       float64                    `json:"averageMood"`
	DailyMoodAverages []DailyMoodAverageResponse `json:"dailyMoodAverages"`
	MoodDistribution  []MoodDistributionResponse `json:"moodDistribution"`
	MoodNotes         []MoodNoteResponse         `json:"moodNotes"`
	CauseSummaries    []CauseSummaryResponse     `json:"causeSummaries"`
}

func toMonthlyAverageMoodResponse(r *services.MonthlyAverageMoodResult) MonthlyAverageMoodResponse {
	return MonthlyAverageMoodResponse{
		Month:             r.Month,
		StartDate:         r.StartDate,
		EndDate:           r.EndDate,
		TotalLogs:         r.TotalLogs,
		AverageMood:       r.AverageMood,
		DailyMoodAverages: toDailyMoodAverageResponses(r.DailyMoodAverages),
	}
}

func toOverviewResponse(r *services.OverviewResult) OverviewResponse {
	return OverviewResponse{
		StartDate:         r.StartDate,
		EndDate:           r.EndDate,
		TotalLogs:         r.TotalLogs,
		AverageMood:       r.AverageMood,
		DailyMoodAverages: toDailyMoodAverageResponses(r.DailyMoodAverages),
		MoodDistribution:  toMoodDistributionResponses(r.MoodDistribution),
		MoodNotes:         toMoodNoteResponses(r.MoodNotes),
		CauseSummaries:    toCauseSummaryResponses(r.CauseSummaries),
	}
}

func toDailyMoodAverageResponses(items []services.DailyMoodAverage) []DailyMoodAverageResponse {
	result := make([]DailyMoodAverageResponse, 0, len(items))
	for _, item := range items {
		causes := make([]DailyCauseDistributionResponse, 0, len(item.CauseDistribution))
		for _, c := range item.CauseDistribution {
			causes = append(causes, DailyCauseDistributionResponse{
				Cause: c.Cause,
				Count: c.Count,
			})
		}
		result = append(result, DailyMoodAverageResponse{
			Date:              item.Date,
			AverageMood:       item.AverageMood,
			TotalLogs:         item.TotalLogs,
			CauseDistribution: causes,
		})
	}
	return result
}

func toMoodDistributionResponses(items []services.MoodDistribution) []MoodDistributionResponse {
	result := make([]MoodDistributionResponse, 0, len(items))
	for _, item := range items {
		result = append(result, MoodDistributionResponse{
			Mood:  item.Mood,
			Count: item.Count,
		})
	}
	return result
}

func toMoodNoteResponses(items []services.MoodNote) []MoodNoteResponse {
	result := make([]MoodNoteResponse, 0, len(items))
	for _, item := range items {
		result = append(result, MoodNoteResponse{
			ID:        item.ID,
			Date:      item.Date,
			Mood:      item.Mood,
			Note:      item.Note,
			Causes:    item.Causes,
			CreatedAt: item.CreatedAt,
		})
	}
	return result
}

func toCauseSummaryResponses(items []services.CauseSummary) []CauseSummaryResponse {
	result := make([]CauseSummaryResponse, 0, len(items))
	for _, item := range items {
		breakdown := make([]MoodDistributionResponse, 0, len(item.MoodBreakdown))
		for _, b := range item.MoodBreakdown {
			breakdown = append(breakdown, MoodDistributionResponse{
				Mood:  b.Mood,
				Count: b.Count,
			})
		}
		result = append(result, CauseSummaryResponse{
			Cause:         item.Cause,
			TotalCount:    item.TotalCount,
			MoodBreakdown: breakdown,
		})
	}
	return result
}
