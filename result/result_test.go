package result_test

import (
	"go-like-rust/result"
	"testing"
)

type StrError struct {
	text string
}

func NewStrErr(text string) StrError {
	return StrError{text}
}

func (err StrError) Error() string {
	return err.text
}

func TestConstructors(t *testing.T) {
	var myInt int = 10
	intResult := result.Ok(myInt)

	if intResult.IsErr() {
		t.Fatal("Result can't be an error")
	}

	intResult = result.Err[int](NewStrErr("Error message"))

	if intResult.IsOk() {
		t.Fatal("Result can't be ok")
	}
}

func TestUnwrap(t *testing.T) {
	myString := "Hello"
	strResult := result.Ok(myString)

	if strResult.Unwrap() != myString {
		t.Fail()
	}

	strResult = result.Err[string](NewStrErr("Some error"))

	if strResult.UnwrapOr("Fallback") != "Fallback" {
		t.Fatal("UnwrapOr fallback value does not match")
	}

	if strResult.UnwrapOrElse(func() string { return "MyString" }) != "MyString" {
		t.Fatal("UnwrapOrElse return fallback function does not match")
	}
}
