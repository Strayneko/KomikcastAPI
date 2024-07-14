package helpers

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/Strayneko/KomikcastAPI/interfaces"
	"github.com/Strayneko/KomikcastAPI/types"
	"github.com/gofiber/fiber/v2"
)

type helper struct {
	controllerHelper interfaces.Helper
}

func New() interfaces.Helper {
	return &helper{}
}

// ValidatePage validates the "page" query parameter from the request context.
// It ensures that the page is a positive integer and converts it to int16.
func (h *helper) ValidatePage(ctx *fiber.Ctx) (int16, *fiber.Error) {
	page := ctx.Query("page", "1")
	curPage, err := strconv.ParseInt(page, 10, 16)

	if err != nil || curPage <= 0 {
		return 0, fiber.NewError(http.StatusBadRequest, "Bad Request: Invalid page")
	}

	return int16(curPage), nil
}

func (h *helper) ResponseError(ctx *fiber.Ctx, err *fiber.Error) error {
	return ctx.Status(err.Code).JSON(&types.ResponseType{
		Status:  false,
		Code:    int16(err.Code),
		Message: err.Message,
	})
}

func (h *helper) ExtractSlug(url string) string {
	slug := strings.TrimSuffix(url, "/")

	splitedSlug := strings.Split(slug, "/")

	if len(splitedSlug) == 0 {
		return ""
	}

	return splitedSlug[len(splitedSlug)-1]
}

// ExtractStarRatingValue extracts the star rating value from a width css attribute Ex: width: 70%, will result 3.5.
// It uses a regular expression to find and return the number in the string.
func (h *helper) ExtractStarRatingValue(starRating string) int8 {
	// Compile the regex pattern to extract the number
	re := regexp.MustCompile(`\d+(\.\d+)?`)

	// Find the match
	match := re.FindString(starRating)

	if len(match) == 0 {
		return 0
	}
	res, err := strconv.Atoi(match)
	if err != nil {
		return 0
	}

	return int8(res / 20)

}
