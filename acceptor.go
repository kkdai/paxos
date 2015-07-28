package paxos

import "log"

func NewAcceptor(id int, nt nodeNetwork, learners ...int) acceptor {
	newAccptor := acceptor{id: id, nt: nt}
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
func (a *acceptor) run() {
	for {
		log.Println("acceptor:", a.id, " wait to recev msg")
		m := a.nt.recev()
		if m == nil {
			continue
		}

		log.Println("acceptor:", a.id, " recev message ", *m)
		switch m.typ {
		case Prepare:
			promiseMsg := a.recevPrepare(*m)
			a.nt.send(*promiseMsg)
			continue
		case Propose:
			accepted := a.recevPropose(*m)
			if accepted {

				for _, lId := range a.learners {
					m.from = a.id
					m.to = lId
					m.typ = Accept
					a.nt.send(*m)
				}
			}
		default:
			log.Fatalln("Unsupport message in accpetor ID:", a.id)
		}
	}
	log.Println("accetor :", a.id, " leave.")
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
	if a.acceptMsg.getProposeSeq() > proposeMsg.getProposeSeq() || a.acceptMsg.getProposeSeq() < proposeMsg.getProposeSeq() {
		log.Println("ID:", a.id, " acceptor not take propose:", proposeMsg.val)
		return false
	}
	return true
}
