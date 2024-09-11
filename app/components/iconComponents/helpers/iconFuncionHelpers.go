package iconHelpers

import (
	"fmt"
	"strings"
)

func SizeParamHandler(defaultClasses string, iconSize int8) string {
	if iconSize == 0 {
		return defaultClasses
	}

	stringifiedIconSize := fmt.Sprintf("%d", iconSize)

	return strings.ReplaceAll(defaultClasses, "8", stringifiedIconSize)
}
