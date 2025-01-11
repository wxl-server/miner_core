package domain

import (
	"miner_core/common/constants"
	"time"
)

type JobDO struct {
	ID          int64
	Name        string
	Description string
	CreatedBy   int64
	UpdatedBy   int64
	Extra       *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type QueryJobListReq struct {
	PageNum  int64
	PageSize int64
	OrderBy  *string
	Order    *constants.Order

	ID             *int64
	CreatedBy      *int64
	CreatedAtStart *time.Time
	CreatedAtEnd   *time.Time
}
