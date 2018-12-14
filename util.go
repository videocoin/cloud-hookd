package hookd

import (
	"fmt"
	"strings"
)

// Common ingest errors
var (
	ErrEmptyStream     = fmt.Errorf("stream is empty")
	ErrInvalidStream   = fmt.Errorf("invalid stream name")
	ErrInvalidUserID   = fmt.Errorf("invalid user id")
	ErrInvalidCameraID = fmt.Errorf("invalid camera id")
)

// StreamInfo used to parsing incoming rtmp stream
type StreamInfo struct {
	UserID   string
	CameraID string
}

// ParseStreamName parses stream info from rtmp url
func ParseStreamName(name string) (*StreamInfo, error) {
	if name == "" {
		return nil, ErrEmptyStream
	}

	parts := strings.Split(name, "-")
	if len(parts) != 2 {
		return nil, ErrInvalidStream
	}

	streamInfo := new(StreamInfo)

	streamInfo.UserID = parts[0]
	streamInfo.CameraID = parts[1]

	fmt.Printf("%+v", parts)

	if len(streamInfo.UserID) == 0 {
		return nil, ErrInvalidUserID
	}

	if len(streamInfo.CameraID) == 0 {
		return nil, ErrInvalidCameraID
	}

	return streamInfo, nil
}
