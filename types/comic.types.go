package types

type ComicListInfoType struct {
	CoverImage  string            `json:"cover_image"`
	ComicType   ComicType         `json:"comic_type"`
	Title       string            `json:"title"`
	Url         string            `json:"url"`
	LastChapter *ComicChapterType `json:"last_chapter"`
	ComicRating *ComicRatingType  `json:"comic_rating"`
}

type ComicRatingType struct {
	StarRating int8   `json:"star_rating"`
	Rating     string `json:"rating"`
}

type ComicChapterType struct {
	LastChapter    string `json:"last_chapter"`
	LastChapterUrl string `json:"last_chapter_url"`
}

type ComicType string

const (
	Manhwa ComicType = "Manhwa"
	Manhua ComicType = "Manhua"
	Manga  ComicType = "Manga"
)

type ComicStatusType string

const (
	Completed ComicStatusType = "Completed"
	Ongoing   ComicStatusType = "Ongoing"
)

type ComicDetailType struct {
	ThumbnailImage string              `json:"thumbnail_image"`
	Title          string              `json:"title"`
	Alias          string              `json:"alias"`
	ComicRating    ComicRatingType     `json:"comic_rating"`
	Genres         []string            `json:"genres"`
	ReleasedDate   string              `json:"released_date"`
	Status         ComicStatusType     `json:"status"`
	TotalChapters  string              `json:"total_chapters"`
	Author         string              `json:"author"`
	ComicType      ComicType           `json:"comic_type"`
	UpdatedAt      string              `json:"updated_at"`
	Summary        string              `json:"summary"`
	Chapters       []ChapterDetailType `json:"chapters"`
}

type ChapterDetailType struct {
	Chapter     string `json:"chapter"`
	LastUpdated string `json:"last_updated"`
}

type ChapterImageType struct {
	Image string `json:"image"`
}
