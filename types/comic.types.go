package types

type ComicListInfoType struct {
	CoverImage  string           `json:"cover_image"`
	ComicType   ComicType        `json:"comic_type"`
	Title       string           `json:"title"`
	Url         string           `json:"url"`
	LastChapter string           `json:"last_chapter"`
	ComicRating *ComicRatingType `json:"comic_rating"`
	Slug        string           `json:"slug"`
}

type ComicRatingType struct {
	StarRating int8   `json:"star_rating"`
	Rating     string `json:"rating"`
}

type ComicType string

const (
	Manhwa ComicType = "Manhwa"
	Manhua ComicType = "Manhua"
	Manga  ComicType = "Manga"
	Novel  ComicType = "Novel"
	Comic  ComicType = "Comic"
)

type ComicStatusType string

const (
	Completed ComicStatusType = "Completed"
	Ongoing   ComicStatusType = "Ongoing"
	Hiatus    ComicStatusType = "Hiatus"
)

type ComicOrderType struct {
	AZ      string
	ZA      string
	Update  string
	Added   string
	Popular string
}

type ComicDetailType struct {
	Title         string              `json:"title"`
	Alias         string              `json:"alias"`
	CoverImage    string              `json:"cover_image"`
	ComicRating   *ComicRatingType    `json:"comic_rating"`
	Summary       string              `json:"summary"`
	FirstChapter  *ChapterDetailType  `json:"first_chapter"`
	LatestChapter *ChapterDetailType  `json:"latest_chapter"`
	ComicInfo     *ComicInfoType      `json:"comic_info"`
	Genres        []GenreType         `json:"genres"`
	Chapters      []ChapterDetailType `json:"chapters"`
	RelatedSeries []ComicListInfoType `json:"related_series"`
}

type ChapterDetailType struct {
	Chapter     string `json:"chapter"`
	ChapterUrl  string `json:"chapter_url"`
	ChapterSlug string `json:"chapter_slug"`
	LastUpdated string `json:"last_updated,omitempty"`
}

type ChapterImageType struct {
	Image string `json:"image"`
}

type ComicInfoType struct {
	Status        ComicStatusType `json:"status"`
	ReleasedDate  string          `json:"released_date"`
	Artist        string          `json:"artist"`
	UpdatedAt     string          `json:"updated_at"`
	ComicType     ComicType       `json:"comic_type"`
	Author        string          `json:"author"`
	Serialization string          `json:"serialization"`
	PostedOn      string          `json:"posted_on"`
	Views         string          `json:"views"`
}

type GenreType struct {
	Genre string `json:"genre"`
	Url   string `json:"url"`
	Slug  string `json:"slug"`
}
