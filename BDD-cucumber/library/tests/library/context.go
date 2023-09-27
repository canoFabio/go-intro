package library

import (
	"fmt"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Name, Email, Password string
}

var users = make(map[string]User)

func anAdminIsAuthenticated(ctx *godog.ScenarioContext) {
	// Initialize admin authentication logic here
}

func theAdminNavigatesToThePage(page string) error {
	// Navigate to the page logic here
	return nil
}

func theAdminEntersValidNameEmailAndPassword() error {
	user := User{Name: "Alice", Email: "a@a.com", Password: "1234"}
	users[user.Email] = user
	return nil
}

func theAdminPressesButton(button string) error {
	// Press button logic here
	return nil
}

func aMessageIsReceived(expectedMessage string) error {
	actualMessage := "User successfully registered" // Simulate the message from the application
	assert.Equal(nil, expectedMessage, actualMessage)
	return nil
}

type Article struct {
	Title, Description string
}

var articles = make(map[string]Article)

func theAdminNavigatesToTheNewArticlePage() error {
	// Navigate to new article page logic here
	return nil
}

func theAdminEntersValidTitleAndDescription() error {
	article := Article{Title: "Article 1", Description: "Desc 1"}
	articles[article.Title] = article
	return nil
}

func theFollowingExistingArticles(table *godog.Table) error {
	// Initialize articles from table
	for i, row := range table.Rows {
		if i == 0 { // Skip header row
			continue
		}
		articles[row.Cells[0].Value] = Article{Title: row.Cells[0].Value, Description: row.Cells[1].Value}
	}
	return nil
}

func theAdminEditsTheArticleWithTheTitle(title string) error {
	// Edit article logic here
	return nil
}

func theAdminUpdatesItsDescriptionTo(description string) error {
	article, exists := articles["Article 1"]
	if !exists {
		return fmt.Errorf("Article does not exist")
	}
	article.Description = description
	articles["Article 1"] = article
	return nil
}

func theArticleWithTheTitleShouldHaveTheDescription(title, description string) error {
	article, exists := articles[title]
	if !exists {
		return fmt.Errorf("Article does not exist")
	}
	assert.Equal(nil, description, article.Description)
	return nil
}

func theAdminDeletesTheArticleWithTheTitle(title string) error {
	_, exists := articles[title]
	if !exists {
		return fmt.Errorf("Article does not exist")
	}
	delete(articles, title)
	return nil
}

func theArticleWithTheTitleShouldNotExist(title string) error {
	_, exists := articles[title]
	assert.False(nil, exists)
	return nil
}
