package main
import (
    "bufio"
    "fmt"
    "log"
    "net"
    "strings"
    "io"
)
func handleConn(conn net.Conn){
     defer conn.Close()
     scanner := bufio.NewScanner(conn)
     i :=0
     for scanner.Scan() {
         ln := scanner.Text()
         fmt.Println(ln)
         if i == 0 {
              method :=strings.Fields(ln)[0]
              fmt.Println("we printed this - method: ",method)
            } else {
            }
            i++
     }

     body := "hello world 2"
     io.WriteString(conn, "http/1.1 200 OK\r\n")
     fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
     io.WriteString(conn, "\r\n")
     io.WriteString(conn, body)
}
func main() {
    server,err := net.Listen("tcp",":8081")
    if err != nil {
       log.Fatalln(err.Error())
    }
    defer server.Close()
    for {
        conn,err := server.Accept()
        if err != nil {
           log.Fatalln(err.Error())
         }
         go handleConn(conn)
     }
  }