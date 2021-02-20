package models

type Film struct {
	ID             int
	name           string
	directorName   string
	productionYear int
	price          int
	length         int
	summary        string
	score          int
	comments       []Comment
}

type Comment struct {
	writer string
	text   string
}
