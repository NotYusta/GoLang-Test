package benchmark

import (
	"fmt"
	"testing"
)

var stringArray []string = []string{
	"something1",
}

func InitBenchmark() {
	DoFindLoopBenchmark()
	println()
}

func DoFindLoopBenchmark() {
	for i := 0; i < 10000; i++ {

		stringArray = append(stringArray, fmt.Sprintf("%v", i))
	}

	stringArray[5000] = "something2"
	stringArray[7999] = "something3"
	stringArray[9999] = "something4"
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

	result := []*string{
		something1,
		something2,
		something3,
		something4,
	}
	
	t.Log(result)
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

	result := []*string{
		something1,
		something2,
		something3,
		something4,
	}
	
	t.Log(result)
}

func findOne(value string) *string {
	for _, stValue := range stringArray {
		if stValue == value {
			return &stValue
		}
	}

	return nil
}
