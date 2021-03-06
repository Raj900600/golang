
package main

import (
	"fmt"
    "net/http"
)

func main() {
links := [] string{
"http://google.com",
"http://facebook.com",
"http://stackverflow.com",
"http://golang.com",
"http://amazon.com",
}

c := make(chan string)
for _,link := range links {
	go checkLink(link,c)
}
for {
	go checkLink(<-c,c)
//fmt.Println(<-c)
}
}
func checkLink(link string,c chan string)  {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link,"might be down")
		c <- link
		return 

	}
	fmt.Println(link,"is up!")
	c <- link
}