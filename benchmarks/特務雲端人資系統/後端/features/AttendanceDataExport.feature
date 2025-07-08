@attendance_data_export
Feature: AttendanceDataExport
  As an HR staff
  I want to export attendance data with filters
  So that I can generate attendance reports for analysis and compliance

  Scenario: HR exports attendance data with filters
    Given scenario "AttendanceRecordAdjustment: HR adjusts agent attendance record"
    When HR exports attendance data with:
      | startDate  | endDate    | department | status   |
      | 2024-01-16 | 2024-01-16 | 工程部     | Adjusted |
    Then the export should contain:
      | agentId | name   | date       | punchIn             | punchOut            | workHours | status   |
      | A001    | Johnny | 2024-01-16 | 2024-01-16T08:30:00 | 2024-01-16T18:00:00 | 9.5       | Adjusted | 