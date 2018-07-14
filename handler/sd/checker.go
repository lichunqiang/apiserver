package sd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"net/http"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func HealthCheck(c *gin.Context) {
	msg := "OK"

	c.String(http.StatusOK, "\n"+msg)
}

//func DiskCheck(c *gin.Context) {
//	u, _ := disk.Usage("/Users/light")
//
//	usedMB := int(u.Used) / MB
//	usedGB := int(u.Used) / GB
//	totalMB := int(u.Total) / MB
//	totalGB := int(u.Total) / GB
//
//	usedPercent := int(u.UsedPercent)
//
//	status := http.StatusOK
//	msg := "OK"
//
//	if usedPercent >= 95 {
//		msg = "CRITICAL"
//	} else if usedPercent >= 90 {
//		status = http.StatusTooManyRequests
//		msg = "WARNING"
//	}
//
//	c.String(
//		status,
//		fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", msg, usedMB, usedGB, totalMB, totalGB, usedPercent),
//	)
//}

func CPUCheck(c *gin.Context) {
	cores, _ := cpu.Counts(false)

	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15

	status := http.StatusOK
	text := "OK"

	if l5 >= float64(cores-1) {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if l5 >= float64(cores-2) {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Load average: %.2f, %.2f, %.2f | Cores: %d", text, l1, l5, l15, cores)
	c.String(status, "\n"+message)
}
