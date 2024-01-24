package domain

import (
	"github.com/jinzhu/gorm"
	"gitlab.iglooinsure.com/axinan/backend/turbo/new-executor/values"
)

type EngineFlowProcess struct {
	Base
	BaseID
	ProductKey    string                `gorm:"index;type:varchar(255);column:product_key"`
	PlanKey       string                `gorm:"index;type:varchar(255);column:plan_key"`
	IsFreezed     bool                  `gorm:"type:tinyint(1);column:is_freezed"`
	TemplateID    uint64                `gorm:"index;type:varchar(255);column:template_id"`
	EngineFlowID  uint64                `gorm:"index;type:varchar(255);column:engine_flow_id"`
	ProcessDetail *values.ProcessDetail `gorm:"type:json;column:process_detail"`
}

func (EngineFlowProcess) TableName() string {
	return "engine_flow_process"
}

func GetEngineFlowProcessByAnyId(db *gorm.DB, anyId string) (*EngineFlowProcess, error) {
	var engine_flow_process EngineFlowProcess
	if err := db.Debug().Where("template_id = ? OR insure_process_id = ? OR external_id = ? OR id = ? OR engine_flow_id = ?", anyId, anyId, anyId, anyId, anyId).First(&engine_flow_process).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &engine_flow_process, nil
}
