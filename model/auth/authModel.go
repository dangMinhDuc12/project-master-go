package authModel

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


type AuthModel struct {
	DB *sql.DB
}

type User struct {
	UserName string
	Email sql.NullString
	Password string
	PhoneNumber string
}

func NewAuthModel(db *sql.DB) *AuthModel {
    return &AuthModel{DB: db}
}


func (am *AuthModel) GetAllUser() ([]User, error) {
	var users []User

	rows, errQueryDb := am.DB.Query("SELECT user_name, email, password, phone_number FROM sso.user LIMIT 30")

	if errQueryDb != nil {
		fmt.Println("errQueryDb", errQueryDb)
		return nil, errQueryDb
	}

	 defer rows.Close()

	for rows.Next() {
		var user User
		errLoopRows := rows.Scan(&user.UserName, &user.Email, &user.Password, &user.PhoneNumber)
		if errLoopRows != nil {
			fmt.Println("errLoopRows", errLoopRows)
			return nil, errLoopRows
		}


		users = append(users, user)
	}

	if errWhenCheckQuery := rows.Err(); errWhenCheckQuery != nil {
		return nil, errWhenCheckQuery
	}

	return users, nil
}


func (am *AuthModel) GetUserByPhoneNumber(phoneNumber string) (User, error) {
	var user User

	errQuery := am.DB.QueryRow("SELECT user_name, email, password, phone_number FROM sso.user where phone_number = $1", phoneNumber).Scan(&user.UserName, &user.Email, &user.Password, &user.PhoneNumber)

	if errQuery != nil {
		return User{}, errQuery
	}

	return user, nil

}



// MarshalJSON customizes the JSON marshaling for the User struct
// func (u *User) MarshalJSON() ([]byte, error) {
//     // Create a map to hold the JSON representation of the User struct
//     userJSON := make(map[string]interface{})

//     // Add Username fields
//     userJSON["username"] = u.UserName

//     // Check if Email value is valid (not NULL)
//     if u.Email.Valid {
//         // Add Email field as a string
//         userJSON["email"] = u.Email.String
//     } else {
//         // Add Email field as an empty string
//         userJSON["email"] = nil
//     }

//     // Marshal the map to JSON
//     return json.Marshal(userJSON)
// }