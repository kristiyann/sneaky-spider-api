package logic

import "testing"

func TestHashAndCheckPassword(t *testing.T) {
	password := "er34jie[u83855_fegji@"

	hashedPassword, _ := hashPassword(password)

	result := checkPasswordHash(password, hashedPassword)

	if !result {
		t.Errorf("TestHashAndCheckPassword: result should be true")
	}
}

func TestHashAndCheckPasswordFail(t *testing.T) {
	password := "er34jie[u83855_fegji@"

	hashedPassword, _ := hashPassword(password)

	password = "er34jie[u83855_fegji"

	result := checkPasswordHash(password, hashedPassword)

	if result {
		t.Errorf("TestHashAndCheckPassword: result should be false")
	}
}
