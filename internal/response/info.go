package response

import "time"

type GetInfo struct {
	Commit  string
	StartAt time.Time
}

type CacheInfo struct {
	TotalHit   int
	TotalStore int
}

type NetworkInfo struct {
	PacketSent uint64
	PacketRecv uint64
	BytesSent  uint64
	BytesRecv  uint64
}
