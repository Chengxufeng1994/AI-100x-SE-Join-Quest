# Task

接著你要進行設計的重構工作。

## Context

重構目標類別： @service/cchinesechessService
重構目標：遵守 OCP, 開放封閉原則

## Folder

- domain: core domain model
- service: application service

## 遵守 BDD 流程

1. 先執行重構
2. 重構完之後要進行回歸測試，下達測試指令，並確保所有測試仍然通過
3. 如果有任一測試不通過，則要修正程式邏輯以確保所有測試都要通過
