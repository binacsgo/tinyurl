package tinyurl

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

var gMD5Impl *MD5Impl

// GetMD5Impl get the gMD5Impl
func GetMD5Impl() *MD5Impl {
	if gMD5Impl == nil {
		gMD5Impl = NewMD5Impl("md5")
	}
	return gMD5Impl
}

// MD5Impl implement the TinyURL by md5 algo.
type MD5Impl struct {
	name string
}

// NewMD5Impl return a pointer to MD5Impl
func NewMD5Impl(name string) *MD5Impl {
	return &MD5Impl{name: name}
}

// Name return the name of MD5Impl
func (impl *MD5Impl) Name() string {
	return impl.name
}

// EncodeURL return the tiny-url string by md5 algo.
func (impl *MD5Impl) EncodeURL(url string) string {
	// 1. get md5 string
	hash := md5.New()
	hash.Write([]byte(url))
	str := hex.EncodeToString(hash.Sum(nil))
	// 2. generate 4 string
	var res [4]string
	for i := 0; i < 4; i++ {
		hexv, _ := strconv.ParseInt(str[i*8:(i+1)*8], 16, 64)
		tmpv := inf & hexv
		turl := []byte{}
		for i := 0; i < 6; i++ {
			turl = append(turl, alpha[mask&tmpv])
			tmpv = tmpv >> 5
		}
		res[i] = string(turl)
	}
	// 3. return the first one. rand() is better.
	return res[0]
}
