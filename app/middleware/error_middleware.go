package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"ozon_test/app/models"
	customErrors "ozon_test/pkg/errors"
)

func CheckError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := errors.Unwrap(c.Errors.Last())
			newErr := new(models.CustomError)
			newErr.CustomErr = err.Error()
			errJson, errMarsh := newErr.MarshalJSON()
			if errMarsh != nil {
				fmt.Println(errMarsh.Error())
				return
			}
			c.Data(customErrors.ConvertErrorToCode(err), "application/json; charset=utf-8", errJson)
			return
		}
	}
}
