package main

import "fmt"
import "net"
import "net/http"

func main() {
    http.HandleFunc("/", hello)
    http.HandleFunc("/ping", ping)
    http.ListenAndServe(":1234", nil)
    fmt.Println(GetLocalIP())
}   

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello, my IP Addrs is "))
    w.Write([]byte(GetLocalIP()))
}

func ping(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong!"))
}

func GetLocalIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, address := range addrs {
        // check the address type and if it is not a loopback the display it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}
