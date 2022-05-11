package dal

import (
	"database/sql"
	"fmt"
	"highload-architect/pkg/constants"
	"highload-architect/pkg/structs"
)

func getFriendsByParams(id uint16, joinParam string, joinBy string) ([]structs.User, error) {
	var users []structs.User

	db, dbErr := sql.Open("mysql", constants.DBConfig)
	if dbErr != nil {
		return users, dbErr
	}

	result, err := db.Query(
		fmt.Sprintf(
			"SELECT `id`, `first_name`, `last_name` FROM friends f LEFT JOIN users u ON u.id=%s WHERE f.%s='%d'",
			joinParam,
			joinBy,
			id,
		),
	)
	if err != nil {
		return users, err
	}

	for result.Next() {
		var user structs.User

		scanErr := result.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
		)

		if scanErr != nil {
			return users, scanErr
		}

		users = append(users, user)
	}

	result.Close()
	db.Close()
	return users, nil
}

func GetFriends(id uint16) ([]structs.User, error) {
	var users []structs.User

	db, dbErr := sql.Open("mysql", constants.DBConfig)
	if dbErr != nil {
		return users, dbErr
	}

	resultUser_1, resultErrUser_1 := getFriendsByParams(id, "user_id_2", "user_id_1")
	if resultErrUser_1 != nil {
		return users, resultErrUser_1
	}

	resultUser_2, resultErrUser_2 := getFriendsByParams(id, "user_id_1", "user_id_2")
	if resultErrUser_2 != nil {
		return users, resultErrUser_2
	}

	users = append(users, resultUser_1...)
	users = append(users, resultUser_2...)

	db.Close()
	return users, nil
}
