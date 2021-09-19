package base32check1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	payload string
	check   string
}{
	{"", "A"},
	{"A", "A"},
	{"AB", "Q"},
	{"ABC", "J"},
	{"ABCD", "V"},
	{"ABCDE", "I"},
	{"ABCDEF", "G"},
	{"ABCDEFG", "A"},
	{"ABCDEFGH", "T"},
	{"ABCDEFGHI", "5"},
	{"ABCDEFGHIJ", "K"},
	{"ABCDEFGHIJK", "A"},
	{"ABCDEFGHIJKL", "F"},
	{"ABCDEFGHIJKLM", "U"},
	{"ABCDEFGHIJKLMN", "M"},
	{"ABCDEFGHIJKLMNO", "R"},
	{"ABCDEFGHIJKLMNOP", "7"},
	{"ABCDEFGHIJKLMNOPQ", "X"},
	{"ABCDEFGHIJKLMNOPQR", "D"},
	{"ABCDEFGHIJKLMNOPQRS", "I"},
	{"ABCDEFGHIJKLMNOPQRST", "5"},
	{"ABCDEFGHIJKLMNOPQRSTU", "U"},
	{"ABCDEFGHIJKLMNOPQRSTUV", "Q"},
	{"ABCDEFGHIJKLMNOPQRSTUVW", "D"},
	{"ABCDEFGHIJKLMNOPQRSTUVWX", "K"},
	{"ABCDEFGHIJKLMNOPQRSTUVWXY", "J"},
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "Y"},
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZ2", "R"},
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZ23", "V"},
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZ234", "U"},
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZ2345", "U"},
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZ23456", "V"},
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZ234567", "V"},
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZ234567ABCDEFGHIJKLMNOPQRSTUVWXYZ2345", "6"},
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZ234567ABCDEFGHIJKLMNOPQRSTUVWXYZ234567ABCDEFGHIJKLMNOPQRSTUVWXYZ234", "K"},
	{"CONSECRATIO", "X"},
	{"CAFEBABE", "N"},
	{"CAFEDEAD", "A"},
	{"DEADBEEF", "L"},
	{"234567", "Z"},
}

func TestCompute(t *testing.T) {
	t.Parallel()
	for _, tt := range cases {
		tt := tt
		t.Run(fmt.Sprintf("Compute %s", tt.payload), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.check, Compute(tt.payload))
		})
	}
}

func TestValidate_AcceptValid(t *testing.T) {
	t.Parallel()
	for _, tt := range cases {
		tt := tt
		t.Run(fmt.Sprintf("Accept Checksum %s", tt.payload), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, true, Validate(tt.payload+tt.check))
		})
	}
}

func TestValidate_RejectInvalid(t *testing.T) {
	t.Parallel()
	for _, tt := range cases {
		tt := tt
		if tt.check != "A" {
			t.Run(fmt.Sprintf("Reject Checksum %s", tt.payload), func(t *testing.T) {
				t.Parallel()
				assert.Equal(t, false, Validate(tt.payload))
			})
		}
	}
}
