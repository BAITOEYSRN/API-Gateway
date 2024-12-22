package trip

import (
	"api-gateway/pkg/models"
	"api-gateway/pkg/trip/usecase"

	"github.com/gofiber/fiber/v2"
)

func Handlers(app *fiber.App) {
	app.Get("/trips", getAllTrips)
	app.Get("/api/trips", getTrips)
	// app.Get("/api/trips/:keyword", getTrips)

}
func getAllTrips(c *fiber.Ctx) error {
	result, err := usecase.GetAllTrips(c)
	if err != nil {
		return models.ResponseError(c, 500, err)
	}
	return models.ResponseSuccess(c, result)
}

func getTrips(c *fiber.Ctx) error {
	result, err := usecase.SelectTrips(c)
	if err != nil {
		return models.ResponseError(c, 500, err)
	}
	return models.ResponseSuccess(c, result)
}
