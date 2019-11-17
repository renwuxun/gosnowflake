package gosnowflake

import "sync/atomic"

type ID32 struct {
	cnt  uint32
	gpid uint32
}

func NewID32(gpid uint32) *ID32 {
	gpid &= 0xff
	id32 := &ID32{
		cnt:  0,
		gpid: gpid,
	}
	return id32
}

func (s *ID32) Gen() uint32 {
	cnt := atomic.AddUint32(&s.cnt, 1)
	return (cnt << 8) | s.gpid
}
