package entity

import (
	"time"
)

type Admin struct {
	ID        int64     `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	Number    int64     `gorm:"column:number;type:bigint;"`
	Name      string    `gorm:"column:name;type:varchar;size:255;"`
	Age       int       `gorm:"column:age;type:int;"`
	Active    bool      `gorm:"column:active;type:tinyint;"`
	Note      string    `gorm:"column:note;type:text;size:65535;"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;"`
}
