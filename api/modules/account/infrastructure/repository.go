package infrastructure

import (
	"github.com/Davidmnj91/MyExpenses/account/domain"
	"github.com/Davidmnj91/MyExpenses/account/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

// NewRepository create new repository
func NewRepository(connection *gorm.DB) domain.AccountRepository {
	return &AccountRepository{db: connection}
}

// Save save given domain object
func (m *AccountRepository) Save(account domain.Account) {
	m.db.Save(convertDomainToEntity(account))
}

// FindNewID find new domain object id
func (m *AccountRepository) FindNewID() (string, error) {
	newID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	m.db.Save(entity.Account{ID: newID.String()})
	return newID.String(), nil
}

// FindByID find domain object using id
func (m *AccountRepository) FindByID(id string) (domain.Account, error) {
	var accountEntity entity.Account
	if err := m.db.First(&accountEntity).Error; err != nil {
		return nil, err
	}

	return convertEntityToDomain(accountEntity), nil
}

func convertEntityToDomain(entity entity.Account) domain.Account {
	return domain.ReconstituteAccount(domain.AccountAnemic{
		ID:   entity.ID,
		Name: entity.Name,
		Password: domain.PasswordAnemic{
			Hashed: entity.Password.Hashed,
			Cost:   entity.Password.Cost,
		},
	})
}

func convertDomainToEntity(account domain.Account) entity.Account {
	return entity.Account{
		ID:   account.ToAnemic().ID,
		Name: account.ToAnemic().Name,
		Password: entity.Password{
			Hashed: account.ToAnemic().Password.Hashed,
			Cost:   account.ToAnemic().Password.Cost,
		},
		CreatedAt:  account.ToAnemic().OpenedAt,
		UpdatedAt: account.ToAnemic().UpdatedAt,
		CloseAt:   account.ToAnemic().ClosedAt,
	}
}
