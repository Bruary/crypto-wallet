package users

import (
	"time"

	"github.com/Bruary/crypto-wallet/db"
	"github.com/Bruary/crypto-wallet/models"
	UUID "github.com/google/uuid"
)

func (*Users) CreateUser(req models.CreateUserRequest) models.BaseResponse {

	user := models.User{
		UUID:      UUID.New().String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		FullName:  req.FirstName,
		Email:     req.Email,
		Age:       req.Age,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := db.AddNewEntryToUsers(user)
	if err != nil {
		return models.BaseResponse{
			Success: false,
			Msg:     err.Error(),
			Error: models.Error{
				Code: 0,
				Msg:  "lol it failed.",
			},
		}
	}

	return models.BaseResponse{
		Success: true,
		Msg:     "Creation done",
		Error: models.Error{
			Code: 0,
			Msg:  "No Err.",
		},
	}
}
