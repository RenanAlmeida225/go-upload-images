package config

import "github.com/joho/godotenv"

func InitializeGodotenv() error {
	err := godotenv.Load()
	return err
}
