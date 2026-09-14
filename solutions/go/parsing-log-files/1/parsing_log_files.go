package parsinglogfiles

import "regexp"

var validLinePattern = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
var logLineSeparator = regexp.MustCompile(`<[~*=\-]*>`)
var quotedPasswordPattern = regexp.MustCompile(`(?i)".*password.*"`)
var endOfLinePattern = regexp.MustCompile(`end-of-line\d+`)
var userNamePattern = regexp.MustCompile(`User +(\S+)`)


func IsValidLine(text string) bool {
	return validLinePattern.MatchString(text)
    
}

func SplitLogLine(text string) []string {
	return logLineSeparator.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		if quotedPasswordPattern.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return endOfLinePattern.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
    	result := make([]string, len(lines))
	for i, line := range lines {
		match := userNamePattern.FindStringSubmatch(line)
		if match == nil {
			result[i] = line
			continue
		}
		result[i] = "[USR] " + match[1] + " " + line
	}
	return result
}
