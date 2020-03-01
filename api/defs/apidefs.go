package defs

//request
type UserCredential struct {
	Username string `json:"user_name"`
}

//data model
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}
