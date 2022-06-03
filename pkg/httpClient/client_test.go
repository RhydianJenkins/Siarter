package httpClient

import (
	"testing"
)

func TestMockResponse(t *testing.T) {
	client := Client{
		"www.doesnt.matter.com/api",
		true,
	}

	expected := getMockResponse()
	got, err := client.Get()

	if err != nil {
		t.Errorf("Error while getting from client: %v", err)
	}

	if !slicesAreEqual(expected, got) {
		t.Errorf("Expected response not equal to mock")
	}
}

func slicesAreEqual(sliceA, sliceB []*Boat) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}

	for i, boat := range sliceA {
		if boat.LAT != sliceA[i].LAT {
			return false
		}
	}

	return true
}
