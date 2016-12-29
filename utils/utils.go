package utils

import "strings"

// IsReplyToBot check reply to bot.
func IsReplyToBot(botID string, replyTo []string) bool {
	for _, u := range replyTo {
		if u == botID {
			return true
		}
	}

	return false
}

// HasKeywords find keywords in the text.
func HasKeywords(text string, keywords ...string) bool {
	for _, k := range keywords {
		if strings.Contains(text, k) {
			return true
		}
	}

	return false
}
