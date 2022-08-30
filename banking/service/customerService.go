package service

import (
	"Desktop/golang/banking/domain"
	"Desktop/golang/banking/dto"
	"Desktop/golang/banking/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
	GetAllByStatus(string) ([]domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, customer := range c {
		response = append(response, customer.ToDto())
	}
	return response, nil
}

func(s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func(s DefaultCustomerService) GetAllByStatus(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.GetByStatus(status)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}