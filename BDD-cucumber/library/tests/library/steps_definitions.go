package library

import "github.com/cucumber/godog"

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(*godog.Scenario) {
		users = make(map[string]User) // Initialize the in-memory map before each scenario
	})

	ctx.Step(`^an admin is authenticated$`, anAdminIsAuthenticated)
	ctx.Step(`^the admin navigates to the "([^"]*)"$`, theAdminNavigatesToThePage)
	ctx.Step(`^the admin enters valid "([^"]*)", "([^"]*)", and "([^"]*)"$`, theAdminEntersValidNameEmailAndPassword)
	ctx.Step(`^the admin presses "([^"]*)"$`, theAdminPressesButton)
	ctx.Step(`^a message "([^"]*)" is received$`, aMessageIsReceived)

	// Para Article Management
	ctx.Step(`^the admin navigates to the "new article" page$`, theAdminNavigatesToTheNewArticlePage)
	ctx.Step(`^the admin enters valid "([^"]*)" and "([^"]*)"$`, theAdminEntersValidTitleAndDescription)
	ctx.Step(`^the following existing articles$`, theFollowingExistingArticles)
	ctx.Step(`^the admin edits the article with the title "([^"]*)"$`, theAdminEditsTheArticleWithTheTitle)
	ctx.Step(`^the admin updates its description to "([^"]*)"$`, theAdminUpdatesItsDescriptionTo)
	ctx.Step(`^the article with the title "([^"]*)" should have the description "([^"]*)"$`, theArticleWithTheTitleShouldHaveTheDescription)
	ctx.Step(`^the admin deletes the article with the title "([^"]*)"$`, theAdminDeletesTheArticleWithTheTitle)
	ctx.Step(`^the article with the title "([^"]*)" should not exist$`, theArticleWithTheTitleShouldNotExist)
}
