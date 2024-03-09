package util

import (
	"github.com/oldjon/gutil"
	"gitlab.com/morbackend/mor_services/mpb"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	com "gitlab.com/morbackend/mor_services/common"
)

func CheckEmailAddr(emailAddr string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$`)
	return regex.MatchString(emailAddr)
}

func GenerateRandomCode(length int) string {
	codeInt := rand.Intn(int(math.Pow(float64(10), float64(length))))
	code := strconv.Itoa(codeInt)
	for len(code) < com.VCodeLen {
		code = "0" + code
	}
	return code
}

func ReadUint32Slice(str string, separator string) []uint32 {
	if str == "" {
		return nil
	}
	out := make([]uint32, 0, 1)
	for _, v := range strings.Split(str, separator) {
		out = append(out, gutil.StrToUint32(v))
	}
	return out
}

const RealPlayerUserIdStart = 10000000

func IsBotUId(userId uint64) bool {
	return userId < RealPlayerUserIdStart
}

// ReadUint32Range if str can not be parsed as uint32 range, will return nil
func ReadUint32Range(str string, sep string) *mpb.Uint32Range {
	if len(str) == 0 {
		return nil
	}
	strs := strings.Split(str, sep)
	if len(strs) != 2 {
		return nil
	}
	min, err := strconv.Atoi(strs[0])
	if err != nil {
		return nil
	}
	max, err := strconv.Atoi(strs[1])
	if err != nil {
		return nil
	}
	return &mpb.Uint32Range{
		Min: uint32(min),
		Max: uint32(max),
	}
}

// ReadFloat64Range if str can not be parsed as uint32 range, will return nil
func ReadFloat64Range(str string, sep string) *mpb.Float64Range {
	if len(str) == 0 {
		return nil
	}
	strs := strings.Split(str, sep)
	if len(strs) != 2 {
		return nil
	}
	min, err := strconv.ParseFloat(strs[0], 64)
	if err != nil {
		return nil
	}
	max, err := strconv.ParseFloat(strs[1], 64)
	if err != nil {
		return nil
	}
	return &mpb.Float64Range{
		Min: min,
		Max: max,
	}
}

func RandomInFloat64Range(r *mpb.Float64Range) float64 {
	min := gutil.Min(r.Min, r.Max)
	max := gutil.Max(r.Min, r.Max)
	return rand.Float64()*(max-min) + min
}
