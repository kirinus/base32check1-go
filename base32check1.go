package base32check1

import (
	"strings"
)

const (
	alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	cardinal = 1 << 5
)

var (
	alphabetList    = strings.Split(alphabet, "")
	alphabetMap     = make(map[string]int)
	primitive       = []int{0x01, 0x11, 0x08, 0x05, 0x03}
	primitivePowers = [cardinal - 1][]int{{}}
)

func init() {
	for i, v := range alphabetList {
		alphabetMap[v] = i
	}
	primitivePowers[1] = primitive
	for i := 2; i <= len(primitivePowers); i++ {
		values := matMul(primitivePowers[i-1], primitive)
		if i < len(primitivePowers) {
			primitivePowers[i] = values
		} else {
			primitivePowers[0] = values
		}
	}
}

func matMul(a []int, b []int) []int {
	mat := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		mat[i] = 0
		for j := 0; j < len(b); j++ {
			if (a[i] & (1 << (len(b) - 1 - j))) != 0 {
				mat[i] ^= b[j]
			}
		}
	}
	return mat
}

// Compute calculates the base32check1 checksum from the given payload
func Compute(payload string) string {
	payloadList := strings.Split(payload, "")
	payloadLength := len(payload)
	sum := 0
	for i := 0; i < payloadLength; i++ {
		value := alphabetMap[payloadList[i]]
		sum ^= matMul([]int{value}, primitivePowers[(i+1)%(cardinal-1)])[0]
	}

	exp := (cardinal - payloadLength - 2) % (cardinal - 1)
	if exp < 0 {
		exp += cardinal - 1
	}
	return alphabetList[matMul([]int{sum}, primitivePowers[exp])[0]]
}

// Validate checks if the given base32check1 payload has a valid checksum
func Validate(payload string) bool {
	return Compute(payload) == "A"
}
