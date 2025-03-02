package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
*
sync.RWMutex
func (rw *RWMutex) Lock(), func (rw *RWMutex) Unlock(): 쓰기 뮤텍스 잠금, 잠금 해제
func (rw *RWMutex) RLock(), func (rw *RWMutex) RUnlock(): 읽기 뮤텍스 잠금 및 잠금 해제
*/
// this is not protected by mutex
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data int = 0

	go func() { // 값을 쓰는 고루틴
		for i := 0; i < 3; i++ {
			data += 1                         // data에 값 쓰기
			fmt.Println("write   : ", data)   // data 값을 출력
			time.Sleep(10 * time.Millisecond) // 10 밀리초 대기
		}
	}()

	go func() { // 값을 읽는 고루틴
		for i := 0; i < 3; i++ {
			fmt.Println("read 1 : ", data) // data 값을 출력(읽기)
			time.Sleep(1 * time.Second)    // 1초 대기
		}
	}()

	go func() { // 값을 읽는 고루틴
		for i := 0; i < 3; i++ {
			fmt.Println("read 2 : ", data) // data 값을 출력(읽기)
			time.Sleep(2 * time.Second)    // 2초 대기
		}
	}()

	time.Sleep(10 * time.Second) // 10초 동안 프로그램 실행
}
