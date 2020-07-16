package generator

import "testing"

func TestListGenerator(t *testing.T) {
	list := []byte("abcde")
	listmap := map[byte]bool{}
	for _, v := range list {
		listmap[v] = true
	}
	g := &ListGenerator{
		List: list,
		Min:  256,
	}
	s, err := g.Generate()
	if err != nil {
		panic(err)
	}
	if len(s) != 256 {
		t.Fatal(s)
	}
	for _, v := range s {
		if listmap[v] == false {
			t.Fatal(s)
		}
	}
	g = &ListGenerator{
		List: list,
		Min:  15,
		Max:  256,
	}
	s, err = g.Generate()
	if err != nil {
		panic(err)
	}
	if len(s) > 256 || len(s) < 15 {
		t.Fatal(s)
	}

}

func TestBytesGenerator(t *testing.T) {
	secret, err := BytesGenerator(15).Generate()
	if len(secret) != 15 || err != nil {
		t.Fatal(secret, err)
	}
}
