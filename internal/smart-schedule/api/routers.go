package api

import (
	"github.com/gofiber/fiber/v2"
	"xx.com/yyy/smart-schedule/internal/smart-schedule/api/scheduler"
)

type Routers interface {
	Initials(app *fiber.App)
}

type routers struct {
	SchedulerRoute scheduler.Router
}

func (r *routers) Initials(app *fiber.App) {
	r.SchedulerRoute.Initial(app)
}

func NewRouters(schedulerRoute scheduler.Router) Routers {
	return &routers{
		SchedulerRoute: schedulerRoute,
	}
}
