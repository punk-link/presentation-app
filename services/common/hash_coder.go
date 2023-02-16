package common

type HashCoder interface {
	Decode(target string) int
	Encode(target int) string
}
