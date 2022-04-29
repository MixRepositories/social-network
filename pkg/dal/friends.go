package dal

import (
	"database/sql"
	"fmt"
	"highload-architect/pkg/constants"
	"highload-architect/pkg/structs"
)

func GetFriends(id uint16) ([]structs.User, error) {
	var users []structs.User

	db, dbErr := sql.Open("mysql", constants.DBConfig)
	if dbErr != nil {
		return users, dbErr
	}

	result, resultErr := db.Query(
		fmt.Sprintf(
			"SELECT `id`, `first_name`, `last_name` FROM friends f LEFT JOIN users u ON u.id = user_id_2 WHERE f.user_id_1 = %d",
			id,
		),
	)

	if resultErr != nil {
		return users, resultErr
	}

	for result.Next() {
		var user structs.User

		scanErr := result.Scan(
			&user.Id,
			&user.Email,
			&user.FirstName,
			&user.LastName,
		)

		if scanErr != nil {
			return users, resultErr
		}

		users = append(users, user)
	}
	// INSERT INTO `friends` (`id`, `user_id_1`, `user_id_2`) VALUES (NULL, '11', '13')
	// SELECT * FROM friends f LEFT JOIN users u ON u.id = user_id_1 WHERE f.user_id_1 = 11;
	// SELECT * FROM friends f LEFT JOIN users u ON u.id = user_id_1 WHERE f.user_id_1 = 11 OR f.user_id_2 = 11;
	result.Close()
	db.Close()
	return users, nil
}
