package util

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rpolnx.com.br/golang-with-ci/src/model/dto"
	"time"
)

const defaultErrorCode = http.StatusInternalServerError

func HandleUnexpectedError(c *gin.Context, err error) {
	dtoError := WrapHttpError(defaultErrorCode, err.Error(), c.FullPath())

	marshal, _ := json.Marshal(dtoError)
	fmt.Println(string(marshal))
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
