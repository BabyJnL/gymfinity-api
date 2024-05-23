package Validator 

import (
	"github.com/go-playground/validator/v10"
)

func ValidateGender(fl validator.FieldLevel) bool {
	gender := fl.Field().String();
	return gender == "Male" || gender == "Female";
}

func ValidateStatus(fl validator.FieldLevel) bool {
	status := fl.Field().String();
	return status == "Active" || status == "Expired";
}

func ValidateRole(fl validator.FieldLevel) bool {
	role := fl.Field().String();
	return role == "Member" || role == "Instructor" || role == "Staff";
}

func RegisterCustomValidators(v *validator.Validate) {
	v.RegisterValidation("validateGender", ValidateGender);
	v.RegisterValidation("validateStatus", ValidateStatus);
	v.RegisterValidation("validateRole", ValidateRole);
}