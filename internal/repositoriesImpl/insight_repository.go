package repositoriesImpl

import (
	"moodly/internal/domain/entities"
	"time"

	"gorm.io/gorm"
)

type InsightRepository struct {
	db *gorm.DB
}

func NewInsightRepository(db *gorm.DB) *InsightRepository {
	return &InsightRepository{db: db}
}

func (r *InsightRepository) FindInsightLogs(
	userID uint,
	selectedDate *time.Time,
) ([]entities.MoodLogEntity, error) {
	var moodLogs []entities.MoodLogEntity

	query := r.db.Where("user_id = ?", userID)

	if selectedDate != nil {
		startDate := *selectedDate
		endDate := startDate.AddDate(0, 0, 1)

		query = query.Where("created_at >= ? AND created_at < ?", startDate, endDate)
	}

	if err := query.Find(&moodLogs).Error; err != nil {
		return nil, err
	}

	return moodLogs, nil
}
