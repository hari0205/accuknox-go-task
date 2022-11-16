package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestDuplicateInArray(t *testing.T) {

	data := []string{"i1", "i3", "i2"}
	datalen := len(data)
	result := duplicateInArray(data)
	fmt.Println(result)
	if datalen != len(result) {
		t.Errorf("Function returned %q, expected %q", result, data)
	}

	data2 := []string{"i1", "i3", "i1"}
	res := duplicateInArray(data2)
	fmt.Println(res)
	if len(data) == len(res) {
		t.Error("Duplicates not returned or no duplicates were detected")
		t.Fail()
	} else {
		t.Log("Duplicates identified successfully")
	}

}

func TestMain(t *testing.T) {

	//Check if file is opening properly
	file, err := os.Open("testtext.txt")
	if err != nil {
		t.Log("Opening file failed")
	}

	buffer := bufio.NewScanner(file)
	if buffer == nil {
		t.Error("Buffer failed to initialize properly")
	}

	mockarr := []string{"i1", "i3", "i1"}
	result := duplicateInArray(mockarr)
	expected := []string{"i1", "i3"}
	if !reflect.DeepEqual(result, expected) {
		t.Error("failed to find duplicates")
		t.Fail()
	}

	// Mock customer log with no duplicates
	c1 := []string{"i1", "i3", "i2", "i4"}
	c2 := []string{"i2", "i1"}
	c3 := []string{"i1"}

	// Merging 3 customer records into a single slice
	c_res := make([]string, 0)
	c_res = append(c_res, c1...)
	c_res = append(c_res, c2...)
	c_res = append(c_res, c3...)

	// top := "i1"
	// second := "i2"
	// third := "i3"

	// Make map to store count
	dict := make(map[string]int)

	for _, num := range c_res {
		dict[num] = dict[num] + 1
	}
	// Checking against manual initlization
	if dict["i1"] != 3 && dict["i2"] != 2 && dict["i3"] != 1 {
		t.Error("Failed to find top 3 dishes")
	}
}
