package model

import "time"

type Group struct {
	ID        string    `json:"id" example:"c6a52781-c927-42ec-bc9e-254ddde5f904"`
	Name      string    `json:"name" example:"MyGroup"`
	CreatedAt time.Time `json:"createdAt" example:"2019-12-23 12:27:37"`
	UpdatedAt time.Time `json:"updatedAt" example:"2019-12-23 12:27:37"`
	ClosedAt  time.Time `json:"closedAt" example:"2019-12-23 12:27:37"`
}
