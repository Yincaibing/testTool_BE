package domain

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"gitlab.iglooinsure.com/axinan/backend/turbo/common/modules/adapter/fact"
	"gitlab.iglooinsure.com/axinan/backend/turbo/new-executor/values"
)

type Reimburse struct {
	Base
	BaseID
	DisplayID        string                   `gorm:"type:varchar(255);column:display_id"`
	ProductKey       string                   `gorm:"type:varchar(255);column:product_key"`
	PlanKey          string                   `gorm:"type:varchar(255);column:plan_key"`
	ClaimID          uint64                   `gorm:"index;column:claim_id"`
	ClaimDisplayID   string                   `gorm:"index;column:claim_display_id"`
	State            string                   `gorm:"type:varchar(255);column:state"`
	Details          *values.EntityDetails    `gorm:"type:json;column:reimburse_details"`
	ApprovedAmount   decimal.Decimal          `gorm:"column:approved_amount"`
	ReimburseMethod  *values.ReimburseMethod  `gorm:"type:json;column:reimburse_method"`
	ReimburseForm    *values.DynamicForm      `gorm:"type:json;column:reimburse_form"`
	ReimburseInfo    *values.ReimburseInfo    `gorm:"type:json;column:reimburse_info"`
	ReimburseOutputs *values.ReimburseOutputs `gorm:"type:json;column:reimburse_outputs"`
}

type ReimburseInfo struct {
	BasicInfo fact.ReimbursementBasicInfoFact `json:"basic_info"`
	values.ReimburseInfo
}

func (Reimburse) TableName() string {
	return "reimburse"
}

func GetReimburseByAnyID(db *gorm.DB, anyId string) (*Reimburse, error) {
	var reimburse Reimburse
	if err := db.Debug().Where("display_id = ? OR external_id = ? OR id = ? OR claim_id = ? OR insure_process_id = ?", anyId, anyId, anyId, anyId, anyId).First(&reimburse).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &reimburse, nil
}
