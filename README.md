# cucumber-go

In this project we use `godog` which is a Golang implementation of Cucumber, a popular tool for Behavior-Driven Development (BDD).


## Installation:

```bash
go get github.com/cucumber/godog/cmd/godog
```

## Example Usage:

```
Feature: Example feature
  Scenario: Add two numbers
    Given I have 5 and 7
    When I add them
    Then the result should be 12
```
```go
package main

import (
	"github.com/cucumber/godog"
)

func iHaveAnd(arg1, arg2 int) error {
	// Store the numbers in a context or struct
	return nil
}

func iAddThem() error {
	// Perform the addition
	return nil
}

func theResultShouldBe(arg1 int) error {
	// Assert the result
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I have (\d+) and (\d+)$`, iHaveAnd)
	ctx.Step(`^I add them$`, iAddThem)
	ctx.Step(`^the result should be (\d+)$`, theResultShouldBe)
}
```
## Run the Tests:
```bash
godog
```