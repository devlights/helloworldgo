package helloworld

import (
	"fmt"
	"testing"
	"time"
)

func TestHello(t *testing.T) {
	// Arrange
	d := time.Date(2019, 10, 2, 11, 12, 13, 0, time.UTC)

	// Act
	actual, err := Hello(&d)

	// Assert
	if err != nil {
		t.Fatalf("failed. [err] %v\n", err)
	}

	want := "Hello World at 2019-10-02 11:12:13"
	if want != actual {
		t.Fatalf("failed. [want] %s\t[actual] %s\n", want, actual)
	}
}

func ExampleHello() {
	d := time.Date(2019, 10, 2, 11, 12, 13, 0, time.UTC)
	message, _ := Hello(&d)

	fmt.Println(message)
	// Output:
	// Hello World at 2019-10-02 11:12:13
}
