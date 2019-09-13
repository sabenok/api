package core

import (
	"os"
	"strconv"
)

type Environment struct {
	DbName          string
	DbUser          string
	DbPassword      string
	DbPoolSize      int
	DbHost          string
	DbPort          int
	BaseCoin        string
	ServerPort      int
	IsDebug         bool
	WsServer        string
	WsBlocksChannel string
	NoahDevApiHost  string
}

func NewEnvironment() *Environment {

	env := Environment{
		DbName:          os.Getenv("DB_NAME"),
		DbUser:          os.Getenv("DB_USER"),
		DbPassword:      os.Getenv("DB_PASSWORD"),
		DbPoolSize:      getEnvAsInt("DB_POOL_SIZE", 10),
		DbHost:          os.Getenv("DB_HOST"),
		DbPort:          getEnvAsInt("DB_PORT", 5432),
		BaseCoin:        getEnv("BASE_COIN", "NOAH"),
		ServerPort:      getEnvAsInt("EXPLORER_API_PORT", 9050),
		IsDebug:         getEnvAsBool("DEBUG", true),
		WsServer:        os.Getenv("WS_SERVER"),
		WsBlocksChannel: os.Getenv("WS_BLOCKS_CHANNEL"),
		NoahDevApiHost:  os.Getenv("NOAH_DEV_API_HOST"),
	}

	return &env
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
