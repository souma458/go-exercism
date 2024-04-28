package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	validLogRegex, _ := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return validLogRegex.MatchString(text)
}

func SplitLogLine(text string) []string {
	regex, _ := regexp.Compile(`<[-~=*]*>`)
	return regex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	regex, _ := regexp.Compile(`(?i)"([^"]*password[^"]*)"`)
	for _, s := range lines {
		if regex.MatchString(s) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	regex, _ := regexp.Compile(`end-of-line\d*`)
	return regex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var taggedLines []string
	regex, _ := regexp.Compile(`User\s+(\w+)`)
	for _, line := range lines {
		match := regex.FindStringSubmatch(line)
		if len(match) > 1 {
			userName := match[1]
			taggedLines = append(taggedLines, fmt.Sprintf("[USR] %s %s", userName, line))
		} else {
			taggedLines = append(taggedLines, line)
		}

	}
	return taggedLines
}
