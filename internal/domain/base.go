package domain

import (
	"strconv"
	"time"
)

// this is a example for domain base struct define
// you can change it according requirement

type Base struct {
	DomainEvent string    `gorm:"-"`
	ID          uint64    `gorm:"primaryKey;autoIncrement;column:id"`
	Version     uint64    `gorm:"column:version"`
	CreatedAt   time.Time `gorm:"->;<-:create;type:DATETIME;default:CURRENT_TIMESTAMP not null;column:created_at"`
	UpdatedAt   time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP;column:updated_at"`
}

type BaseID struct {
	InsureProcessID string `json:"insure_process_id" gorm:"index;type:varchar(255);column:insure_process_id"`
	ExternalID      string `json:"external_id" gorm:"index;type:varchar(255);column:external_id"`
}

func (b *Base) IDString() string {
	if b == nil {
		return ""
	}
	return strconv.FormatUint(b.ID, 10)
}
