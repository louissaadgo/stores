package controllers

import (
	"stores/models"
	"unsafe"

	"github.com/gofiber/fiber/v2"
)

func StoreSignup(c *fiber.Ctx) error {

	signupModel := models.StoreSignup{}
	err := c.BodyParser(&signupModel)
	if err != nil {
		return c.SendString("Invalid data(JSON) sent")
	}

	store := models.Store{
		Name:        signupModel.Name,
		Address:     signupModel.Address,
		CountryCode: signupModel.CountryCode,
		Phone:       signupModel.Phone,
		PublicEmail: signupModel.PublicEmail,
		Email:       signupModel.Email,
		Password:    signupModel.Password,
	}

	storeResponse := *(*models.StoreResponse)(unsafe.Pointer(&store))

	response := models.Response{
		Type:    "success",
		Message: "signup successfully",
		Data:    storeResponse,
	}

	return c.JSON(response)
}
