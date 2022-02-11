package entity

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.TagMap["chackPadItemStatus"] = govalidator.Validator(func(str string) bool {
		res := str == CHECK_PAD_ITEM_PENDING.String()
		res = res || str == CHECK_PAD_ITEM_CANCELED.String()
		res = res || str == CHECK_PAD_ITEM_PREPARING.String()
		res = res || str == CHECK_PAD_ITEM_ON_THE_WAY.String()
		res = res || str == CHECK_PAD_ITEM_DELIVERED.String()
		return res
	})
}

type CheckPadItemStatus int

const (
	CHECK_PAD_ITEM_PENDING CheckPadItemStatus = iota + 1
	CHECK_PAD_ITEM_CANCELED
	CHECK_PAD_ITEM_PREPARING
	CHECK_PAD_ITEM_ON_THE_WAY
	CHECK_PAD_ITEM_DELIVERED
)

func (t CheckPadItemStatus) String() string {
	switch t {
	case CHECK_PAD_ITEM_PENDING:
		return "PENDING"
	case CHECK_PAD_ITEM_CANCELED:
		return "CANCELED"
	case CHECK_PAD_ITEM_PREPARING:
		return "PREPARING"
	case CHECK_PAD_ITEM_ON_THE_WAY:
		return "ON_THE_WAY"
	case CHECK_PAD_ITEM_DELIVERED:
		return "DELIVERED"
	}
	return ""
}
