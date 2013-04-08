package main

import (
	"fmt"
	"github.com/TShadwell/NHTGD2013/dbabstract"
	"github.com/TShadwell/NHTGD2013/markov"
	"github.com/TShadwell/NHTGD2013/secrets"
	"github.com/TShadwell/NHTGD2013/twfy"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const numwords = 300

func main() {
	http.HandleFunc(
		"/markov",
		MarkovMP,
	)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

var API = twfy.API{
	Key: secrets.TWFYKey,
}

func MarkovMP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var pid twfy.PersonID

	pidstr := r.Form["pid"]
	if pidstr == nil && len(pidstr) <= 0 {
		fmt.Fprint(w, "pid must be provided")
		return
	}

	n, err := fmt.Sscan(
		" "+pidstr[0]+" ",
		&pid,
	)

	switch {
	case err != nil:
		log.Println(err)
		fallthrough
	case n == 0:
		fmt.Fprint(w, "Invalid PersonID :(")
		return
	}

	m, err := database.RetrieveChain(
		pid,
	)

	//var m *markov.Chain

	if err != nil {
		fmt.Fprint(w, "Error reading from database.")
		log.Println(err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	if m == nil {
		m = markov.NewChain(3)

		err := API.WriteTexts(
			m,
			pid,
		)

		if err != nil {
			fmt.Fprint(w, "Error getting twfy")
			log.Println(err)
			return
		}

		//store the markov chain
		err = database.StoreChain(
			*m,
			pid,
		)

		if err != nil {
			fmt.Fprint(w, "Error storing chain in database.")
			log.Println(err)
		}
	}
	fmt.Fprint(w, m.Generate(numwords))

}
