package manipulator

import (
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"html"
	"net/url"
	"regexp"
	"strconv"
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
