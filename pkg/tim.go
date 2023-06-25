package pkg

import (
	"fmt"
	"time"
)

func NowTimeStr() string {
	return fmt.Sprintf("[%s]",time.Now().Format("2006-01-02 15:04:05"))
}