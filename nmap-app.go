package main

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func scanPort(protocol, hostname string, port int, results chan int) {
    address := net.JoinHostPort(hostname, fmt.Sprintf("%d", port))
    conn, err := net.DialTimeout(protocol, address, 1*time.Second)
    if err == nil {
        results <- port
        conn.Close()
    } else {
        results <- 0 // send 0 to maintain expected count
    }
}

func main() {
    host := "scanme.nmap.org"
    ports := make(chan int, 100)
    results := make(chan int)
    var openPorts []int

    // Spawn 100 workers
    for i := 0; i < cap(ports); i++ {
        go func() {
            for p := range ports {
                scanPort("tcp", host, p, results)
            }
        }()
    }

    // Feed 1024 ports
    go func() {
        for i := 1; i <= 1024; i++ {
            ports <- i
        }
        close(ports)
    }()

    // Read exactly 1024 results
    for i := 1; i <= 1024; i++ {
        port := <-results
        if port != 0 {
            openPorts = append(openPorts, port)
        }
    }

    sort.Ints(openPorts)
    fmt.Println("Open ports:")
    for _, port := range openPorts {
        fmt.Printf("%d open\n", port)
    }
}
