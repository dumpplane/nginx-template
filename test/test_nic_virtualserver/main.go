package main

import (
    "fmt"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    conf_v1 "github.com/nginxinc/kubernetes-ingress/pkg/apis/configuration/v1"
    "encoding/json"
)


func main() {

    
    typeMeta := metav1.TypeMeta{
        APIVersion: "k8s.nginx.org/v1",
        Kind: "VirtualServer",
    }

    objectMeta := metav1.ObjectMeta{
        Name: "test",
        Namespace: "test",
    }

    tea := conf_v1.Upstream{
        Name: "tea",
        Service: "tea-svc",
        Port: 80,
    }

    coffee := conf_v1.Upstream{
        Name: "coffee",
        Service: "coffee-svc",
        Port: 80,
    }
 
    action_tea := conf_v1.Action{
        Pass: "tea",
    }    

    action_coffee := conf_v1.Action{
        Pass: "coffee",
    }
 
    route_tea := conf_v1.Route{
        Path: "/tea",
        Action: &action_tea,
    }

    route_coffee := conf_v1.Route{
        Path: "/coffee",
        Action: &action_coffee,
    }

    spec := conf_v1.VirtualServerSpec{
        IngressClass: "nginx",
        Host: "test.example.com",
        Upstreams: []conf_v1.Upstream{tea, coffee},
        Routes: []conf_v1.Route{route_tea, route_coffee},
    }

    vs := conf_v1.VirtualServer{
        TypeMeta: typeMeta,
        ObjectMeta: objectMeta,
        Spec: spec,
    }


    jsonData, err := json.Marshal(vs)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(string(jsonData))
}
