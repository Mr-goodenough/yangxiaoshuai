package db

import (
	"time"

	"gorm.io/gorm"
)

func InsertCompetition(db *gorm.DB, date time.Time, cname string, sex bool, isgroup bool) {
	newConpetition := Competition{CID: 0, Cname: cname, StartTime: date, Sex: sex, Isgroup: isgroup}
	db.Create(&newConpetition)
}
func InsertStudent(db *gorm.DB, name string, sex bool, class int16, major string) {
	newStudent := Student{SID: 0, Class: class, Sex: sex, Sname: name, Major: major}
	db.Create(&newStudent)
}
func InsertTeacher(db *gorm.DB, name string) {
	newTeacher := Teacher{TID: 0, Tname: name}
	db.Create(&newTeacher)
}
func InsertVolunteer(db *gorm.DB, sid int, cid int, position string) {
	newVolunteer := Volunteer{SID: sid, CID: cid, Position: position}
	db.Create(&newVolunteer)
}
func InsertManager(db *gorm.DB, tid int, cid int, position string) {
	newManager := Manager{TID: tid, CID: cid, Position: position}
	db.Create(&newManager)
}
func InsertPalyer(db *gorm.DB, sid int, cid int) {
	newPlayer := Player{SID: sid, CID: cid}
	db.Create(&newPlayer)
}
func InsertPlayergroup(db *gorm.DB, sid int, cid int) {
	newPlayergroup := Playergroup{GID: 0, SID: sid, CID: cid}
	db.Create(&newPlayergroup)
}
func InsertResult(db *gorm.DB, cid int, sgid int, isgroup bool, ranking int16, detail string) {
	newResult := Result{CID: cid, Isgroup: isgroup, Ranking: ranking, Detail: detail}
	db.Create(&newResult)
}
