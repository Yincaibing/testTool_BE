package domain

import (
	"github.com/jinzhu/gorm"
	"gitlab.iglooinsure.com/axinan/backend/turbo/new-executor/values"
)

type QuotationForm struct {
	FastQuote    *values.DynamicForm   `gorm:"type:json;column:fast_quote"`
	Addons       *values.BenefitParams `gorm:"type:json;column:add_ons"`
	DiscountInfo *values.DiscountInfo  `gorm:"type:json;column:discount_info"`
}

type Application struct {
	QuotationID   uint64              `gorm:"column:quotation_id"`
	Details       *values.DynamicForm `gorm:"type:json;column:details"`
	State         string              `gorm:"type:varchar(255);column:state"`
	Meta          *values.DynamicForm `json:"meta" gorm:"type:json;column:meta"`
	PlanKey       string              `gorm:"type:varchar(255);column:plan_key"`
	ProductKey    string              `gorm:"type:varchar(255);column:product_key"`
	OriginDetails *values.DynamicForm `gorm:"type:json;column:origin_details"`
	Base
	BaseID
	QuotationForm
	Outputs     *values.QuotationOutputs `gorm:"type:json;column:outputs"`
	PremiumInfo *values.PremiumInfo      `gorm:"type:json;column:premium_info"`
}

func (Application) TableName() string {
	return "application"
}

func GetApplicationByAnyId(db *gorm.DB, anyId string) (*Application, error) {
	var application Application
	if err := db.Debug().Where("quotation_id = ? OR insure_process_id = ? OR external_id = ? OR id = ?", anyId, anyId, anyId, anyId).First(&application).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &application, nil
}
