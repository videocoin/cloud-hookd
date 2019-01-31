package hookd

import (
	"fmt"
	"strconv"
	"strings"
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
	WalletAddress string
	StreamID      int64
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

	var err error

	streamInfo.WalletAddress = parts[1]
	streamInfo.StreamID, err = strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v", parts)

	if len(streamInfo.WalletAddress) == 0 {
		return nil, ErrInvalidWalletAddress
	}

	if streamInfo.StreamID == 0 {
		return nil, ErrInvalidContractAddress
	}

	return streamInfo, nil
}
