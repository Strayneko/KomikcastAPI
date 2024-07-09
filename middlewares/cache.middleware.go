package middlewares

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Strayneko/KomikcastAPI/configs"
	"github.com/Strayneko/KomikcastAPI/types"
	"github.com/gofiber/fiber/v2"
)

// CacheMiddleware creates a middleware for caching GET requests based on page type.
// It caches responses for 60 minutes, using the request path and query parameters to generate cache keys.
func CacheMiddleware(pageType string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if ctx.Method() != "GET" {
			// Only cache GET requests
			return ctx.Next()
		}
		params := "?"
		page := ctx.Query("page", "1")
		if pageType == "list" {
			params += "page=" + page
		} else if pageType == "search" {
			params += "query=" + ctx.Query("query", "") + "&page=" + page
		}

		cacheKey := ctx.Path() + params // Generate a cache key from the request path and query parameters
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

		// Cache the response for 60 minutes
		configs.Cache.Set(cacheKey, data, 60*time.Minute)
		return nil
	}
}
