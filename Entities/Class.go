package Entities;

type Class struct {
	ClassID		int		`json:"classId"`
	Name		string	`json:"name" validate:"required,max=64"`
	Description	string	`json:"description" validate:"required"`
	Quota		int		`json:"quota" validate:"required"`
}