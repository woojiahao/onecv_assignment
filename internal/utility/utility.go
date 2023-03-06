package utility

import (
	"net/mail"
	"strings"
)

func GetMentionsFromNotification(notification string) []string {
	var mentions []string
	for _, part := range strings.Fields(notification) {
		if part[0:1] != "@" {
			continue
		}

		if _, err := mail.ParseAddress(part[1:]); err == nil {
			mentions = append(mentions, part[1:])
		}
	}

	return mentions
}
