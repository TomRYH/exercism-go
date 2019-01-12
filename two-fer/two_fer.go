// Package twofer contains sharing functions
package twofer

import (
	"fmt"
)

// ShareWith shares with the given name
func ShareWith(name string) string {

	retStr := "One for %s, one for me."

	if len(name) > 0 {
		return fmt.Sprintf(retStr, name)
	}
	return fmt.Sprintf(retStr, "you")
}
