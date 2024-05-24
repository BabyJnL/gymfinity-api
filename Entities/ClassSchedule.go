package Entities

type ClassSchedule struct {
	ScheduleID		int		`json:"scheduleId"`
	ClassID			int		`json:"classId" validate:"required"`
	InstructorID	int		`json:"instructorId" validate:"required"`	
	Date			string	`json:"date" validate:"required"`
	StartTime		string	`json:"startTime" validate:"required"`
	EndTime			string	`json:"endTime" validate:"required"`
}