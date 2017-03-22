// Copyright © 2017 Karthik Nayak <Karthik.188@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package uber

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/karthiknayak/cab/cab"
	"github.com/karthiknayak/cab/geocode"

	"net/http"
)

//Uber type exposes the uber cab and holds the Server Auth key.
type Uber struct {
	key string
}

type errorMsg struct {
	Message string `json:"message"`
}

//CabFare is a struct to parse the JSON returned via the Uber API when
//fetching the fare estimate for a single class of vehicle.
type CabFare struct {
	Currency string  `json:"currency_code"`
	Name     string  `json:"localized_display_name"`
	Estimate string  `json:"estimate"`
	Minimum  int     `json:"minimum"`
	EstLow   float64 `json:"low_estimate"`
	EstHigh  float64 `json:"high_estimate"`
	Surge    float64 `json:"surge_multiplier"`
	Duration int     `json:"duration"`
	Dist     float64 `json:"distance"`
}

//CabFares is a struct to parse the JSON returned via the Uber API
//when fetching the fare estimate for all vehicle classes.
type CabFares struct {
	Fares []CabFare `json:"prices"`
}

//CabTime is a struct to parse the JSON returned via the Uber API when
//fetching the time to arrival estimate for a single class of vehicle.
type CabTime struct {
	Name     string `json:"localized_display_name"`
	Estimate int    `json:"estimate"`
}

//CabTimes is a struct to parse the JSON returned via the Uber API when
//fetching the time to arrival estimate for all vehicle classes.
type CabTimes struct {
	Times []CabTime `json:"times"`
}

//Init function assigns the Server Auth key
func (u *Uber) Init() {
	u.key = "xX9rSwC86iTW14KhhMBsz-cPKMufTQkgJ0ONCBh-"
}

func addParamToURL(url string, str map[string]string) string {
	i := len(str)
	for key, val := range str {
		url = url + key + "=" + val
		i = i - 1
		if i > 0 {
			url = url + "&"
		}
	}
	return url
}

func (u *Uber) getUberData(url string, d interface{}) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Uber: Couldn't create new HTTP request", err)
	}
	req.Header.Set("Accept-Language", "en_US")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Token "+u.key)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Uber: Couldn't retireve estimate time", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Uber: Can't parse http body", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var eMsg errorMsg
		err := json.Unmarshal(body, &eMsg)
		if err != nil {
			log.Fatal("Uber couldn't Umarshal the data", err)
		}
		log.Fatal(eMsg.Message)
	}

	err = json.Unmarshal(body, d)
	if err != nil {
		log.Fatal("Uber couldn't Umarshal the data", err)
	}
}

//GetFare gets the Fares for all classes of vehicles for a given
//'from' and 'to' address.
func (u *Uber) GetFare(a, b string) *cab.Fares {
	from := geocode.GetCode(a)
	to := geocode.GetCode(b)

	str := make(map[string]string)
	str["start_latitude"] = strconv.FormatFloat(from.Lat, 'f', -1, 64)
	str["start_longitude"] = strconv.FormatFloat(from.Lng, 'f', -1, 64)
	str["end_latitude"] = strconv.FormatFloat(to.Lat, 'f', -1, 64)
	str["end_longitude"] = strconv.FormatFloat(to.Lng, 'f', -1, 64)

	url := addParamToURL("https://api.uber.com/v1.2/estimates/price?", str)
	var data = new(CabFares)

	u.getUberData(url, data)

	var fares = new(cab.Fares)

	for _, val := range data.Fares {
		fares.Fare = append(fares.Fare, cab.Fare{Currency: val.Currency, Name: val.Name, Est: val.Estimate, Surge: val.Surge, Duration: val.Duration, Dist: val.Dist})
	}

	fares.From = a
	fares.To = b

	return fares
}

//GetTime gets the estimated time of arrival for all classes of
//vehicles for a given address.
func (u *Uber) GetTime(a string) *cab.Times {
	loc := geocode.GetCode(a)

	str := make(map[string]string)
	str["start_latitude"] = strconv.FormatFloat(loc.Lat, 'f', -1, 64)
	str["start_longitude"] = strconv.FormatFloat(loc.Lng, 'f', -1, 64)

	url := addParamToURL("https://api.uber.com/v1.2/estimates/time?", str)
	var data = new(CabTimes)

	u.getUberData(url, data)

	var times = new(cab.Times)
	for _, val := range data.Times {
		times.Time = append(times.Time, cab.Time{Name: val.Name, Est: val.Estimate})
	}

	times.Loc = a

	return times
}
