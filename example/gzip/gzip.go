package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var name string

type helloWorldResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	http.Handle("/helloworld",
		NewGzipHandler(http.HandlerFunc(helloWorldHandler)),
	)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(rw http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello " + name}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

func NewGzipHandler(next http.Handler) http.Handler {
	return &GZipHandler{next}
}

type GZipHandler struct {
	next http.Handler 
	//압축이 완료되거나, 압축이 필요없어 리턴받은 핸들러로 http.Handle에 전달
}

func (h *GZipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encodings := r.Header.Get("Accept-Encoding")

	if strings.Contains(encodings, "gzip") {
		h.serveGzipped(w, r) //gizp 압축
	} else if strings.Contains(encodings, "deflate") {
		panic("Deflate not implemented") //deflate시 패닉
	} else {
		h.servePlain(w, r) //gizp이 아닐시 그냥 리턴
	}
}

func (h *GZipHandler) serveGzipped(w http.ResponseWriter, r *http.Request) {
	gzw := gzip.NewWriter(w) //gzip 객체갱성
	defer gzw.Close() // 반환

	w.Header().Set("Content-Encoding", "gzip") //gzip 헤더 추가
	h.next.ServeHTTP(GzipResponseWriter{gzw, w}, r) //압축시작
}

func (h *GZipHandler) servePlain(w http.ResponseWriter, r *http.Request) {
	h.next.ServeHTTP(w, r) //입력받은 핸들러 함수 콜 여기에서는 헬로우월드핸들러
}

type GzipResponseWriter struct {
	gw *gzip.Writer
	http.ResponseWriter
}

func (w GzipResponseWriter) Write(b []byte) (int, error) {
	if _, ok := w.Header()["Content-Type"]; !ok {
		// 콘텐츠  타입이 없을 시 Content-Type 추가
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return w.gw.Write(b)  // 압축 후 리턴
}

func (w GzipResponseWriter) Flush() {
	w.gw.Flush()
	if fw, ok := w.ResponseWriter.(http.Flusher); ok {
		fw.Flush()
	}
}