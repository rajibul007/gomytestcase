package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {

		fmt.Println("error occured ", err)
		os.Exit(1)

	}
	/*defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(resp)
	fmt.Println(reflect.TypeOf(resp))
	fmt.Println(string(body))
	fmt.Println(reflect.TypeOf(body))
	*/

	//bs := make([]byte, 9999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))
	io.Copy(os.Stdout, resp.Body)

}
