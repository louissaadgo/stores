package views

import "stores/models"

type AllAdmins struct {
	Admins []models.Admin `json:"admins"`
}

type CurrentAdmin struct {
	Admin models.Admin `json:"admin"`
}
