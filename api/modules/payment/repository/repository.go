package repository

import (
	"database/sql"
	entity2 "github.com/Davidmnj91/MyExpenses/modules/payment/entity"
	model2 "github.com/Davidmnj91/MyExpenses/modules/payment/model"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
	"time"
)

// Repository repository interface
type Repository interface {
	Create(paymentID string, concept string, amount float32, paymentConfiguration model2.PaymentConfiguration) (entity2.Payment, error)
	Update(paymentID string, concept string, amount float32, paymentConfiguration model2.PaymentConfiguration) (entity2.Payment, error)
	FindByID(paymentID string) (entity2.Payment, error)
}

// PaymentRepository repository for profile data
type PaymentRepository struct {
	db *sql.DB
}

// New create repository instance
func New(db *sql.DB) Repository {
	return &PaymentRepository{db: db}
}

// Create create payment data in database
func (repository PaymentRepository) Create(paymentID string, concept string, amount float32, paymentConfiguration model2.PaymentConfiguration) (entity2.Payment, error) {
	payment := entity2.Payment{ID: paymentID}

	findPaymentTx := repository.db.First(&payment)

	if err := findPaymentTx.Error; err != nil {
		return entity2.Payment{}, err
	}

	paymentConfigurationJson := pgtype.JSONB{}
	if err := pgtype.JSONB.Scan(&paymentConfigurationJson, paymentConfiguration); err != nil {
		return entity2.Payment{}, err
	}

	paymentEntity := entity2.Payment{
		ID:                   uuid.NewString(),
		Concept:              concept,
		Amount:               amount,
		PaymentConfiguration: paymentConfigurationJson,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	createPaymentTx := repository.db.Create(paymentEntity)

	if err := createPaymentTx.Error; err != nil {
		return entity2.Payment{}, err
	}

	return paymentEntity, nil
}

// Update update payment data by paymentID
func (repository PaymentRepository) Update(paymentID string, concept string, amount float32, paymentConfiguration model2.PaymentConfiguration) (entity2.Payment, error) {
	paymentEntity := entity2.Payment{ID: paymentID}

	paymentConfigurationJson := pgtype.JSONB{}
	if err := pgtype.JSONB.Scan(&paymentConfigurationJson, paymentConfiguration); err != nil {
		return entity2.Payment{}, err
	}

	paymentUpdate := entity2.Payment{
		Concept:              concept,
		Amount:               amount,
		PaymentConfiguration: paymentConfigurationJson,
		UpdatedAt:            time.Now(),
	}
	updatedPaymentTx := repository.db.Model(&paymentEntity).Updates(paymentUpdate)

	if err := updatedPaymentTx.Error; err != nil {
		return entity2.Payment{}, err
	}

	return paymentEntity, nil
}

// FindByID find payment data by payment id
func (repository PaymentRepository) FindByID(paymentID string) (entity2.Payment, error) {
	paymentEntity := entity2.Payment{ID: paymentID}

	findPaymentTx := repository.db.First(&paymentEntity)

	if err := findPaymentTx.Error; err != nil {
		return entity2.Payment{}, err
	}

	return paymentEntity, nil
}
