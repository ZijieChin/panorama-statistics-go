package database

import (
	"time"
)

func InsertStat(stat Statistics) error {
	result := DB.Create(&stat)
	return result.Error
}

func GetStat(user, page string, start, end time.Time) (stats []Statistics) {
	stat := &Statistics{
		User: user,
		Page: page,
	}
	DB.Select("user", "page", "created_at").Where(stat).Where("created_at BETWEEN ? AND ?", start, end).Find(&stats)
	return
}
