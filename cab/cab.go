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

package cab

import (
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

//Fare holds the fare related data for a particular class of vehicles.
type Fare struct {
	Currency string
	Name     string
	Est      string
	Surge    float64
	Duration int
	Dist     float64
}

//Fares holds the fares for all classes of vehicles with the From and
//to address.
type Fares struct {
	Fare []Fare
	From string
	To   string
}

//Time holds the estimated time of arrival for a particular class of vehicle.
type Time struct {
	Name string
	Est  int
}

//Times holds the times for all class of vehicles with the pickup
//Location.
type Times struct {
	Time []Time
	Loc  string
}

//Cab interface is used to get the Fare and Time of arrival for any
//given cab type.
type Cab interface {
	Init()
	GetFare(a, b string) *Fares
	GetTime(a string) *Times
}

//PrintFares pretty-prints the fares for the given 'from' and 'to'
//address.
func (f *Fares) PrintFares() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAlignment(1)
	data := [][]string{
		[]string{"From:", strings.ToUpper(f.From)},
		[]string{"To:", strings.ToUpper(f.To)},
	}
	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Est", "Surge", "Distance", "Duration"})
	for _, val := range f.Fare {
		table.Append([]string{val.Name, val.Est, strconv.FormatFloat(val.Surge, 'f', -1, 64), strconv.FormatFloat(val.Dist*1.60934, 'f', 2, 64) + " km", strconv.FormatFloat(float64(val.Duration)/60.0, 'f', -1, 64) + " min"})
	}
	table.Render()
}

//PrintTimes pretty-prints the time of arrival for the given 'location'.
func (t *Times) PrintTimes() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAlignment(1)

	table.Append([]string{"Location", t.Loc})
	table.Render()

	table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "ETA"})
	for _, val := range t.Time {
		table.Append([]string{val.Name, strconv.FormatFloat(float64(val.Est)/60.0, 'f', -1, 64) + " min"})
	}
	table.Render()
}
