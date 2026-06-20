package utils

import (
	"encoding/json"
	"moodly/internal/domain/entities"
	"strconv"
	"strings"
)

func CalculateAverageMood(logs []entities.MoodLogEntity) float64 {
	if len(logs) == 0 {
		return 0
	}

	total := 0
	for _, log := range logs {
		total += log.Mood
	}

	return RoundOneDecimal(float64(total) / float64(len(logs)))
}

func CalculateMoodDistributionRecord(logs []entities.MoodLogEntity) map[string]int {
	result := map[string]int{}

	for _, log := range logs {
		moodKey := strconv.Itoa(log.Mood)
		result[moodKey]++
	}

	return result
}

func CalculateCauseAnalysis(logs []entities.MoodLogEntity) map[string]map[string]int {
	result := map[string]map[string]int{}

	for _, log := range logs {
		causes := ParseCauses(log.Causes)
		moodKey := strconv.Itoa(log.Mood)

		for _, cause := range causes {
			cause = strings.TrimSpace(cause)
			if cause == "" {
				continue
			}

			if result[cause] == nil {
				result[cause] = map[string]int{}
			}

			result[cause][moodKey]++
		}
	}

	return result
}

func ParseCauses(rawCauses string) []string {
	rawCauses = strings.TrimSpace(rawCauses)

	if rawCauses == "" {
		return []string{}
	}

	var causes []string

	// กรณีเก็บเป็น JSON string เช่น ["work","money"]
	if err := json.Unmarshal([]byte(rawCauses), &causes); err == nil {
		return causes
	}

	// กรณีเก็บเป็น string ธรรมดา เช่น work,money
	parts := strings.Split(rawCauses, ",")

	cleanCauses := []string{}

	for _, part := range parts {
		cause := strings.TrimSpace(part)
		cause = strings.Trim(cause, `"'[]{} `)

		if cause != "" {
			cleanCauses = append(cleanCauses, cause)
		}
	}

	return cleanCauses
}

func RoundOneDecimal(value float64) float64 {
	result, _ := strconv.ParseFloat(strconv.FormatFloat(value, 'f', 1, 64), 64)
	return result
}
