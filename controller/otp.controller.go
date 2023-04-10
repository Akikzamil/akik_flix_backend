package controller

import (
	"akikflix/util"
	"time"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

const appTimeout = time.Minute

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
    Username: util.GetVariable("TWILIO_SID"),
    Password: util.GetVariable("TWILIO_TOKEN"),
})

func twilioSendOTP(phoneNumber string) (string, error) {
    params := &twilioApi.CreateVerificationParams{}
    params.SetTo(phoneNumber)
    params.SetChannel("sms")

    resp, err := client.VerifyV2.CreateVerification(util.GetVariable("TWILIO_SERVICE_SID"), params)
    if err != nil {
        return "", err
    }

    return *resp.Sid, nil
}

func twilioVerifyOTP(phoneNumber string, code string) error {
    params := &twilioApi.CreateVerificationCheckParams{}
    params.SetTo(phoneNumber)
    params.SetCode(code)

    resp, err := client.VerifyV2.CreateVerificationCheck(util.GetVariable("TWILIO_SERVICE_SID"), params)
    if err != nil {
        return err
    } else if *resp.Status == "approved" {
        return nil
    }

    return nil
}