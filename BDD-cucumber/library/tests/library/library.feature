Feature: User and Article Management

  Background:
    Given the database is empty
    And an admin is authenticated

  # User Management
  Scenario: A user is successfully registered
    Given the admin navigates to the "registration" page
    When the admin enters valid "name", "email", and "password"
    And the admin presses "Register"
    Then a message "User successfully registered" is received

  Scenario Outline: User registration fails due to invalid data
    Given the admin navigates to the "registration" page
    When the admin enters "<name>", "<email>", and "<password>"
    And the admin presses "Register"
    Then an error message "<error_message>" is received

    Examples:
      | name  | email    | password | error_message         |
      | ""    | a@a.com  | 1234     | "Name is required"    |
      | Alice | ""       | 1234     | "Email is required"   |
      | Alice | a@a.com  | ""       | "Password is required"|

    # Article Management
  Scenario: An article is successfully published
    Given the admin navigates to the "new article" page
    When the admin enters valid "title" and "description"
    And the admin presses "Publish"
    Then a message "Article successfully published" is received

  Scenario: An existing article is successfully updated
    Given the following existing articles
      | title      | description |
      | Article 1  | Desc 1      |
    When the admin edits the article with the title "Article 1"
    And the admin updates its description to "New Desc"
    Then the article with the title "Article 1" should have the description "New Desc"

  Scenario: An existing article is successfully deleted
    Given the following existing articles
      | title      | description |
      | Article 1  | Desc 1      |
    When the admin deletes the article with the title "Article 1"
    Then the article with the title "Article 1" should not exist