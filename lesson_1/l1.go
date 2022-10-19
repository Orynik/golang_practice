package main

import "fmt"

func GenerateCheck() (result []HealthCheck) {
	for index := 1; index <= 5; index++ {
		var tempStatus string = FailStatus

		if index%2 == 0 {
			tempStatus = PassStatus
		}

		var structTemp = HealthCheck{
			ServiceId: index,
			Pass:      tempStatus,
		}

		result = append(result, structTemp)
	}

	return
}

const PassStatus string = "pass"
const FailStatus string = "fail"

type HealthCheck struct {
	ServiceId int
	Pass      string
}

func main() {
	var resultCheck []HealthCheck = GenerateCheck()

	for index := 0; index < len(resultCheck); index++ {
		if resultCheck[index].Pass == PassStatus {
			fmt.Println(resultCheck[index].ServiceId)
		}
	}
}
