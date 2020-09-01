package tinyurl

// TinyURL interface def.
type TinyURL interface {
	Name() string
	EncodeURL(string) string
}

// Check
var _ TinyURL = &MD5Impl{}
