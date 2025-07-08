@agent_information_query
Feature: AgentInformationQuery
  As an HR staff
  I want to query agent complete information
  So that I can verify and review agent data accuracy

  Scenario: HR queries agent complete information
    Given scenario "AgentManualRegistration: HR manually creates new agent profile"
    When HR queries agent 'A001' information
    Then the system returns:
      | agentId | name   | joinDate   | idNumber   | position | status |
      | A001    | Johnny | 2024-01-15 | A123456789 | 技術駭客 | Active |
    And includes emergency contact:
      | name | relationship | phone      |
      | Jane | Sister       | 0987654321 |
    And includes documents:
      | type            | status   |
      | profilePhoto    | Uploaded |
      | idCardFront     | Uploaded |
      | idCardBack      | Uploaded | 