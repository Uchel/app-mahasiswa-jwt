package model

type Student struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Major    string `json:"major"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
