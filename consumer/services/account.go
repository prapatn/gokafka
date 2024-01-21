package services

import (
	"consumer/repositories"
	"encoding/json"
	"events"
	"log"
	"reflect"
)

type EventHandler interface {
	Handle(topic string, eventBytes []byte)
}

type accountEventHandler struct {
	accountRepo repositories.AccountRepository
}

func NewAccountEventHandler(accountRepo repositories.AccountRepository) EventHandler {
	return accountEventHandler{accountRepo: accountRepo}
}

func (h accountEventHandler) Handle(topic string, eventBytes []byte) {
	log.Println(topic)
	switch topic {
	case reflect.TypeOf(events.OpenAccountEvent{}).Name():
		event := events.OpenAccountEvent{}
		err := json.Unmarshal(eventBytes, &event)
		if err != nil {
			log.Println(err)
			return
		}

		bankAccount := repositories.BankAccount{
			ID:            event.ID,
			AccountHolder: event.AccountHolder,
			AccountType:   event.AccountType,
			Balance:       event.OpeningBalance,
		}

		err = h.accountRepo.Save(bankAccount)
		if err != nil {
			log.Println(err)
			return
		}
	case reflect.TypeOf(events.DepositFundEvent{}).Name():
		event := events.DepositFundEvent{}
		err := json.Unmarshal(eventBytes, &event)
		if err != nil {
			log.Println(err)
			return
		}

		bankAccount, err := h.accountRepo.FindById(event.ID)
		if err != nil {
			log.Println(err)
			return
		}

		bankAccount.Balance += event.Amount
		err = h.accountRepo.Save(bankAccount)
		if err != nil {
			log.Println(err)
			return
		}
	case reflect.TypeOf(events.WithDrawFundEvent{}).Name():
		event := events.WithDrawFundEvent{}
		err := json.Unmarshal(eventBytes, &event)
		if err != nil {
			log.Println(err)
			return
		}

		bankAccount, err := h.accountRepo.FindById(event.ID)
		if err != nil {
			log.Println(err)
			return
		}

		bankAccount.Balance -= event.Amount
		err = h.accountRepo.Save(bankAccount)
		if err != nil {
			log.Println(err)
			return
		}
	case reflect.TypeOf(events.CloseAccountEvent{}).Name():
		event := events.CloseAccountEvent{}
		err := json.Unmarshal(eventBytes, &event)
		if err != nil {
			log.Println(err)
			return
		}
		err = h.accountRepo.Delete(event.ID)
		if err != nil {
			log.Println(err)
			return
		}
	default:
		log.Println("no event handler")
	}
}
