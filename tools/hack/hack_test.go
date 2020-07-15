package hack

import (
	"bytes"
	"testing"
)

func TestString(t *testing.T) {
	b := []byte("hello world")
	a := String(b)

	if a != "hello world" {
		t.Fatal(a)
	}

	b[0] = 'a'

	if a != "aello world" {
		t.Fatal(a)
	}

	b = append(b, "abc"...)
	if a != "aello world" {
		t.Fatalf("a:%v, b:%v", a, b)
	}
}

func TestByte(t *testing.T) {
	a := "hello world"

	b := Slice(a)

	if !bytes.Equal(b, []byte("hello world")) {
		t.Fatal(string(b))
	}
}

func TestMutable(t *testing.T) {
	a := []byte{'a', 'b', 'c'}
	b := String(a) // b is a mutable string.
	c := string(b) // Warn, c is a mutable string
	if c != "abc" {
		t.Fatalf("assert fail")
	}

	// c changed after a is modified
	a[0] = 's'
	if c != "sbc" {
		t.Fatal("test mutable string fail")
	}
}
