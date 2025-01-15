package util

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	GetParamInt   = GetParam[int]
	GetParamInt64 = GetParam[int64]
)

func GetParam[T int | int64](c echo.Context, param string) (T, error) {
	value, err := strconv.ParseInt(c.Param(param), 10, 64)
	if err != nil {
		return 0, err
	}

	return T(value), nil
}
