package Entities

type User struct {
	UserID		int		`json:"userId"`
	Firstname	string	`json:"firstname"`
	Lastname	string	`json:"lastname"`
	Gender		string	`json:"gender"`
	Address		string	`json:"address"`
	PhoneNumber	string	`json:"phoneNumber"`
	Email		string	`json:"email"`
	JoinDate	string	`json:"joinDate"`
	Status		string	`json:"status"`
	ValidUntil	string	`json:"validUntil"`
	Role		string	`json:"role"`
	PhotoPath	string	`json:"photoPath"`
}