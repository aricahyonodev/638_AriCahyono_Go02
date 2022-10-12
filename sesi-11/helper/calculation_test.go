package helper

import (
	"fmt"
	"testing"
)

func TestFailSumWithFail(t *testing.T)  {
	result := Sum(10, 10)

	if result != 40 {
		t.Fail()
	}

	fmt.Println("TestFailSum Eksekusi Terhenti")
}
func TestSumWithFailNow(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		t.FailNow()
	}

	fmt.Println("TestSum Eksekusi Terhenti")
}

func TestFailSumWithError(t *testing.T)  {
	result := Sum(10, 10)

	if result != 40 {
		t.Error()
	}

	fmt.Println("TestFailSum Eksekusi Terhenti")
}
func TestSumWithFatal(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		t.Fatal()
	}

	fmt.Println("TestSum Eksekusi Terhenti")
}