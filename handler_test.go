package lowlatencystreaming

import "testing"

func TestParseHLSPath(t *testing.T) {
	// 테스트 케이스들을 구조체로 정의하는 것은 일반적인 테스트 패턴 (Go)
	testCases := []struct {
		name         string
		path         string
		expectedType FileType
		expectedOk   bool
	}{
		{"재생목록 요청", "/live/stream.m3u8", PlaylistType, true},
		{"영상 조각 요청", "/live/stream.ts", SegmentType, true},
		{"잘못된 확장자", "/live/video.mp4", UnknownType, false},
		{"경로 없음", "/", UnknownType, false},
		{"숨겨진 파일", "/.DS_Store", UnknownType, false},
	}

	for _, tc := range testCases {
		// t.Run을 사용하면 각 테스트 케이스들을 독립적으로 실행하고 결과를 볼 수 있다.
		t.Run(tc.name, func(t *testing.T) {
			fileType, ok := parseHLSPath(tc.path)

			if fileType != tc.expectedType {
				t.Errorf("예상 타입: %v, 실제 타입: %v", tc.expectedType, fileType)
			}

			if ok != tc.expectedOk {
				t.Errorf("예상ㄴ 성공 여부: %v, 실제 성공 여부: %v", tc.expectedOk, ok)
			}
		})
	}
}
