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

package cmd

import (
	"fmt"

	"github.com/karthiknayak/cab/cab"
	"github.com/karthiknayak/cab/uber"
	"github.com/spf13/cobra"
)

var loc string

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Check time required for cab arrival",
	Long:  `Check time required for cab arrival.`,
	Run: func(cmd *cobra.Command, args []string) {
		if loc == "" {
			fmt.Println("No location provided")
			return
		}

		var cabs []cab.Cab
		cabs = append(cabs, &uber.Uber{})

		for _, val := range cabs {
			val.Init()
			time := val.GetTime(loc)
			time.PrintTimes()
		}
	},
}

func init() {
	RootCmd.AddCommand(timeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	timeCmd.PersistentFlags().StringVarP(&loc, "loc", "l", "", "Location for getting cab ETA")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
