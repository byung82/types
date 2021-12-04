package types

import "testing"

func TestSecureCodeSecureString(t *testing.T) {
	tests := []struct {
		name    string
		s       SecureCode
		want    string
		wantErr bool
	}{
		{
			"성공",
			SecureCode("123456"),
			"****",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Secure()
			if (err != nil) != tt.wantErr {
				t.Errorf("secureString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("secureString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
