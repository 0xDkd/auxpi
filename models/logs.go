package models

import (
	"auxpi/auxpiAll"
)

type Log struct {
	Model

	Type    string `json:"type" gorm:"size:32"`
	Content string `json:"content" gorm:"size:255"`
	Part    string `json:"part" gorm:"size:32"`
	Level   string `json:"level" gorm:"size:32"`
}

//创建日志
func AddLog(t, content, part, level string) {
	db.Create(&Log{
		Type:    t,
		Content: content,
		Part:    part,
		Level:   level,
	})
}

//查询日志
func GetLogs(offset, limit int, maps interface{}) (log []Log, count int) {
	err := db.Model(&Log{}).
		Where(maps).
		Count(&count).
		Offset(offset).
		Limit(limit).
		Find(&log).Error

	modelsError(auxpi.ErrorToString(err))

	return
}

//查询 api 使用情况
func GetApiInfo() (apis []Report) {
	err := db.Model(&Log{}).
		Where("type=?", "API Call").
		Select("COUNT(*) AS `number` , created_day AS `date`").
		Order("created_day ASC").
		Group("`created_day`").
		Limit(7).
		Scan(&apis).Error
	modelsError(auxpi.ErrorToString(err))

	return
}

//增加 API 调用记录
func AddApiLog(content string) {
	db.Create(&Log{
		Type:    "API Call",
		Content: content,
		Part:    "SYSTEM",
		Level:   "NONE",
	})
}

//迁移Logs 数据表
func MigrateLogs() error {
	if db.HasTable(&Log{}) {
		err := db.DropTable(&Log{}).Error
		err = db.CreateTable(&Log{}).Error
		return err
	} else {
		err := db.CreateTable(&Log{}).Error
		return err
	}
}
