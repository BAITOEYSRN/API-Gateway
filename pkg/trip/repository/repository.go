package repository

import (
	"encoding/json"
	"fmt"
	"os"
)

type Trip struct {
	Title       string   `json:"title"`
	Eid         string   `json:"eid"`
	URL         string   `json:"url"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Tags        []string `json:"tags"`
}

// Define the structure for the response which contains an array of trips
type TripsResponse struct {
	Trips []Trip `json:"trips"`
}

func TripsData(filePath string) (*TripsResponse, error) {
	// ตรวจสอบว่าไฟล์มีอยู่จริงหรือไม่

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file does not exist: %v", err)
	}

	// อ่านไฟล์ JSON
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	// แปลงข้อมูล JSON เป็น struct
	var tripsResponse TripsResponse
	err = json.Unmarshal(data, &tripsResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %v", err)
	}

	// คืนค่าผลลัพธ์
	return &tripsResponse, nil
}
