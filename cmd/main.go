package main

import (
	"JampiCrm/internal/database"

	"JampiCrm/internal/delivery/logger"
	"JampiCrm/internal/delivery/rest"

	"JampiCrm/internal/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	logger.Init()
	e := echo.New()
	rest.LoadMiddlewares(e)

	db := database.GetDb()

	container := usecase.NewContainer(db)
	h := rest.NewHandler(container.AuthUsecase, container.Middleware, db)

	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start((":4000")))

}
