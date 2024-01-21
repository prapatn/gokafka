package services

import (
	"errors"
	"events"
	"log"
	"producer/commands"

	"github.com/google/uuid"
)

type AccountService interface {
	OpneAccount(command commands.OpenAccountCommand) (id string, err error)
	DepositFund(command commands.DepositFundCommand) error
	WithDrawFund(command commands.WithDrawFundCommand) error
	CloseAccount(command commands.CloseAccountCommand) error
}

type accountService struct {
	eventProducer EventProducer
}

func NewAccountService(eventProducer EventProducer) AccountService {
	return accountService{eventProducer: eventProducer}
}

func (s accountService) OpneAccount(command commands.OpenAccountCommand) (id string, err error) {
	if command.AccountHolder == "" || command.AccountType == 0 || command.OpeningBalance == 0 {
		return "", errors.New("Bad Request")
	}
	event := events.OpenAccountEvent{
		ID:             uuid.NewString(),
		AccountHolder:  command.AccountHolder,
		AccountType:    command.AccountType,
		OpeningBalance: command.OpeningBalance,
	}
	log.Printf("%#v", event)

	err = s.eventProducer.Produe(event)
	return event.ID, err
}

func (s accountService) DepositFund(command commands.DepositFundCommand) error {
	if command.ID == "" || command.Amount == 0 {
		return errors.New("Bad Request")
	}
	event := events.DepositFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}
	log.Printf("%#v", event)

	return s.eventProducer.Produe(event)
}

func (s accountService) WithDrawFund(command commands.WithDrawFundCommand) error {
	if command.ID == "" || command.Amount == 0 {
		return errors.New("Bad Request")
	}
	event := events.WithDrawFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}
	log.Printf("%#v", event)

	return s.eventProducer.Produe(event)
}

func (s accountService) CloseAccount(command commands.CloseAccountCommand) error {
	if command.ID == "" {
		return errors.New("Bad Request")
	}
	event := events.CloseAccountEvent{
		ID: command.ID,
	}
	log.Printf("%#v", event)

	return s.eventProducer.Produe(event)
}
