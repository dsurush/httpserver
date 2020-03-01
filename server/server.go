package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	host := "0.0.0.0"
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "9999"
	}
	err := start(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("cant start %s:%s", host, port)
	}
}

func start(addr string) (err error){
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("can't listen on %s: %s", addr, err)
		return err
	}
	defer func() {
		err := listener.Close()
		if err != nil {
			log.Fatalf("Get listener %e", err)
		}
	}()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("can't accept connection %v", conn)
			continue
		}
		handleConn(conn)
	}
	return nil
}

func handleConn(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("can't close connection:%v", err)
		}
	}()
	reader := bufio.NewReader(conn)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("can't read %v", err)
	}
	log.Printf(requestLine)
	MRP := strings.Split(strings.TrimSpace(requestLine), " ")
	if len(MRP) != 3{
		return           /////////////////----->>> Что это??
	}
	method := MRP[0]
	request := MRP[1]
	protocol := MRP[2]

	if method == "GET" && protocol == `HTTP/1.1` {
		switch request {
		case "/":
			writer := bufio.NewWriter(conn)
			fileBytes, err := ioutil.ReadFile("html/index.html")
			if err != nil {
				log.Fatalf("cant read from file %s: %er", "html/index.html", err)
			}
			_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
			_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(fileBytes)))
			_, _ = writer.WriteString("Content-Type: text/html" + "\r\n")
			_, _ = writer.WriteString("Connection: Close\r\n")
			_, _ = writer.WriteString("\r\n")
			_, _ = writer.Write(fileBytes)
			err = writer.Flush()
			if err != nil {
				log.Printf("can't sent response: %v", err)
			}
		case "/images/taxi.jpeg":
			writer := bufio.NewWriter(conn)
			fileBytes, err := ioutil.ReadFile("images/taxi.jpeg")
			if err != nil {
				log.Fatalf("cant read from file %s: %er", "images/taxi.jpeg", err)
			}
			_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
			_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(fileBytes)))
			_, _ = writer.WriteString("Content-Type: taxi.jpeg" + "\r\n")
			_, _ = writer.WriteString("Connection: Close\r\n")
			_, _ = writer.WriteString("\r\n")
			_, _ = writer.Write(fileBytes)
			err = writer.Flush()
			if err != nil {
				log.Printf("can't sent response: %v", err)
			}
		case "/justTXT.html":
			writer := bufio.NewWriter(conn)
			fileBytes, err := ioutil.ReadFile("html/justTXT.html")
			if err != nil {
				log.Fatalf("cant read from file %s: %er", "html/justTXT.html", err)
			}
			_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
			_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(fileBytes)))
			_, _ = writer.WriteString("Content-Type: taxi.jpeg" + "\r\n")
			_, _ = writer.WriteString("Connection: Close\r\n")
			_, _ = writer.WriteString("\r\n")
			_, _ = writer.Write(fileBytes)
			err = writer.Flush()
			if err != nil {
				log.Printf("can't sent response: %v", err)
			}
		case "/pdf/HTTP.pdf":
			writer := bufio.NewWriter(conn)
			fileBytes, err := ioutil.ReadFile("pdf/HTTP.pdf")
			if err != nil {
				log.Fatalf("cant read from file %s: %er", "pdf/HTTP.pdf", err)
			}
			_, err = writer.WriteString("HTTP/1.1 200 OK\r\n")
			if err != nil {
				log.Fatalf("can't send http/1.1 200 ok : %e", err)
			}
			_, err = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(fileBytes)))
			if err != nil {
				log.Fatalf("can't send length %e", err)
			}
			_, err = writer.WriteString("Content-Type: application/pdf" + "\r\n")
			if err != nil {
				log.Fatalf("can't send content-type %e", err)
			}
			_, err = writer.WriteString("Connection: Close\r\n")
			if err != nil {
				log.Fatalf("cant send conn close %e", err)
			}
			_, err = writer.WriteString("\r\n")
			if err != nil {
				log.Fatalf("cant send `/r/n` %e", err)
			}
			_, err = writer.Write(fileBytes)
			if err != nil {
				log.Fatalf("cant send bytes %e", err)
			}
			err = writer.Flush()
			if err != nil {
				log.Printf("can't sent response: %v", err)
			}
		case "/text/text.txt":
			writer := bufio.NewWriter(conn)
			fileBytes, err := ioutil.ReadFile("text/text.txt")
			if err != nil {
				log.Fatalf("cant read from file %s: %er", "text/text.txt", err)
			}
			_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
			_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(fileBytes)))
			_, _ = writer.WriteString("Content-Type: text" + "\r\n")
			_, _ = writer.WriteString("Connection: Close\r\n")
			_, _ = writer.WriteString("\r\n")
			_, _ = writer.Write(fileBytes)
			err = writer.Flush()
			if err != nil {
				log.Printf("can't sent response: %v", err)
			}
		case "/text/text.txt?download":
			writer := bufio.NewWriter(conn)
			fileBytes, err := ioutil.ReadFile("text/text.txt")
			if err != nil {
				log.Fatalf("cant read from file %s: %er", "text/text.txt", err)
			}
			_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
			_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(fileBytes)))
			_, _ = writer.WriteString("Content-Disposition: attachment; filename=text/text.txt\r\n")
			_, _ = writer.WriteString("Connection: Close\r\n")
			_, _ = writer.WriteString("\r\n")
			_, _ = writer.Write(fileBytes)
			err = writer.Flush()
			if err != nil {
				log.Printf("can't sent response: %v", err)
			}
		}
	}
}





