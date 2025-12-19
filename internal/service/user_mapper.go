package service

import (
	"github.com/sq1754/user-age-api/db/sqlc"
	"github.com/sq1754/user-age-api/internal/models"
)

func ToUserModel(u sqlc.User) models.User {
	return models.User{
		ID:   int64(u.ID), // ✅ explicit conversion
		Name: u.Name,
		DOB:  u.Dob,
		Age:  CalculateAge(u.Dob),
	}
}
