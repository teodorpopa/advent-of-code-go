package utils

import (
	"github.com/gookit/color"
	"time"
)

const TARGET_TIME = 10 * time.Millisecond

func RunDay(days map[int]interface{}, day int, parts []int) {
	v, ext := days[day]
	if !ext {
		color.Errorf("Day %d has not been implemented yet\n\n", day)
		return
	}

	for _, part := range parts {
		color.Cyanf("\nRunning solution for day %d part %d\n", day, part)

		startTime := time.Now()
		answer := v.(func(int) int)(part)
		elapsedTime := time.Since(startTime)

		color.Successf("\nPart %d: %d\n\n", part, answer)

		if elapsedTime > TARGET_TIME {
			color.Warnf("Solution took %s to run.\n\n", elapsedTime)
		} else {
			color.Grayf("Solution took %s to run.\n\n", elapsedTime)
		}
	}

}
