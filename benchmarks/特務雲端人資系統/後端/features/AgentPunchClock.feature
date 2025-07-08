@agent_punch_clock
Feature: AgentPunchClock
  身為特務，我要在上下班時各做一次打卡紀錄，這樣才能證明我的出勤起訖時間。

  Scenario: Agent punches in for the first time
    Given scenario AgentManualRegistration: HR manually creates new agent profile
    And today is '2024-01-16'
    And agent 'A001' has no punch records for today
    When agent 'A001' punches in with:
      | timestamp           | location    | ipAddress     | device     |
      | 2024-01-16T09:00:00 | Main Office | 192.168.1.100 | Desktop-01 |
    Then the punch record should be created:
      | agentId | date       | punchType | timestamp           | location    | status |
      | A001    | 2024-01-16 | IN        | 2024-01-16T09:00:00 | Main Office | Valid  |
    And agent 'A001' work status should be:
      | status    | lastPunch           |
      | Punched In| 2024-01-16T09:00:00 |

  Scenario: Agent punches out after working
    Given scenario AgentPunchClock: Agent punches in for the first time
    When agent 'A001' punches out with:
      | timestamp           | location    | ipAddress     | device     |
      | 2024-01-16T18:00:00 | Main Office | 192.168.1.100 | Desktop-01 |
    Then the punch record should be created:
      | agentId | date       | punchType | timestamp           | location    | status |
      | A001    | 2024-01-16 | OUT       | 2024-01-16T18:00:00 | Main Office | Valid  |
    And agent 'A001' should have daily attendance:
      | date       | punchIn             | punchOut            | workHours |
      | 2024-01-16 | 2024-01-16T09:00:00 | 2024-01-16T18:00:00 | 9.0       |

  Scenario: Agent attempts duplicate punch in
    Given scenario "AgentPunchClock: Agent punches in for the first time"
    When agent 'A001' attempts to punch in again with:
      | timestamp           | location    |
      | 2024-01-16T09:30:00 | Main Office |
    Then the punch attempt fails with:
      | reason                        |
      | Already punched in today      |
    And shows existing punch record:
      | punchType | timestamp           |
      | IN        | 2024-01-16T09:00:00 |

  Scenario: Agent has leave request and skips punch
    Given scenario "LeaveRequestSubmission: Agent submits approved sick leave"
    And today is '2024-01-17'
    When agent 'A001' views punch status
    Then the system shows:
      | message              | reason    |
      | No punch required    | On leave  |

  Scenario Outline: Punch time validation fails
    Given scenario "AgentManualRegistration: HR manually creates new agent profile"
    And today is '2024-01-16'
    When agent 'A001' attempts to punch in with:
      | timestamp    |
      | <timestamp>  |
    Then the punch fails with:
      | reason   |
      | <reason> |

    Examples:
      | timestamp           | reason                    |
      | 2024-01-15T23:59:00 | Cannot punch for past day |
      | 2024-01-17T00:01:00 | Cannot punch for future   | 