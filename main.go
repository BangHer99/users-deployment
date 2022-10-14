package main

import (
	"be12/deploy/config"
	"be12/deploy/migration"

	"be12/deploy/factory"
	"be12/deploy/utils/database/mysql"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	cfg := config.GetConfig()
	db := mysql.InitDBmySql(cfg)

	migration.InitMigrate(db)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))

}
