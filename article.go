package main

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type Article struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Content     string `json:"content"`
}
