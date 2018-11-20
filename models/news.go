package models

import "fmt"

func retrieveNews(news string) string {
	var statement = fmt.Sprintf("retrieve this ", news)
	return statement
}
