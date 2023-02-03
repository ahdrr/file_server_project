package model

type Dir_stat struct {
	Dmsg
	Data []File `json:"data"`
}

type File struct {
	Name     string `json:"name"`
	Time     string `json:"time"`
	Size     string `json:"size"`
	Filetype string `json:"filetype"`
}

type Mverr struct {
	Name string `json:"name"`
	Err  string `json:"err"`
}
