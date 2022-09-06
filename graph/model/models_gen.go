// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Repo struct {
	ID              int      `json:"id"`
	NodeID          string   `json:"nodeID"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	FullName        string   `json:"fullName"`
	HTMLURL         string   `json:"htmlUrl"`
	URL             string   `json:"url"`
	StargazersCount int      `json:"stargazersCount"`
	WatchersCount   int      `json:"watchersCount"`
	Visibility      string   `json:"visibility"`
	ImageURL        string   `json:"imageUrl"`
	Languages       []string `json:"languages"`
}
