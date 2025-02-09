package main

import "fmt"

func main() {
	// 부동소수점 값
	floatNum := 10.334

	// 1000을 곱해서 정수로 변환 (소수점 3자리까지 보존)
	intNum := int64(floatNum * 1000) // 10334

	// 결과를 다시 1000으로 나누어서 표시
	fixedNum := float64(intNum) / 1000

	fmt.Printf("원래 값(부동소수점): %v\n", floatNum)
	fmt.Printf("정수로 변환: %v\n", intNum)
	fmt.Printf("고정 소수점 결과: %v\n", fixedNum)
}
