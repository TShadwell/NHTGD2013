package main

import (
	"net/http"
	"fmt"
	"github.com/TShadwell/nhtgd2013/twfy"
	"log"
	"github.com/TShadwell/nhtgd2013/dbabstract"
	"github.com/TShadwell/nhtgd2013/markov"
	"github.com/TShadwell/nhtgd2013/secrets"
	"strings"
	"time"
	"math/rand"
)

const numwords = 2000

const numtwo = 300

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

func MarkovMP(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	var pid twfy.PersonID

	pidstr := r.Form["pid"]
	if pidstr == nil&& len(pidstr) <= 0{
		fmt.Fprint(w, "pid must be provided")
		return
	}

	n, err := fmt.Sscan(
		" " + pidstr[0] + " ",
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

	if err != nil{
		fmt.Fprint(w, "Error reading from database.")
		log.Println(err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	if m == nil{
		m = markov.NewChain(numwords)

		texts, err := API.GetTexts(
			pid,
		)

		if err != nil{
			fmt.Fprint(w, "Error getting twfy")
			log.Println(err)
			return
		}
		textss := strings.Join(texts, " ")
		//fmt.Println("\nGOT TEXTS: " + textss)
		_, e := m.Write([]byte(textss))
		if e != nil{
			panic(e)
		}

		//store the markov chain
		/*err = database.StoreChain(
			*m,
			pid,
		)*/

		/*if err != nil{
			fmt.Fprint(w, "Error storing chain in database.")
			log.Println(err)
		}*/
	}
	fmt.Fprint(w, m.Generate(numtwo))

}
