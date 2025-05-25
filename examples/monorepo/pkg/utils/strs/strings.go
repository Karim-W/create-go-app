package strs

import (
	"encoding/base64"
	"regexp"
	"strings"
)

func Base64Encode(str string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(str))
}

func Base64EncodeBuffer(buff []byte) string {
	return base64.RawURLEncoding.EncodeToString(buff)
}

func Base64Decode(str string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(str)
}

func IsAlphanumeric(str string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9]+$")
	return re.MatchString(str)
}

func Escape(str string) string {
	str = strings.ReplaceAll(str, "-", `\\-`)
	str = strings.ReplaceAll(str, "/", `\\/`)
	return str
}

func Build(strs ...string) string {
	builder := strings.Builder{}
	for _, str := range strs {
		builder.WriteString(str)
	}

	return builder.String()
}
