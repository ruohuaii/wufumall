package persistence

import (
	"github.com/ruohuaii/wufumall/domain/entity"
	"github.com/ruohuaii/wufumall/domain/repository"
	"gorm.io/gorm"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/4/24
* Time: 10:36
**/

var _ repository.RequestRecordRepository = &RequestRecordRepo{}

type RequestRecordRepo struct {
	db *gorm.DB
}

func NewRequestRecordRepo(db *gorm.DB) *RequestRecordRepo {
	return &RequestRecordRepo{db: db}
}

func (r *RequestRecordRepo) Save(record *entity.RequestRecord) error {
	return r.db.Create(record).Error
}

func (r *RequestRecordRepo) GetByRequestId(requestId string) (*entity.RequestRecord, error) {
	var record = &entity.RequestRecord{}
	err := r.db.Where(&entity.RequestRecord{RequestId: requestId}).Find(record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}
