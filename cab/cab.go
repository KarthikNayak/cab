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

type Fare struct {
	Currency string
	Name     string
	Est      string
	Surge    float64
	Duration int
	Dist     float64
}

type Fares struct {
	Fare []Fare
	From string
	To   string
}

type Time struct {
	Name string
	Est  int
}

type Times struct {
	Time []Time
	Loc  string
}

type Cab interface {
	Init()
	GetFare(a, b string) *Fares
	GetTime(a string) *Times
}

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
