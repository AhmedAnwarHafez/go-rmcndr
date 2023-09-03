package main

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetSearchHandler(c *fiber.Ctx) error {

	q := c.Query("q")

	// means that htmx triggered this request
	if c.Get("Hx-Trigger-Name") != "q" {

		return c.Render("search", fiber.Map{
			"Title": "rcmndr",
			"Query": q,
		}, "layouts/main")
	}

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		return err
	}
	fmt.Println("q", q)
	rows, err := db.Query(`SELECT u.nickname as Nickname, u.is_public as IsPublic, g.name as GenreName, COUNT(r.id) AS GenereCount
FROM recommendations r
JOIN users u ON r.user_id = u.id
JOIN genres g ON r.genre_id = g.id
WHERE u.nickname LIKE ?
GROUP BY g.id`, "%"+q+"%")

	if err != nil {
		return err
	}

	defer rows.Close()

	type User struct {
		Id         int64
		Nickname   string
		IsPublic   bool
		Genres     []GenreSummary
		Similarity string
	}

	usersMap := make(map[string]*User)

	for rows.Next() {
		var user User
		var genre GenreSummary
		err := rows.Scan(&user.Nickname, &user.IsPublic, &genre.GenreName, &genre.Count)
		if err != nil {
			return err
		}
		// Check if user already exists in the map
		if _, exists := usersMap[user.Nickname]; !exists {
			usersMap[user.Nickname] = &User{Nickname: user.Nickname}
		}

		// Append the genre to the user's genres
		usersMap[user.Nickname].Genres = append(usersMap[user.Nickname].Genres, genre)
	}

	// Get the user's songs
	session, err := store.Get(c)
	if err != nil {
		return err
	}

	// Get the current user
	currentUserId, ok := session.Get("user_id").(int)
	if !ok {
		return fmt.Errorf("could not get user_id from session")
	}

	rows, err = db.Query(`SELECT  g.name as GenreName, COUNT(r.id) AS GenereCount
FROM recommendations r
JOIN users u ON r.user_id = u.id
JOIN genres g ON r.genre_id = g.id
WHERE u.id = ?
GROUP BY g.id`, currentUserId)

	if err != nil {
		return err
	}

	var myGenres []GenreSummary
	for rows.Next() {
		var genre GenreSummary
		err := rows.Scan(&genre.GenreName, &genre.Count)
		if err != nil {
			return err
		}

		myGenres = append(myGenres, genre)
	}

	for i, user := range usersMap {
		val := CosineSimilarity(myGenres, user.Genres)
		usersMap[i].Similarity = fmt.Sprintf("%.1f", val)
	}

	return c.Render("search-result", fiber.Map{
		"Title": "rcmndr",
		"Query": q,
		"Users": usersMap,
	})
}

func cosineSimilarity(a []string, b []string) float64 {
	return 0.0
}
