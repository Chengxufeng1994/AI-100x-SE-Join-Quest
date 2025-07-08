@leave_request_approval
Feature: LeaveRequestApproval
  As a Manager
  I want to approve or reject leave requests
  So that team members' leave can be properly managed

  Scenario: Manager approves agent leave request
    Given scenario "LeaveRequestSubmission: Agent submits full day vacation leave"
    When manager 'A000' reviews leave request 'LR001':
      | action   | comments           |
      | Approve  | Approved for PTO   |
    Then the leave request 'LR001' should be:
      | status   | approvedBy | comments          |
      | Approved | A000       | Approved for PTO  |
    And agent 'A001' leave balance should be updated:
      | leaveType | availableDays | usedDays |
      | 特休      | 14.0          | 1.0      |

  Scenario: Manager rejects agent leave request
    Given scenario "LeaveRequestSubmission: Agent submits full day vacation leave"
    When manager 'A000' reviews leave request 'LR001':
      | action  | comments                     |
      | Reject  | Insufficient staff coverage |
    Then the leave request 'LR001' should be:
      | status   | rejectedBy | comments                     |
      | Rejected | A000       | Insufficient staff coverage |
    And agent 'A001' leave balance should remain:
      | leaveType | availableDays | usedDays |
      | 特休      | 15.0          | 0.0      | 