package paxos

import "log"

type learner struct {
	id           int
	acceptedMsgs map[int]message
	nt           nodeNetwork
}

//Initilize learner and prepare message pool.
func NewLearner(id int, nt nodeNetwork, acceptorIDs ...int) *learner {
	newLearner := &learner{id: id, nt: nt}
	newLearner.acceptedMsgs = make(map[int]message)
	for _, aceptId := range acceptorIDs {
		newLearner.acceptedMsgs[aceptId] = message{}
	}
	return newLearner
}

//Run learner process and will return learn value if reach majority.
func (l *learner) run() string {

	for {
		m := l.nt.recev()
		if m == nil {
			continue
		}

		log.Println("Learner: recev msg:", *m)
		l.handleRecevAccept(*m)
		learnMsg, isLearn := l.chosen()
		if isLearn == false {
			continue
		}
		return learnMsg.getProposeVal()
	}
}

//Check acceptor message and compare with local accpeted proposal to make sure it is most updated.
func (l *learner) handleRecevAccept(acceptMsg message) {
	hasAcceptedMsg := l.acceptedMsgs[acceptMsg.from]
	if hasAcceptedMsg.getProposeSeq() < acceptMsg.getProposeSeq() {
		//get bigger num will replace it to keep it updated.
		l.acceptedMsgs[acceptMsg.from] = acceptMsg
	}
}

//Every acceptor might send different proposal ID accept mesage to learner.
//Learner only chosen if the accept count reach majority.
func (l *learner) chosen() (message, bool) {
	acceptCount := make(map[int]int)
	acceptMsgMap := make(map[int]message)

	//Separate each acceptor message according their proposal ID and count it
	for _, msg := range l.acceptedMsgs {
		proposalNum := msg.getProposeSeq()
		acceptCount[proposalNum]++
		acceptMsgMap[proposalNum] = msg
	}

	//Check count if reach majority will return as chosen value.
	for chosenNum, chosenMsg := range acceptMsgMap {
		if acceptCount[chosenNum] > l.majority() {
			return chosenMsg, true
		}
	}
	return message{}, false
}

//Count for majority, need initilize the count at constructor.
func (l *learner) majority() int {
	return len(l.acceptedMsgs)/2 + 1
}
