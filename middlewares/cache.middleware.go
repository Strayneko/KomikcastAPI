package middlewares

import (
	"context"
	"encoding/json"
	"github.com/Strayneko/KomikcastAPI/types"
	"log"
	"net/http"
	"time"

	"github.com/Strayneko/KomikcastAPI/configs"
	"github.com/gofiber/fiber/v2"
)

var rdbCtx = context.Background()

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
		cached, err := configs.Cache.Get(rdbCtx, cacheKey).Result()
		if err != nil {
			log.Printf("Can't retrieve cache key %s: %v", cacheKey, err)
		}

		// unmarshal cached value and return it as a response
		var responseBody *types.ResponseType
		err = json.Unmarshal([]byte(cached), &responseBody)
		if len(cached) > 0 && err == nil {
			return ctx.JSON(responseBody)
		}

		// skip caching when error happens
		err = ctx.Next()
		if err != nil || ctx.Response().StatusCode() != http.StatusOK {
			return err
		}

		body := ctx.Response().Body()
		// Cache the response for 60 minutes
		if err := configs.Cache.Set(rdbCtx, cacheKey, body, 60*time.Minute).Err(); err != nil {
			panic(err)
		}
		return nil
	}
}
