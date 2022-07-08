package benchmark

import (
	"fmt"
	"testing"
)

var stringArray []string = []string{
	"something1",
	"something2",
	"something3",
}

func InitBenchmark() {
	for i := 0; i < 100000000; i++ {
		stringArray = append(stringArray, "something2a" + fmt.Sprintf("%v", i))
	}
	
	fmt.Println("string array len: ", len(stringArray))
	normalFindRes := testing.Benchmark(BenchmarkNormalFind)
	explicitFindRes := testing.Benchmark(BenchmarkExplicitFind)

	fmt.Println("Normal: ", normalFindRes.T)
	fmt.Println("Explicit: ", explicitFindRes.T)
}

func BenchmarkNormalFind(t *testing.B) {
	something1 := findOne("something1")
	something2 := findOne("something2")
	something3 := findOne("something3")
	var something4 *string

	fmt.Println(something1, something2, something3, something4)
}

func BenchmarkExplicitFind(t *testing.B) {
	var something1 *string
	var something2 *string
	var something3 *string
	var something4 *string

	for _, stValue := range stringArray {
		if stValue == "something1" {
			something1 = &stValue
		} else if stValue == "something2" {
			something2 = &stValue
		} else if stValue == "something3" {
			something3 = &stValue
		} else if stValue == "something4" {
			something4 = &stValue
		}
	}

	fmt.Println(something1, something2, something3, something4)

}

func findOne(value string) *string {
	for _, stValue := range stringArray {
		if stValue == value {
			return &stValue
		}
	}

	return nil
}
