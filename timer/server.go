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
		return fmt.Sprintf("<fc=#e9d460>%s</fc>", "0"+strconv.Itoa(i))
	default:
		return fmt.Sprintf("<fc=#87d37c>%s</fc>", strconv.Itoa(i))
	}
}

func formatTime(i int) string {
	min := digit(i / 60)
	sec := digit(i % 60)
	return fmt.Sprintf("<fc=#d64541>TIMER</fc> [%s:%s]\n", min, sec)
}

func timer(t int, c, p chan string, q, done chan bool) {
	tick := time.Tick(time.Second)
Loop:
	for {
		if t <= -1 {
			p <- "end"
			done <- true
			break Loop
		}
		if t == 30 {
			select {
			case p <- "30":
			default:
			}

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

func audio(c chan string) {
	for {
		switch <-c {
		case "start":
			play("start2")
		case "end":
			play("end")
		case "30":
			play("count_down_30")
		}
	}
}

func main() {
	t, _ := strconv.Atoi(os.Args[1])
	p := make(chan string)
	c := make(chan string)
	q := make(chan bool)
	done := make(chan bool)
	go SocketServer(3333, c, q)
	go timer(60*t, c, p, q, done)
	go audio(p)
	p <- "start"
	<-done

}
