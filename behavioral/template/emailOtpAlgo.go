package template

import "fmt"

type EmailOtpAlgo struct {
}

func (s *EmailOtpAlgo) genOTP(n int) string {
	baseOTP := "72958278492"
	otp := baseOTP[:n]
	fmt.Printf("Email: Generated random otp %s\n", otp)
	return otp
}

// this is default across all varinats of algortihm
func (s *EmailOtpAlgo) saveOTPToCache(otp string) {
	fmt.Println("Email: Saving otp to cache")
}

func (s *EmailOtpAlgo) prepareMessage(otp string) string {
	msg := fmt.Sprintf("Your OTP code is %s", otp)
	return msg
}

func (s *EmailOtpAlgo) sendNotification(msg string) {
	fmt.Println("Email: Sending notification to user -", msg)
}
