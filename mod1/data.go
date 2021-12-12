package main

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type post struct {
	Sql string `json:"sql"`
}
