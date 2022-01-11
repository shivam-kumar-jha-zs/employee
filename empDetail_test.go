package main

import (
	"testing"
)

// TestSpeak is used to check the testcases
func TestFindEmpDetail(t *testing.T) {
	testcases := []struct {
		input           empDetails
		ExpectedOutput  empDetails
		ExpectedOutput1 bool
	}{
		{empDetails{"shivam", 23, "Golang"}, empDetails{"shivam", 23,
			"Golang"}, true},
		{empDetails{"sumit", 19, "java"}, empDetails{}, false},
	}
	for _, tc := range testcases {
		ActualOutput, boolValue := checkEmpDetails(tc.input)
		if tc.ExpectedOutput != ActualOutput {
			t.Errorf(" %v is expectedoutput %v is actucaloutput", tc.ExpectedOutput, ActualOutput)
		} else if boolValue != tc.ExpectedOutput1 {
			t.Errorf(" %v is expectedoutput %v is actucaloutput", tc.ExpectedOutput1, boolValue)
		}

	}
}
