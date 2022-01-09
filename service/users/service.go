package users

import (
	"github.com/Bruary/crypto-wallet/service"
)

type Users struct {
}

func NewUsers() service.Service {
	return &Users{}
}
