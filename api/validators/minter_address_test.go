package validators

import (
	"testing"
)

type tCase struct {
	Address string
	Valid   bool
}

func TestValidateNoahAddress(t *testing.T) {
	tCases := []tCase{
		{
			Address: "Noahxce542add0391b893d58c5fad21339f0f312cfa30",
			Valid:   true,
		},
		{
			Address: "NoahxCE542ADD0391B893D58C5FAD21339F0F312CFA30",
			Valid:   true,
		},
		{
			Address: "NoahxCE542ADD0391B893D58C5FAD21339F0F312CFA301",
			Valid:   false,
		},
		{
			Address: "NoahXCE542ADD0391B893D58C5FAD21339F0F312CFA30",
			Valid:   false,
		},
		{
			Address: "NoahxHE542ADD0391B893D58C5FAD21339F0F312CFA30",
			Valid:   false,
		},
		{
			Address: "Noahxce542add0391b893d58c5fad21339f0f312cfa3",
			Valid:   false,
		},
	}

	for _, c := range tCases {
		if isValidNoahAddress(c.Address) != c.Valid {
			t.Fatalf("Address validation failed. For %s expected %t, got %t", c.Address, c.Valid, !c.Valid)
		}
	}
}
