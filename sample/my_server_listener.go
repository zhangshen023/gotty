package sample

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/funnycode-org/gotty/base"
	"github.com/funnycode-org/gotty/server/listener"
	"sync"
	"time"
)

const _ListenerName = "my_server_listener"

func init() {
	listener.RegisterListener(_ListenerName, func() (listener.Listener, error) {
		return &MyServerListener{
			Sessions: make(map[int]base.Session),
			lock:     sync.RWMutex{},
		}, nil
	})
}

type MyServerListener struct {
	Sessions map[int]base.Session
	lock     sync.RWMutex
}

func (d *MyServerListener) FactoryConstruct() error {
	panic("implement me")
}

// 用户在这里能得到session，然后去发送数据
func (d *MyServerListener) OnOpen(session base.Session) error {
	fmt.Printf("session %d is opened!", session.SessionId())
	d.lock.RLock()
	d.Sessions[session.SessionId()] = session
	d.lock.RUnlock()

	time.AfterFunc(time.Second*2, func() {
		d.TestSendByte()
	})
	return nil
}

//
func (d *MyServerListener) TestSendByte() error {
	for i, session := range d.Sessions {
		content := fmt.Sprintf("%d: I announce that I am a handsome boy!", i)
		cpblf := CustomizeProtocolBasedLengthField{
			Type:   1,
			Flag:   2,
			Length: uint64(len(content)),
			Body:   content,
		}
		myBytes := make([]byte, 1+1+8+len(cpblf.Body))
		buf := bytes.NewBuffer(myBytes)
		buf.WriteByte(cpblf.Type)
		buf.WriteByte(cpblf.Flag)
		binary.Write(buf, binary.LittleEndian, cpblf.Length)
		buf.WriteString(cpblf.Body)
		session.Send(buf.Bytes())
	}
	return nil
}

func (d *MyServerListener) OnClose(session base.Session) error {
	fmt.Printf("session %d is closed!", session.SessionId())
	return nil
}

func (d *MyServerListener) OnSend(session base.Session) error {
	fmt.Printf("session %d is sending bytes: %x!", session.SessionId(), session)
	return nil
}

func (d *MyServerListener) OnReceive(session base.Session, bytes []byte) ([]byte, error) {
	panic("implement me")
}

func (d *MyServerListener) OnError(session base.Session) error {
	panic("implement me")
}

func (d *MyServerListener) GetRegistryListenerName() string {
	return _ListenerName
}
