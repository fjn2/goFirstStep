package logic

import (
	"container/list"
	"test/dao"
	"test/dto"
)

const bufferSize = 1

func GetAll(customerDTO dto.Customer) []dto.Customer {
	var customerDAO = dao.Customer{}

	// Prepare the buffer to pass data from the DB
	channel := make(chan dao.Customer, bufferSize)

	go customerDAO.GetAll(channel)

	var customerList = list.New()

	// Read until the channel is closed
	for {
		if customer, ok := <-channel; ok {
			customerList.PushBack(customer)
		} else {
			break
		}
	}

	// Prepare output
	response := make([]dto.Customer, customerList.Len())
	var index int

	// Convert elements from DAO to DTO
	for element := customerList.Front(); element != nil; element = element.Next() {
		cDAO := element.Value.(dao.Customer)
		response[index].Id = cDAO.Id
		response[index].Name = cDAO.Name
		index = index + 1
	}

	return response
}
