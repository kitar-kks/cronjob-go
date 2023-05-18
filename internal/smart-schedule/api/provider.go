package api

import (
	"xx.com/yyy/smart-schedule/internal/smart-schedule/api/scheduler"
	"xx.com/yyy/smart-schedule/internal/smart-schedule/database"
)

func CreateAPI(dbDriver database.Drivers) API {
	schedulerTask := scheduler.NewTask()
	schedulerRepo := scheduler.NewRepository(dbDriver)
	schedulerUseCase := scheduler.NewUseCase(schedulerRepo, schedulerTask)
	schedulerHandler := scheduler.NewHandler(schedulerUseCase)
	schedulerRouter := scheduler.NewRouter(schedulerHandler)
	apiRouters := NewRouters(schedulerRouter)
	apiAPI := NewAPI(apiRouters)
	return apiAPI
}
