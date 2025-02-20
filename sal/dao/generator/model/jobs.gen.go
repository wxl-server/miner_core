// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameJobPO = "jobs"

// JobPO mapped from table <jobs>
type JobPO struct {
	ID          int64          `gorm:"column:id;primaryKey" json:"id"`
	Name        string         `gorm:"column:name;not null" json:"name"`
	Description string         `gorm:"column:description;not null" json:"description"`
	CreatedBy   int64          `gorm:"column:created_by;not null" json:"created_by"`
	UpdatedBy   int64          `gorm:"column:updated_by;not null" json:"updated_by"`
	Extra       *string        `gorm:"column:extra" json:"extra"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName JobPO's table name
func (*JobPO) TableName() string {
	return TableNameJobPO
}
