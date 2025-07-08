# 特務雲端人資系統 API 規格

本目錄包含完整的 OpenAPI 3.0 規格文件，將原始的 `api.md` 轉換為標準的 OpenAPI 格式，並按功能模組分成多個檔案。


## 使用方式

### 1. 整合完整規格
主檔案 `openapi.yaml` 包含所有 API 的完整規格，可以直接用於：
- Swagger UI 文檔生成
- API 用戶端程式碼生成
- API 測試工具

### 2. 工具支援
支援主流的 OpenAPI 工具：
- Swagger UI
- Postman
- Insomnia

## 開發建議

### 1. 維護策略
- 先更新對應的模組檔案
- 確保主檔案的 `$ref` 引用正確
- 保持版本一致性

### 2. 驗證工具
推薦使用以下工具進行規格驗證：
```bash
# 使用 swagger-validator
swagger-validator openapi.yaml

# 使用 redoc-cli
redoc-cli validate openapi.yaml
```

### 3. 文檔生成
```bash
# 生成 HTML 文檔
redoc-cli build openapi.yaml --output api-docs.html

# 啟動開發伺服器
swagger-ui-serve openapi.yaml
```
