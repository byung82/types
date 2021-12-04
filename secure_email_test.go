package types

import "testing"

func TestSecureEmailSecureString(t *testing.T) {
	tests := []struct {
		name    string
		s       SecureEmail
		want    string
		wantErr bool
	}{
		{
			"성공",
			SecureEmail("test@gmail.com"),
			"***@***.com",
			false,
		},
		{
			"실패",
			SecureEmail("test"),
			"",
			true,
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
