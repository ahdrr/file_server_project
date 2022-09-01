package systeminfo

import (
	"github.com/DataDog/gopsutil/disk"
	"filrserver/pkgs/model"

)


var diskdeDetails model.Diskjson
func GetDiskinfo(path string) (model.Diskjson, error) {
	a, err := disk.Usage(path)
	if err != nil {
		return model.Diskjson{}, err
	}
	diskdeDetails = model.Diskjson{
		Disktotal:   a.Total,
		Diskfree:    a.Free,
		Diskuse:     a.Used,
		UsedPercent: (uint64(a.UsedPercent) + 1),
	}
	return diskdeDetails, err
}
