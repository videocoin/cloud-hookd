package hookd

import (
	"fmt"
)

// Common ingest errors
var (
	ErrEmptyStream            = fmt.Errorf("stream is empty")
	ErrInvalidStream          = fmt.Errorf("invalid stream name")
	ErrInvalidWalletAddress   = fmt.Errorf("invalid user id")
	ErrInvalidContractAddress = fmt.Errorf("invalid contract address")
)

// StreamInfo used to parsing incoming rtmp stream
type StreamInfo struct {
	JobID string
}

// ParseStreamName parses stream info from rtmp url
func ParseStreamName(id string) (*StreamInfo, error) {
	if id == "" {
		return nil, ErrEmptyStream
	}

	return &StreamInfo{JobID: id}, nil
}
