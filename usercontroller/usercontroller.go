package usercontroller

import (
	"time"
)

// User query record
type queryRecord struct {
	// TODO it should be replace by my own api key
	ip   string
	time time.Time
}

// TODO should be replace by database
var queryRecords = []queryRecord{}

// MaxTime means max query times per day
var MaxTime = 0

// QuerryAcquire will check is user reach it's limit of daily query
func QuerryAcquire(ip string) bool {
	if MaxTime == 0 || queryIn24Hours(ip) < MaxTime {
		queryRecords = append(queryRecords, queryRecord{ip, time.Now()})
		return true
	}

	return false
}

func queryIn24Hours(ip string) int {
	now := time.Now()
	count := 0
	for _, record := range queryRecords {
		if record.ip == ip && now.Sub(record.time) < 24*time.Hour {
			count++
		}
	}

	return count
}

// AddQueryRecord is a test only function
func AddQueryRecord(ip string, time time.Time) {
	queryRecords = append(queryRecords, queryRecord{ip, time})
}
