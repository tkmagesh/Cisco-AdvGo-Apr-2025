package utils

import "testing"

func Test_IsPrime(t *testing.T) {
	// t.Skip()
	// Arrange
	no := 13
	expected := true

	// Act
	actual := IsPrime(no)

	// Assert
	// Considered to be "Passed" if not reported as "Failed"
	if actual != expected {
		/*
			t.Logf("Expected : %v, but actual : %v\n", expected, actual)
			t.Fail()
		*/
		t.Errorf("Expected : %v, but actual : %v\n", expected, actual)
	}
}

// Data driven tests
func Test_All_IsPrime(t *testing.T) {
	test_data := []struct {
		name     string
		no       int
		expected bool
	}{
		{name: "Test-13", no: 13, expected: true},
		// {name: "Test-15", no: 15, expected: false},
		{name: "Test-17", no: 17, expected: true},
		{name: "Test-19", no: 19, expected: true},
	}
	for _, td := range test_data {
		t.Run(td.name, func(t *testing.T) {
			actual := IsPrime(td.no)
			if actual != td.expected {
				t.Errorf("Expected : %v, but actual : %v\n", td.expected, actual)
			}
		})
	}

}
