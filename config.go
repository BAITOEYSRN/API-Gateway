package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type ICnf struct {
	BasePath string
}

func AppConfig() ICnf {
	basepath := os.Getenv("BASE_PATH")
	if basepath == "" {
		_, b, _, _ := runtime.Caller(0)
		basepath = filepath.Dir(b)
	}
	err := godotenv.Load(basepath + "/.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}
	return ICnf{
		BasePath: basepath,
	}
}
