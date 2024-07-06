package types

type ComicType struct {
	CoverImage  string           `json:"cover_image"`
	ComicType   string           `json:"comic_type"`
	Title       string           `json:"title"`
	Url         string           `json:"url"`
	LastChapter string           `json:"last_chapter"`
	ComicRating ComicRatingTypes `json:"comic_rating"`
}

type ComicRatingTypes struct {
	StarRating int     `json:"star_rating"`
	Rating     float32 `json:"rating"`
}
