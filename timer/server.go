package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func SocketServer(port int, c chan string, q chan bool) {
	listen, err := net.Listen("tcp4", "127.0.0.1:"+strconv.Itoa(port))
	defer listen.Close()
	if err != nil {
		panic(err)
	}
	// log.Printf("server start at :%d", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handler(conn, c, q)
	}

}

func handler(conn net.Conn, s chan string, q chan bool) {
	defer conn.Close()
	w := bufio.NewWriter(conn)
	q <- true
	w.Write([]byte(<-s))
	w.Flush()

}

func digit(i int) string {
	switch {
	case i < 10:
		return "0" + strconv.Itoa(i)
	default:
		return strconv.Itoa(i)
	}
}

func formatTime(i int) string {
	min := digit(i / 60)
	sec := digit(i % 60)
	return fmt.Sprintf("<fc=#d64541>TIMER</fc> [%s:%s]\n", min, sec)
}

func timer(t int, c chan string, q, done chan bool) {
	tick := time.Tick(time.Second)
Loop:
	for {
		if t <= 0 {
			done <- true
			break Loop
		}
		select {
		case <-q:
			c <- formatTime(t)
		case <-tick:
			t--
		default:
		}
		time.Sleep(10 * time.Millisecond)
	}

}

func main() {
	go play("start")
	t, _ := strconv.Atoi(os.Args[1])
	c := make(chan string)
	q := make(chan bool)
	done := make(chan bool)
	go SocketServer(3333, c, q)
	go timer(60*t, c, q, done)
	// go printer(c, q)
	<-done

}
