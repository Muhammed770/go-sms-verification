package api
import {
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

}

type jsonResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data any `json:"data"`
}

var validate = validator.New()

func (app *Config) validateBody(c *gin.Context, data any) error {
	if err := c.BindJSON(data); err != nil {
		return err
	}
	if err := validate.Struct(data); err != nil {
		return err
	}
	return nil
}
