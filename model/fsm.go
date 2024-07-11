package model

import (
	"time"
)

// Data 业务字段表，必须，根据具体业务设计字段
type Data struct {
	ID              uint64  `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint(20) unsigned;not null"`
	TaskID          string  `gorm:"index:idx_task_id;column:task_id;type:char(32);not null"`
	Symbol          string  `gorm:"column:symbol;type:varchar(20);not null;default:''"`
	Quantity        float64 `gorm:"column:quantity;type:decimal(50,15);not null;default:0.000000000000000"`
	Amount          float64 `gorm:"column:amount;type:decimal(50,15);not null;default:0.000000000000000"`
	Operator        string  `gorm:"column:operator;type:varchar(20);not null;default:''"`
	TransactionTime uint64  `gorm:"column:transaction_time;type:bigint(20) unsigned;not null;default:0;comment:'业务时间'"` // 业务时间
	Comment         string  `gorm:"column:comment;type:varchar(128);not null;default:'';comment:'备注说明'"`                // 备注说明
}

// TableName get sql table name.获取数据库表名
func (m *Data) TableName() string {
	return "data"
}

// Task 任务主表，必须，维护状态驱动执行
type Task struct {
	ID         string    `gorm:"primaryKey;column:id;type:char(32);not null"`
	RequestID  string    `gorm:"unique;column:request_id;type:char(32);not null;comment:'初始请求ID'"`       // 初始请求ID
	Type       string    `gorm:"column:type;type:varchar(128);not null;comment:'业务类型'"`                  // 业务类型
	State      string    `gorm:"index:idx_state;column:state;type:varchar(128);not null;comment:'任务状态'"` // 任务状态
	Version    uint32    `gorm:"column:version;type:int(10) unsigned;not null;default:1"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;default:current_timestamp()"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;not null;default:current_timestamp()"`
}

// TableName get sql table name.获取数据库表名
func (m *Task) TableName() string {
	return "task"
}

// UniqueRequest 防重表，必须，创建更新操作对成功幂等
type UniqueRequest struct {
	ID         uint64    `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint(20) unsigned;not null"`
	RequestID  string    `gorm:"unique;column:request_id;type:char(32);not null;comment:'对成功幂等'"` // 对成功幂等
	TaskID     string    `gorm:"column:task_id;type:char(32);not null"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;default:current_timestamp()"`
}

// TableName get sql table name.获取数据库表名
func (m *UniqueRequest) TableName() string {
	return "unique_request"
}
