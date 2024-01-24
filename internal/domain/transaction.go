package domain

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"gitlab.iglooinsure.com/axinan/backend/integpay/gateway/entity"
	"gorm.io/datatypes"
	"time"
)

type Transaction struct {
	Base
	State                 entity.TxState `gorm:"size:64"`
	Title                 string         `gorm:"size:256"`
	UserID                string         `gorm:"size:64"`
	Amount                decimal.Decimal
	PayAction             entity.PayAction `gorm:"size:64"`
	Service               string           `gorm:"size:64"`
	Currency              string           `gorm:"size:16"`
	ObjectID              string           `gorm:"size:64;uniqueIndex:idx_svc_txnid,priority:1"`
	ProductKey            string           `gorm:"size:64;uniqueIndex:idx_svc_txnid,priority:2"`
	ObjectType            string           `gorm:"size:64;uniqueIndex:idx_svc_txnid,priority:3"`
	Description           string
	FailedMsg             string
	NotifyURL             string
	SucceedAt             *time.Time `gorm:"null"`
	UsedProvider          string     `gorm:"size:64"`
	TransactionID         string     `gorm:"size:64;uniqueIndex"`
	InsureProcessID       string     `gorm:"size:64;index"`
	FailedRedirectURL     string
	SucceedRedirectURL    string
	PendingRedirectURL    string
	ProviderData          datatypes.JSON `gorm:"type:json"`
	TargetPrice           decimal.Decimal
	PromotionRedemptionId string `gorm:"size:64"`
	PromotionStatus       string `gorm:"size:64"`
}

func (Transaction) TableName() string {
	return "transaction"
}

func GetTransactionByAnyID(db *gorm.DB, anyId string) (*Transaction, error) {
	var transaction Transaction
	if err := db.Debug().Where("object_id = ? OR transaction_id = ? OR user_id = ? OR insure_process_id = ?", anyId, anyId, anyId, anyId).First(&transaction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &transaction, nil
}
