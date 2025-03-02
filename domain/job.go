package domain

import (
	"miner_core/common/consts"
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

type QueryJobListReqDO struct {
	PageNum  int64
	PageSize int64
	OrderBy  *string
	Order    *consts.Order

	ID             *int64
	Name           *string
	CreatedBy      *int64
	CreatedAtStart *time.Time
	CreatedAtEnd   *time.Time
}
