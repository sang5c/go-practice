package _interface

import (
	"example/syntax/interface/fedex"
	"example/syntax/interface/ko"
	"testing"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func TestSend(t *testing.T) {
	fedexSender := &fedex.FedexSender{}
	SendBook("어린 왕자", fedexSender)
	SendBook("그리스인 조르바", fedexSender)

	postSender := &ko.PostSender{}
	SendBook("어린 왕자", postSender)
	SendBook("그리스인 조르바", postSender)
}