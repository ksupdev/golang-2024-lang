package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"strings"
	"time"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{accRepo: accRepo}
}

func (service accountService) NewAccount(customerID int, request NewAccountRequest) (*AccountResponse, error) {

	// validate
	if request.Amount < 5000 {
		return nil, errs.NewValidationError("amount at least 5,0000")
	}

	if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
		return nil, errs.NewValidationError("account type should be saving or checking")
	}

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      1,
	}
	newAcc, err := service.accRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	response := AccountResponse{
		AccountID:   newAcc.AccountID,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}

	return &response, nil
}
func (service accountService) GetAccounts(customerID int) ([]AccountResponse, error) {
	accounts, err := service.accRepo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	responses := []AccountResponse{}
	for _, v := range accounts {
		responses = append(responses, AccountResponse{
			AccountID:   v.AccountID,
			OpeningDate: v.OpeningDate,
			AccountType: v.AccountType,
			Amount:      v.Amount,
			Status:      v.Status,
		})
	}
	return responses, nil
}
