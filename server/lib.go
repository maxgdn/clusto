package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"text/tabwriter"
	"os"
	"time"
)

type clientHandle struct {
	clients []Info
}


type Info struct {
	Key string
	Hostname string
	IP string
	Updated string
}
//updates view
func updateUI(clients []Info) {
	//prints out with color to term
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "NUM | Key | Domain | IP | Updated")
	for i, c := range clients {
		row := "[" + string(i) + "]" + " \t " + c.Key + " \t " + c.Hostname + " \t " + c.IP + " \t " + c.Updated + " \t "
		fmt.Fprintln(w, row)
    }
	fmt.Fprintln(w)
	w.Flush()
	//cycle by 10 at a time
}

//handles updating of list
func upsertClient(clients []Info, sent Info) []Info {

	//if in list update the entry
	for i, c := range clients {
		if(c.Key == sent.Key) {
			clients[i] = sent
			return clients
		}
    }

	//else append
	clients = append(clients, sent)
	return clients
}

func (h *clientHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var info Info

	err := json.NewDecoder(r.Body).Decode(&info)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	h.clients = upsertClient(h.clients, info)
	updateUI(h.clients)
	//fmt.Println(h.clients)
}

func uiLoop(handle *clientHandle) {
	for {
		updateUI(handle.clients)
		time.Sleep(5000 * time.Millisecond)
	}
}

func main() {
	var handle = &clientHandle{clients: make([]Info, 0)}
	http.Handle("/update", handle)
	//ui call
	go uiLoop(handle)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
