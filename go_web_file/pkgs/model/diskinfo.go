package model

type Diskstruct struct {
	Dmsg
	Data Diskjson `json:"data"`
}

type Diskjson struct {
	Disktotal   uint64 `json:"disktotal"`
	Diskfree    uint64 `json:"diskfree"`
	Diskuse     uint64 `json:"diskuse"`
	UsedPercent uint64 `json:"usedPercent"`
}
