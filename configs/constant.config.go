package configs

import "github.com/Strayneko/KomikcastAPI/types"

var ComicOrderBy types.ComicOrderType = types.ComicOrderType{
	AZ:      "title",
	ZA:      "titlereverse",
	Update:  "update",
	Added:   "latest",
	Popular: "popular",
}

var ComicOrderParams []string = []string{"az", "za", "update", "added", "popular"}

func GetComicOrderBy(order string) string {
	switch order {
	case "az":
		return ComicOrderBy.AZ
	case "za":
		return ComicOrderBy.ZA
	case "update":
		return ComicOrderBy.Update
	case "added":
		return ComicOrderBy.Added
	case "popular":
		return ComicOrderBy.Popular
	default:
		return ""
	}
}
