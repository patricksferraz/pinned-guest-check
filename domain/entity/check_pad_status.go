package entity

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.TagMap["checkPadStatus"] = govalidator.Validator(func(str string) bool {
		res := str == CHECK_PAD_OPENED.String()
		res = res || str == CHECK_PAD_CANCELED.String()
		res = res || str == CHECK_PAD_AWAITING_PAYMENT.String()
		res = res || str == CHECK_PAD_PAID.String()
		return res
	})
}

type CheckPadStatus int

const (
	CHECK_PAD_OPENED CheckPadStatus = iota + 1
	CHECK_PAD_CANCELED
	CHECK_PAD_AWAITING_PAYMENT
	CHECK_PAD_PAID
)

func (t CheckPadStatus) String() string {
	switch t {
	case CHECK_PAD_OPENED:
		return "OPENED"
	case CHECK_PAD_CANCELED:
		return "CANCELED"
	case CHECK_PAD_AWAITING_PAYMENT:
		return "AWAITING_PAYMENT"
	case CHECK_PAD_PAID:
		return "PAID"
	}
	return ""
}
