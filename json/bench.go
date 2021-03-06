package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	content, _ := ioutil.ReadFile(os.Args[1])
	N := int(1e5)
	beg := time.Now()
	for i := 0; i < N; i++ {
		var v interface{}
		_ = json.Unmarshal(content, &v)
	}
	end := time.Now()
	fmt.Printf("json.Unmarshal: %s\n", end.Sub(beg))
}
