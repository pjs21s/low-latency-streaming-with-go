package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const hlsPath = "./hls_files"

func main() {
	// HTTP 요청을 처리할 핸들러 함수를 등록
	http.HandleFunc("/", hlsHandler)

	log.Println("HLS 서버를 8080 포트에서 시작")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

func hlsHandler(w http.ResponseWriter, r *http.Request) {
	fileName := filepath.Base(r.URL.Path)

	fileType, ok := parseHLSPath(fileName)
	if !ok {
		http.NotFound(w, r)
		return
	}

	filePath := filepath.Join(hlsPath, fileName)

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
		log.Printf("파일 읽기 에러: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	switch fileType {
	case PlaylistType:
		w.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
	case SegmentType:
		w.Header().Set("Content-Type", "video/mp2t")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(fileBytes)
}
