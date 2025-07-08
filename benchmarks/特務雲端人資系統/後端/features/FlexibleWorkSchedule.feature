@flexible_work_schedule
Feature: FlexibleWorkSchedule
  As an Agent
  I want to set and modify my flexible work schedule
  So that I can enjoy flexible working hours while complying with labor regulations

  Scenario: Agent sets flexible work schedule
    Given scenario "AgentManualRegistration: HR manually creates new agent profile"
    And today is '2024-01-16'
    When agent 'A001' sets work schedule for '2024-01-16':
      | workStart | workEnd | breakStart | breakEnd | totalHours |
      | 09:30     | 18:30   | 12:00      | 13:00    | 8.0        |
    Then the work schedule should be saved:
      | agentId | date       | workStart | workEnd | breakStart | breakEnd | status |
      | A001    | 2024-01-16 | 09:30     | 18:30   | 12:00      | 13:00    | Active |
    And daily work plan should comply with:
      | maxDailyHours | maxWeeklyHours | minBreakAfterHours |
      | 8.0           | 40.0           | 4.0                |

  Scenario: Agent modifies next day schedule before 10AM
    Given scenario "FlexibleWorkSchedule: Agent sets flexible work schedule"
    And current time is '2024-01-17T08:30:00'
    When agent 'A001' modifies work schedule for '2024-01-16':
      | workStart | workEnd | breakStart | breakEnd | totalHours |
      | 10:00     | 19:00   | 12:30      | 13:30    | 8.0        |
    Then the work schedule should be updated:
      | agentId | date       | workStart | workEnd | status   |
      | A001    | 2024-01-16 | 10:00     | 19:00   | Modified |
    And modification history should record:
      | agentId | date       | field     | oldValue | newValue | modifiedAt          |
      | A001    | 2024-01-16 | workStart | 09:30    | 10:00    | 2024-01-17T08:30:00 |

  Scenario Outline: Agent fails to set invalid work schedule
    Given scenario "AgentManualRegistration: HR manually creates new agent profile"
    And today is '2024-01-16'
    When agent 'A001' sets work schedule for '2024-01-16':
      | workStart   | workEnd   | breakStart   | breakEnd   |
      | <workStart> | <workEnd> | <breakStart> | <breakEnd> |
    Then the schedule setting fails with:
      | reason   |
      | <reason> |

    Examples:
      | workStart | workEnd | breakStart | breakEnd | reason                              |
      | 09:00     | 18:00   | 12:00      | 12:00    | Break duration must be at least 30 minutes |
      | 09:00     | 18:30   | 12:00      | 13:00    | Daily work hours exceed 8 hours limit      |
      | 10:00     | 14:00   |            |          | Break required for 4+ hour shifts         |
      | 20:00     | 04:00   | 23:00      | 00:00    | Invalid time range                         | 