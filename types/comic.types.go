package types

type ComicType struct {
	CoverImage  string           `json:"cover_image"`
	ComicType   string           `json:"comic_type"`
	Title       string           `json:"title"`
	Url         string           `json:"url"`
	LastChapter ComicChapterType `json:"last_chapter"`
	ComicRating ComicRatingType  `json:"comic_rating"`
}

type ComicRatingType struct {
	StarRating int8   `json:"star_rating"`
	Rating     string `json:"rating"`
}

type ComicChapterType struct {
	LastChapter    string `json:"last_chapter"`
	LastChapterUrl string `json:"last_chapter_url"`
}
