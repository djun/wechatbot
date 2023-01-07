package rule

import (
	"sync"
	"time"
)

const (
	STARTTIME = 9
	ENDTIME   = 21
)

type Rule struct{}

var isWork = true
var Grule = &Rule{}
var lock sync.Mutex

func (r *Rule) SetWork(work bool) {
	lock.Lock()
	defer lock.Unlock()
	isWork = work
	return
}
func (r *Rule) GetWork() bool {
	lock.Lock()
	defer lock.Unlock()
	return isWork
}

// 判断时间在今天的早上 9点到 晚上 9 点区间内
func (r *Rule) IsWorkTime(s int, e int) bool {
	if s < 0 || s > 24 {
		s = STARTTIME
	}
	if e < 0 || e > 24 || e <= s {
		e = ENDTIME
	}
	t := time.Now()
	startTime := time.Date(t.Year(), t.Month(), t.Day(), s, 0, 0, 0, time.Local)
	endTime := time.Date(t.Year(), t.Month(), t.Day(), e, 0, 0, 0, time.Local)
	// 判断当前时间是否在当天的 STARTTIME  --  ENDTIME
	if t.Unix() > startTime.Unix() && t.Unix() < endTime.Unix() {
		return true
	}
	return false
}

func (r *Rule) InSlice(str string, sli []string) bool {
	for _, v := range sli {
		if v == str {
			return true
		}
	}
	return false
}
