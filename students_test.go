package coverage

import (
	"errors"
	"os"
	"strconv"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
func TestPeople_Len(t *testing.T) {
	var people People

	people = append(people, Person{})
	people = append(people, Person{})
	people = append(people, Person{})

	expect := 3
	got := people.Len()

	if expect != got {
		t.Errorf("Expected: %d, got %d", expect, got)
	}
}

func TestPeople_Less(t *testing.T) {
	var people People
	people = append(people, Person{})
	people = append(people, Person{})

	birthDay := time.Now()
	people[0].birthDay = birthDay
	people[1].birthDay = birthDay.Add(time.Minute * 1000)

	expect := false
	got := people.Less(0, 1)
	if expect != got {
		t.Errorf("Separate dates. Expected: %t, got %t", expect, got)
	}

	people[0].birthDay = birthDay
	people[1].birthDay = birthDay
	people[0].firstName = "xxx"
	people[1].firstName = "yyy"

	expect = true
	got = people.Less(0, 1)
	if expect != got {
		t.Errorf("Same dates. Separate fn. Expected: %t, got %t", expect, got)
	}

	people[0].firstName = "xxx"
	people[1].firstName = "xxx"
	people[0].lastName = "xxx"
	people[1].lastName = "yyy"

	expect = true
	got = people.Less(0, 1)
	if expect != got {
		t.Errorf("Same dates. Same fn. Separate sn. Expected: %t, got %t", expect, got)
	}
}

func TestPeople_Swap(t *testing.T) {
	var people People
	people = append(people, Person{})
	people = append(people, Person{})

	people[0].firstName = "0"
	people[1].firstName = "1"

	people.Swap(0, 1)

	expect := "1"
	got := people[0].firstName

	if expect != got {
		t.Errorf("Expected person name: %s, got %s", expect, got)
	}
}

func TestNew(t *testing.T) {
	_, err := New("1 2 3 4 5\n3 4 5 4")
	expect := "Rows need to be the same length"
	if err == nil {
		t.Errorf("Expected error: %s", expect)
	} else {
		got := err.Error()
		if expect != got {
			t.Errorf("Expected: %s, got: %s", expect, got)
		}
	}

	_, err = New("1 2 3 4 5\n1 2 3 4 a")
	if !errors.Is(err, strconv.ErrSyntax) {
		t.Errorf("Expected: %s, got: %T", strconv.ErrSyntax.Error(), err)
	}

	_, err = New("1 2 3 4 5")
	if err != nil {
		t.Errorf("Expected: nil, got %s", err.Error())
	}
}

func TestMatrix_Rows(t *testing.T) {
	m, _ := New("1 2 3 4 5\n10 20 30 40 50\n100 200 300 400 500")

	rows := m.Rows()

	if rows[0][0] != 1 || rows[1][0] != 10 || rows[2][4] != 500 {
		t.Errorf("Unexpected rows")
	}
}

func TestMatrix_Cols(t *testing.T) {
	m, _ := New("1 2 3 4 5\n10 20 30 40 50\n100 200 300 400 500")

	cols := m.Cols()

	if cols[0][0] != 1 || cols[1][0] != 2 || cols[4][1] != 50 {
		t.Errorf("Unexpected cols")
	}
}

func TestMatrix_Set(t *testing.T) {
	m, _ := New("1 2 3 4 5\n10 20 30 40 50\n100 200 300 400 500")

	ok := m.Set(0, 0, 666)

	if !ok || m.data[0] != 666 {
		t.Errorf("Matrix set error")
	}

	ok = m.Set(42324, 234, 666)
	if ok == true {
		t.Errorf("Unexpected error (rows and cols outrange)")
	}
}
