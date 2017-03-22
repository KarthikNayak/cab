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

var from, to string

// fareCmd represents the fare command
var fareCmd = &cobra.Command{
	Use:   "fare",
	Short: "Check current cab fares",
	Long:  `Check current cab fares.`,
	Run: func(cmd *cobra.Command, args []string) {
		if from == "" {
			fmt.Println("No start location provided")
			return
		}
		if to == "" {
			fmt.Println("No end location provided")
			return
		}

		var cabs []cab.Cab
		cabs = append(cabs, &uber.Uber{})

		for _, val := range cabs {
			val.Init()
			fare := val.GetFare(from, to)
			fare.PrintFares()
		}
	},
}

func init() {
	RootCmd.AddCommand(fareCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fareCmd.PersistentFlags().String("foo", "", "A help for foo")
	fareCmd.PersistentFlags().StringVarP(&from, "from", "f", "", "Start location for the cab")
	fareCmd.PersistentFlags().StringVarP(&to, "to", "t", "", "End location for the cab")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
