package repositoriesImpl

import (
	"errors"
	"moodly/internal/domain/entities"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) FindByEmail(email string) (*entities.UserEntity, error) {
	var user entities.UserEntity

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) CreateUser(user *entities.UserEntity) error {
	return r.db.Create(user).Error
}

func (r *AuthRepository) FindOrCreateOAuthAccount(
	userID uint,
	provider string,
	providerAccountID string,
) (*entities.OAuthAccountEntity, error) {
	var account entities.OAuthAccountEntity

	err := r.db.
		Where("provider = ? AND provider_account_id = ?", provider, providerAccountID).
		First(&account).Error

	if err == nil {
		if account.UserID != userID {
			return nil, errors.New("oauth account already linked to another user")
		}

		return &account, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	account = entities.OAuthAccountEntity{
		UserID:            userID,
		Provider:          provider,
		ProviderAccountID: providerAccountID,
	}

	if err := r.db.Create(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}
