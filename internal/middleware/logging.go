package middleware

import (
	"log"
	"time"

	"github.com/labstack/echo/v5"
)

func Logging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		start := time.Now()
		log.Print(start)
		
		err := next(c)
		if err != nil {
			log.Print(err)
		}
		
		end := time.Now()
		log.Print(end)

		return err
	}
}
