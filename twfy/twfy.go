package twfy

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func (a *API) get(endpoint string, args url.Values) (bytes []byte, err error){
	args.Add("output", "js")
	resp, err := http.Get(
		"http://www.theyworkforyou.com/api/" +
		endpoint + "?" +
		args.Encode(),
	)

	if err != nil{
		return
	}
	defer resp.Body.Close()
	bytes, err = ioutil.ReadAll(
		resp.Body,
	)

	return
}

//func (a *API) GetMPById(i PersonID) 
func (a *API) GetMpsForParty(p Party) (ms []Member, err error){
	bytes, err := a.get(
		"getMPs",
		url.Values{
			"party":{string(p)},
		},
	)

	if err != nil{
		return
	}

	err = json.Unmarshal(bytes, &ms)

	return

}
func (a *API) GetMPById(i PersonID)
func (a *API) GetMPById(i PersonID)
func (a *API) GetMPById(i PersonID)

