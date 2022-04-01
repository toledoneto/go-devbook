package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorID       uint64    `json:"authorId,omitempty"`
	AuthorUsername string    `json:"authorPostname,omitempty"`
	Likes          uint64    `json:"likes"`
	CreatedAt      time.Time `json:"autoCreateTime"`
	UpdatedAt      time.Time `json:"autoUpdateTime"`
}

func (post *Post) Prepare() error {
	if err := post.validate(); err != nil {
		return err
	}

	post.format()
	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("title must not be empty")
	}

	if post.Content == "" {
		return errors.New("content must not be empty")
	}

	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
