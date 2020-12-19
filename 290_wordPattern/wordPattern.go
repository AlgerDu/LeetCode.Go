package leetcode

import "strings"

func wordPattern(pattern string, s string) bool {

	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	mapPV := make(map[byte]string)
	mapVP := make(map[string]byte)

	for offset := 0; offset < len(pattern) && offset < len(words); offset++ {
		desireV := words[offset]
		desireP := pattern[offset]

		p, okp := mapVP[desireV]

		if okp && p != desireP {
			return false
		}

		v, okv := mapPV[desireP]

		if okv && v != desireV {
			return false
		}

		if !okv && !okp {
			mapVP[desireV] = desireP
			mapPV[desireP] = desireV
		}
	}

	return true
}
