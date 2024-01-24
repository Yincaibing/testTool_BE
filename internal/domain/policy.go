package domain

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"gitlab.iglooinsure.com/axinan/backend/turbo/new-executor/values"
	"time"
)

type Policy struct {
	Base
	BaseID
	ProductKey         string              `gorm:"type:varchar(255);column:product_key"`
	PlanKey            string              `gorm:"type:varchar(255);column:plan_key"`
	QuotationID        uint64              `gorm:"column:quotation_id"`
	ApplicationID      uint64              `gorm:"index;column:application_id"`
	UnderWritingID     uint64              `gorm:"column:under_writing_id"`
	CoiNo              string              `gorm:"type:varchar(100);column:coi_no"`
	CoiDoc             string              `gorm:"type:varchar(500);column:coi_doc"`
	IssuedAt           *time.Time          `gorm:"type:DATETIME;default:null;column:issued_at"`
	ExpireAt           *time.Time          `gorm:"type:DATETIME;default:null;column:expire_at"`
	StartAt            *time.Time          `gorm:"type:DATETIME;default:null;column:start_at"`
	EndAt              *time.Time          `gorm:"type:DATETIME;default:null;column:end_at"`
	CancelAt           *time.Time          `gorm:"type:DATETIME;default:null;column:cancel_at"`
	RefundAt           *time.Time          `gorm:"type:DATETIME;default:null;column:refund_at"`
	TerminateAt        *time.Time          `gorm:"type:DATETIME;default:null;column:terminate_at"`
	ClaimableStartTime *time.Time          `gorm:"type:DATETIME;default:null;column:claimable_start_time"`
	ClaimableEndTime   *time.Time          `gorm:"type:DATETIME;default:null;column:claimable_end_time"`
	CoverStartTime     *time.Time          `gorm:"type:DATETIME;default:null;column:cover_start_time"`
	CoverEndTime       *time.Time          `gorm:"type:DATETIME;default:null;column:cover_end_time"`
	CustomerID         string              `gorm:"column:customer_id"`
	State              string              `gorm:"type:varchar(100);column:state"`
	DisplayID          string              `gorm:"uniqueIndex;type:varchar(100);column:display_id"`
	PolicyHolder       *values.DynamicForm `gorm:"type:json;column:policy_holder"`
	PolicyBaseInfo     *values.DynamicForm `gorm:"type:json;column:policy_base_info"`
	InsuredInfo        *values.DynamicForm `gorm:"type:json;column:insured_info"`
	Beneficiary        *values.DynamicForm `gorm:"type:json;column:beneficiary"`
	SOIInfo            *values.DynamicForm `gorm:"type:json;column:soi_info"`
	PaymentInfo        *values.DynamicForm `gorm:"type:json;column:payment_info"`
	Preference         *values.DynamicForm `gorm:"type:json;column:preference"`
	Meta               *values.DynamicForm `gorm:"type:json;column:meta"`
	PaymentState       string              `gorm:"type:varchar(500);payment_state"`
	InsuredAmount      decimal.Decimal     `gorm:"column:insured_amount"`
	PremiumAmount      decimal.Decimal     `gorm:"column:premium_amount"`
	// cancel reason/ refund reason or other comments
	Reason       string               `gorm:"type:varchar(500);column:reason"`
	ClaimDetails *values.ClaimDetails `gorm:"type:json;column:claim_details"`
	claimAble    bool                 `json:"-"`
}

func (Policy) TableName() string {
	return "policy"
}

func GetPolicyByAnyID(db *gorm.DB, anyId string) (*Policy, error) {
	var policy Policy
	if err := db.Debug().Where("display_id = ? OR insure_process_id = ? OR external_id = ? OR id = ? OR application_id = ? OR under_writing_id = ?", anyId, anyId, anyId, anyId, anyId, anyId).First(&policy).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &policy, nil
}
