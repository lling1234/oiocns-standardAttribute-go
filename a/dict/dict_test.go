package dict

import "testing"

func TestDictParse(t *testing.T) {
	tests := []struct {
		name string
	}{
		struct{name string}{""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s:=DictParse()
			t.Log(s)
			t.Log("------------------")
		})
	}
}
