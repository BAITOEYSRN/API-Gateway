package main

import (
	config "api-gateway"
	trip "api-gateway/pkg/trip/handler"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	appPort := config.AppConfig()
	app := fiber.New()

	// การตั้งค่า CORS ที่ปลอดภัย
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization,x-apikey",
		AllowOrigins:     "*",                                      // กำหนด origin ที่อนุญาต
		AllowCredentials: false,                                    // หากต้องการอนุญาตให้มีการส่งคุกกี้ หรือข้อมูลการยืนยันตัวตน
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS", // กำหนด HTTP methods ที่อนุญาต
	}))

	// Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "ok",
			"message": "Service is healthy",
		})
	})
	trip.Handlers(app)
	port := appPort.Port
	log.Fatal(app.Listen(port))
}
