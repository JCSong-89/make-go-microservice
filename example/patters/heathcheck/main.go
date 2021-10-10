package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/VividCortex/ewma"
)

var ma = ewma.MovingAverage
var threshold = 1000 * time.Millisecond
var timeout = 1000 * time.Microsecond
var resetting = false
var resetMutex = sync.RWMutex{}

func main(){
	ma = ewma.NewMovingAverage()

	http.HandleFunc("/", mainHanlder) //메인 서비스 경로
	http.HandleFunc("/health", healthHandler) // 상태점검 경로

	http.ListenAndServe(":8000", nil) 
}

func mainHanlder(rw http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	if !isHealthy() {
		respondServiceUnhealthy(rw)
		return
	} // 지정된 임계치보다 높은 비정상 상태라면 상태불량 리턴

	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Average request time: %f (ms)\n", ma.Value()/1000000)

	duration := time.Now().Sub(startTime)
	ma.Add(float64(duration))
}

func respondServiceUnhealthy(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusServiceUnavailable) // 서비스 불가 헤더 추가

	resetMutex.RLock() // 서비스 복구되는 동안 대기하고 평균 재설정을 위한 잠ㄱㄱ,ㅁ
	defer resetMutex.Unlock() // 함수 실행 끝에 락 풀음

	if !resetting {
		go sleepAndResetAverage() // 고루틴 생성 
	}
}

func sleepAndResetAverage() {
	resetMutex.Lock() // 잠금
	resetting = true // fleg 변경 -> 복구하는 동안 레이스 컨디션을 피하기 위해 위에서 잠금
	resetMutex.Unlock() // 잠금 해제

	time.Sleep(timeout) // 복구 대기 시작 
	ma = ewma.NewMovingAverage() 

	resetMutex.Lock()
	resetting = false // 복구완료로 fleg 변경
	resetMutex.Unlock()
}

func healthHandler(rw http.ResponseWriter, r *http.Request) {
	if !isHealthy() {
		rw.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	fmt.Fprint(rw, "OK")
}

func isHealthy() bool {
	return (ma.Value() < float64(threshold)) // 인계값이 넘 으면 false
}


