package auth

import (
	"errors"
	"net/http"
	"strings"
)

// get api key from auth header
// Example:
// Authorization: ApiKey {insert key here}
func GetApiKey (headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no autorization found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2{
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed auth header 1")
	}

	return vals[1], nil
}