package utils

import "testing"

func TestHashPassword(t *testing.T) {
	data := []struct {
		password string
	}{
		{password: "123456"},
		{password: "hello world"},
		{password: ""},
	}

	for i, datum := range data {
		result, err := HashPassword(datum.password)
		if err != nil {
			t.Error(err)
		}
		if result == "" {
			t.Errorf("#%d: got %s", i, result)
		}
	}
}

func TestVerifyPassword(t *testing.T) {

	data := []struct {
		password string
		hash     string
		expected bool
	}{
		{
			password: "123456",
			hash:     "$2a$10$dlNzzonmjHsbi8sjBVRXSOiYo46Z6juG.Yn61bfyV3x/efluQ3Qpy",
			expected: true,
		},
		{
			password: "123456",
			hash:     "$2a$10$WaKKtJw4cwwVGjvwg0vWIeJzOTSLUYdHbqLa9cbqr0CqpXaFMv7EG",
			expected: false,
		},
	}

	for i, datum := range data {
		result := VerifyPassword(datum.password, datum.hash)
		if result != datum.expected {
			t.Errorf("#%d: expected %v, got %v", i, datum.expected, result)
		}
	}
}
