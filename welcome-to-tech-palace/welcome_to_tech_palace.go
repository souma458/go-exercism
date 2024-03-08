package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	newLine := "\n"
	return buildStarBorder(numStarsPerLine) + newLine + welcomeMsg + newLine + buildStarBorder(numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}

func buildStarBorder(numStars int) string {
	return strings.Repeat("*", numStars)
}
