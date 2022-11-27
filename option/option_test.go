package option_test

import (
	"go-like-rust/option"
	"strconv"
	"testing"
)

func TestConstructors(t *testing.T) {
	var myInt int = 10
	someInt := option.Some(myInt)

	if someInt.IsNone() {
		t.Fatal("Option must have a value")
	}

	someInt = option.None[int]()

	if someInt.IsSome() {
		t.Fatal("Option can't have a value")
	}
}

func TestUnwrap(t *testing.T) {
	var myInt int = 10
	someInt := option.Some(myInt)

	if someInt.Unwrap() != myInt {
		t.Fail()
	}

	someInt = option.None[int]()
	if someInt.UnwrapOr(0) != 0 {
		t.Fatal("UnwrapOr fallback value does not match")
	}

	if someInt.UnwrapOrElse(func() int { return 100 }) != 100 {
		t.Fatal("UnwrapOrElse return fallback function does not match")
	}
}

func TestMap(t *testing.T) {
	someString := option.Some("123")

	someInt := option.Map(&someString, func(value string) int {
		v, err := strconv.Atoi(value)
		if err != nil {
			return 0
		}

		return v
	})

	if someInt.Unwrap() != 123 {
		t.Fatal("Error at Map, values do not match")
	}

	someString = option.None[string]()
	someInt = option.MapOr(&someString, 200, func(value string) int {
		v, err := strconv.Atoi(value)
		if err != nil {
			return 0
		}

		return v
	})

	if someInt.Unwrap() != 200 {
		t.Fatal("Error at MapOr, values do not match")
	}

}
