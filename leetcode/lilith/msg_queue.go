package lilith

import (
	"sync"
)

type Message struct {
	Msg   string
	MsgId int
}

type MsgQueue struct {
	Msgs []*Message
	Lock *sync.Mutex
}

func (self *MsgQueue) WriteMessage(msg string, id int) {
	self.Lock.Lock()
	defer self.Lock.Unlock()

	message := &Message{Msg: msg}
	self.Msgs = append(self.Msgs, message)

}

func (self *MsgQueue) ReadMessage(id int) string {
	self.Lock.Lock()
	defer self.Lock.Unlock()

	message := self.Msgs[0]
	self.Msgs = self.Msgs[1:]
	return message.Msg
}
