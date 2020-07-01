// The Go standard library comes with excellent support
// for HTTP clients and servers in the `net/http`
// package. In this example we'll use it to issue simple
// HTTP requests.
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	//"time"
	"regexp"
	"os"
	"strconv"
)

func main() {

	count, _ := strconv.Atoi(os.Args[2])
	uri := "http://" + os.Args[1]

	var responseSlice []*http.Response

	for i := 0; i <= count - 1; i += 1 {

		// time.Sleep(time.Duration(i*30) * time.Second)

		resp, err := http.Get(uri)
		if err != nil {
			panic(err)
		}

		responseSlice = append(responseSlice, resp)

		printBody(responseSlice[i])

	}

	for i := 0; i <= count - 1; i += 1 {

		responseSlice[i].Body.Close()

	}

}


func printBody(response *http.Response) {

	rip, _ := regexp.Compile("<p>Server IP address: (.+)</p>")
	rtime, _ := regexp.Compile("<p>Server local time: (.+)</p>")

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
                panic(err)
        }

	bodyString := string(bodyBytes)

	if response.StatusCode == 200 {
		fmt.Println("Response status:", response.Status, "Server IP: ",rip.FindStringSubmatch(bodyString)[1]
,"\t","Time: ",rtime.FindStringSubmatch(bodyString)[1])
	} else {
		fmt.Println("Response status:", response.Status)
	}
}
