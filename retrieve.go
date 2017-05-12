package main    

import (
  "fmt"
  "net/http"    
  "io/ioutil"  
  "time"
)               

func time_page_load (url string, c chan string) {
	
	t1 := time.Now()
	unixNano := t1.UnixNano()                                                                      
    umillisec := unixNano / 1000000                                                                 
   
	resp, err := http.Get(url)  											
	fmt.Println("http transport error is:", err)
	body, err := ioutil.ReadAll(resp.Body)  
	fmt.Println("read error is:", err)

	fmt.Println(string(body)[0:100]) 
	
	t2 := time.Now()
	unixNano2 := t2.UnixNano()                                                                      
    umillisec2 := unixNano2 / 1000000  
                                                                    
	diff:=umillisec2 - umillisec
	fmt.Println("\nRetrieve time was: ",diff, " Milliseconds.\n")
	
	c<-fmt.Sprintf("\n%s - time: %d\n",url,diff)
}

func printer(c chan string) {
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}


func main() {
	
	var c chan string = make(chan string)
	
	go time_page_load("https://rails-hku-demo.herokuapp.com/", c)
  
	go time_page_load("https://gohkudemo.herokuapp.com/", c)
 
	go printer(c)
	
	var input string
	fmt.Scanln(&input)
}                            



