Feature: Calculator

  Scenario: Add two numbers
    Given I have a calculator
    When I add 2 and 3
    Then the result should be 5

  Scenario: Add two numbers via HTTP
    Given I send a POST request to "/add" with body:
      """
      {
        "a": 2,
        "b": 3
      }
      """
    Then the response status code should be 200
    And the response body should be:
      """
      {
        "result": 5
      }
      """