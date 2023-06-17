package template

// structure of the algorithm remains same
type OtpTemplate struct {
	algo OTPAlgo
}

func (o *OtpTemplate) GenAndSendOTP(n int) {
	// series of steps to execute
	otp := o.algo.genOTP(n)
	o.algo.saveOTPToCache(otp)
	msg := o.algo.prepareMessage(otp)
	o.algo.sendNotification(msg)
}

func (o *OtpTemplate) setAlgo(algo OTPAlgo) {
	o.algo = algo
}
