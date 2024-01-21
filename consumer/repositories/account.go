package repositories

import (
	"log"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Save(bankAccout BankAccount) error
	Delete(id string) error
	FindAll() ([]BankAccount, error)
	FindById(id string) (BankAccount, error)
}

type BankAccount struct {
	ID            string
	AccountHolder string
	AccountType   int
	Balance       float64
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	err := db.Table("bank_accounts").AutoMigrate(&BankAccount{})
	if err != nil {
		log.Println(err)
	}
	return accountRepository{db: db}
}

func (r accountRepository) Save(bankAccout BankAccount) error {
	return r.db.Save(bankAccout).Error
}

func (r accountRepository) Delete(id string) error {
	log.Println(id)
	return r.db.Delete(BankAccount{ID: id}).Error
}

func (r accountRepository) FindAll() (bankAccout []BankAccount, err error) {
	err = r.db.Find(&bankAccout).Error
	return bankAccout, err
}
func (r accountRepository) FindById(id string) (bankAccout BankAccount, err error) {
	err = r.db.Where("id = ?", id).First(&bankAccout).Error
	return bankAccout, err
}
