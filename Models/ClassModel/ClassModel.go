package ClassModel 

import (
	"database/sql"
	DB "gymfinity-backend-api/Connection"
	"gymfinity-backend-api/Entities"
)

func isExists(classId *int) bool {
	classExists := false;
	DB.Connection.QueryRow("SELECT EXISTS (SELECT * FROM classes WHERE class_id = ?)", classId).Scan(&classExists);

	return classExists;
}

func GetAll() ([]Entities.Class, error) {
	rows, err := DB.Connection.Query("SELECT * FROM classes");
	if err != nil {
		return nil, err
	}

	defer rows.Close();
	
	classes := []Entities.Class{};
	for rows.Next() {
		var class Entities.Class;
		if err := rows.Scan(&class.ClassID, &class.Name, &class.Description, &class.Quota); err != nil {
			return nil, err;
		}

		classes = append(classes, class);
	}

	if len(classes) == 0 {
		return nil, sql.ErrNoRows;
	}

	if err = rows.Err(); err != nil {
        return nil, err;
    }

	return classes, nil;
}

func GetById(classId *int) (*Entities.Class, error) {
	rows := DB.Connection.QueryRow("SELECT * FROM classes WHERE class_id = ?", classId);

	var class Entities.Class;
	if err := rows.Scan(&class.ClassID, &class.Name, &class.Description, &class.Quota); err != nil {
		return nil, err;
	}

	return &class, nil;
}

func Create(classData *Entities.Class) error {
	query := "INSERT INTO classes (name, description, quota) VALUES (?, ?, ?)";

	_, err := DB.Connection.Exec(query, classData.Name, classData.Description, classData.Quota);
	if err != nil {
		return err;
	}

	return nil;
}

func Update(classId *int, updatedData *Entities.Class) error {
	classExists := isExists(classId);
	if !classExists {
		return sql.ErrNoRows;
	}

	query := "UPDATE classes SET name = ?, description = ?, quota = ? WHERE class_id = ?";

	_, err := DB.Connection.Exec(query, updatedData.Name, updatedData.Description, updatedData.Quota, classId);
	if err != nil {
		return err;
	}

	return nil;
}

func Delete(classId *int) error {
	classExists := isExists(classId);
	if !classExists {
		return sql.ErrNoRows;
	}

	query := "DELETE FROM classes WHERE class_id = ?";

	_, err := DB.Connection.Exec(query, classId);
	if err != nil {
		return err;
	}

	return nil;
}