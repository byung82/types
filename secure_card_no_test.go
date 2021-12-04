package types

import "testing"

func TestSecureCardNumberSecure(t *testing.T) {
	tests := []struct {
		name string
		d    SecureCardNumber
		want string
	}{
		{
			"성공",
			SecureCardNumber("1111-2222-3333-4444"),
			"1111-22**-****-4444",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Secure(); got != tt.want {
				t.Errorf("Secure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecureCardNumberCardTxt(t *testing.T) {
	tests := []struct {
		name string
		d    SecureCardNumber
		want string
	}{
		{
			"성공",
			SecureCardNumber("1111-2222-3333-4444"),
			"4444",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.CardTxt(); got != tt.want {
				t.Errorf("CardTxt() = %v, want %v", got, tt.want)
			}
		})
	}
}
