package cron

import (
	"github.com/go-co-op/gocron"
	"time"
)

var Schedulers = map[string]*gocron.Scheduler{}

func New() *gocron.Scheduler {
	tbkk, _ := time.LoadLocation("Asia/Bangkok")
	return gocron.NewScheduler(tbkk)
}
