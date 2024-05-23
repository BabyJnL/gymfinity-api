package Entities

type User struct {
	UserID		int		`json:"userId"`
	Firstname	string	`json:"firstname" validate:"required,max=64"`
	Lastname	string	`json:"lastname" validate:"required,max=64"`
	Gender		string	`json:"gender" validate:"required,validateGender"`
	Address		string	`json:"address" validate:"required"`
	PhoneNumber	string	`json:"phoneNumber" validate:"required,max=12"`
	Email		string	`json:"email" validate:"required,email"`
	JoinDate	string	`json:"joinDate" validate:"required"`
	Status		string	`json:"status" validate:"required,validateStatus"`
	ValidUntil	string	`json:"validUntil"`
	Role		string	`json:"role" validate:"required,validateRole"`
	PhotoPath	string	`json:"photoPath"`
}