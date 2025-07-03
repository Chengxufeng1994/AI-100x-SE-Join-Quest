@double11
Feature: Double 11 Promotion - Bulk Same Product Discount
  As a shopper during the double eleven festival
  I want to receive a 20% discount on every set of 10 same products
  So that I can save more when buying in bulk

  Scenario: Buying 12 identical socks at 100 each
    Given the double eleven promotion is active
    When a customer places an order with:
      | productName | quantity | unitPrice |
      | 襪子        | 12       | 100       |
    Then the order summary should be:
      | totalAmount |
      | 1000        |
    And the customer should receive:
      | productName | quantity |
      | 襪子        | 12       |

  Scenario: Buying 27 identical socks at 100 each
    Given the double eleven promotion is active
    When a customer places an order with:
      | productName | quantity | unitPrice |
      | 襪子        | 27       | 100       |
    Then the order summary should be:
      | totalAmount |
      | 2300        |
    And the customer should receive:
      | productName | quantity |
      | 襪子        | 27       |

  Scenario: Buying 10 different products at 100 each
    Given the double eleven promotion is active
    When a customer places an order with:
      | productName | quantity | unitPrice |
      | A           | 1        | 100       |
      | B           | 1        | 100       |
      | C           | 1        | 100       |
      | D           | 1        | 100       |
      | E           | 1        | 100       |
      | F           | 1        | 100       |
      | G           | 1        | 100       |
      | H           | 1        | 100       |
      | I           | 1        | 100       |
      | J           | 1        | 100       |
    Then the order summary should be:
      | totalAmount |
      | 1000        |
    And the customer should receive:
      | productName | quantity |
      | A           | 1        |
      | B           | 1        |
      | C           | 1        |
      | D           | 1        |
      | E           | 1        |
      | F           | 1        |
      | G           | 1        |
      | H           | 1        |
      | I           | 1        |
      | J           | 1        |
