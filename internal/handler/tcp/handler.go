package tcp

import (
	"bufio"
	"fmt"
	"github.com/taranovegor/com.jurnalo/internal/collector"
	"net"
)

type Handler struct {
	listener  net.Listener
	collector *collector.Collector
}

func NewHandler(
	listener net.Listener,
	collector *collector.Collector,
) Handler {
	return Handler{
		listener:  listener,
		collector: collector,
	}
}

func (hdlr Handler) Handle() {
	defer hdlr.listener.Close()

	for {
		conn, err := hdlr.listener.Accept()
		if err != nil {
			fmt.Println("error accepting: ", err.Error())

			continue
		}

		go hdlr.handleConnection(conn)
	}
}

func (hdlr Handler) handleConnection(conn net.Conn) {
	sc := bufio.NewScanner(conn)
	it := collector.NewEmptyItem()
	for sc.Scan() {
		line := sc.Text()
		if len(line) == 0 {
			hdlr.collector.Collect(it)
			it = collector.NewEmptyItem()

			continue
		}

		it.ParseLine(line)
	}
}
