package template

// Skeleton of OTP Algorithm
type OTPAlgo interface {
	genOTP(int) string
	saveOTPToCache(string)
	prepareMessage(string) string
	sendNotification(string)
}
