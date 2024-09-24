package vo

type SettingsVO struct {
	AboutMe           string         `json:"aboutMe"`
	ICP               string         `json:"icp"`
	AboutMePageSongId string         `json:"aboutMePageSongId"`
	IndexName         string         `json:"indexName"`
	IndexUrl          string         `json:"indexUrl"`
	Description       string         `json:"description"`
	LogoUrl           string         `json:"logoUrl"`
	Announcement      string         `json:"announcement"`
	VisitorCount      uint           `json:"visitorCount"`
	RandomBlogVOs     []RandomBlogVO `json:"randomBlogs"`
	RencentMoments    []any          `json:"recentMoments"`
}
