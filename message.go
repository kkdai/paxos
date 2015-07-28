package paxos

import "log"

type msgType int

const (
	Prepare msgType = iota + 1
	Propose
	Promise
	Accept
)

type message struct {
	from   int
	to     int
	typ    msgType
	seq    int
	preSeq int
	val    string
}

func (m *message) getProposeVal() string {
	return m.val
}

func (m *message) getProposeSeq() int {

	switch m.typ {
	case Promise:
		return m.preSeq
	case Accept:
	case Propose:
	case Prepare:
		return m.seq
	default:
		log.Println("Don't support message typ:", m.typ, " from:", m.from)
		//panic("Don't support message typ:", m.typ, " from:", m.from)

	}
	return 0
}
