package repository

import "github.com/ruohuaii/wufumall/domain/entity"

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/4/24
* Time: 10:34
**/

type RequestRecordRepository interface {
	Save(record *entity.RequestRecord) error
	GetByRequestId(requestId string) (*entity.RequestRecord, error)
}
