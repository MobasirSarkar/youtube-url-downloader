package models

type YouTubeResponse struct {
	Kind  string `json:"kind"`
	Etag  string `json:"etag"`
	Items []Item `json:"items"`
}

type Item struct {
	Id             string         `json:"id"`
	Kind           string         `json:"kind"`
	Etag           string         `json:"etag"`
	Snippet        Snippet        `json:"snippet"`
	ContentDeatils ContentDeatils `json:"contentDeatails"`
	Statistics     Stats          `json:"statistics"`
}

type Snippet struct {
	PublishedAt string     `json:"publishedAt"`
	ChannelId   string     `json:"channelId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Thumbnails  Thumbnails `json:"thumbnails"`
	CategoryId  string     `json:"categoryId"`
}

type ContentDeatils struct {
	Duration string `json:"duration"`
}

type Thumbnails struct {
	Default Thumbnail `json:"default"`
	Medium  Thumbnail `json:"medium"`
	High    Thumbnail `json:"high"`
}

type Thumbnail struct {
	Url string `json:"url"`
}

type Stats struct {
	ViewCount     string `json:"viewCount"`
	LikeCount     string `json:"likeCount"`
	DislikeCount  string `json:"dislikeCount"`
	FavoriteCount string `json:"favoriteCount"`
	CommentCount  string `json:"commentCount"`
}
