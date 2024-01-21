package events

import "reflect"

var (
	Topics = []string{
		reflect.TypeOf(OpenAccountEvent{}).Name(),
		reflect.TypeOf(DepositFundEvent{}).Name(),
		reflect.TypeOf(WithDrawFundEvent{}).Name(),
		reflect.TypeOf(CloseAccountEvent{}).Name(),
	}
)

type Event interface{}

type OpenAccountEvent struct {
	ID             string
	AccountHolder  string
	AccountType    int
	OpeningBalance float64
}

type DepositFundEvent struct {
	ID     string
	Amount float64
}

type WithDrawFundEvent struct {
	ID     string
	Amount float64
}

type CloseAccountEvent struct {
	ID string
}
