package db

import (
	"fmt"

	"github.com/Bruary/crypto-wallet/models"
)

func AddNewEntryToUsers(user models.User) error {

	query := fmt.Sprintf(`INSERT INTO users VALUES('%s','%s','%s','%s','%s',%d, %v, %v, %v);`,
		user.UUID, user.FirstName, user.LastName, user.FullName, user.Email, user.Age, user.CreatedAt, user.UpdatedAt, user.DeletedAt)

	fmt.Println("query: ", query)

	db := getDB()

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
