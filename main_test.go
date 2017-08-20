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
		b := xor(a)
		t.Logf("%d -> %d -> %d", c, a, b)
	}
}

func Test_saltxor(t *testing.T) {
	var key int64 = 19850818
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
		a := saltxor(c, key)
		b := saltxor(a, key)
		t.Logf("%d -> %d -> %d", c, a, b)
	}
}

func Test_saltxorBase58(t *testing.T) {
	var key int64 = 19850818
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
		a := saltxor(c, key)
		ah := intToBase58(a)
		bh, err := base58ToInt(ah)
		if err != nil {
			t.Fatal(err)
		}
		b := saltxor(bh, key)
		t.Logf("%d -> %d -> %s -> %d -> %d", c, a, ah, bh, b)
	}
}

func Test_hashid(t *testing.T) {
	hdata := NewData()
	hdata.MinLength = 8
	hdata.Salt = "this is my sacred key"
	encoder, err := NewWithData(hdata)
	if err != nil {
		t.Fatal(err)
	}

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
		s, err := encoder.EncodeInt64([]int64{c})
		if err != nil {
			t.Fatal(err)
		}
		ss, err := encoder.DecodeInt64WithError(s)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%d -> %s -> %d", c, s, ss)
	}
}
