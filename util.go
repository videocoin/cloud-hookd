package hookd

import (
	"fmt"
	"strings"
)

// Common ingest errors
var (
	ErrEmptyStream            = fmt.Errorf("stream is empty")
	ErrInvalidStream          = fmt.Errorf("invalid stream name")
	ErrInvalidUserID          = fmt.Errorf("invalid user id")
	ErrInvalidContractAddress = fmt.Errorf("invalid contract address")
)

// StreamInfo used to parsing incoming rtmp stream
type StreamInfo struct {
	UserID   string
	StreamID string
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
	streamInfo.StreamID = parts[1]

	fmt.Printf("%+v", parts)

	if len(streamInfo.UserID) == 0 {
		return nil, ErrInvalidUserID
	}

	if len(streamInfo.StreamID) == 0 {
		return nil, ErrInvalidContractAddress
	}

	return streamInfo, nil
}
