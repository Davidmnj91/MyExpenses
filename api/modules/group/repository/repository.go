package repository

import (
	"database/sql"
	entity2 "github.com/Davidmnj91/MyExpenses/modules/account/entity"
	entity3 "github.com/Davidmnj91/MyExpenses/modules/group/entity"
	"github.com/google/uuid"
	"time"
)

// Repository repository interface
type Repository interface {
	Create(accountID string, groupName string) (entity3.Group, error)
	Update(groupID string, newGroupName string) (entity3.Group, error)
	Close(groupID string) (entity3.Group, error)
	FindByID(groupID string) (entity3.Group, error)
}

// GroupRepository repository for profile data
type GroupRepository struct {
	db *sql.DB
}

// New create repository instance
func New(db *sql.DB) Repository {
	return &GroupRepository{db: db}
}

// Create create group data in database
func (repository GroupRepository) Create(accountID string, groupName string)(entity3.Group, error)  {
	accountEntity := entity2.Account{ID: accountID}

	findAccountTx := repository.db.First(&accountEntity)

	if err := findAccountTx.Error; err != nil {
		return entity3.Group{}, err
	}

	groupEntity := entity3.Group{
		ID:        uuid.NewString(),
		Name:      groupName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ClosedAt:   time.Time{},
	}

	createGroupTx := repository.db.Create(groupEntity)

	if err := createGroupTx.Error; err != nil {
		return entity3.Group{}, err
	}

	return groupEntity, nil
}

// Update update group data by groupID
func (repository GroupRepository) Update(groupID string, newGroupName string) (entity3.Group, error) {
	groupEntity := entity3.Group{ID: groupID}

	updatedGroupTx := repository.db.Model(&groupEntity).Updates(entity3.Group{Name: newGroupName})

	if err := updatedGroupTx.Error; err != nil {
		return entity3.Group{}, err
	}

	return groupEntity, nil
}

// Close close group by groupID
func (repository GroupRepository) Close(groupID string) (entity3.Group, error) {
	groupEntity := entity3.Group{ID: groupID}

	findGroupTx := repository.db.First(&groupEntity)

	if err := findGroupTx.Error; err != nil {
		return entity3.Group{}, err
	}

	closeGroupTx := repository.db.Model(&groupEntity).Updates(entity3.Group{ClosedAt: time.Now()})

	if err := closeGroupTx.Error; err != nil {
		return entity3.Group{}, err
	}

	return groupEntity, nil
}


// FindByID find group data by group id
func (repository GroupRepository) FindByID(groupID string) (entity3.Group, error) {
	groupEntity := entity3.Group{ID: groupID}

	findGroupTx := repository.db.First(&groupEntity)

	if err := findGroupTx.Error; err != nil {
		return entity3.Group{}, err
	}

	return groupEntity, nil
}
