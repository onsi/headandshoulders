package headandshoulders

import (
	"fmt"

	"github.com/onsi/ginkgo"
)

func RIt(text string, body interface{}, count ...int) bool {
	for i := 1; i <= tries(count); i++ {
		description := fmt.Sprintf("%s (attempt %d)", text, i)
		ginkgo.It(description, body)
	}

	return true
}

const defaultTries = 3

func tries(count []int) int {
	if len(count) > 0 {
		return count[0]
	}

	return defaultTries
}
