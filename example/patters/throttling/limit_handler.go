package throttling

import "net/http"

type LimitHandler struct {
	connections chan struct{} // 채널
	handler http.Handler
}

func NewLimitHandler(connections int, next http.Handler) *LimitHandler {
	cons := make(chan struct{}, connections) // 허용되는 연결수
	for i := 0; i < connections; i++ {
		cons <- struct{}{}
	}

	return &LimitHandler{
		connections: cons,
		handler: next,
	}
}

func (l *LimitHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	select {
	case <-l.connections:
		l.handler.ServeHTTP(rw, r) // 응답처리
		l.connections <- struct{}{} // 빈메세지 초기화
	default:
		http.Error(rw, "Busy", http.StatusTooManyRequests)
	} // 커넥션이 되지 않았으면  busy
}

