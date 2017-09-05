package lib

import (
	"time"
)

func TsToShortdate(ts string) (rlt string) {
	rlt = ""
	if dtX, err := time.Parse("2006-01-02", ts[:10]); err == nil {
		rlt = dtX.Format("20060102")
	}
	return
}
