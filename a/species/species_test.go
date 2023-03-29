package species

import "testing"

func TestQueryStandardId(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test",args{token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiZWxvbmdJZCI6IjM4MDY2MzQ1NTQ1NzM0OTYzMyIsImNvbXBhbnlJZCI6IjM4MDY2MzQ1NTQ1NzM0OTYzMyIsImV4cCI6MTY4MDA4NDU5MCwiaWF0IjoxNjgwMDc3MzkwLCJ1c2VySWQiOiIzNTg2MjY1Nzg2MTc0NzA5NzYifQ.JyLVDDCFs54zb3UNErOvvRkyINX36lmmlUXN6FfF5Bc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QueryStandardId(tt.args.token)
		})
	}
}
