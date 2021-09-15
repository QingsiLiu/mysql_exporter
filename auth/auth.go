package auth

import (
	"mysql_exporter/config"
	"mysql_exporter/utils"
	"strings"
)

func isAuth(secret string, config *config.AuthConfig) bool {
	if config == nil || config.User == "" && config.Password == "" {
		return true
	}

	plaintext := utils.Base64Encode(strings.TrimPrefix(secret, "Basic "))

	elements := strings.SplitN(plaintext, ":", 2)
	if len(elements) != 2 {
		return false
	}

	return elements[0] == config.User && utils.Md5Text(elements[1]) == config.Password
}
