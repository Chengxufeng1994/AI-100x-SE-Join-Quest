@leave_request_submission
Feature: LeaveRequestSubmission
  As an Agent
  I want to submit leave requests
  So that I can request time off and have it properly documented and approved

  Scenario: Agent submits full day vacation leave
    Given scenario "AgentManualRegistration: HR manually creates new agent profile"
    And agent 'A001' has leave balance:
      | leaveType | availableDays | usedDays |
      | 特休      | 15.0          | 0.0      |
      | 事假      | 14.0          | 0.0      |
      | 病假      | 30.0          | 0.0      |
    When agent 'A001' submits leave request:
      | leaveType | startDate  | endDate    | duration | reason           | isFullDay |
      | 特休      | 2024-01-20 | 2024-01-20 | 1.0      | Personal matters | true      |
    Then the leave request should be created:
      | requestId | agentId | leaveType | startDate  | endDate    | duration | status  |
      | LR001     | A001    | 特休      | 2024-01-20 | 2024-01-20 | 1.0      | Pending |
    And notification should be sent to manager:
      | managerId | message                               |
      | A000      | Agent A001 submitted leave request   |

  Scenario: Agent submits partial day sick leave with medical certificate
    Given scenario "AgentManualRegistration: HR manually creates new agent profile"
    When agent 'A001' submits leave request:
      | leaveType | startDate  | startTime | endTime | duration | reason    | medicalCert |
      | 病假      | 2024-01-20 | 13:00     | 17:00   | 4.0      | Flu       | cert001.pdf |
    Then the leave request should be created:
      | requestId | agentId | leaveType | startDate  | duration | status  | hasDocument |
      | LR002     | A001    | 病假      | 2024-01-20 | 4.0      | Pending | true        |

  Scenario: Agent submits approved sick leave
    Given scenario "LeaveRequestSubmission: Agent submits partial day sick leave with medical certificate"
    When manager 'A000' reviews leave request 'LR002':
      | action   | comments              |
      | Approve  | Medical cert verified |
    Then the leave request 'LR002' should be:
      | status   | approvedBy | comments              |
      | Approved | A000       | Medical cert verified |
    And agent 'A001' leave balance should be updated:
      | leaveType | availableDays | usedDays |
      | 病假      | 29.5          | 0.5      |

  Scenario: Agent cancels pending leave request
    Given scenario "LeaveRequestSubmission: Agent submits full day vacation leave"
    When agent 'A001' cancels leave request 'LR001'
    Then the leave request 'LR001' should be:
      | status    | cancelledAt         |
      | Cancelled | 2024-01-19T09:00:00 |
    And agent 'A001' leave balance should remain:
      | leaveType | availableDays | usedDays |
      | 特休      | 15.0          | 0.0      |

  Scenario Outline: Agent fails to submit invalid leave request
    Given scenario "AgentManualRegistration: HR manually creates new agent profile"
    When agent 'A001' submits leave request:
      | leaveType   | startDate   | endDate   | duration   | reason   |
      | <leaveType> | <startDate> | <endDate> | <duration> | <reason> |
    Then the leave request fails with:
      | reason   |
      | <error>  |

    Examples:
      | leaveType | startDate  | endDate    | duration | reason     | error                           |
      | 特休      | 2024-01-15 | 2024-01-20 | 5.0      | Vacation   | Cannot request past dates       |
      | 特休      | 2024-01-25 | 2024-01-20 | -3.0     | Invalid    | End date before start date     |
      | 病假      | 2024-01-25 | 2024-01-25 | 1.0      |            | Sick leave requires medical cert|
      | 無效假    | 2024-01-25 | 2024-01-25 | 1.0      | Test       | Invalid leave type             |

  Scenario Outline: Leave balance validation
    Given scenario "AgentManualRegistration: HR manually creates new agent profile"
    And agent 'A001' has leave balance:
      | leaveType   | availableDays   |
      | <leaveType> | <availableDays> |
    When agent 'A001' submits leave request:
      | leaveType   | startDate  | duration   |
      | <leaveType> | 2024-01-25 | <duration> |
    Then the leave request fails with:
      | reason   |
      | <error>  |

    Examples:
      | leaveType | availableDays | duration | error                        |
      | 事假      | 2.0           | 3.0      | Insufficient leave balance   |
      | 特休      | 0.0           | 1.0      | No available vacation days   | 