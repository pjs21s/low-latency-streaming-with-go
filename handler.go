package main

import "strings"

// HLS 요청 종류를 나타내는 상수 정의
const (
	UnknownType FileType = iota
	PlaylistType
	SegmentType
)

type FileType int

func parseHLSPath(path string) (FileType, bool) {
	if strings.HasSuffix(path, ".m3u8") {
		return PlaylistType, true
	}

	if strings.HasSuffix(path, "ts") {
		return SegmentType, true
	}

	return UnknownType, false
}
