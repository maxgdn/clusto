package main

import (
	"fmt"
	"log"
	"net/http"
	"text/tabwriter"
	"os"

	"github.com/maxgdn/clusto/internal/info"
)

type clientHandle struct {
	clients []Info
}

//updates view
func updateUI(clients []Info) {
	//prints out with color to term
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Key \t Domain \t IP \t Updated Last\t.")
	fmt.Fprintln(w, "a\tb\tc")
	fmt.Fprintln(w)
	w.Flush()
	//cycle by 10 at a time
}

//handles updating of list
func upsertClient(clients []Info) []Info {

	//if in list update the entry

	//else append
	clients = append(clients, "DOG")
	return clients
}

func (h *clientHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	h.clients = upsertClient(h.clients)
	updateUI(h.clients)
	fmt.Println(h.clients)
}

func main() {
	var handle = &clientHandle{clients: make([]Info, 0)}
	http.Handle("/update", handle)
	//ui call
	log.Fatal(http.ListenAndServe(":8080", nil))

}
