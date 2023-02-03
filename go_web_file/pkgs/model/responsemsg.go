package model

type Dmsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type MvResoponse struct {
	Dmsg
	Err []Mverr `json:"err"`
}
