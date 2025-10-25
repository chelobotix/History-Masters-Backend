package config

import (
	"fmt"
	"myapp/internal/database"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewConfig(e *echo.Echo) (*gorm.DB, error) {
	// DB Connection
	db, err := database.NewConnection()
	if err != nil {
		return nil, err
	}

	fmt.Println("connected to DB")

	return db, nil

}

// func startMiddlewares(e *echo.Echo) {
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
// 		StackSize: 1 << 10, // 1 KB
// 		LogLevel:  log.ERROR,
// 	}))

// 	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
// 		AllowOrigins:     []string{"https://*", "http://*"},
// 		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
// 		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
// 		AllowCredentials: true,
// 		MaxAge:           300,
// 	}))

// 	config := middleware.RateLimiterConfig{
// 		Skipper: middleware.DefaultSkipper,
// 		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
// 			middleware.RateLimiterMemoryStoreConfig{
// 				Rate:      rate.Limit(0.25),
// 				Burst:     5,
// 				ExpiresIn: 1 * time.Minute,
// 			},
// 		),

// 		IdentifierExtractor: func(c echo.Context) (string, error) {
// 			return c.RealIP(), nil
// 		},
// 		ErrorHandler: func(c echo.Context, err error) error {
// 			return c.JSON(http.StatusTooManyRequests, map[string]string{
// 				"error": "Too many requests",
// 			})
// 		},
// 		DenyHandler: func(c echo.Context, identifier string, err error) error {
// 			return c.JSON(http.StatusTooManyRequests, map[string]string{
// 				"error":      "Too many requests",
// 				"identifier": identifier,
// 			})
// 		},
// 	}

// 	e.Use(middleware.RateLimiterWithConfig(config))
// }

// func cronJobs(db *gorm.DB, c *cron.Cron, logger *zap.Logger) {
// 	err := c.AddFunc("* 0 3 * * *", func() {
// 		eis, err := updater_service.NewUpdater(db, logger)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		logger.Info("starting economy indexers cron job...")
// 		eis.PerformUpdate()
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	c.Start()
// }

// func testRedis(redisDB *redis.Client, logger *zap.Logger) error {
// 	response, err := redisDB.Ping(context.Background()).Result()
// 	if err != nil {
// 		return err
// 	}

// 	logger.Info(fmt.Sprintf("Redis says ...%s", response))

// 	return nil
// }
