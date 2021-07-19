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

func GetTimes(user string) int {
	stat := []Statistics{}
	now := time.Now()
	// time interval: 10 minutes
	interval, _ := time.ParseDuration("-10m")
	justnow := now.Add(interval)
	DB.Select("id").Where("user = ? AND created_at < ? AND created_at > ?", user, now, justnow).Find(&stat)
	aaa := len(stat)
	return aaa
}
