package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const hlsPath = "./hls_files"

func main() {
	// ===================== 디버깅 코드 추가 시작 =====================

	// 1. 현재 작업 디렉토리(CWD)를 확인합니다.
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("현재 작업 디렉토리를 가져올 수 없습니다: %v", err)
	}
	log.Printf("현재 작업 디렉토리(CWD): %s", cwd)

	// 2. 우리 코드가 사용하려는 hlsPath를 절대 경로로 변환해 봅니다.
	absHlsPath, err := filepath.Abs(hlsPath)
	if err != nil {
		log.Fatalf("hls_files 경로를 절대 경로로 변환하는데 실패했습니다: %v", err)
	}
	log.Printf("파일 서버가 찾으려는 루트 디렉토리(절대 경로): %s", absHlsPath)

	// 3. 해당 경로가 실제로 존재하는지 확인합니다.
	if _, err := os.Stat(absHlsPath); os.IsNotExist(err) {
		log.Fatalf("치명적 에러: 서버가 %s 디렉토리를 찾을 수 없습니다!", absHlsPath)
	}
	log.Printf("확인 완료: %s 디렉토리가 존재합니다.", absHlsPath)

	// ====================== 디버깅 코드 추가 끝 ======================

	fs := http.FileServer(http.Dir(hlsPath))
	http.Handle("/", loggingMiddleware(addCorsHeaders(fs)))

	log.Println("ABR HLS 서버를 8080 포트에서 시작")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func loggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("요청 시작: %s %s", r.Method, r.URL.Path)

		h.ServeHTTP(w, r)

		log.Printf("요청 완료: %s %s (%v)", r.Method, r.URL.Path, time.Since(start))
	})
}

func addCorsHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	})
}
