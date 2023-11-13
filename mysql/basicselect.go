package db

import (
	"time"

	"gorm.io/gorm"
)

func SelectCompetition(db *gorm.DB, cname string, sex bool, startTime time.Time, isgroup bool) (cid int, err error) {
	competition := &Competition{}
	result := db.Where("cname = ? AND sex=? AND startTime =? AND isgroup=?", cname, sex, startTime, isgroup).First(&competition)
	return competition.CID, result.Error
} //检索得到主键
