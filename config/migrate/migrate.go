package main

import (
	"moodly/config/initializers"
	"moodly/internal/domain/entities"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	//ถ้ามีการเพิ่ม model หรือแก้ไข อย่าลืม migrate
	initializers.DB.AutoMigrate(
		&entities.UserEntity{},
		&entities.OAuthAccountEntity{},
		&entities.MoodLogEntity{},
		&entities.CustomCauseEntity{},
	)
}
