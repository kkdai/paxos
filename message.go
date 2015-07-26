package paxos

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

func (m *message) getSeqNumber() int {
	return m.seq
}

func (m *message) getProposeVal() string {
	return m.val
}

func (m *message) getProposeSeq() int {

	switch m.typ {
	case Promise:
		return m.preSeq
	case Accept:
		return m.seq
	default:
		panic("Don't support message typ")

	}
}
