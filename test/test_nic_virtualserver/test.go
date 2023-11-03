package main

import (
    "fmt"
    "io/ioutil"
    conf_v1 "github.com/nginxinc/kubernetes-ingress/pkg/apis/configuration/v1"
    "reflect"
    "encoding/json"
)

func listStructKeys(s interface{}) {
	valueOfS := reflect.ValueOf(s)

	if valueOfS.Kind() != reflect.Struct {
		fmt.Println("Input is not a struct.")
		return
	}

	typeOfS := valueOfS.Type()

	for i := 0; i < valueOfS.NumField(); i++ {
		field := typeOfS.Field(i)
		fmt.Println("Key:", field.Name)
	}
}

func main() {

    yamlFile, err := ioutil.ReadFile("vs.yaml")
    if err != nil {
        fmt.Println("Error reading YAML file:", err)
        return
    }

    fmt.Println(string(yamlFile))

    jsonData, err := json.Marshal(yamlFile)
    if err != nil {
        fmt.Println("Error marshaling JSON:", err)
        return
    }

    fmt.Println("reflect.TypeOf(yamlFile): %s ", reflect.TypeOf(yamlFile))
    fmt.Println("reflect.TypeOf(jsonData): %s ", reflect.TypeOf(jsonData))
   
    var vs conf_v1.VirtualServer

    err = json.Unmarshal([]byte(jsonData), &vs)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    listStructKeys(vs) 
    fmt.Printf("%v\n", vs.TypeMeta)
    fmt.Printf("%v\n", vs.ObjectMeta)
    fmt.Printf("%v\n", vs.Spec)
    fmt.Printf("%v\n", vs.Status)
}
