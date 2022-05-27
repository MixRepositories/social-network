package dal

import (
	"database/sql"
	"fmt"
	"highload-architect/pkg/config"
	"highload-architect/pkg/structs"
)

func GetAuthUser(email string, hashPassword string) (structs.User, error) {
	var user structs.User

	db, dbErr := sql.Open("mysql", config.GetDbConfig())
	if dbErr != nil {
		return user, dbErr
	}

	result, resultErr := db.Query(
		fmt.Sprintf(
			"SELECT `id`, `email`, `first_name`, `last_name` FROM `users` WHERE email='%s' AND password='%s'",
			email,
			hashPassword,
		),
	)

	if resultErr != nil {
		return user, resultErr
	}

	for result.Next() {
		scanErr := result.Scan(
			&user.Id,
			&user.Email,
			&user.FirstName,
			&user.LastName,
		)

		if scanErr != nil {
			panic(scanErr)
		}
	}

	result.Close()
	db.Close()
	return user, nil
}

func GetUserById(id uint16) (structs.User, error) {
	var user structs.User

	db, dbErr := sql.Open("mysql", config.GetDbConfig())
	if dbErr != nil {
		return user, dbErr
	}

	result, resultErr := db.Query(
		fmt.Sprintf("SELECT * FROM `users` WHERE id=%d", id),
	)
	if resultErr != nil {
		return user, resultErr
	}

	for result.Next() {
		scanErr := result.Scan(
			&user.Id,
			&user.Email,
			&user.Password,
			&user.FirstName,
			&user.LastName,
			&user.City,
			&user.Gender,
			&user.Birthday,
		)

		if scanErr != nil {
			panic(scanErr)
		}
	}

	result.Close()
	db.Close()
	return user, nil
}

func CreateUser(
	email string,
	hashPassword string,
	firstName string,
	lastName string,
	birthday string,
	gender string,
	age string,
	city string,
) error {
	db, err := sql.Open("mysql", config.GetDbConfig())
	if err != nil {
		return err
	}

	insert, insertErr := db.Query(
		fmt.Sprintf(
			"INSERT INTO `users` (`email`, `password`, `first_name`, `last_name`, `birthday`, `gender`, `city`) value ('%s', '%s', '%s', '%s', '%s', '%s', '%s')",
			email,
			hashPassword,
			firstName,
			lastName,
			birthday,
			gender,
			city,
		),
	)

	if insertErr != nil {
		return insertErr
	}

	insert.Close()
	db.Close()
	return nil
}

func GetUsers(exception []uint16) ([]structs.User, error) {
	var users []structs.User

	var exceptionStr string

	for i := 0; i < len(exception); i++ {
		if i == len(exception)-1 {
			exceptionStr = exceptionStr + fmt.Sprintf("%d", exception[i])

		} else {
			exceptionStr = exceptionStr + fmt.Sprintf("%d, ", exception[i])
		}
	}

	db, dbErr := sql.Open("mysql", config.GetDbConfig())
	if dbErr != nil {
		return users, dbErr
	}

	result, resultErr := db.Query(
		fmt.Sprintf("SELECT `id`, `email`, `first_name`, `last_name` FROM `users` WHERE id not in (%s)", exceptionStr),
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

	result.Close()
	db.Close()
	return users, nil
}
