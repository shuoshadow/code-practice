package lock

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func GetProcessID() int {
	return os.Getgid()
}

func GetGoroutineID() string {
	buf := make([]byte, 128)
	stk := buf[:runtime.Stack(buf, false)]
	stkStr := string(stk)
	return strings.TrimSpace(strings.Split(strings.Split(stkStr, "[running]")[0], " ")[1])
}

func GetOwnerId() string {
	pid := GetProcessID()
	gid := GetGoroutineID()
	return fmt.Sprintf("%d-%s", pid, gid)
}
