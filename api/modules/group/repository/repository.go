package repository

import (
	account "github.com/Davidmnj91/MyExpenses/account/entity"
	"github.com/Davidmnj91/MyExpenses/group/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Repository repository interface
type Repository interface {
	Create(accountID string, groupName string) (entity.Group, error)
	Update(groupID string, newGroupName string) (entity.Group, error)
	Close(groupID string) (entity.Group, error)
	FindByID(groupID string) (entity.Group, error)
}

// GroupRepository repository for profile data
type GroupRepository struct {
	db *gorm.DB
}

// New create repository instance
func New(db *gorm.DB) Repository {
	return &GroupRepository{db: db}
}

// Create create group data in database
func (repository GroupRepository) Create(accountID string, groupName string)(entity.Group, error)  {
	accountEntity := account.Account{ID: accountID}

	findAccountTx := repository.db.First(&accountEntity)

	if err := findAccountTx.Error; err != nil {
		return entity.Group{}, err
	}

	groupEntity := entity.Group{
		ID:        uuid.NewString(),
		Name:      groupName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ClosedAt:   time.Time{},
	}

	createGroupTx := repository.db.Create(groupEntity)

	if err := createGroupTx.Error; err != nil {
		return entity.Group{}, err
	}

	return groupEntity, nil
}

// Update update group data by groupID
func (repository GroupRepository) Update(groupID string, newGroupName string) (entity.Group, error) {
	groupEntity := entity.Group{ID: groupID}

	updatedGroupTx := repository.db.Model(&groupEntity).Updates(entity.Group{Name: newGroupName})

	if err := updatedGroupTx.Error; err != nil {
		return entity.Group{}, err
	}

	return groupEntity, nil
}

// Close close group by groupID
func (repository GroupRepository) Close(groupID string) (entity.Group, error) {
	groupEntity := entity.Group{ID: groupID}

	findGroupTx := repository.db.First(&groupEntity)

	if err := findGroupTx.Error; err != nil {
		return entity.Group{}, err
	}

	closeGroupTx := repository.db.Model(&groupEntity).Updates(entity.Group{ClosedAt: time.Now()})

	if err := closeGroupTx.Error; err != nil {
		return entity.Group{}, err
	}

	return groupEntity, nil
}


// FindByID find group data by group id
func (repository GroupRepository) FindByID(groupID string) (entity.Group, error) {
	groupEntity := entity.Group{ID: groupID}

	findGroupTx := repository.db.First(&groupEntity)

	if err := findGroupTx.Error; err != nil {
		return entity.Group{}, err
	}

	return groupEntity, nil
}
