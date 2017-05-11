

package main    

import (
  "fmt"
  "net/http"    
  "io/ioutil"  
  "time"
)               

func main() {
	rails_time:=time_page_load("https://rails-hku-demo.herokuapp.com/")
	fmt.Println("rails response time: ", rails_time)   
	golang_time:=time_page_load("https://gohkudemo.herokuapp.com/")
	fmt.Println("golang response time: ", golang_time)  
	}                            

func time_page_load (url string) int64{
	
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
	
	return diff
}
