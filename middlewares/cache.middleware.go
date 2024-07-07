package middlewares

import (
	"encoding/json"
	"github.com/Strayneko/KomikcastAPI/configs"
	"github.com/Strayneko/KomikcastAPI/types"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

func CacheMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if ctx.Method() != "GET" {
			// Only cache GET requests
			return ctx.Next()
		}

		cacheKey := ctx.Path() + "?page=" + ctx.Query("page", "1") // Generate a cache key from the request path and query parameters
		// Check if the response is already in the cache
		if cached, found := configs.Cache.Get(cacheKey); found {
			return ctx.JSON(cached)
		}
		err := ctx.Next()
		if err != nil || ctx.Response().StatusCode() != http.StatusOK {
			return err
		}
		var data types.ResponseType

		body := ctx.Response().Body()
		err = json.Unmarshal(body, &data)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		// Cache the response for 10 minutes
		configs.Cache.Set(cacheKey, data, 60*time.Minute)
		return nil
	}
}
