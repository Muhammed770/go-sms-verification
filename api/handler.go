package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Muhammed770/go-sms-verification/data"
	"github.com/gin-gonic/gin"
	"github.com/muhammed770/go-sms-verification/data"
)

func (app *Config) sendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		appTimeout := time.Duration(5) * time.Second
		_, cancel := context.WithTimeout(context.Background(),appTimeout)
		var payload data.OTPData
		defer cancel()

		app.validateBody(c,&payload)

		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}

		_, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errorJSON(c, err)
			return
		}

		app.writeJSON(c, http.statusAccepted, "OTP sent successfully")
	}
}

func (app *Config) verifySMS() gin.HandlerFunc {}
