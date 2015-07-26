package paxos

import "testing"

func TestBasicNetwork(t *testing.T) {
	nt := CreateNetwork(1, 3, 5, 2, 4)
	go func() {
		nt.recevFrom(5)
		nt.recevFrom(1)
		nt.recevFrom(3)
		nt.recevFrom(2)
		nt.recevFrom(4)
	}()

	m1 := message{from: 3, to: 1, typ: Prepare, seq: 1, preSeq: 0, val: "m1"}
	nt.sendTo(m1)
	m2 := message{from: 5, to: 3, typ: Accept, seq: 2, preSeq: 1, val: "m2"}
	nt.sendTo(m2)
	m3 := message{from: 4, to: 2, typ: Promise, seq: 3, preSeq: 2, val: "m3"}
	nt.sendTo(m3)
}
