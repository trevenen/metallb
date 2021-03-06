// +build gofuzz

package ndp

func Fuzz(data []byte) int {
	m, err := ParseMessage(data)
	if err != nil {
		return 0
	}

	b1, err := m.MarshalBinary()
	if err != nil {
		panic(err)
	}

	if err := m.UnmarshalBinary(b1); err != nil {
		panic(err)
	}

	b2, err := MarshalMessage(m)
	if err != nil {
		panic(err)
	}

	if _, err := ParseMessage(b2); err != nil {
		panic(err)
	}

	return 1
}
