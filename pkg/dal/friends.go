package dal

import (
	"database/sql"
	"fmt"
	"highload-architect/pkg/config"
	"highload-architect/pkg/structs"
)

func getFriendsByParams(id uint16, joinParam string, joinBy string) ([]structs.User, error) {
	var users []structs.User

	db, dbErr := sql.Open("mysql", config.GetDbConfig())
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

func deleteFriendsById(friendId string, side string) error {
	db, dbErr := sql.Open("mysql", config.GetDbConfig())
	if dbErr != nil {
		return dbErr
	}

	result, err := db.Query(
		fmt.Sprintf(
			"DELETE FROM `friends` WHERE `%s` = %s",
			side,
			friendId,
		),
	)
	if err != nil {
		return err
	}

	result.Close()
	db.Close()
	return nil
}

func GetFriends(id uint16) ([]structs.User, error) {
	var users []structs.User

	db, dbErr := sql.Open("mysql", config.GetDbConfig())
	if dbErr != nil {
		return users, dbErr
	}

	resultUser_1, resultErrUser_1 := getFriendsByParams(id, "target", "initiator")
	if resultErrUser_1 != nil {
		return users, resultErrUser_1
	}

	resultUser_2, resultErrUser_2 := getFriendsByParams(id, "initiator", "target")
	if resultErrUser_2 != nil {
		return users, resultErrUser_2
	}

	users = append(users, resultUser_1...)
	users = append(users, resultUser_2...)

	db.Close()
	return users, nil
}

func CreateFriends(id uint16, friendId string) error {

	db, dbErr := sql.Open("mysql", config.GetDbConfig())
	if dbErr != nil {
		return dbErr
	}

	_, err := db.Query(
		fmt.Sprintf(
			"INSERT INTO friends (initiator, target) VALUES (%d, %s)",
			id,
			friendId,
		),
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteFriend(selfId uint16, friendId string) error {
	db, dbErr := sql.Open("mysql", config.GetDbConfig())
	if dbErr != nil {
		return dbErr
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err_1 := tx.Exec(
		fmt.Sprintf(
			"DELETE FROM `friends` WHERE `initiator` = %d AND `target` = %s",
			selfId,
			friendId,
		),
	)
	if err_1 != nil {
		tx.Rollback()
		return err_1
	}

	_, err_2 := tx.Exec(
		fmt.Sprintf(
			"DELETE FROM `friends` WHERE `initiator` = %s AND `target` = %d",
			friendId,
			selfId,
		),
	)
	if err_2 != nil {
		tx.Rollback()
		return err_2
	}

	db.Close()
	err = tx.Commit()
	return err
}
