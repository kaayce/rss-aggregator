package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey - extracts an APIKey from headers of HTTP Reques
// eg: Authorization: ApiKey {insert key here}
func GetAPIKey(headers http.Header) (string, error) {
	headerAuthKey := headers.Get("Authorization")

	if headerAuthKey == "" {
		return headerAuthKey, errors.New("no authentication info found")
	}

	values := strings.Split(headerAuthKey, " ")
	if len(values) != 2 {
		return "", errors.New("malformed authentication header string")
	}
	if values[0] != "ApiKey" {
		return "", errors.New("malformed first path of authentication header string")
	}
	return values[1], nil
}
