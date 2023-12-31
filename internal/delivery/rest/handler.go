package rest

import (
	"JampiCrm/internal/usecase/auth"
	"net/http"

	"strconv"

	midUtil "JampiCrm/internal/delivery/auth"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Declare meta and data at the package level
var meta interface{}
var data interface{}

type handler struct {
	authUsecase auth.Usecase
	db          *gorm.DB
	middleware  midUtil.MidlewareInterface
}

type handlerRespont struct {
	Status  int
	Message string
	Meta    interface{} `json:"meta"`
	Data    interface{} `json:"data"`
}

func NewHandler(authUsecase auth.Usecase, middleware midUtil.MidlewareInterface, db *gorm.DB) *handler {

	return &handler{
		authUsecase: authUsecase,
		db:          db,
		middleware:  middleware,
	}
}

func handleError(c echo.Context, statusCode int, err error, meta interface{}, data interface{}) error {
	var response handlerRespont

	if statusCode != http.StatusOK {
		response = handlerRespont{
			Status:  statusCode,
			Message: "internal error: " + err.Error(),
			Meta:    meta,
			Data:    data,
		}
	} else {
		response = handlerRespont{
			Status:  statusCode,
			Message: "PROCESS SUCCESS: " + err.Error(),
			Meta:    meta,
			Data:    data,
		}
	}

	return c.JSON(statusCode, response)
}

func transformIdToInt(c echo.Context) int {
	// get param
	ID := c.Param("id")
	Id, err := strconv.Atoi(ID)
	if err != nil {
		panic(err)
	}
	return Id
}

func HandlingLimitAndOffset(c echo.Context) (int, int) {
	// Get query parameters with default values
	limitStr := c.QueryParam("limit")
	if limitStr == "" {
		limitStr = "15" // Default value
	}
	offsetStr := c.QueryParam("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}

	// Convert to integers with error handling
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		panic("error converting 'limit' to integer")
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		panic("error converting 'offset' to integer")
	}

	// Return the values
	return limit, offset
}

func WriteLogErorr(msg string, err error) {
	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logrus.SetFormatter(formatter)
	logrus.WithFields(logrus.Fields{
		"err": err,
	}).Error(msg, err.Error())
}

func WriteLogInfo(msg string) {
	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logrus.SetFormatter(formatter)
	logrus.Info(msg)
}
