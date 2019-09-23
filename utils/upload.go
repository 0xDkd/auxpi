package utils

import (
	"regexp"
)

func CheckUrl(url string) (bool, bool) {
	var re = regexp.MustCompile(`(?m)https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&\/\/=]*)`)
	result := re.MatchString(url)
	//logs.Alert(result, url)
	isLocal := false
	var rex = regexp.MustCompile(`(?m)localhost`)
	isLocal = rex.MatchString(url)

	return result, isLocal
}
