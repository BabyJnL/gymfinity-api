package ReservationModel

import (
	// "database/sql"
	"database/sql"
	"fmt"
	DB "gymfinity-backend-api/Connection"
	"gymfinity-backend-api/Entities"
)

func isExists(reservationId *int) bool {
	reservationExists := false;
	DB.Connection.QueryRow("SELECT EXISTS (SELECT * FROM reservations WHERE reservation_id = ?)", reservationId).Scan(&reservationExists);

	return reservationExists;
}

func GetAll() ([]Entities.Reservation, error) {
	rows, err := DB.Connection.Query(`SELECT r.reservation_id, CONCAT(m.firstname, ' ', m.lastname) AS member_name, c.name AS class_name, rs.status_name, r.date 
	FROM reservations r 
	JOIN users m ON r.member_id = m.user_id
	JOIN classes c ON r.class_id = c.class_id
	JOIN reservation_statuses rs ON r.status_id = rs.status_id`);
	if err != nil {
		return nil, err;
	}
	defer rows.Close();

	var reservations []Entities.Reservation;
	for rows.Next() {
		var reservation Entities.Reservation;
		if err := rows.Scan(&reservation.ReservationID, &reservation.MemberName, &reservation.ClassName, &reservation.Status, &reservation.Date); err != nil {
			return nil, err;
		}
		fmt.Println(reservation);
		reservations = append(reservations, reservation);
	}

	return reservations, nil;
}

func GetById(reservationId *int) (*Entities.Reservation, error) {
	row := DB.Connection.QueryRow("SELECT r.reservation_id, r.member_id, r.class_id, rs.status_name, r.date FROM reservations r JOIN reservation_statuses rs ON r.status_id = rs.status_id WHERE r.reservation_id = ?", reservationId);

	var reservation Entities.Reservation;
	if err := row.Scan(&reservation.ReservationID, &reservation.MemberName, &reservation.ClassName, &reservation.Status, &reservation.Date); err != nil {
		return nil, err;
	}

	return &reservation, nil;
}

func Create(reservationData *Entities.CreateReservation) error {
	query := "INSERT INTO reservations (member_id, class_id, status_id, date) VALUE (?, ?, ?, ?)";

	_, err := DB.Connection.Exec(query, reservationData.MemberId, reservationData.ClassId, reservationData.StatusId, reservationData.Date);
	if err != nil {
		return err;
	}

	return nil;
}

func Update(reservationId *int, updatedStatus *Entities.UpdateReservationStatus) error {
	reservationExists := isExists(reservationId);
	if !reservationExists {
		return sql.ErrNoRows;
	}

	query := "UPDATE reservations SET status_id = ? WHERE reservation_id = ?";

	_, err := DB.Connection.Exec(query, updatedStatus.StatusId, reservationId);
	if err != nil {
		return err;
	}

	return nil;
}