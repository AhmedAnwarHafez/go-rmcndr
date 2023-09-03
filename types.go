package main

type Song struct {
	Id      int64  `json:"id"`
	UserId  int64  `json:"user_id"`
	Title   string `json:"title"`
	Artist  string `json:"artist"`
	GenreId int64  `json:"genre"`
	Link    string `json:"link"`
	Comment string `json:"comment"`
}

type User struct {
	Id       string
	Nickname string
	IsPublic bool
	Songs    []Song
}
