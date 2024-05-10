package twofer

import "fmt"

// Determines what to say when we're given an extra cookie
func ShareWith(name string) string {
	template := "One for %s, one for me."
	if len(name) > 0 {
		return fmt.Sprintf(template, name)
	} else {
		return fmt.Sprintf(template, "you")
	}
}
