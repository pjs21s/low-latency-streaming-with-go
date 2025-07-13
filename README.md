Simple HLS Streaming Server based on Go Language.

# 1. 폴더 구조

```
├── hls_files/ # HLS 세그먼트(.ts) 및 플레이리스트(.m3u8) 파일 디렉토리
├── .gitignore
├── README.md
├── go.mod
├── handler.go # HLS 경로 파싱 및 파일 서빙 핸들러
├── handler_test.go # parseHLSPath 단위 테스트
├── main.go # 서버 진입점
└── player.html # 간단한 HTML HLS 플레이어 예시
```

# 2. 사전 준비
```markdown
- Go 1.18+
- ngrok CLI
1. Download & Install ngrok from ngrok website
- macOS (Homebrew): `brew install -cask ngrok`
- Linux: extract zip after download package from https://ngrok.com/download
- Windows: add PATH after download offical zip
```

# 3. 설치 및 실행

의존성 설치
```bash
go mod download
```

서버 실행
```bash
go run main.go
```
기본 포트는 `:8080` 입니다.

ngrok tunneling
```bash
ngrok http 8080
```

change URL in player.html
```javascript
var videoSrc = 'https://a79d19f49963.ngrok-free.app/master.m3u8';
```

or just using 'localhost'
```javascript
var videoSrc = 'http://localhost:8080/master.m3u8';
```


# 4. 테스트
```bash
go test -v
```