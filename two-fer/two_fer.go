package twofer

import "fmt"

// ShareWith returns a custom version for sentence One for you, one for me
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
