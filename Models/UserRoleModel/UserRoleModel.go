package UserRoleModel 

import (
	"database/sql"
	DB "gymfinity-backend-api/Connection"
	"gymfinity-backend-api/Entities"
)

func GetAll() ([]Entities.UserRole, error) {
	rows, err := DB.Connection.Query("SELECT * FROM user_roles");
	if err != nil {
		return nil, err;
	}
	defer rows.Close();

	var userRoles []Entities.UserRole;
	for rows.Next() {
		var userRole Entities.UserRole;
		if err := rows.Scan(&userRole.RoleID, &userRole.Name); err != nil {
			return nil, err;
		}
		userRoles = append(userRoles, userRole);
	}

	if len(userRoles) == 0 {
		return nil, sql.ErrNoRows;
	}

	if err = rows.Err(); err != nil {
        return nil, err;
    }

	return userRoles, nil;
}