package Entities

type CreatePayment struct {
	PaymentID			int		`json:"paymentId"`
	MemberID			int		`json:"memberId" validate:"required"`
	StaffID				int		`json:"staffId" validate:"required"`
	MembershipTypeID	int		`json:"membershipTypeId" validate:"required"`
	Amount				float64	`json:"amount" validate:"required"`
	Date				string	`json:"date" validate:"required"`
	PhotoPath			string	`json:"photoPath" validate:"required"`
}

type Payment struct {
	PaymentID			int		`json:"paymentId"`
	MemberName			string	`json:"memberName"`
	StaffName			string	`json:"staffName"`
	MembershipType		string	`json:"membershipTypeId"`
	Amount				float64	`json:"amount"`
	Date				string	`json:"date"`
	PhotoPath			string	`json:"photoPath"`
}