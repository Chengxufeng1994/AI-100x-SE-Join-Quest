@attendance_exception_alert
Feature: AttendanceExceptionAlert
  As an HR staff
  I want to view attendance exception alerts
  So that I can monitor and address potential labor regulation violations

  Scenario: System shows attendance exception alerts
    Given scenario "AgentManualRegistration: HR manually creates new agent profile"
    And agent 'A001' works continuously for:
      | date       | workStart | workEnd | breakHours |
      | 2024-01-16 | 08:00     | 20:00   | 0.0        |
    When HR views attendance exceptions
    Then the system shows alerts:
      | agentId | date       | alertType              | description                |
      | A001    | 2024-01-16 | Excessive working hours | 12 hours without break    | 