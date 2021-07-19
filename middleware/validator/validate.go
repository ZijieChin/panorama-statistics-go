package middleware

import (
	"Panorama-Statistics/database"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type StatPost struct {
	User string `validate:"required"`
	Page string `validate:"required"`
}

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "POST" {
			c.Next()
			return
		}
		user := c.PostForm("user")
		if database.GetTimes(user) >= 10 {
			c.AbortWithStatusJSON(400, map[string]string{
				"msg": "Too many duplicate visits, you can still visit the site normally, but no visit count will be recorded.",
			})
			return
		}
		page := c.PostForm("page")
		err := StatPostValid(&StatPost{
			User: user,
			Page: page,
		})
		if err != nil {
			c.AbortWithError(400, err)
		} else {
			c.Next()
		}
	}
}

func StatPostValid(s *StatPost) (err error) {
	validate := validator.New()
	err = validate.Struct(s)
	return
}
