package basemodel

import (
	"gsadmin/core/utils/assertion"
	"time"
)

type Model struct {
	CreateId  int       `json:"create_id" gorm:"create_id;size:64"`
	UpdateId  int       `json:"update_id" gorm:"update_id;size:64"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP on update current_timestamp;"`
}

func (m *Model) SetCreate(userId uint) {
	m.CreateId = assertion.AnyToInt(userId)
	m.UpdateId = assertion.AnyToInt(userId)
}

func (m *Model) SetUpdate(userId uint) {
	m.UpdateId = assertion.AnyToInt(userId)
}
