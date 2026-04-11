package utils

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}
	return string(ret), nil
}

func GenerateSlug(input string) string {
	input = strings.ToLower(strings.TrimSpace(input))

	var b strings.Builder
	lastDash := false

	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			b.WriteRune(r)
			lastDash = false
		} else if !lastDash {
			b.WriteRune('-')
			lastDash = true
		}
	}

	slug := strings.Trim(b.String(), "-")
	return slug
}

func GenerateTrackingNumber() *string {
	str, _ := GenerateRandomString(4)
	trackNum := "TRK-" + time.Now().Format("20060102") + strings.ToUpper(str)
	return &trackNum
}

func GetUintParam(c *gin.Context, name string) (uint, error) {
	idStr := c.Param(name)

	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, errors.New("invalid id")
	}

	return uint(id64), nil
}

func ToJSONB(v any) (datatypes.JSON, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return datatypes.JSON(b), nil
}

func FromJSONB(data datatypes.JSON, v any) error {
    err := json.Unmarshal([]byte(data), v)
    if err != nil {
        return err
    }
    return nil
}

func ParseETDToDays(etd string) int {
	// "1-2 days" → 2
	// "3 days" → 3

	re := regexp.MustCompile(`\d+`)
	nums := re.FindAllString(etd, -1)

	if len(nums) == 0 {
		return 0
	}

	max := 0
	for _, n := range nums {
		val, _ := strconv.Atoi(n)
		if val > max {
			max = val
		}
	}

	return max
}

// GenerateOTP generates a cryptographically secure numeric OTP of a given length.
func GenerateOTP(length int) (string, error) {
	const digits = "0123456789"
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		b[i] = digits[num.Int64()]
	}
	return string(b), nil
}
