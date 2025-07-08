@attendance_history_query
Feature: AttendanceHistoryQuery
  As an Agent
  I want to view my attendance history
  So that I can track my work hours and attendance records

  Scenario: Agent views attendance history
    Given scenario "AttendanceRecordAdjustment: HR adjusts agent attendance record"
    When agent 'A001' queries attendance history for:
      | startDate  | endDate    |
      | 2024-01-15 | 2024-01-17 |
    Then the system returns attendance records:
      | date       | punchIn             | punchOut            | workHours | status   |
      | 2024-01-16 | 2024-01-16T08:30:00 | 2024-01-16T18:00:00 | 9.5       | Adjusted | 