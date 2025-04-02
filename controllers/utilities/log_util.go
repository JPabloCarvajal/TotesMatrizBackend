package utilities

import (
	"errors"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type LogUtil struct {
	LogService *services.UserLogService
}

func NewLogUtil(logService *services.UserLogService) *LogUtil {
	return &LogUtil{LogService: logService}
}

func (l *LogUtil) RegisterLog(c *gin.Context, logMessage string) error {
	userEmail := c.GetHeader("Username")
	if userEmail == "" {
		return errors.New("missing Username header")
	}

	_, err := l.LogService.CreateUserLog(userEmail, logMessage)
	if err != nil {
		return err
	}

	return nil
}
