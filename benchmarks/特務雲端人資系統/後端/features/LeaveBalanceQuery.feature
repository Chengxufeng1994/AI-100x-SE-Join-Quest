@leave_balance_query
Feature: LeaveBalanceQuery
  As an Agent
  I want to view my leave balances and history
  So that I can plan my time off effectively

  Scenario: Agent views current leave balances
    Given scenario "LeaveRequestApproval: Manager approves agent leave request"
    When agent 'A001' queries leave balances
    Then the system returns:
      | leaveType | totalDays | usedDays | availableDays | expiryDate |
      | 特休      | 15.0      | 1.0      | 14.0          | 2024-12-31 |
      | 事假      | 14.0      | 0.0      | 14.0          | 2024-12-31 |
      | 病假      | 30.0      | 0.0      | 30.0          | 2024-12-31 |

  Scenario: Agent views leave history with detailed breakdown
    Given scenario "LeaveRequestApproval: Manager approves agent leave request"
    When agent 'A001' queries leave history for:
      | year |
      | 2024 |
    Then the system returns leave history:
      | leaveType | requestDate | startDate  | duration | status   | approvedBy |
      | 特休      | 2024-01-18  | 2024-01-20 | 1.0      | Approved | A000       | 