package main

import (
	"fmt"
	"os"
	"time"
)

type Storage struct {
	P int
	A [1000000]int
}

func (s *Storage) produce() {
	if s.P == 999999 {
		return
	}
	if s.P > -1 && s.A[s.P] != 1 {
		fmt.Println("producer error")
		os.Exit(1)
	}
	s.P++
	if s.A[s.P] != 0 {
		fmt.Println("producer error")
		os.Exit(1)
	}
	s.A[s.P] = 1
	//fmt.Println("produced", s.P)
}

func (s *Storage) consume() {
	if s.P < 0 {
		return
	}
	if s.A[s.P] != 1 {
		fmt.Println("consumer error")
		os.Exit(1)
	}
	s.A[s.P] = 0
	s.P--
	if s.P > -1 && s.A[s.P] != 1 {
		fmt.Println("consumer error")
		os.Exit(1)
	}
	//fmt.Println("consumed", s.P)
}


func producer(s *Storage) {
	for {
		s.produce()
	}
}

func consumer(s *Storage) {
	for {
		s.consume()
	}
}

func main() {
	var a [1000000]int
	s := Storage{P:-1, A:a}
	go producer(&s)
	go consumer(&s)
	time.Sleep(100 * time.Second)
}

