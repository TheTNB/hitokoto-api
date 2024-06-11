package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// Return Fiber configuration.
	return fiber.Config{
		AppName:     "Hitokoto API",
		Prefork:     true,
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
		JSONDecoder: sonic.Unmarshal,
		JSONEncoder: sonic.Marshal,
	}
}
