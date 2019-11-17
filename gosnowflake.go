package gosnowflake

import (
	"sync/atomic"
	"time"
)

const delta = 1538006400000

// Snowflake
// ttttttttttttttttttttttttttttttttttttttttttccccccccccccmmmmmmmmmm
// m:8,gpid t:42,time ms c:14,time ms counter
type Snowflake struct {
	statems uint64
	gpid    uint64
}

func nowms() uint64 {
	return uint64(time.Now().UnixNano()/1e6) - delta
}

// New create uint64 id generator
func New(gpid uint64) *Snowflake {
	gpid &= 0xff
	sf := &Snowflake{
		statems: 0,
		gpid:    gpid,
	}
	return sf
}

// Gen generate uint64 id
func (s *Snowflake) Gen() uint64 {
	if oldstatems, newstatems := s.statems, nowms()<<14; oldstatems < newstatems {
		atomic.CompareAndSwapUint64(&s.statems, oldstatems, newstatems-1)
	}

	// 必须使用atomic的返回值，s.statems随时被其他线程改变
	statems := atomic.AddUint64(&s.statems, 1)

	// 把gpid放到最后的8个位，利于排序
	return (statems << 8) | s.gpid
}
