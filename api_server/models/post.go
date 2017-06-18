package models

import (
	"encoding/json"

	"../strip"
)

type Post struct {
	Title      string
	Slug       string
	URL        string
	Summary    string
	Date       string
	Sentiment  float32
	Image      string
	CategoryID Category
}

type PostJSON struct {
	Title      string   `json:"title, omitempty"`
	Slug       string   `json:"slug, omitempty"`
	URL        string   `json:"url, omitempty"`
	Summary    string   `json:"summary"`
	Date       string   `json:"date, omitempty"`
	Sentiment  float32  `json:"sentiment"`
	Image      string   `json:"image"`
	CategoryID Category `json:"category_id, omitempty"`
}

func (p *Post) MarshalJSON() ([]byte, error) {
	return json.Marshal(PostJSON{
		p.Title,
		p.Slug,
		p.URL,
		strip.StripTags(p.Summary),
		p.Date,
		p.Sentiment,
		p.Image,
		p.CategoryID,
	})
}

func (p *Post) UnmarshalJSON(b []byte) error {
	temp := &PostJSON{}

	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}

	p.Title = temp.Title
	p.Slug = temp.Slug
	p.URL = temp.URL
	p.Summary = temp.Summary
	p.Date = temp.Date
	p.Sentiment = temp.Sentiment
	p.Image = temp.Image
	p.CategoryID = temp.CategoryID

	return nil
}