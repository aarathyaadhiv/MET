package helper

import (
	"fmt"

	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient

func TwillioSetup(accountsId, authToken string) {
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountsId,
		Password: authToken,
	})
}

func SendOtp(phNo, serviceId string) (string, error) {
	params := &verify.CreateVerificationParams{}
	params.SetTo(phNo)
	params.SetChannel("sms")

	res, err := client.VerifyV2.CreateVerification(serviceId, params)
	if err != nil {
		return "", err
	}
	return *res.Sid, nil
}

func ValidateOtp(otp models.OtpVerify, serviceId string) error {
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo("+91" + otp.PhNo)
	params.SetCode(otp.Code)

	resp, err := client.VerifyV2.CreateVerificationCheck(serviceId, params)
	if err != nil {
		return err
	} else if *resp.Status == "approved" {
		return nil
	}

	return fmt.Errorf("cannot authenticate the otp because %s", *resp.Status)
}
