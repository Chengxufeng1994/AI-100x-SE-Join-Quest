@attendance_record_adjustment
Feature: AttendanceRecordAdjustment
  As an HR staff
  I want to adjust agent attendance records
  So that attendance data can be corrected when necessary with proper audit trail

  Scenario: HR adjusts agent attendance record
    Given scenario "AgentPunchClock: Agent punches out after working"
    When HR adjusts attendance for agent 'A001' on '2024-01-16':
      | field     | oldValue            | newValue            | reason            | operator |
      | punchIn   | 2024-01-16T09:00:00 | 2024-01-16T08:30:00 | Forgot to punch   | HR001    |
    Then the attendance record should be updated:
      | agentId | date       | punchIn             | punchOut            | workHours | status   |
      | A001    | 2024-01-16 | 2024-01-16T08:30:00 | 2024-01-16T18:00:00 | 9.5       | Adjusted |
    And audit trail should record:
      | agentId | date       | field   | oldValue            | newValue            | operator | reason          |
      | A001    | 2024-01-16 | punchIn | 2024-01-16T09:00:00 | 2024-01-16T08:30:00 | HR001    | Forgot to punch | 