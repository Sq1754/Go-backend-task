package service

import "time"

// CalculateAge calculates age from date of birth
func CalculateAge(dob time.Time) int {
	now := time.Now()

	age := now.Year() - dob.Year()

	// If birthday has not occurred yet this year, subtract 1
	if now.YearDay() < dob.YearDay() {
		age--
	}

	return age
}
