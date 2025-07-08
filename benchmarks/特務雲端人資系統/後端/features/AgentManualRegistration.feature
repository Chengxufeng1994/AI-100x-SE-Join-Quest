@agent_manual_registration
Feature: AgentManualRegistration
  身為人事，我想人工填寫方式註冊特務，這樣就能同時保存勞工名卡並正式建立新特務。

  Scenario: HR manually creates new agent profile successfully
    Given no existing agent records
    When HR creates an agent with:
      | name  | email | joinDate   | idNumber     | birthDate  | gender | address      | homePhone   | mobilePhone | education | nationality | position    | department | managerId | monthlySalary | workExperience | militaryStatus |
      | Johnny   | johnny@waterballsa.tw |2024-01-15 | A123456789   | 1990-01-01 | M      | Taipei City  | 02-12345678 | 0912345678  | Master    | Taiwan      | 技術駭客    | 工程部     | A000    | 80000         | 一年實習經驗 | 已服役 |
    And uploads documents:
      | documentType        | fileName        |
      | profilePhoto        | johnny_photo.jpg |
      | idCardFront         | id_front.jpg    |
      | idCardBack          | id_back.jpg     |
    And sets emergency contact:
      | name   | relationship | phone      | address     |
      | Jane   | Sister       | 0987654321 | Taichung    |
    And sets labor insurance:
      | startDate  |
      | 2024-01-15 |
    Then the agent 'A001' should be created with:
      | agentId | name   | status |
      | A001    | Johnny | Active |
    And the agent 'A001' should have profile:
      | name   | email | joinDate   | idNumber   | birthDate  | gender | position | monthlySalary |
      | Johnny | johnny@waterballsa.tw | 2024-01-15 | A123456789 | 1990-01-01 | M      | 技術駭客 | 80000         |

  Scenario Outline: HR creates agent with invalid data
    Given no existing agent records
    When HR creates an agent with:
      | agentId   | name   | email   | joinDate   | idNumber   | birthDate   | gender   | monthlySalary   |
      | <agentId> | <name> | <email> | <joinDate> | <idNumber> | <birthDate> | <gender> | <monthlySalary> |
    Then the request fails with:
      | reason   |
      | <reason> |
    Examples:
      | agentId | name | email | joinDate   | idNumber   | birthDate  | gender | monthlySalary | reason                    |
      | A001    |      | john@waterballsa.tw | 2024-01-15 | A123456789 | 1990-01-01 | M      | 80000         | Name is required          |
      | A001    | John | john@waterballsa.tw | 2025-100-12 | A123456789 | 1990-01-01 | M      | 80000         | Invalid join date format  |
      | A001    | John | john@waterballsa.tw | 2024-01-15 | 123        | 1990-01-01 | M      | 80000         | Invalid ID number format  |
      | A001    | John | john@waterballsa.tw | 2024-01-15 | A123456789 | 2030-01-01 | M      | 80000         | Birth date cannot be future |
      | A001    | John | john@waterballsa.tw | 2024-01-15 | A123456789 | 1990-01-01 | X      | 80000         | Invalid gender value      |
      | A001    | John | john@waterballsa.tw | 2024-01-15 | A123456789 | 1990-01-01 | M      | -1000         | Salary must be positive   |
      | A001    | John | invalid-email | 2024-01-15 | A123456789 | 1990-01-01 | M      | 80000         | Invalid email format      |
      | A001    | John | john@gmail.com | 2024-01-15 | A123456789 | 1990-01-01 | M      | 80000         | Only waterballsa.tw email domain is allowed |