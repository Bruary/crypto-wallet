package models

type BaseResponse struct {
	Success string
	Msg     string
	Error   Error
}

type Error struct {
	Code int
	Msg  string
}
