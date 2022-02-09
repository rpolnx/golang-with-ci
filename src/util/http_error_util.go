package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpolnx.com.br/golang-with-ci/src/model/dto"
	"time"
)

const defaultErrorCode = http.StatusInternalServerError

func HandleUnexpectedError(c *gin.Context, err error) {
	dtoError := WrapHttpError(defaultErrorCode, err.Error(), c.FullPath())

	Logger.Errorf("Got error %s at path %s", dtoError.Message, dtoError.Path)

	c.JSON(defaultErrorCode, dtoError)
}

func WrapHttpError(httpCode int, message string, path string) *dto.ErrorDTO {
	return &dto.ErrorDTO{
		Timestamp: time.Now(),
		Status:    httpCode,
		Error:     http.StatusText(httpCode),
		Message:   message,
		Path:      path,
	}
}
