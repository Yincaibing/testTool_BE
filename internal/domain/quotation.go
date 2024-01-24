package domain

import (
	"github.com/jinzhu/gorm"
	"gitlab.iglooinsure.com/axinan/backend/turbo/new-executor/values"
	"time"
)

type Quotation struct {
	Base
	BaseID
	PlanID     uint64              `gorm:"column:plan_id"`
	ProductKey string              `gorm:"type:varchar(255);column:product_key"`
	PlanKey    string              `gorm:"type:varchar(255);column:plan_key"`
	Meta       *values.DynamicForm `gorm:"type:json;column:meta"`
	State      string              `gorm:"type:varchar(255);column:state"`
	QuotationForm
	Outputs     *values.QuotationOutputs `gorm:"type:json;column:outputs"`
	PremiumInfo *values.PremiumInfo      `gorm:"type:json;column:premium_info"`
	ExpireAt    *time.Time               `gorm:"type:DATETIME;default:null;column:expire_at"`
	CancelAt    *time.Time               `gorm:"type:DATETIME;default:null;column:cancel_at"`
	ExpireDays  int                      `gorm:"type:int;default:0;column:expire_days"`
}

func (Quotation) TableName() string {
	return "quotation"
}

func GetQuotationByAnyID(db *gorm.DB, anyId string) (*Quotation, error) {
	var quotation Quotation
	if err := db.Debug().Where("external_id = ? OR id = ? OR insure_process_id = ?", anyId, anyId, anyId).First(&quotation).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &quotation, nil
}
