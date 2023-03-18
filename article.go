package main

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Declare a global articles array that we can then populate in our main
// function to simulate a database

var articles = []Article{
	{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}
