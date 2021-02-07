package manipulator

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

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
