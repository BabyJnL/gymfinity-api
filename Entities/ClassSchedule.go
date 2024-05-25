package Entities

type CreateClassSchedule struct {
	ScheduleID		int		`json:"scheduleId"`
	ClassID			int		`json:"classId" validate:"required"`
	InstructorID	int		`json:"instructorId" validate:"required"`	
	Date			string	`json:"date" validate:"required"`
	StartTime		string	`json:"startTime" validate:"required"`
	EndTime			string	`json:"endTime" validate:"required"`
}

type ClassSchedule struct {
	ScheduleID			int		`json:"scheduleId"`
	ClassName			string	`json:"className"`
	ClassDescription	string	`json:"classDescription"`
	ClassQuota			int		`json:"classQuota"`
	InstructorName		string	`json:"instructorName"`	
	Date				string	`json:"date"`
	StartTime			string	`json:"startTime"`
	EndTime				string	`json:"endTime"`
}