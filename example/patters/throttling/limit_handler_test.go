package throttling

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newTestHandler(ctx context.Context) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		<-r.Context().Done() //요청 대기를 위한 컨텍스트 채널에 Done이 포함되기 까지 대기함, WithCancel 컨텍스트 첫번째 
	})
}

func setup(ctx context.Context) (*httptest.ResponseRecorder, *http.Request) {
	r := httptest.NewRequest("GET", "/health", nil)
	r = r.WithContext(ctx)
	return httptest.NewRecorder(), r
} // Test를 위한 Get 요청 생성

func TestReturnsBusyWhenConnectionsExhausted(t *testing.T) {
	ctx, cancel1 := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())
	handler := NewLimitHandler(1, newTestHandler(ctx)) // 신규 LimitHandler 생성하고 모의 핸들러를 전달
	rw, r := setup(ctx)

	time.AfterFunc(10*time.Millisecond, func() {
		cancel1()
		cancel2()
	}) // 10초 후 컨텍스트 대한 채널  cancel 메서드 호출

	waitGroup := sync.WaitGroup{}
	waitGroupAdd(2) // 채널 두개 생성

	go func() {
		handler.ServeHTTP(rw, r) // 응답처리
		waitGroup.Done() // 채널 실행 완료
	}() // 즉시 실행 함수

	go func() {
		handler.ServeHTTP(rw2, r2) // 응답처리
		waitGroup.Done() // 채널 실행 완료
	}()// 즉시 실행 함수

	waitGroup.Wait() // 채널이 다 완료될때까지 기다림 

	if rw.Code == http.StatusOK && rw2.Code == http.StatusOK {
		t.Fatalf("One request should have been busy, request 1: %v, request 2: %v", rw.Code, rw2.Code)
 	} //두개의 응답이 모두 OK이 일때 페이탈 에러 발생
}

