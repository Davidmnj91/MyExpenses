package repository

import (
	"github.com/Davidmnj91/MyExpenses/payment/entity"
	"github.com/Davidmnj91/MyExpenses/payment/model"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
	"time"
)

// Repository repository interface
type Repository interface {
	Create(paymentID string, concept string, amount float32, paymentConfiguration model.PaymentConfiguration) (entity.Payment, error)
	Update(paymentID string, concept string, amount float32, paymentConfiguration model.PaymentConfiguration) (entity.Payment, error)
	FindByID(paymentID string) (entity.Payment, error)
}

// PaymentRepository repository for profile data
type PaymentRepository struct {
	db *gorm.DB
}

// New create repository instance
func New(db *gorm.DB) Repository {
	return &PaymentRepository{db: db}
}

// Create create payment data in database
func (repository PaymentRepository) Create(paymentID string, concept string, amount float32, paymentConfiguration model.PaymentConfiguration) (entity.Payment, error) {
	payment := entity.Payment{ID: paymentID}

	findPaymentTx := repository.db.First(&payment)

	if err := findPaymentTx.Error; err != nil {
		return entity.Payment{}, err
	}

	paymentConfigurationJson := pgtype.JSONB{}
	if err := pgtype.JSONB.Scan(&paymentConfigurationJson, paymentConfiguration); err != nil {
		return entity.Payment{}, err
	}

	paymentEntity := entity.Payment{
		ID:                   uuid.NewString(),
		Concept:              concept,
		Amount:               amount,
		PaymentConfiguration: paymentConfigurationJson,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	createPaymentTx := repository.db.Create(paymentEntity)

	if err := createPaymentTx.Error; err != nil {
		return entity.Payment{}, err
	}

	return paymentEntity, nil
}

// Update update payment data by paymentID
func (repository PaymentRepository) Update(paymentID string, concept string, amount float32, paymentConfiguration model.PaymentConfiguration) (entity.Payment, error) {
	paymentEntity := entity.Payment{ID: paymentID}

	paymentConfigurationJson := pgtype.JSONB{}
	if err := pgtype.JSONB.Scan(&paymentConfigurationJson, paymentConfiguration); err != nil {
		return entity.Payment{}, err
	}

	paymentUpdate := entity.Payment{
		Concept:              concept,
		Amount:               amount,
		PaymentConfiguration: paymentConfigurationJson,
		UpdatedAt:            time.Now(),
	}
	updatedPaymentTx := repository.db.Model(&paymentEntity).Updates(paymentUpdate)

	if err := updatedPaymentTx.Error; err != nil {
		return entity.Payment{}, err
	}

	return paymentEntity, nil
}

// FindByID find payment data by payment id
func (repository PaymentRepository) FindByID(paymentID string) (entity.Payment, error) {
	paymentEntity := entity.Payment{ID: paymentID}

	findPaymentTx := repository.db.First(&paymentEntity)

	if err := findPaymentTx.Error; err != nil {
		return entity.Payment{}, err
	}

	return paymentEntity, nil
}
