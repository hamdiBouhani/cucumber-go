package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/cucumber/godog"
)

type calculator struct {
	result int

	response *http.Response
	body     []byte
}

func (c *calculator) iHaveACalculator() error {
	// Initialize the calculator
	return nil
}

func (c *calculator) iAddAnd(arg1, arg2 int) error {
	c.result = arg1 + arg2
	return nil
}

func (c *calculator) theResultShouldBe(expected int) error {
	if c.result != expected {
		return godog.ErrPending
	}
	return nil
}

func (h *calculator) iSendAPOSTRequestToWithBody(path string, body *godog.DocString) error {
	url := "http://localhost:8080" + path
	fmt.Println("Sending POST request to:", url) // Add this line

	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(body.Content)))
	if err != nil {
		return fmt.Errorf("failed to send POST request: %v", err)
	}

	h.response = resp
	h.body, err = io.ReadAll(resp.Body)
	return err
}
func (h *calculator) theResponseStatusCodeShouldBe(statusCode int) error {
	if h.response.StatusCode != statusCode {
		return fmt.Errorf("expected status code %d, but got %d", statusCode, h.response.StatusCode)
	}
	return nil
}

func (h *calculator) theResponseBodyShouldBe(expectedBody *godog.DocString) error {
	var expected, actual map[string]interface{}

	// Parse the expected response body
	if err := json.Unmarshal([]byte(expectedBody.Content), &expected); err != nil {
		return fmt.Errorf("failed to parse expected response body: %v", err)
	}

	// Parse the actual response body
	if err := json.Unmarshal(h.body, &actual); err != nil {
		return fmt.Errorf("failed to parse actual response body: %v", err)
	}

	// Compare the response bodies
	if fmt.Sprint(actual) != fmt.Sprint(expected) {
		return fmt.Errorf("expected response body %v, but got %v", expected, actual)
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	c := &calculator{}

	ctx.Step(`^I have a calculator$`, c.iHaveACalculator)
	ctx.Step(`^I add (\d+) and (\d+)$`, c.iAddAnd)
	ctx.Step(`^the result should be (\d+)$`, c.theResultShouldBe)

	ctx.Step(`^I send a POST request to "([^"]*)" with body:$`, c.iSendAPOSTRequestToWithBody)
	ctx.Step(`^the response status code should be (\d+)$`, c.theResponseStatusCodeShouldBe)
	ctx.Step(`^the response body should be:$`, c.theResponseBodyShouldBe)
}

func TestFeatures(t *testing.T) {

	dir, _ := os.Getwd()
	t.Logf("Current working directory: %s", dir)

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{fmt.Sprintf("%s", dir)},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
