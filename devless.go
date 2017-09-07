package main

import (
	"net/http"
	"io/ioutil" 
	"bytes"
	"fmt"
	"strings"
	"encoding/json"
	"strconv"
)

type Devless struct {
	url, devlessToken string
}


func main() {
	var init = Devless{"http://localhost:8080","0cc39661d0556719cedbda4bfa463ac1"}
	
	//get data 
	// results := GetData(init,"contacts", "contact_tabler",
	//  [] string {param("greaterThan", "id,5")})
	
	//  adding data 
	// type record struct {
	// 	name string 
	// }

	// data := record{name:"edmond"}

	// results := PostData(init, "contacts", "contact_tabler", data)
	
	// //updating data 
	// type Record struct {
	// 	Name string 
	// }

	// data := Record{Name:"kofi"}

	// results := UpdateData(init, "contacts", "contact_tabler", 9, data)
	

	// updateData();
	// postData()
	// results := DeleteData(init, "contacts", "contact_tabler", 9)
	// var params  []interface{}
	// params = append(params, 1)
	// results := call(init, "devless", "getUserProfile", params)

	// fmt.Printf(results)
	// call()
}

func GetData(d Devless, service string, table string, params []string ) string {
	url := d.url+"/api/v1/service/"+service+"/db?table="+table
	for _, param := range params { url += param }
	return requestProcessor(url, d.devlessToken, "GET", `{}`)
}

func param(name string, value string) string {
	return "&"+name+"="+value
}

func PostData(d Devless, service string, table string, data interface{}) string{
	url := d.url+"/api/v1/service/"+service+"/db"
	jsonData, _ := json.Marshal(data)
	stringData := string(jsonData)
    var payload = (`{
	  "resource":[{
	    "name":"`+table+`",
	    "field":[`+stringData+`]
	  }]
	}`)
    return requestProcessor(url,d.devlessToken, "POST", payload)
}

func UpdateData(d Devless, service string, table string, id int, data interface{}) string{
	url := d.url+"/api/v1/service/"+service+"/db"
	jsonData, _ := json.Marshal(data)
	stringData := string(jsonData)
	var payload = (`{
	  "resource":[{
	    "name":"`+table+`",
	    "params":[{
	      "where":"id,contact_tabler",
	      "data":[`+stringData+`]
	    }]
	  }]
	}`)
    
    return requestProcessor(url,d.devlessToken, "PATCH", payload)
}

func DeleteData(d Devless, service string, table string, id int) string {
	url := d.url+"/api/v1/service/"+service+"/db"
    var payload = (`{  
   "resource":[  
      {  
         "name":"`+table+`",
         "params":[  
            {  
               "delete":"true",         
               "where":"id,=,`+strconv.Itoa(id)+`"

            }
         ]
      }

    ]
}   `)
    return requestProcessor(url, d.devlessToken, "DELETE", payload)
}




func call(d Devless, service string, method string, params []interface{} ) string {
	url := d.url+"/api/v1/service/"+service+"/rpc?action="+method
    jsonData, _ := json.Marshal(params)
    stringData := string(jsonData)
    fmt.Printf(stringData)
    var payload = (`{
    "jsonrpc": "2.0",
    "method": "`+service+`",
    "id": "1000",
    "params": `+stringData+`
	}`)

    return requestProcessor(url, d.devlessToken, "POST", payload)
} 


func requestProcessor(url string,devlessToken string, method string, payload string) string {
	
    var jsonStr = []byte(payload)
    req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("devless-token", devlessToken)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    return string(body)
}  


//docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:latest go build -v

