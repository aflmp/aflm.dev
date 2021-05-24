package server

import (
	"embed"
	"encoding/json"
)

type Config struct {
	Port     string
	PostList []byte
	Blog     embed.FS
	Assets   embed.FS
}

// postConfig represents posts.json
type postConfig struct {
	Posts posts `json:"posts"`
}

func (p *postConfig) UnmarshalJSON(data []byte) error {
	var pc struct {
		Posts []struct {
			Title string `json:"title"`
			Date  string `json:"date"`
		} `json:"Posts"`
	}

	err := json.Unmarshal(data, &pc)
	if err != nil {
		return err
	}

	p.Posts = make(posts, len(pc.Posts))
	for i, post := range pc.Posts {
		p.Posts[i].Title = post.Title
		p.Posts[i].Date = post.Date
		p.Posts[i].ID = idFromTitle(post.Title)
	}

	return nil
}
