package models


type Book struct{
	Id string   `json:"id"`         // this json tag tel go that when converting struct to json use this name and vice versa
	Title string  `json:"title"` 
	Author string   `json:"author"`    // name start with capital becuase we must  make it public and exportable because we can use it in other package
	Quantity  int    `json:"quantity"`
}