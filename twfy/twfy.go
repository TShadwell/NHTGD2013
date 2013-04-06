package twfy

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"strings"
	"net/http"
	"net/url"
	"regexp"
)

var tagRemover = regexp.MustCompile("<[^>]*>")

func (a *API) get(endpoint string, args url.Values) (bytes []byte, err error) {
	args.Add("output", "/js")
	args.Add("key", string(a.Key))
	resp, err := http.Get(
		"http://www.theyworkforyou.com/api/" +
			endpoint + "?" +
			args.Encode(),
	)

	fmt.Println("Args: ", args.Encode())

	if err != nil {
		return
	}
	defer resp.Body.Close()
	bytes, err = ioutil.ReadAll(
		resp.Body,
	)

	//var errTest jsonResponse

	//Check for errors before we go
	//err = json.Unmarshal(bytes, &errTest)
	/*
	if err != nil {
		return
	}
	*/

	//Got TWFY error.
	/*if errTest.Error != "" {
		err = errTest.Error
	}*/

	return
}

/*
	Created specifically for RobotMP
*/
func (a *API) GetTexts(p PersonID) (op []string, err error) {
	bytes, err := a.get(
		"getHansard",
		url.Values{
			"person": {
				fmt.Sprint(p),
			},
			"num": {
				"450",
			},
		},
	)

	if err != nil{
		return
	}

	var marshalled JsonHansard
	err = json.Unmarshal(bytes, &marshalled)
	if err != nil{
		return
	}
	fmt.Println(marshalled)
	for _, v := range marshalled.Rows {
		t := sanitiseTexts(v.Body)
		if !strings.HasPrefix(t, "To"){
			op = append(op, t)
		}
	}
	return
}

func sanitiseTexts(t string) string {
	return html.UnescapeString(tagRemover.ReplaceAllLiteralString(t, ""))
}

func (a *API) GetMembers() (ms []Member, err error) {
	bytes, err := a.get(
		"getMPs",
		url.Values{},
	)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &ms)
	return
}
func (a *API) GetMpsForParty(p Party) (ms []Member, err error) {
	bytes, err := a.get(
		"getMPs",
		url.Values{
			"party": {string(p)},
		},
	)

	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &ms)

	return

}
