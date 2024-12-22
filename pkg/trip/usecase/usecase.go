package usecase

import (
	config "api-gateway"
	"api-gateway/pkg/trip/repository"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAllTrips(c *fiber.Ctx) ([]repository.Trip, error) {
	path := config.AppConfig()
	dir := fmt.Sprintf("%s/assets/trips.json", path.BasePath)
	result, err := repository.TripsData(dir)
	if err != nil {
		log.Printf("- Error =======>%s", err)
		return nil, err
	}

	return result.Trips, err
}
func SelectTrips(c *fiber.Ctx) (*repository.TripsResponse, error) {
	path := config.AppConfig()
	// keyword := c.Params("keyword")
	keyword := c.Query("keyword")
	if keyword == "" {
		log.Printf("- keyword is empty")

		return nil, fiber.NewError(fiber.StatusBadRequest, "Keyword parameter is required")
	}
	log.Println("- keyword is ", keyword)
	dir := fmt.Sprintf("%s/assets/trips.json", path.BasePath)
	result, err := repository.TripsData(dir)
	if err != nil {
		log.Printf("- Error =======>%s", err)
		return nil, err
	}
	// กรองข้อมูลจาก Trips โดยดูว่า tags มี keyword หรือไม่
	var filteredTrips []repository.Trip
	for _, trip := range result.Trips {
		for _, tag := range trip.Tags {
			if tag == keyword {
				filteredTrips = append(filteredTrips, trip)
				break
			}
		}
	}
	response := &repository.TripsResponse{
		Trips: filteredTrips,
	}
	log.Println("- response ====> ", response)
	// หากไม่เจอ trip ที่ตรงกับ keyword ให้คืนค่าผลลัพธ์เป็นว่าง
	if len(filteredTrips) == 0 {
		log.Printf("- No trips found for keyword: %s", keyword)
		return nil, fiber.NewError(fiber.StatusNotFound, "No trips found for the specified keyword")
	}
	return response, nil
}
