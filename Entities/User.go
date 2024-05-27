package Entities

type User struct {
	UserID		int		`json:"userId"`
	Firstname	string	`json:"firstname" validate:"required,max=64"`
	Lastname	string	`json:"lastname" validate:"required,max=64"`
	Gender		string	`json:"gender" validate:"required"`
	Address		string	`json:"address" validate:"required"`
	PhoneNumber	string	`json:"phoneNumber" validate:"required,max=12"`
	Email		string	`json:"email" validate:"required,email"`
	Pin			int		`json:"pin" validate:"required"`
	JoinDate	string	`json:"joinDate"`
	Status		string	`json:"status" validate:"required"`
	ValidUntil	string	`json:"validUntil"`
	Role		string	`json:"role" validate:"required"`
	PhotoPath	string	`json:"photoPath"`
}

type UserStatus struct {
	StatusID	int		`json:"statusId"`
	Name		string	`json:"name"`
}

type UserRole	struct {
	RoleID		int		`json:"roleId"`
	Name		string	`json:"name"`
}

type UserVerify	struct {
	Email		string	`json:"email" validate:"required,email"`
	Pin			int		`json:"pin" validate:"required"`
}