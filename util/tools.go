package util

import (
	"Panorama-Statistics/database"
	"time"
)

var MaxTime, _ = time.ParseInLocation("20060102", "21001230", time.Local)
var MinTime, _ = time.ParseInLocation("20060102", "19000101", time.Local)

// start: 0, end: 1
func TimeHandler(ts string, typ int) time.Time {
	t, err := time.ParseInLocation("20060102", ts, time.Local)
	if err != nil {
		if typ == 0 {
			t = MinTime
		} else {
			t = MaxTime
		}
	}
	return t
}

func ResFormatter(result []database.Statistics) []map[string]string {
	res := make([]map[string]string, 0)
	for _, r := range result {
		item := make(map[string]string)
		item["user"] = r.User
		item["page"] = r.Page
		item["visit_time"] = r.CreatedAt.Format("2006-01-02 15:03:04")
		res = append(res, item)
	}
	return res
}
