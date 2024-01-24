package domain

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"gitlab.iglooinsure.com/axinan/backend/turbo/common/modules/adapter/fact"
	"gitlab.iglooinsure.com/axinan/backend/turbo/new-executor/values"
	"time"
)

type ClaimInfo struct {
	BasicInfo fact.ClaimBasicInfoFact `json:"basic_info"`
	values.ClaimInfo
}

type Claim struct {
	Base
	BaseID
	ClaimAmount     decimal.Decimal       `gorm:"column:claim_amount"`
	ApproveAmount   decimal.Decimal       `gorm:"column:approve_amount"`
	ApprovedAt      *time.Time            `gorm:"type:DATETIME;default:null;column:approved_at"`
	ClaimForm       *values.DynamicForm   `gorm:"type:json;column:claim_form"`
	DisplayID       string                `gorm:"type:varchar(255);column:display_id"`
	ProductKey      string                `gorm:"type:varchar(255);column:product_key"`
	PlanKey         string                `gorm:"type:varchar(255);column:plan_key"`
	PolicyID        uint64                `gorm:"index;column:policy_id"`
	PolicyDisplayID string                `gorm:"column:policy_display_id"`
	State           string                `gorm:"index;type:varchar(255);column:state"`
	Details         *values.EntityDetails `gorm:"type:json;column:claim_details"`
	//SelectedBenefits *values.SelectedBenefits `gorm:"type:json;column:selected_benefits"`
	ClaimInfo    *values.ClaimInfo    `gorm:"type:json;column:claim_info"`
	ClaimOutPuts *values.ClaimOutputs `gorm:"type:json;column:outputs"`
	// For submission advance rule, it's advance rule's input
	SubmissionExternalParams *values.DynamicForm `gorm:"type:json;column:submission_external_params"`
}

func (Claim) TableName() string {
	return "claim"
}

func GetClaimByAnyId(db *gorm.DB, anyId string) (*Claim, error) {
	var claim Claim
	if err := db.Debug().Where("display_id = ? OR insure_process_id = ? OR external_id = ? OR policy_id = ? OR id = ?", anyId, anyId, anyId, anyId, anyId).First(&claim).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &claim, nil
}
