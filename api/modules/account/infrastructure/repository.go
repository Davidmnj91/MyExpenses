package infrastructure

import (
	"database/sql"
	"github.com/Davidmnj91/MyExpenses/modules/account/domain"
	"github.com/Davidmnj91/MyExpenses/modules/account/entity"
)

type AccountRepository struct {
	db *sql.DB
}

// NewRepository create new repository
func NewRepository(connection *sql.DB) domain.AccountRepository {
	return &AccountRepository{db: connection}
}

// Save save given domain object
func (m *AccountRepository) Save(account domain.Account) error {
	accountEntity := convertDomainToEntity(account)

	query := `INSERT INTO accounts(id, name, password_hashed, password_cost) VALUES (?, ?, ?, ?)`

	stmt, err := m.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err = stmt.Exec(accountEntity.ID, accountEntity.Name, accountEntity.Password.Hashed, accountEntity.Password.Cost); err != nil {
		return err
	}

	return nil
}

func (m *AccountRepository) Update(account domain.Account) error {
	accountEntity := convertDomainToEntity(account)

	query := `UPDATE accounts SET name = ?, password_hashed = ?, password_cost = ?, updated_at = ?, closed_at = ? WHERE id = ?`

	stmt, err := m.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err = stmt.Exec(accountEntity.Name, accountEntity.Password.Hashed, accountEntity.Password.Cost, accountEntity.UpdatedAt, accountEntity.CloseAt, accountEntity.ID); err != nil {
		return err
	}

	return nil
}

// FindByID find domain object using id
func (m *AccountRepository) FindByID(id string) (domain.Account, error) {
	query := `SELECT * FROM accounts WHERE id=?`

	stmt, err := m.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Query(id)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	var accountEntity = &entity.Account{}
	if err := result.Scan(&accountEntity.ID, &accountEntity.Name, &accountEntity.CreatedAt, &accountEntity.UpdatedAt); err != nil {
		return nil, err
	}

	return convertEntityToDomain(*accountEntity), nil
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
		CreatedAt: account.ToAnemic().OpenedAt,
		UpdatedAt: account.ToAnemic().UpdatedAt,
		CloseAt:   account.ToAnemic().ClosedAt,
	}
}
