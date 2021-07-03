package entity

import (
	"time"
)

type Admin struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	//[ 1] number                                         bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	Number int64 `gorm:"column:number;type:bigint;"`
	//[ 2] name                                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name string `gorm:"column:name;type:varchar;size:255;"`
	//[ 3] age                                            int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Age int `gorm:"column:age;type:int;"`
	//[ 4] active                                         tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	Active bool `gorm:"column:active;type:tinyint;"`
	//[ 5] note                                           text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Note string `gorm:"column:note;type:text;size:65535;"`
	//[ 6] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;"`
	//[ 7] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;"`
}
