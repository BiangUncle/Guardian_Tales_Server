package service

import (
	"errors"
	"fmt"
	"sync"
)

type Service struct {
	server_map  map[string]string
	address_map map[string]string
}

var _service *Service
var _serviceOnce sync.Once

func NewService() *Service {
	_serviceOnce.Do(func() {
		_service = &Service{
			server_map:  make(map[string]string),
			address_map: make(map[string]string),
		}
	})
	return _service
}

func (s *Service) Register(address, reg_ser string) error {

	if !s.check(address, reg_ser) {
		return errors.New("This service / address has registered! Plz use other address.")
	}

	fmt.Println("Register successfully")
	s.server_map[reg_ser] = address
	s.address_map[address] = reg_ser
	return nil
}

func (s *Service) FindServer(reg_ser string) string {
	if ser, ok := s.server_map[reg_ser]; ok {
		return ser
	}
	return ""
}

func (s *Service) check(address, reg_ser string) bool {
	if _, ok := s.server_map[reg_ser]; ok {
		fmt.Println("Register failed")
		return false
	}
	if _, ok := s.address_map[address]; ok {
		fmt.Println("Register failed")
		return false
	}
	return true
}
