package builtin

import (
	"encoding/json"
	"fmt"
	"time"
	"xx.com/yyy/smart-schedule/configuration"
	"xx.com/yyy/smart-schedule/internal/smart-schedule/api/scheduler"
	"xx.com/yyy/smart-schedule/pkg/cron"
)

type BuildIn interface {
	Register()
}

type buildIn struct {
}

func (b *buildIn) Register() {
	cr := cron.New()
	ts := scheduler.NewTask()

	config := configuration.Config
	for _, s := range config.Schedulers {
		// Cron expressions supported: https://crontab.guru/
		_, err := cr.Cron(s.Cron).Do(func(s configuration.Scheduler) {
			go func(s configuration.Scheduler) {
				// Parse body & header
				body := map[string]string{}
				header := map[string]string{}
				_ = json.Unmarshal([]byte(s.Task.Body), &body)
				_ = json.Unmarshal([]byte(s.Task.Header), &header)

				fmt.Println(time.Now().Format(time.DateTime), "Task", s.Job, "is running...")
				ts.ApiRequest(scheduler.CreateScheduler{
					Job:  s.Job,
					Cron: s.Cron,
					Task: scheduler.CreateSchedulerTask{
						Type: "api",
						Config: scheduler.CreateSchedulerConfig{
							URL:    s.Task.URL,
							Method: s.Task.Method,
							Body:   body,
							Header: header,
						},
					},
				})
			}(s)
		}, s)

		if err != nil {
			fmt.Println(time.Now().Format(time.DateTime), "Create task is error:", err)
		} else {
			fmt.Println(time.Now().Format(time.DateTime), "Task", s.Job, "is created")
		}
	}

	// Starts the scheduler blocking
	cr.StartBlocking()
}

func New() BuildIn {
	return &buildIn{}
}
