package main

import (
	"fmt"
	"encoding/json"
)
// how to make the json data in go lang
type Course struct{
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Platform string `json:"platform"`
	Tags     []string `json:"tags,omitempty"`
}

func main(){
	EncodeJsonData()
	fmt.Println("JSON data created")
}

func EncodeJsonData(){
	lcoCourses := []Course{
		{"ReactJS Bootcamp", 299, "learncodeonline.in", []string{"web-dev" , "js"}},
		{"Angular Bootcamp", 199, "learncodeonline.in", []string{"web-dev" , "js"}},
		{"VueJS Bootcamp", 299, "learncodeonline.in", nil},
	}

	finalJson , err := json.MarshalIndent(lcoCourses , "" , "\t")
	if(err!= nil){
		panic(err)
	}
	fmt.Println("Final JSON data:")
	fmt.Println(string(finalJson))
}

func DecodeJsonData(){
	jsonDataFromWeb := []byte(`
	{
		"name": "ReactJS Bootcamp",		
}

		"price": 299,
		"platform": "learncodeonline.in",
		"tags": ["web-dev" , "js"]
	}
	`)				

	var courseObj Course
	checkValid := json.Valid(jsonDataFromWeb)		
	if(checkValid){
		fmt.Println("JSON data is valid")
		json.Unmarshal(jsonDataFromWeb , &courseObj)
		fmt.Printf("%#v\n" , courseObj)
	}
	else{
		fmt.Println("JSON data is not valid")
	}		
}					
