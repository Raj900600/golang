
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
	checkLink(link,c)
}
for i:=0; i<len(links);i++{
	
fmt.Println(<-c)
}
}
func checkLink(link string,c chan string)  {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link,"might be down")
		c <- "Might be down"
		return 

	}
	fmt.Println(link,"is up!")
	c <- "yes is up"
}