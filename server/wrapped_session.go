package server

import (
	"github.com/funnycode-org/gotty/base"
	"reflect"
)

type WrappedSession struct {
	session *Session
}

func (w *WrappedSession) GetListener() base.Listener {
	panic("can't be used!")
}

func (w *WrappedSession) GetWrappedSession() interface{} {
	panic("can't be used!")
}

func newWrappedSession(session *Session) base.Session {
	return &WrappedSession{
		session: session,
	}
}

func (w *WrappedSession) Close() error {
	return w.session.Close()
}

func (w *WrappedSession) SessionId() int {
	return w.session.SessionId()
}

func (w *WrappedSession) Send(bytes []byte) error {
	return w.session.Send(bytes)
}

func (w *WrappedSession) GetRegistryProtocol() reflect.Type {
	return w.session.GetRegistryProtocol()
}

func (w *WrappedSession) GetSendChannel() <-chan []byte {
	panic("can't be called!")
}
