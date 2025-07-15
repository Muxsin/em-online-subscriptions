package repositories

import (
	"effective-mobile/online-subscriptions/internal/models"
	"gorm.io/gorm"
)

type SubscriptionRepositoryInterface interface {
	Insert(subscription *models.Subscription) error
	List() ([]*models.Subscription, error)
	GetByID(id uint) (*models.Subscription, error)
	Update(subscription *models.Subscription) error
	Delete(subscription *models.Subscription) error
}

type SubscriptionRepository struct {
	DB *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepositoryInterface {
	return &SubscriptionRepository{
		DB: db,
	}
}

func (repository *SubscriptionRepository) Insert(subscription *models.Subscription) error {
	return repository.DB.Create(subscription).Error
}

func (repository *SubscriptionRepository) List() ([]*models.Subscription, error) {
	var subscriptions []*models.Subscription

	err := repository.DB.Find(&subscriptions).Error

	return subscriptions, err
}

func (repository *SubscriptionRepository) GetByID(id uint) (*models.Subscription, error) {
	var subscription models.Subscription

	err := repository.DB.First(&subscription, id).Error

	return &subscription, err
}

func (repository *SubscriptionRepository) Update(subscription *models.Subscription) error {
	return repository.DB.Save(subscription).Error
}

func (repository *SubscriptionRepository) Delete(subscription *models.Subscription) error {
	return repository.DB.Delete(subscription).Error
}
