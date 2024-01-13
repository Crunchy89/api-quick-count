package main

import (
	"net/http"
	"os"

	"github.com/centrifugal/centrifuge"
	"github.com/crunchy89/api-quick-count/app/socket/handler"
	"github.com/crunchy89/api-quick-count/app/socket/service"
	"github.com/crunchy89/api-quick-count/utils/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initLog() *logrus.Entry {
	if os.Getenv("ENVIRONMENT") != "production" {
		logrus.SetFormatter(&logrus.TextFormatter{})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	return logrus.WithField("service", "api-techcode")
}

func main() {
	logger := initLog()

	// Load .env
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("can't load .env with error %s", err)
	}
	logrus.Infof("this program running in %s", os.Getenv("ENVIRONMENT"))

	dsn := os.Getenv("MYSQL_DB_HOST")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	e := echo.New()
	if os.Getenv("ENVIRONMENT") == "production" {
		e.Debug = false
	} else {
		e.Debug = true
	}

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			logData := logger.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
				"host":   values.Host,
			})
			if values.Status >= 300 {
				logData.Warnf("failed request with error %s", values.Error)
			} else {
				logData.Infof("success request with return %d", values.Status)
			}

			return nil
		},
	}), middleware.CORS())

	e.Validator = validator.InitValidator()
	e.HTTPErrorHandler = validator.InitErrorHandler()

	socket, err := centrifuge.New(centrifuge.Config{})
	if err != nil {
		logger.Fatal("error ", err)
	}

	if err := socket.Run(); err != nil {
		logger.Fatal("error ", err)
	}

	service.Connect(socket, logger)

	var config centrifuge.WebsocketConfig
	config.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	v1route := InitV1Routes(socket, db, logger, e)
	v1route.Routes()

	socketWsHandler := centrifuge.NewWebsocketHandler(socket, config)
	e.Any("/socket", echo.WrapHandler(handler.ApiChannel(socketWsHandler)))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server Health ✅.")
	})

	logger.Fatal(e.Start(os.Getenv("PORT")))
}
