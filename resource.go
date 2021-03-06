package main

type (
	Field struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	CustomField struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Multiple bool   `json:"multiple"`
		Value    string `json:"value"`
	}
)
