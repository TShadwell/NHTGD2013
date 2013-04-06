package twfy

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func (a *API) get(endpoint string, args url.Values) (bytes []byte, err error){
	args.Add("output", "js")
	args.Add("key", string(a.Key))
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

	var errTest jsonResponse

	//Check for errors before we go
	err = json.Unmarshal(bytes, &errTest)
	if err != nil{
		return
	}

	//Got TWFY error.
	if errTest.Error != ""{
		err = errTest.Error
	}

	return
}

//func (a *API) GetMPById(i PersonID) 
/*
	Created specifically for RobotMP
*/
func (a *API) GetTexts(p PersonID) (op []string, err error){
	bytes, err := a.get(
		"getHansard",
		url.Values{
			"person":{
				fmt.Sprint(p),
			},
			"num":{
				"1000",
			},
		},
	)

	var marshalled jsonHansard
	err = json.Unmarshal(bytes, &marshalled)
	for _, v := range marshalled.Rows{
		op = append(op, sanitiseTexts(v.Body))
	}
	return
}

func sanitiseTexts(t string) string{
	return t
}

func (a *API) GetMembers() (ms []Member, err error){
	bytes, err := a.get(
		"getMPs",
		url.Values{},
	)
	if err != nil{
		return
	}

	err = json.Unmarshal(bytes, &ms)
	return
}
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

