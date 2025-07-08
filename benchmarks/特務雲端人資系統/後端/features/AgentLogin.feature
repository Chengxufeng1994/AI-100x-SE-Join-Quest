@agent_login
Feature: AgentLogin
  身為特務，我必須登入此系統，此系統才知道我是誰，才能幫我執行各項人事業務。

  Scenario: Agent login failed due to invalid email domain
    When an agent attempts login with:
      | email           | displayName |
      | sean@gmail.com      | Johnny Pan  |
    Then the request fails with:
      | reason                    |
      | Only waterballsa.tw email domain is allowed |

  Scenario: Agent login failed due to no agent record found
    When an agent attempts login with:
      | email           | 
      | johnny@waterballsa.tw |
    Then the request fails with:
      | reason                    |
      | Agent record not found    |

  Scenario: Agent login succeeds after HR creates profile
    Given scenario AgentManualRegistration: HR manually creates new agent profile
    When an agent attempts login with:
      | email           |
      | johnny@waterballsa.tw |
    Then the login succeeds with:
      | agentId | name   | email   | 
      | A001    | Johnny | johnny@waterballsa.tw |

  Scenario Outline: Agent login fails with invalid Google domain
    When an agent attempts login with:
      | email   | 
      | <email>       | 
    Then the login fails with:
      | reason                           |
      | Only waterballsa.tw email domain is allowed |
    Examples:
      | email                |
      | test@gmail.com       |
      | user@yahoo.com       |
      | admin@company.com    |
      | agent@waterball.tw   | 