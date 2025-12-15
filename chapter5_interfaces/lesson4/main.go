package main

func getExpenseReport(e expense) (string, float64) {
	email, isEmail := e.(email)
	emailCost := email.cost()
	sms, isSms := e.(sms)
	smsCost := sms.cost()
	if isEmail {
		return email.toAddress, emailCost
	}
	if isSms {
		return sms.toPhoneNumber, smsCost
	}
	return "", 0.0
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}
