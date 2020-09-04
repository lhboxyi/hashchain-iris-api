package util

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	PREFIX      = "{"
	SUFFIX      = "}"
	keyLength   = 32
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)



func Matches(rawPassword, encodedPassword string) bool {
	salt := extractSalt(encodedPassword)
	rawPasswordEncoded := digest(salt, rawPassword)
	if rawPasswordEncoded == encodedPassword {
		return true
	}
	return false
}

func Encode(rawPassword string) string {
	salf := fmt.Sprintf("%s%s%s", PREFIX, generateKey(), SUFFIX)
	return digest(salf, rawPassword)
}

func generateKey() string {
	randomString := getRandomString(keyLength)
	var coder = base64.NewEncoding(base64Table)
	base64EncodedKey := coder.EncodeToString([]byte(randomString))
	return base64EncodedKey
}
func extractSalt(string2 string) string {
	i := strings.Index(string2, "}")
	salt := string2[0 : i+1]
	return fmt.Sprintf("%s", salt)
}

func digest(salt, rawPassword string) string {
	saltedPassword := fmt.Sprintf("%s%s", rawPassword, salt)
	s_ob := sha512.New()
	s_ob.Write([]byte(saltedPassword))
	r := s_ob.Sum(nil)
	encodedStr := hex.EncodeToString(r)
	return fmt.Sprintf("%s%s", salt, encodedStr)
}

func getRandomString(l int) string {
	bytes := []byte(base64Table)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
