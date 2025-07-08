@agent_batch_registration
Feature: AgentBatchRegistration
  As an HR staff
  I want to upload Excel files to batch register multiple agents
  So that I can efficiently process large numbers of new employees

  Scenario: HR uploads Excel employee data
    Given no existing agent records
    When HR uploads employee Excel file:
      | fileName         | records |
      | employees.xlsx   | 3       |
    And the Excel contains:
      | name  | joinDate   | idNumber   | birthDate  | gender | position | monthlySalary |
      | Alice | 2024-01-01 | B123456789 | 1985-05-15 | F      | 成長駭客 | 75000         |
      | Bob   | 2024-01-10 | C123456789 | 1992-03-20 | M      | 客服     | 45000         |
      | Carol | 2024-01-20 | D123456789 | 1988-12-10 | F      | 美編     | 50000         |
    And assigns positions and managers:
      | name  | position | manager |
      | Alice | 成長駭客 | A000    |
      | Bob   | 客服     | A001    |
      | Carol | 美編     | A001    |
    Then 3 agents should be created with status 'Pending'
    And agent IDs should be auto-generated as:
      | name  | agentId |
      | Alice | A002    |
      | Bob   | A003    |
      | Carol | A004    | 