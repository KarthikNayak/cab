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

package geocode

import (
	"context"
	"log"

	"googlemaps.github.io/maps"
)

var (
	key = "AIzaSyC9bUXljlN24C4fywUls1MUq0xPYKzv3hM"
)

//GetCode returns the Latitude and Longitude values given an address
//string. This is done using Google's geocode API.
func GetCode(address string) maps.LatLng {
	conn, err := maps.NewClient(maps.WithAPIKey(key))
	if err != nil {
		log.Fatal("Couldn't initiate connection with Google Geocode API", err)
	}

	req := &maps.GeocodingRequest{
		Address: address,
	}

	res, err := conn.Geocode(context.Background(), req)
	if err != nil {
		log.Fatal("Couldn't find any location: " + address)
	}

	return res[0].Geometry.Location
}
