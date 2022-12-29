package util

import (
	"encoding/json"
	"fmt"
	"log"
)

func Pretty(obj interface{}) string {
    b, err := json.MarshalIndent(obj, "", "  ")
    if err != nil {
        log.Fatal(err)
    }
    return string(b)
}

func PrettyPrint(obj interface{}) {
    fmt.Println(Pretty(obj))
}
