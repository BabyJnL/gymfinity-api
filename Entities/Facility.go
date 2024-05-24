package Entities 

type Facility struct {
	FacilityID		int		`json:"facilityId"`
	Name			string	`json:"name" validate:"required,max=64"`
	Description		string	`json:"description" validate:"required"`
	Photo			string	`json:"photo" validate:"required"`
}

type FacilityStatus struct {
	StatusID		int		`json:"statusId"`
	Name			string	`json:"name"`
}