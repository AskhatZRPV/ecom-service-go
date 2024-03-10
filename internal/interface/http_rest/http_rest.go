package http_rest

import (
	"context"
	"ecomsvc/internal/core/config"
	"ecomsvc/internal/interface/http_rest/common"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// func New(lc fx.Lifecycle, config *config.Config, handlers []common.Handler) {
// 	f := fiber.New(
// 		fiber.Config{
// 			ErrorHandler: common.ErrorHandler,
// 			ReadTimeout:  time.Second * 3,
// 		},
// 	)
// 	for _, h := range handlers {
// 		f.Add(h.Method(), h.Pattern(), append(h.Middleware(), h.Handle)...)
// 	}
//
// 	lc.Append(fx.Hook{
// 		OnStart: func(_ context.Context) error {
// 			go func() {
// 				err := f.Listen(fmt.Sprintf("localhost:%d", config.HttpPort))
// 				if err != nil {
// 					panic(err)
// 				}
// 			}()
//
// 			return nil
// 		},
// 		OnStop: func(ctx context.Context) error {
// 			return f.ShutdownWithContext(ctx)
// 		},
// 	})
// }

func New(ctx context.Context, config *config.Config, handlers []common.Handler) {
	f := fiber.New(
		fiber.Config{
			ReadTimeout: time.Second * 3,
		},
	)

	f.Use(logger.New())

	for _, h := range handlers {
		f.Add(h.Method(), h.Pattern(), append(h.Middleware(), h.Handle)...)
	}

	go func() {
		// err := f.Listen(fmt.Sprintf("%s:%d", config.HttpServer.Host, config.HttpServer.Port))
		if err := f.Listen(":8080"); err != nil {
			panic(err)
		}
	}()

	<-ctx.Done()
	log.Info("Gracefully stopping Fiber Server")
	f.Shutdown()
}
