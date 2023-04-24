package entity

import "time"

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/4/24
* Time: 10:27
**/

type RequestRecord struct {
	Id uint64 `gorm:"pk;autoIncrement" json:"id"`
	//UserId 用户标记
	UserId uint64 `json:"user_id"`
	//RequestId 请求标记
	RequestId string `gorm:"size:64;not null" json:"request_id"`
	//RequestIP 请求方IP地址,支持IPV6
	RequestIP string `gorm:"size:40;not null" json:"request_ip"`
	//RequestURI 请求地址
	RequestURI string `gorm:"size:255;not null" json:"request_uri"`
	//RequestData 请求参数
	RequestData string `gorm:"type:longtext" json:"request_data"`
	//CreatedAt 数据创建时间
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
