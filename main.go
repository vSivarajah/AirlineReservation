/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vSivarajah/AirlineReservation/cmd"
)

func main() {
	cmd.Execute()
	//app.StartApplication()

	goodMorningHandler := func(w http.ResponseWriter, req *http.Request) {
		
		_, err := fmt.Fprintf(w, "Good morning! You are using verb: %s, URL: %s\n", req.Method, req.URL)
		if err != nil {
			fmt.Println(http.StatusBadRequest)
		}else{
			fmt.Println(http.StatusOK)
		}
	}
	http.HandleFunc("/goodmorning", goodMorningHandler)
	log.Fatal(http.ListenAndServe(":1234", nil))

}
