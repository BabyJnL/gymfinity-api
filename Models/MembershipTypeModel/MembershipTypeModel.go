package MembershipTypeModel

import (
	"database/sql"
	DB "gymfinity-backend-api/Connection"
	"gymfinity-backend-api/Entities"
)

func GetAll() ([]Entities.Membership, error) {
	rows, err := DB.Connection.Query("SELECT * FROM membership_types");
	if err != nil {
		return nil, err;
	}
	defer rows.Close();

	var membership_types []Entities.Membership;
	for rows.Next() {
		var membership_type Entities.Membership;
		if err := rows.Scan(&membership_type.MembershipID, &membership_type.Name); err != nil {
			return nil, err;
		}
		membership_types = append(membership_types, membership_type);
	}

	if len(membership_types) == 0 {
		return nil, sql.ErrNoRows;
	}

	if err = rows.Err(); err != nil {
        return nil, err;
    }

	return membership_types, nil;
}