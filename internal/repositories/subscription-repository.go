package repositories

import (
	"effective-mobile/online-subscriptions/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type SubscriptionRepositoryInterface interface {
	Insert(subscription *models.Subscription) error
	List() ([]*models.Subscription, error)
	GetByID(id string) (*models.Subscription, error)
	Update(subscription *models.Subscription) error
	Delete(subscription *models.Subscription) error
	CalculateTotalCost(user_ID *uuid.UUID, service_name string, period_start, period_end *time.Time) (int, error)
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

func (repository *SubscriptionRepository) GetByID(id string) (*models.Subscription, error) {
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

func (repository *SubscriptionRepository) CalculateTotalCost(
	user_ID *uuid.UUID,
	service_name string,
	period_start, period_end *time.Time,
) (int, error) {
	var total_cost *int
	query := repository.DB.Model(&models.Subscription{})

	if user_ID != nil {
		query = query.Where("user_id = ?", *user_ID)
	}

	if service_name != "" {
		query = query.Where("service_name ILIKE ?", "%"+service_name+"%")
	}

	date_parse_function := "TO_DATE(start_date, 'MM-YYYY')"

	if period_start != nil {
		query = query.Where(date_parse_function+" >= ?", *period_start)
	}

	if period_end != nil {
		query = query.Where(date_parse_function+" <= ?", *period_end)
	}

	if err := query.Select("SUM(price)").Scan(&total_cost).Error; err != nil {
		return 0, err
	}

	if total_cost != nil {
		return *total_cost, nil
	}

	return 0, nil
}
