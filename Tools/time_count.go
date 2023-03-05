package tools

import (
	"time"
	"fmt"
)

func Time_count(start_time time.Time) string {
	elapsed := time.Since(start_time)
	return fmt.Sprintf("Total time = %v", elapsed)
}
