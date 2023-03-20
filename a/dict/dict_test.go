package dict

import "testing"

func TestDictParse(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"111"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DictParse()
			t.Log("------------------")
			t.Log(s)
			t.Log("------------------")
			// for _, v := range s {
			// 	t.Log("``````````````````````")
			// 	t.Log(v)
			// 	t.Log("``````````````````````")

			// }
			ExcelWrite(s)

		})
	}
}
