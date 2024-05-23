package UserModel

import (
	"database/sql"
	DB "gymfinity-backend-api/Connection"
	"gymfinity-backend-api/Entities"
)

func GetAll(role *string) ([]Entities.User, error) {
	var rows *sql.Rows;
	var err error;

	if *role == "" {
		rows, err = DB.Connection.Query("SELECT * FROM users");
	} else {
		rows, err = DB.Connection.Query("SELECT * FROM users WHERE role = ?", role);
	}

	if err != nil {
		return nil, err;
	}

	defer rows.Close();

	users := []Entities.User{};

	for rows.Next() {
		var user Entities.User;
		if err := rows.Scan(&user.UserID, &user.Firstname, &user.Lastname, &user.Gender, &user.Address, &user.PhoneNumber, &user.Email, &user.JoinDate, &user.Status, &user.ValidUntil, &user.Role, &user.PhotoPath); err != nil {
			return nil, err;
		}
		users = append(users, user);
	}

	if len(users) == 0 {
		return nil, sql.ErrNoRows;
	}

	if err = rows.Err(); err != nil {
        return nil, err;
    }

	return users, nil;
}

func GetById(userId *int) (*Entities.User, error) {
	row := DB.Connection.QueryRow("SELECT * FROM users WHERE user_id = ?", userId);

	var user Entities.User
	err := row.Scan(&user.UserID, &user.Firstname, &user.Lastname, &user.Gender, &user.Address, &user.PhoneNumber, &user.Email, &user.JoinDate, &user.Status, &user.ValidUntil, &user.Role, &user.PhotoPath);
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err;
		}

		return nil, err
	}

	return &user, nil;
}