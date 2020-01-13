package manipulator

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"html"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
)

func HandleBase64(encode bool, data string) (string, error) {
	if encode {
		return b64.StdEncoding.EncodeToString([]byte(data)), nil
	}

	decoded, err := b64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func HandleURL(encode bool, data string) (string, error) {
	if encode {
		return url.QueryEscape(data), nil
	}

	decoded, err := url.QueryUnescape(data)
	if err != nil {
		return "", err
	}
	return decoded, nil
}

func HandleHTML(encode bool, data string) string {
	if encode {
		return html.EscapeString(data)
	}

	return html.UnescapeString(data)
}

func HandleHex(encode bool, data string) (string, error) {
	if encode {
		return hex.EncodeToString([]byte(data)), nil
	}

	decoded, err := hex.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func HandleBinary(encode bool, data string) (string, error) {
	if encode {
		res := ""
		for _, c := range data {
			res = fmt.Sprintf("%s%.8b", res, c)
		}
		return res, nil
	}

	re := regexp.MustCompile("[0|1]{8}")

	match := re.FindAllString(data, -1)
	b := make([]byte, len(match))
	for i, s := range match {
		n, err := strconv.ParseUint(s, 2, 8)
		if err != nil {
			return "", err
		}

		b[i] = byte(n)
	}
	return string(b), nil
}

func HandleSha1(value string) string {
	rValue := sha1.Sum([]byte(value))
	return fmt.Sprintf("%x", string(rValue[:]))
}

func HandleSha256(value string) string {
	rValue := sha256.Sum256([]byte(value))
	return fmt.Sprintf("%x", string(rValue[:]))
}

func HandleSha512(value string) string {
	rValue := sha512.Sum512([]byte(value))
	return fmt.Sprintf("%x", string(rValue[:]))
}

func HandleMD5(value string) string {
	rValue := md5.Sum([]byte(value))
	return fmt.Sprintf("%x", string(rValue[:]))
}

func HandleUpper(value string) string {
	return strings.ToUpper(value)
}

func HandleLower(value string) string {
	return strings.ToLower(value)
}

func HandleSnake(value string) string {
	return strcase.ToSnake(value)
}

func HandleCamel(value string) string {
	return strcase.ToCamel(value)
}

func HandleLowerCamel(value string) string {
	return strcase.ToLowerCamel(value)
}
