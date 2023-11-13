package db

import "time"

type Competition struct {
	CID       int `gorm:"primaryKey"`
	StartTime time.Time
	Cname     string `sql:"type:varchar(20)"`
	Sex       bool
	Isgroup   bool
}

type Student struct {
	SID   int    `gorm:"primaryKey"`
	Sname string `sql:"type:varchar(12)"`
	Sex   bool
	Class int16
	Major string `sql:"type:varchar(20)"`
}

type Teacher struct {
	TID   int    `gorm:"primaryKey"`
	Tname string `sql:"type:varchar(12)"`
}

type Volunteer struct {
	CID      int
	SID      int
	Position string `sql:"type:varchar(20)"`
}

type Manager struct {
	TID      int
	CID      int
	Position string `sql:"type:varchar(20)"`
}

type Player struct {
	SID int
	CID int
}

type Playergroup struct {
	GID int `gorm:"primaryKey"`
	SID int
	CID int
}

type Result struct {
	CID     int
	Isgroup bool
	SGID    int
	Ranking int16  // assuming smallint is 16-bit wide, adjust accordingly if it's not the case
	Detail  string `sql:"type:varchar(20)"`
}
