package web

type ResoponseLogin struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Gender   string `json:"gender"`
	Token    string `json:"token"`
}
