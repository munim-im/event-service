package utils

import "github.com/speps/go-hashids/v2"

// hash generator
func GenerateHash(content int, hashLength int) string {
	hd := hashids.NewData()
	hd.Salt = GetEnvironmentVariable("HASHING_SECRET", "this-is-my-salt")
	hd.MinLength = hashLength
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([] int{content,})
	return e
}
