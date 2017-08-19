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

func Test_intToBase58(t *testing.T) {
	cases := []int64{
		1,
		15,
		16,
		100,
		10011,
		12345678,
		912345678,
		938182374237,
	}
	for _, c := range cases {
		t.Logf("%d -> %s", c, intToBase58(c))
	}
}

func Test_base58ToInt(t *testing.T) {
	cases := []string{
		"2",
		"g",
		"h",
		"2J",
		"3YB",
		"26gWw",
		"2oC1zU",
		"qDnXU5v",
	}
	for _, c := range cases {
		i, err := base58ToInt(c)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%s -> %d", c, i)
	}
}

func Test_xor(t *testing.T) {
	cases := []int64{
		1,
		15,
		16,
		100,
		10011,
		12345678,
		912345678,
		938182374237,
	}
	for _, c := range cases {
		a := xor(c)
		t.Logf("%d -> %d", c, a)
		b := xor(a)
		t.Logf("%d -> %d", c, b)
	}
}
