package FacilityStatusModel

import (
	"database/sql"

	DB "gymfinity-backend-api/Connection"
	"gymfinity-backend-api/Entities"
)

func GetAll() ([]Entities.FacilityStatus, error) {
	rows, err :=  DB.Connection.Query("SELECT * FROM facility_statuses");
	if err != nil {
		return nil, err;
	}
	defer rows.Close();

	var facilityStatuses []Entities.FacilityStatus
	for rows.Next() {
		var facilityStatus Entities.FacilityStatus;
		if err := rows.Scan(&facilityStatus.StatusID, &facilityStatus.Name); err != nil {
			return nil, err;
		}
		facilityStatuses = append(facilityStatuses, facilityStatus);
	}

	if len(facilityStatuses) == 0 {
		return nil, sql.ErrNoRows;
	}
	
	if err := rows.Err(); err != nil {
		return nil, err;
	}

	return facilityStatuses, nil;
}