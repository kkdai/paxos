package paxos

type msgType int

const (
	Prepare msgType = iota + 1
	Promise         //2
	Propose         //3
	Accept          //4
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
		return m.seq
	case Accept:
	case Propose:
	case Prepare:
		return m.seq
	default:
		//log.Println("Don't support message typ:", m.typ, " from:", m.from)
	}
	return 0
}
