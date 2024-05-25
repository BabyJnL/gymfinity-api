package UserStatusModel

import (
	"database/sql"
	DB "gymfinity-backend-api/Connection"
	"gymfinity-backend-api/Entities"
)

func GetAll() ([]Entities.UserStatus, error) {
	rows, err := DB.Connection.Query("SELECT * FROM user_statuses");
	if err != nil {
		return nil, err;
	}
	defer rows.Close();

	var userStatuses []Entities.UserStatus;
	for rows.Next() {
		var userStatus Entities.UserStatus;
		if err := rows.Scan(&userStatus.StatusID, &userStatus.Name); err != nil {
			return nil, err;
		}
		userStatuses = append(userStatuses, userStatus);
	}

	if len(userStatuses) == 0 {
		return nil, sql.ErrNoRows;
	}

	if err = rows.Err(); err != nil {
        return nil, err;
    }

	return userStatuses, nil;
}