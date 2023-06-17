package template

import "fmt"

type SMSOtpAlgo struct {
}

func (s *SMSOtpAlgo) genOTP(n int) string {
	baseOTP := "534546344332"
	otp := baseOTP[:n]
	fmt.Printf("SMS: Generated random otp %s\n", otp)
	return otp
}

// this is default across all varinats of algortihm
func (s *SMSOtpAlgo) saveOTPToCache(otp string) {
	fmt.Println("SMS: Saving otp to cache")
}

func (s *SMSOtpAlgo) prepareMessage(otp string) string {
	msg := fmt.Sprintf("Your OTP for login is %s", otp)
	return msg
}

func (s *SMSOtpAlgo) sendNotification(msg string) {
	fmt.Println("SMS: Sending notification to user -", msg)
}
