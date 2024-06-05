package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	comp "api/computation"
	"api/errors"

	"github.com/gin-gonic/gin"
)

func FizzbuzzHandler() func(c *gin.Context) error {
	return func(c *gin.Context) error {
		value := c.Param("number")
		number, err := strconv.Atoi(value)
		if err != nil {
			return errors.NewInputError(c, fmt.Sprintf("Failed to parse number: %v", err))
		}
		response := comp.CalculateFizzbuzz(number)

		c.JSON(http.StatusOK, response)
		return nil
	}
}
