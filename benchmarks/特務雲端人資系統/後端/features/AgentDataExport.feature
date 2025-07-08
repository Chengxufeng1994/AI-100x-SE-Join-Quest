@agent_data_export
Feature: AgentDataExport
  As an HR staff
  I want to export agent data with filters
  So that I can generate reports and analyze agent information

  Scenario: HR exports agent data
    Given scenario "AgentBatchRegistration: HR uploads Excel employee data"
    When HR exports agent data with filters:
      | department | status |
      | 工程部     | Active |
    Then the export should contain:
      | agentId | name   | department | status |
      | A001    | Johnny | 工程部     | Active | 