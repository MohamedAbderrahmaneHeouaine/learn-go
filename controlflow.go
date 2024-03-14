package main

import (
	"fmt"
	"time"
)

type ServerNew struct {
	quitch   chan struct{} // 0 bytes
	msgchnew chan string
}

func (s *ServerNew) quit() {
	//close(s.quitch)
	s.quitch <- struct{}{}
}
func newServer() *ServerNew {
	return &ServerNew{
		quitch:   make(chan struct{}),
		msgchnew: make(chan string, 128),
	}
}

func (s *ServerNew) start() {
	fmt.Println("server starting")
	s.loop() // block
}
func (s *ServerNew) loop() {
mainloop:
	for {
		select {
		case <-s.quitch:
			fmt.Println("quiting server")
			break mainloop
		case msg := <-s.msgchnew:
			s.handleMessage(msg)
		default:
			//fmt.Println("start")
		}
	}
	fmt.Println("server is shutting down")
}
func (s *ServerNew) handleMessage(msg string) {
	fmt.Println("we recieve a message", msg)
}
func (s *ServerNew) sendMessage(msg string) {
	s.msgchnew <- msg
}
func control() {
	server := newServer()
	//go server.start()
	//server.sendMessage("hey")
	go func() {
		time.Sleep(time.Second * 5)
		server.quit()
	}()
	server.start()
}
