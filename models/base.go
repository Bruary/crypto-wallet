package models

type BaseResponse struct {
	Success bool
	Msg     string
	Error   Error
}

type Error struct {
	Code int
	Msg  string
}
