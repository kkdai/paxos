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

//Acceptor process detail logic.
func (a *accept) run() {
	for {
		m := a.nt.receive()
		switch m.typ {
		case Prepare:
			promiseMsg := a.recevPrepare(m)
			a.nt.send(promiseMsg)
			continue
		case Propose:
			accepted := a.recevPropose(m)
			if accepted {

				for _, learner := range learners {
					m.aafrom = a.id
					m.to = learner
					m.typ = Accept
					a.nt.send(m)
				}
			}
		default:
			log.Fatalln("Unsupport message in accpetor ID:", a.id)
		}
	}
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
func (a *acceptor) recevPropose(proposeMsg message) bool {
	//Already accept bigger propose before
	if a.acceptMsg.getProposeSeq() > proposeMsg.getProposeSeq() || a.acceptMsg.getProposeSeq() < promiseMsg.getProposeSeq() {
		log.Println("ID:", a.id, " acceptor not take propose:", proposeMsg.val)
		return false
	}
	return true
}
