package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	done := make(chan int, 2) // 버퍼가 2개인 비동기 채널 생성
	count := 4                // 반복할 횟수

	go func() {
		for i := 0; i < count; i++ {

			fmt.Println("고루틴 : ", i) // 반복문의 변수 출력
			done <- i
		}
	}()

	for i := 0; i < count; i++ {
		fmt.Println("메인 함수 : ", <-done) // 반복문의 변수 출력
	}
	//fmt.Scanln()
}
