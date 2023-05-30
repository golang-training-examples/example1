package hello_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/golang-training-examples/example1/pkg/hello"
)

func TestSayHello(t *testing.T) {
	got := hello.SayHello("Dela")
	want := "Hello Dela!"
	if got != want {
		t.Errorf(`Hello("Dela") = "%s"; want "%s"`, got, want)
	}
}

func TestSayHello1000(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	for i := 0; i < 1000; i++ {
		got := hello.SayHello("Dela" + strconv.Itoa(i))
		want := "Hello Dela" + strconv.Itoa(i) + "!"
		if got != want {
			t.Errorf(`Hello("Dela") = "%s"; want "%s"`, got, want)
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func TestPrintSayHello(t *testing.T) {
	hello.PrintSayHello("Dela")
}
