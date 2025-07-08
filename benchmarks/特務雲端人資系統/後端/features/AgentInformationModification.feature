@agent_information_modification
Feature: AgentInformationModification
  As an HR staff
  I want to modify agent information with audit trail
  So that agent data can be kept accurate and all changes are tracked

  Scenario: HR modifies agent information with audit trail
    Given scenario "AgentManualRegistration: HR manually creates new agent profile"
    When HR modifies agent 'A001' with:
      | field         | oldValue   | newValue   |
      | mobilePhone   | 0912345678 | 0987654321 |
      | monthlySalary | 80000      | 85000      |
    And the modification is performed by:
      | operatorId | operatorName |
      | HR001      | HR Manager   |
    Then the agent 'A001' should be updated with:
      | mobilePhone | monthlySalary |
      | 0987654321  | 85000         |
    And an audit log should be created:
      | agentId | field         | oldValue   | newValue   | operator   | timestamp          |
      | A001    | mobilePhone   | 0912345678 | 0987654321 | HR Manager | 2024-01-16T10:30:00 |
      | A001    | monthlySalary | 80000      | 85000      | HR Manager | 2024-01-16T10:30:00 | 