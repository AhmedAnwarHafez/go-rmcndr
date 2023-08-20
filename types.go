package main

type Song struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Artist  string `json:"artist"`
	Genre   string `json:"genre"`
	Link    string `json:"link"`
	Comment string `json:"comment"`
}

type User struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	IsPublic bool   `json:"is_public"`
	Songs    []Song `json:"songs"`
}
