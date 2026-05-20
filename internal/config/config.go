package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server    ServerConfig
	JWT       JWTConfig
	Services  ServicesConfig
	Redis     RedisConfig
	RateLimit RateLimitConfig
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type JWTConfig struct {
	Secret          string
	ExpirationHours int
}

type ServicesConfig struct {
	UserServiceURL    string
	OrderServiceURL   string
	ProductServiceURL string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

type RateLimitConfig struct {
	Requests int
	Duration time.Duration
}

func Load() (*Config, error) {
	godotenv.Load()

	return &Config{
		Server: ServerConfig{
			Port:         getEnv("SERVER_PORT", "8080"),
			ReadTimeout:  getDurationEnv("SERVER_READ_TIMEOUT", 30*time.Second),
			WriteTimeout: getDurationEnv("SERVER_WRITE_TIMEOUT", 30*time.Second),
		},
		JWT: JWTConfig{
			Secret:          getEnv("JWT_SECRET", "default-secret"),
			ExpirationHours: getIntEnv("JWT_EXPIRATION_HOURS", 24),
		},
		Services: ServicesConfig{
			UserServiceURL:    getEnv("USER_SERVICE_URL", "http://localhost:8081"),
			OrderServiceURL:   getEnv("ORDER_SERVICE_URL", "http://localhost:8082"),
			ProductServiceURL: getEnv("PRODUCT_SERVICE_URL", "http://localhost:8083"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
		},
		RateLimit: RateLimitConfig{
			Requests: getIntEnv("RATE_LIMIT_REQUESTS", 100),
			Duration: getDurationEnv("RATE_LIMIT_DURATION", 1*time.Minute),
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}
