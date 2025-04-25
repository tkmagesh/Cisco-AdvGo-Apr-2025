package utils

import (
	"testing"
	"time"
)

/*
TimeService Mock for Morning
*/
type MorningTimeService struct {
}

func (mts MorningTimeService) GetCurrent() time.Time {
	return time.Date(2025, time.April, 25, 10, 0, 0, 0, time.UTC)
}
func Test_Greeter_4_Morning(t *testing.T) {
	// Arrange
	userName := "Magesh"
	expected := "Hi Magesh, Good Morning!"
	timeServiceMock := MorningTimeService{}

	sut := NewGreeter(userName, timeServiceMock)
	// Act
	actual := sut.Greet()

	// Assert
	if expected != actual {
		t.Errorf("expected : %q, actual : %q\n", expected, actual)
	}
}

type AfterMorningTimeService struct {
}

func (mts AfterMorningTimeService) GetCurrent() time.Time {
	return time.Date(2025, time.April, 25, 15, 0, 0, 0, time.UTC)
}
func Test_Greeter_4_After_Morning(t *testing.T) {
	// Arrange
	userName := "Magesh"
	expected := "Hi Magesh, Good Day!"
	timeServiceMock := AfterMorningTimeService{}
	sut := NewGreeter(userName, timeServiceMock)
	// Act
	actual := sut.Greet()

	// Assert
	if expected != actual {
		t.Errorf("expected : %q, actual : %q\n", expected, actual)
	}
}
