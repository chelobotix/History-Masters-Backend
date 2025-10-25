// package config

// import (
// 	"calculadora_rico/internal/database"
// 	"calculadora_rico/internal/services/error_handler_service"
// 	"calculadora_rico/internal/services/logger_service"
// 	"calculadora_rico/internal/services/updater_service"
// 	"context"
// 	"fmt"
// 	_ "github.com/joho/godotenv/autoload"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// 	"github.com/labstack/gommon/log"
// 	"github.com/redis/go-redis/v9"
// 	"github.com/robfig/cron"
// 	"go.uber.org/zap"
// 	"golang.org/x/time/rate"
// 	"gorm.io/gorm"
// 	"net/http"
// 	"os"
// 	"time"
// )

// func NewConfig(
// 	e *echo.Echo,
// ) (
// 	*gorm.DB,
// 	*redis.Client,
// 	*zap.Logger,
// 	error_handler_service.ErrorHandler,
// 	error,
// ) {
// 	// Zap Logger
// 	logger, err := logger_service.NewZapLogger(e)
// 	if err != nil {
// 		log.Panic(err)
// 		return nil, nil, nil, nil, err
// 	}

// 	// ErrorHandler Service
// 	errorHandler, err := error_handler_service.New(logger)
// 	if err != nil {
// 		return nil, nil, nil, nil, err
// 	}

// 	// DB Connection
// 	db, err := database.NewConnection()
// 	if err != nil {
// 		return nil, nil, nil, nil, err
// 	}

// 	// Redis Connection
// 	redisDB := redis.NewClient(&redis.Options{
// 		Addr:     os.Getenv("REDIS_URL"),
// 		Password: "",
// 		DB:       0,
// 	})

// 	err = testRedis(redisDB, logger)
// 	if err != nil {
// 		return nil, nil, nil, nil, err
// 	}

// 	// Migrate DB
// 	err = database.AutoMigrate(db)
// 	if err != nil {
// 		return nil, nil, nil, nil, err
// 	}

// 	// Create Fixtures
// 	err = database.LoadFixtures(db)
// 	if err != nil {
// 		return nil, nil, nil, nil, err
// 	}

// 	// Middlewares
// 	startMiddlewares(e)

// 	// Cron Jobs
// 	c := cron.New()
// 	cronJobs(db, c, logger)

// 	return db, redisDB, logger, errorHandler, err
// }

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
