package service

import "github.com/Bruary/crypto-wallet/models"

type Service interface {
	CreateUser(req interface{}) models.BaseResponse
}
