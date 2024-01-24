package models

import (
	"github.com/labstack/echo"
)

type QueryFizzBuzz struct {
	Int1  int    `query:"int1"`
	Int2  int    `query:"int2"`
	Limit int    `query:"limit"`
	Str1  string `query:"str1"`
	Str2  string `query:"str2"`
	Count int    `query:"_"`
}

type PayloadFizzBuzz struct {
	Data    []string `json:"data,omitempty"`
	Success bool     `json:"success"`
	Message string   `json:"message"`
}

type PayloadStatistics struct {
	Data    QueryFizzBuzz `json:"data,omitempty"`
	Success bool          `json:"success"`
	Message string        `json:"message"`
}

func ResponseFizzBuzz(c echo.Context, status int, data []string, success bool, s string) error {
	response := PayloadFizzBuzz{
		Data:    data,
		Success: success,
		Message: s,
	}
	return c.JSON(status, response)
}

func ResponseStatistics(c echo.Context, status int, data QueryFizzBuzz, success bool, s string) error {
	response := PayloadStatistics{
		Data:    data,
		Success: success,
		Message: s,
	}
	return c.JSON(status, response)
}