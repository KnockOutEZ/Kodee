package systemUsage

import (
	"fmt"
	"time"

	"github.com/KnockOutEZ/Kodee/backend/utils"
	"github.com/shirou/gopsutil/cpu"
)

func GetCpuUsage() string{
	cpuPercent, err := cpu.Percent(time.Second, false)
	utils.CheckErr(err)
	usedPercent := fmt.Sprintf("%.2f", cpuPercent[0])
	return usedPercent + "%"
}