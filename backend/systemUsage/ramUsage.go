package systemUsage

import (
	"fmt"

	"github.com/KnockOutEZ/Kodee/backend/utils"
	"github.com/shirou/gopsutil/mem"
)

func GetRamUsage() []string {
	m, err := mem.VirtualMemory()
	utils.CheckErr(err)
	usedMessage := fmt.Sprintf(
		"%s (%.2f%%)",
		getReadableSize(m.Used),
		m.UsedPercent,
	)
	return []string{usedMessage, getReadableSize(m.Total), getReadableSize(m.Available), getReadableSize(m.Free)}
}

func getReadableSize(sizeInBytes uint64) (readableSizeString string) {
	var (
		units = []string{"B", "KB", "MB", "GB", "TB", "PB"}
		size  = float64(sizeInBytes)
		i     = 0
	)
	for ; i < len(units) && size >= 1024; i++ {
		size = size / 1024
	}
	readableSizeString = fmt.Sprintf("%.2f %s", size, units[i])
	return
}