package template

import "fmt"

/*
Template Method is a behavioral design pattern that allows you to defines a skeleton of an algorithm in a base class
and let subclasses override the steps without changing the overall algorithm’s structure.
*/

// Lets take the algo for generating and sending OTPs.

func RunDemo() {
	// Send and generate OTP to a customer
	smsOTP := &SMSOtpAlgo{}
	emailOtp := &EmailOtpAlgo{}

	otp := &OtpTemplate{algo: smsOTP}
	otp.GenAndSendOTP(6)
	fmt.Println()
	otp.setAlgo(emailOtp)
	otp.GenAndSendOTP(5)
}
