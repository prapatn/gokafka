package controller

import (
	"log"
	"producer/commands"
	"producer/services"

	"github.com/gofiber/fiber/v2"
)

type AccountController interface {
	OpneAccount(c *fiber.Ctx) error
	DepositFund(c *fiber.Ctx) error
	WithDrawFund(c *fiber.Ctx) error
	CloseAccount(c *fiber.Ctx) error
}

type accountController struct {
	accountService services.AccountService
}

func NewAccountController(accountService services.AccountService) AccountController {
	return accountController{accountService: accountService}
}

func (a accountController) OpneAccount(c *fiber.Ctx) error {
	command := commands.OpenAccountCommand{}
	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	id, err := a.accountService.OpneAccount(command)
	if err != nil {
		log.Println(err)
		return err
	}

	c.Status(fiber.StatusCreated)

	return c.JSON(fiber.Map{
		"msg": "open account success",
		"id":  id,
	})
}

func (a accountController) DepositFund(c *fiber.Ctx) error {
	command := commands.DepositFundCommand{}
	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = a.accountService.DepositFund(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"msg": "Deposit fund success",
	})
}

func (a accountController) WithDrawFund(c *fiber.Ctx) error {
	command := commands.WithDrawFundCommand{}
	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = a.accountService.WithDrawFund(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"msg": "WithDraw fund success",
	})
}

func (a accountController) CloseAccount(c *fiber.Ctx) error {
	command := commands.CloseAccountCommand{}
	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = a.accountService.CloseAccount(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"msg": "close account success",
	})
}
