package paxos

import "log"

func NewAcceptor(id int, nt nodeNetwork, learners ...int) *acceptor {
	newAccptor := &acceptor{id: id, nt: nt}
	newAccptor.learners = learners
	return newAccptor
}

type acceptor struct {
	id         int
	learners   []int
	acceptMsg  message
	promiseMsg message
	nt         nodeNetwork
}

//After acceptor receive prepare message.
//It will check  prepare number and return acceptor if it is bigest one.
func (a *acceptor) recevPrepare(prepare message) *message {
	if a.promiseMsg.getProposeSeq() >= prepare.getProposeSeq() {
		log.Println("ID:", a.id, "Already accept bigger one")
		return nil
	}
	log.Println("ID:", a.id, "Accept ")

	a.promiseMsg = message{typ: Promise,
		from: a.id,
		to:   prepare.from,
		seq:  prepare.seq}
	return &a.promiseMsg
}

//Recev Propose only check if acceptor already accept bigger propose before.
//Otherwise, will just forward this message out and change its type to "Accept" to learning later.
func (a *acceptor) recevPropose(proposeMsg message) *message {
	//Already accept bigger propose before
	if a.acceptMsg.getProposeSeq() > proposeMsg.getProposeSeq() {
		log.Println("ID:", a.id, " acceptor not take propose:", proposeMsg.val)
		return nil
	}
	a.acceptMsg = proposeMsg
	proposeMsg.typ = Accept
	return &proposeMsg
}
