package utils

import (
	"log/slog"
	"regexp"
)

func NormalizePhone(phone string, logger *slog.Logger) (string, bool) {
	re := regexp.MustCompile(`\D`)
	cleaned := re.ReplaceAllString(phone, "")

	if len(cleaned) < 11 {
		logger.Info("phone number len: %d", len(cleaned))
		return "", false
	}

	if cleaned[0] == '8' {
		cleaned = "7" + cleaned[1:]
	}

	validFormat := regexp.MustCompile(`7\d{10}$`)
	if !validFormat.MatchString(cleaned) {
		logger.Info("phone format err [%s]", cleaned)
		return "", false
	}
	return cleaned, true
}
