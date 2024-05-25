package PaymentModel

import (
	"database/sql"

	DB "gymfinity-backend-api/Connection"
	"gymfinity-backend-api/Entities"
)

func GetAll() ([]Entities.Payment, error) {
	rows, err := DB.Connection.Query(`SELECT p.payment_id, CONCAT(m.firstname,' ', m.lastname) AS member_name, CONCAT(s.firstname,' ', s.lastname) AS staff_name, mb.type_name, p.amount, p.date, p.photo_path
	FROM payments p
	JOIN users m ON p.member_id = m.user_id
	JOIN users s ON p.staff_id = s.user_id
	JOIN membership_types mb ON p.membership_type_id = mb.type_id`);
	if err != nil {
		return nil, err;
	}
	defer rows.Close();

	var payments []Entities.Payment;
	for rows.Next() {
		var payment Entities.Payment;
		if err := rows.Scan(&payment.PaymentID, &payment.MemberName, &payment.StaffName, &payment.MembershipType, &payment.Amount, &payment.Date, &payment.PhotoPath); err != nil {
			return nil, err;
		}
		payments = append(payments, payment);
	}

	if len(payments) == 0 {
		return nil, sql.ErrNoRows;
	}
	
	if err := rows.Err(); err != nil {
		return nil, err;
	}

	return payments, nil;
}

func GetById(paymentId *int) (*Entities.Payment, error) {
	row := DB.Connection.QueryRow(`SELECT p.payment_id, CONCAT(m.firstname,' ', m.lastname) AS member_name, CONCAT(s.firstname,' ', s.lastname) AS staff_name, mb.type_name, p.amount, p.date, p.photo_path
	FROM payments p
	JOIN users m ON p.member_id = m.user_id
	JOIN users s ON p.staff_id = s.user_id
	JOIN membership_types mb ON p.membership_type_id = mb.type_id WHERE p.payment_id = ?`, paymentId);

	var payment Entities.Payment;
	if err := row.Scan(&payment.PaymentID, &payment.MemberName, &payment.StaffName, &payment.MembershipType, &payment.Amount, &payment.Date, &payment.PhotoPath); err != nil {
		return nil, err;
	}

	return &payment, nil;
}

func Create(paymentData *Entities.CreatePayment) error {
	query := "INSERT INTO payments (member_id, staff_id, membership_type_id, amount, date, photo_path) VALUE (?, ?, ?, ?, ?, ?)"

	_, err := DB.Connection.Exec(query, paymentData.MemberID, paymentData.StaffID, paymentData.MembershipTypeID, paymentData.Amount, paymentData.Date, paymentData.PhotoPath);
	if err != nil {
		return err;
	}

	return nil;
}