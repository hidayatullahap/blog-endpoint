package models

import (
	"time"
)

type Post struct {
	ID        int       `db:"id"`
	Title     string    `db:"title"`
	Slug      string    `db:"slug"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func GetPosts() []Post {
	db := DbConnect()
	rows, err := db.Queryx("SELECT * FROM posts")
	defer db.Close()

	if err != nil {
		panic(err)
	}

	posts := make([]Post, 0) // 0 is the initial size of the slice
	for rows.Next() {
		var post Post
		err = rows.StructScan(&post)
		posts = append(posts, post)

		if err != nil {
			panic(err)
		}
	}

	return posts
}
