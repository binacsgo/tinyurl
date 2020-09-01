package tinyurl

const (
	inf  int64 = 0x3FFFFFFF
	mask int64 = 0x0000003D
)

var alpha = []byte("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
