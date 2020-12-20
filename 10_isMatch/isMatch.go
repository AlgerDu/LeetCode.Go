package leetcode

func isMatch(s string, p string) bool {

	sIndex := 0
	pKey, pValue, pIndex := getParttenKV(-1, p)
	repeatCount := 0
	unstablseCount := 0

	for pIndex < len(p) {

		v := s[sIndex]

		if sOffset >= len(s) {
			if pKey == '*' && pOffset == len(p) {
				return true
			}
			return false
		}

		switch pKey {
		case '*':
			if pValue != '.' && s[sOffset] != pValue {
				pKey, pValue, pOffset = getParttenKV(pOffset, p)
			} else {
				sOffset++
			}
		case '.':
			pKey, pValue, pOffset = getParttenKV(pOffset, p)
			sOffset++
		default:

			if lastV == pValue {
				pKey, pValue, pOffset = getParttenKV(pOffset, p)
				continue
			}

			if v != pValue {
				return false
			}

			pKey, pValue, pOffset = getParttenKV(pOffset, p)
			sOffset++

			lastV = v
			v = s[sOffset]
		}
	}

	if sOffset < len(s) {
		return false
	}

	return true
}

func getParttenKV(offset int, p string) (byte, byte, int) {

	offset++

	if offset >= len(p) {
		return 0, 0, offset
	}

	v := p[offset]

	if offset+1 < len(p) && p[offset+1] == '*' {
		return '*', v, offset + 1
	}

	return v, v, offset
}
