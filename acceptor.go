package paxos

import "log"

//Create a accetor and also assign learning IDs into acceptor.
//Acceptor: Will response request from proposer, promise the first and largest seq number propose.
//          After proposer reach the majority promise.  Acceptor will pass the proposal value to learner to confirn and choose.
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
		//	log.Println("acceptor:", a.id, " wait to recev msg")
		m := a.nt.recev()
		if m == nil {
			continue
		}

		//	log.Println("acceptor:", a.id, " recev message ", *m)
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
	log.Println("ID:", a.id, " Promise")
	prepare.to = prepare.from
	prepare.from = a.id
	prepare.typ = Promise
	a.acceptMsg = prepare
	return &prepare
}

//Recev Propose only check if acceptor already accept bigger propose before.
//Otherwise, will just forward this message out and change its type to "Accept" to learning later.
func (a *acceptor) recevPropose(proposeMsg message) bool {
	//Already accept bigger propose before
	log.Println("accept:check propose. ", a.acceptMsg.getProposeSeq(), proposeMsg.getProposeSeq())
	if a.acceptMsg.getProposeSeq() > proposeMsg.getProposeSeq() || a.acceptMsg.getProposeSeq() < proposeMsg.getProposeSeq() {
		log.Println("ID:", a.id, " acceptor not take propose:", proposeMsg.val)
		return false
	}
	log.Println("ID:", a.id, " Accept")
	return true
}
