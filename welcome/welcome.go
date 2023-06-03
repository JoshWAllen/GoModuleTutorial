package welcome

import "fmt"

func Phrase(name string) string {
	welcome_message := fmt.Sprintf("Welcome %s", name)
	return welcome_message
}