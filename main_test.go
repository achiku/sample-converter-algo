package main

import "testing"

func Test_intToHex(t *testing.T) {
	cases := []int64{
		1,
		15,
		16,
		12345678,
		912345678,
		938182374237,
	}
	for _, c := range cases {
		t.Logf("%d -> %s", c, intToHex(c))
	}
}

func Test_hexToInt(t *testing.T) {
	cases := []string{
		"1",
		"f",
		"10",
		"bc614e",
		"36614a4e",
		"da7006f35d",
	}
	for _, c := range cases {
		i, err := hexToInt(c)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%s -> %d", c, i)
	}
}
