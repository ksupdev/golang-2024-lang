package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
)

type customerService struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) customerService {
	return customerService{customerRepo: customerRepo}
}

func (service customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := service.customerRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	customerResponses := []CustomerResponse{}
	for _, customer := range customers {
		customerResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		customerResponses = append(customerResponses, customerResponse)
	}
	return customerResponses, nil
}

func (service customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := service.customerRepo.GetById(id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotfoundError("customer not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	customerResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &customerResponse, nil
}
