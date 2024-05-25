package Entities

type Reservation struct {
	ReservationID		int		`json:"reservationId"`
	MemberId			int		`json:"memberId"`
	ClassId				int		`json:"classId"`
	Status				string	`json:"status"`
	Date				string	`json:"date"`
}

type CreateReservation struct {
	ReservationID		int		`json:"reservationId"`
	MemberId			int		`json:"memberId" validate:"required"`
	ClassId				int		`json:"classId" validate:"required"`
	StatusId			int		`json:"statusId" validate:"required"`
	Date				string	`json:"date" validate:"required"`
}

type UpdateReservationStatus struct {
	StatusId			int		`json:"statusId" validate:"required"`
}