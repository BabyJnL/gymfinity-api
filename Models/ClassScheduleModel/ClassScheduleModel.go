package ClassScheduleModel

import (
	"fmt"
	"database/sql"

	DB "gymfinity-backend-api/Connection"
	"gymfinity-backend-api/Entities"
)

func isExists(scheduleId *int) bool {
	userExists := false;
	DB.Connection.QueryRow("SELECT EXISTS (SELECT * FROM class_schedules WHERE schedule_id = ?)", scheduleId).Scan(&userExists);

	return userExists;
}

func GetAll() ([]Entities.ClassSchedule, error) {
	rows, err := DB.Connection.Query(`SELECT cs.schedule_id, c.name, c.description, c.quota, CONCAT(i.firstname, ' ', i.lastname) AS instructor_name, cs.date, cs.start_time, cs.end_time 
	FROM class_schedules cs
	JOIN classes c ON cs.class_id = c.class_id
	JOIN users i ON i.user_id = cs.instructor_id`);
	if err != nil {
		return nil, err;
	}
	defer rows.Close();

	var classSchedules []Entities.ClassSchedule;
	for rows.Next() {
		var classSchedule Entities.ClassSchedule;
		fmt.Println(classSchedule);
		if err := rows.Scan(&classSchedule.ScheduleID, &classSchedule.ClassName, &classSchedule.ClassDescription, &classSchedule.ClassQuota, &classSchedule.InstructorName, &classSchedule.Date, &classSchedule.StartTime, &classSchedule.EndTime); err != nil {
			return nil, err;
		}
		classSchedules = append(classSchedules, classSchedule);
	}

	if len(classSchedules) == 0 {
		return nil, sql.ErrNoRows;
	}
	
	if err := rows.Err(); err != nil {
		return nil, err;
	}

	return classSchedules, nil;
}

func GetById(scheduleId *int) (*Entities.ClassSchedule, error) {
	row := DB.Connection.QueryRow(`SELECT cs.schedule_id, c.name, c.description, c.quota, CONCAT(i.firstname, ' ', i.lastname) AS instructor_name, cs.date, cs.start_time, cs.end_time 
	FROM class_schedules cs
	JOIN classes c ON cs.class_id = c.class_id
	JOIN users i ON i.user_id = cs.instructor_id
	WHERE schedule_id = ?`, scheduleId);

	var classSchedule Entities.ClassSchedule;
	if err := row.Scan(&classSchedule.ScheduleID, &classSchedule.ClassName, &classSchedule.ClassDescription, &classSchedule.ClassQuota, &classSchedule.InstructorName, &classSchedule.Date, &classSchedule.StartTime, &classSchedule.EndTime); err != nil {
		return nil, err;
	}

	return &classSchedule, nil;
}

func Create(classScheduleData *Entities.CreateClassSchedule) error {
	query := "INSERT INTO class_schedules (class_id, instructor_id, date, start_time, end_time) VALUES (?, ?, ?, ?, ?)";

	_, err := DB.Connection.Exec(query, classScheduleData.ClassID, classScheduleData.InstructorID, classScheduleData.Date, classScheduleData.StartTime, classScheduleData.EndTime);
	if err != nil {
		return err;
	}

	return nil;
}

func Update(scheduleId *int, updatedData *Entities.CreateClassSchedule) error {
	classScheduleExists := isExists(scheduleId);
	if !classScheduleExists {
		return sql.ErrNoRows;
	}

	query := "UPDATE class_schedules SET class_id = ?, instructor_id = ?, date = ?, start_time = ?, end_time = ?";

	_, err := DB.Connection.Exec(query, updatedData.ClassID, updatedData.InstructorID, updatedData.Date, updatedData.StartTime, updatedData.EndTime);
	if err != nil {
		return err;
	}

	return nil
}

func Delete(scheduleId *int) error {
	classScheduleExists := isExists(scheduleId);
	if !classScheduleExists {
		return sql.ErrNoRows;
	}

	query := "DELETE FROM class_schedules WHERE schedule_id = ?";

	_, err := DB.Connection.Exec(query, scheduleId);
	if err != nil {
		return err;
	}

	return nil;
}