@leave_rule_management
Feature: LeaveRuleManagement
  As an HR staff
  I want to manage leave rules by position
  So that different positions have appropriate leave policies

  Scenario: HR sets leave rules for specific position
    Given the system has positions:
      | position | department |
      | 技術駭客 | 工程部     |
      | 客服     | 客服部     |
    When HR sets leave rules for position '技術駭客':
      | leaveType | annualDays | carryOver | requireApproval |
      | 特休      | 15.0       | true      | true           |
      | 事假      | 14.0       | false     | true           |
      | 病假      | 30.0       | false     | false          |
    Then the position '技術駭客' should have leave rules:
      | leaveType | annualDays | carryOver | requireApproval |
      | 特休      | 15.0       | true      | true           |
      | 事假      | 14.0       | false     | true           |
      | 病假      | 30.0       | false     | false          |

  Scenario: System automatically applies leave rules when agent is created
    Given scenario "LeaveRuleManagement: HR sets leave rules for specific position"
    When HR creates an agent with position '技術駭客':
      | agentId | name   | position |
      | A002    | Alice  | 技術駭客 |
    Then agent 'A002' should automatically have leave balances:
      | leaveType | totalDays | availableDays |
      | 特休      | 15.0      | 15.0          |
      | 事假      | 14.0      | 14.0          |
      | 病假      | 30.0      | 30.0          | 