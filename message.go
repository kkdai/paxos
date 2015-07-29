package paxos

type msgType int

const (
	Prepare msgType = iota + 1 // Send from proposer -> acceptor
	Promise                    // Send from acceptor -> proposer
	Propose                    // Send from proposer -> acceptor
	Accept                     // Send from acceptor -> learner
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
	return m.seq
}
