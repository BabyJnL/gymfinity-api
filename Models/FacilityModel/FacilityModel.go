package FacilityModel 

import (
	"database/sql"

	DB "gymfinity-backend-api/Connection"
	"gymfinity-backend-api/Entities"
)

func isExists(facilityId *int) bool {
	userExists := false;
	DB.Connection.QueryRow("SELECT EXISTS (SELECT * FROM facilities WHERE facility_id = ?)", facilityId).Scan(&userExists);

	return userExists;
}

func GetAll() ([]Entities.Facility, error) {
	rows, err := DB.Connection.Query("SELECT * FROM facilities");
	if err != nil {
		return nil, err;
	}
	defer rows.Close();

	var facilities []Entities.Facility;
	for rows.Next() {
		var facility Entities.Facility
		if err := rows.Scan(&facility.FacilityID, &facility.Name, &facility.Description, &facility.Photo); err != nil {
			return nil, err;
		}
		facilities = append(facilities, facility);
	}

	if len(facilities) == 0 {
		return nil, sql.ErrNoRows;
	}
	
	if err := rows.Err(); err != nil {
		return nil, err;
	}

	return facilities, nil;
}

func GetById(facilityId *int) (*Entities.Facility, error) {
	row := DB.Connection.QueryRow("SELECT * FROM facilities WHERE facility_id = ?", facilityId);

	var facility Entities.Facility;
	if err := row.Scan(&facility.FacilityID, &facility.Name, &facility.Description, &facility.Photo); err != nil {
		return nil, err;
	}

	return &facility, nil;
}

func Create(facilityData *Entities.Facility) error {
	query := "INSERT INTO facilities (name, description, photo) VALUES (?, ?, ?)";

	_, err := DB.Connection.Exec(query, facilityData.Name, facilityData.Description, facilityData.Photo);
	if err != nil {
		return err;
	}

	return nil;
}

func Update(facilityId *int, updatedData *Entities.Facility) error {
	facilityExists := isExists(facilityId);
	if !facilityExists {
		return sql.ErrNoRows
	}

	query := "UPDATE facilities SET name = ?, description = ?, photo = ? WHERE facility_id = ?";

	_, err := DB.Connection.Exec(query, updatedData.Name, updatedData.Description, updatedData.Photo, facilityId);
	if err != nil {
		return err;
	}

	return nil;
}

func Delete(facilityId *int) error {
	facilityExists := isExists(facilityId);
	if !facilityExists {
		return sql.ErrNoRows
	}

	query := "DELETE FROM facilities WHERE facility_id = ?";

	_, err := DB.Connection.Exec(query, facilityId);
	if err != nil {
		return err;
	}

	return nil;
}