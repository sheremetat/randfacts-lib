package facts

import (
	"testing"
)

func TestWrongDir(t *testing.T) {
	_, err := Init("./facts_wrong")

	if err == nil {
		t.Error("Error should be not nil")
	}
}

func TestGetFactSuccess(t *testing.T) {
	lib, _ := Init("./facts")

	key := "go"
	fact, _ := lib.GetFact(key)
	if len(fact) == 0 {
		t.Error("Fact should be found")
	}

	key = "Go"
	fact, _ = lib.GetFact(key)
	if len(fact) == 0 {
		t.Error("Fact should be found")
	}
}

func TestFindFactSuccess(t *testing.T) {
	lib, _ := Init("./facts")

	text := "go is awesome"
	fact, _ := lib.FindFact(text)
	if len(fact) == 0 {
		t.Error("Fact should be found")
	}

	text = "Go is awesome"
	fact, _ = lib.FindFact(text)
	if len(fact) == 0 {
		t.Error("Fact should be found")
	}

}

func TestGetFactFail(t *testing.T) {
	lib, _ := Init("./facts")

	key := "test"
	fact, _ := lib.GetFact(key)
	if len(fact) > 0 {
		t.Error("Fact should not be found")
	}
}

func TestFindFactFail(t *testing.T) {
	lib, _ := Init("./facts")

	text := "test is awesome"
	fact, _ := lib.FindFact(text)
	if len(fact) > 0 {
		t.Error("Fact should not be found")
	}
}
