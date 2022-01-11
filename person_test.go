package main

import (
	"testing"
)

type Output struct {
	out    bool
	person Person
}

func TestCheckAge(t *testing.T) {
	testCase := []struct {
		desc   string
		input  Person
		output Output
	}{
		{desc: "error case", input: Person{"person1", 21},
			output: Output{}},
		{desc: "success", input: Person{"person2", 24},
			output: Output{true,
				Person{"person2", 24}}},
	}

	for i, tc := range testCase {
		ok, person := CheckAge(tc.input)

		if ok != tc.output.out {
			t.Errorf("index %v %v failed expected %v got %v", i, tc.desc, tc.output, ok)
		}
		if person != tc.output.person {
			t.Errorf("index %v %v failed expected %v got %v", i, tc.desc, tc.output, ok)
		}
	}
}

func BenchmarkCheckAge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckAge(Person{name: "Ram", age: 21})
	}
}
