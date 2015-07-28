package paxos

import (
	"log"
	"testing"
	"time"
)

func TestBasicNetwork(t *testing.T) {
	log.Println("TestBasicNetowk........................")
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

	m4 := message{from: 4, to: 2, typ: Promise, seq: 3, preSeq: 2, val: "m4"}
	nt.sendTo(m4)
}

func TestProserFunction(t *testing.T) {
	log.Println("TestProserFunction........................")
	//Three acceptor and one proposer
	network := CreateNetwork(100, 1, 2, 3)

	//Create acceptors
	var acceptors []acceptor
	aId := 1
	for aId <= 3 {
		acctor := NewAcceptor(aId, network.getNodeNetwork(aId), 0)
		acceptors = append(acceptors, acctor)
		aId++
	}

	//Create proposer
	proposer := NewProposer(100, "value1", network.getNodeNetwork(100), 1, 2, 3)

	//Run proposer and acceptors
	go proposer.run()

	for index, _ := range acceptors {
		go acceptors[index].run()
	}
	time.Sleep(10 * time.Second)
}
