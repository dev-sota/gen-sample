package entity

import (
	"time"
)

type Task struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	//[ 1] admin_id                                       bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	AdminID int64 `gorm:"column:admin_id;type:bigint;"`
	//[ 2] name                                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name string `gorm:"column:name;type:varchar;size:255;"`
	//[ 3] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;"`
	//[ 4] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;"`
}
