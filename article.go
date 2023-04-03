package restserver

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type Article struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Content     string `json:"content"`
}
