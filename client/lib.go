package main

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

var url string

type Info struct {
	Key      string
	Hostname string
	IP       string
	Updated  string
}

func getIP(hostname string) string {
	addr, err := net.LookupIP(hostname)
	if err != nil {
		return ""
	} else {
		return net.IP.String(addr[0])
	}
}

func getMetrics(id string) Info {
	host, _ := os.Hostname()
	ip := getIP(host)
	t := time.Now().Format(time.RFC850)
	return Info{Key: id, Hostname: host, IP: ip, Updated: t}

}

func submit(id string) {
	var info = getMetrics(id)
	content, _ := json.Marshal(info)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(content))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func main() {
	url = "http://localhost:8080/update"
	var id = uuid.New().String()
	for {
		submit(id)
		time.Sleep(30000 * time.Millisecond)
	}
}
