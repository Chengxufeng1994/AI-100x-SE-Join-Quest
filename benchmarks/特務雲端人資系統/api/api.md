# HR System API Specification

## Authentication APIs

### API Agent Login: [POST] /auth/login
- Behavior: "An agent logs into the HR system using their Google account."
- Properties:
  - googleEmail(str): Google account email (must be @waterballsa.tw domain)
  - displayName(str): Display name from Google account
- Response Format:
  ```json
  {
    "agentId": "string",        // Agent identifier
    "name": "string",           // Agent name
    "position": "string",       // Agent position
    "department": "string",     // Agent department
    "token": "string",          // JWT token
    "expiresAt": "string"       // Token expiration time
  }
  ```

## Agent Management APIs

### API Create Agent: [POST] /agents
- Behavior: "HR staff creates a new agent profile manually."
- Properties:
  - agentId(str): Agent identifier
  - name(str): Full name
  - joinDate(date): Join date
  - idNumber(str): ID card number
  - birthDate(date): Birth date
  - gender(enum): Gender (M/F)
  - address(str): Home address
  - homePhone(str): Home phone number
  - mobilePhone(str): Mobile phone number
  - education(str): Education level
  - nationality(str): Nationality
  - position(str): Position title
  - department(str): Department name
  - managerId(str): Manager's agent ID
  - monthlySalary(decimal): Monthly salary
- Response Format:
  ```json
  {
    "agentId": "string",
    "name": "string",
    "status": "Active|Pending|Inactive",
    "message": "string"
  }
  ```

### API Batch Create Agents: [POST] /agents/batch
- Behavior: "HR staff uploads Excel file to batch register multiple agents."
- Properties:
  - file(file): Excel file containing agent data
- Response Format:
  ```json
  {
    "totalProcessed": "number",
    "successCount": "number",
    "failCount": "number",
    "createdAgents": [
      {
        "agentId": "string",
        "name": "string",
        "status": "string"
      }
    ],
    "errors": [
      {
        "row": "number",
        "error": "string"
      }
    ]
  }
  ```

### API Get Agent: [GET] /agents/{agentId}
- Behavior: "Retrieve complete information of a specific agent."
- Properties:
  - agentId(str): Agent identifier
- Response Format:
  ```json
  {
    "agentId": "string",
    "name": "string",
    "googleEmail": "string",
    "joinDate": "date",
    "idNumber": "string",
    "birthDate": "date",
    "gender": "M|F",
    "address": "string",
    "homePhone": "string",
    "mobilePhone": "string",
    "education": "string",
    "nationality": "string",
    "position": "string",
    "department": "string",
    "managerId": "string",
    "monthlySalary": "decimal",
    "status": "Active|Pending|Inactive",
    "emergencyContacts": [
      {
        "contactId": "number",
        "name": "string",
        "relationship": "string",
        "phone": "string",
        "address": "string"
      }
    ],
    "documents": [
      {
        "documentId": "number",
        "documentType": "string",
        "fileName": "string",
        "status": "Uploaded|Verified|Rejected",
        "uploadedAt": "timestamp"
      }
    ]
  }
  ```

### API Update Agent: [PUT] /agents/{agentId}
- Behavior: "HR staff modifies agent information with audit trail."
- Properties:
  - agentId(str): Agent identifier
  - updates(object): Fields to update
  - operatorId(str): Operator's ID
  - operatorName(str): Operator's name
- Response Format:
  ```json
  {
    "agentId": "string",
    "updatedFields": ["string"],
    "message": "string",
    "auditLogIds": ["number"]
  }
  ```

### API Search Agents: [GET] /agents
- Behavior: "Search agents with various filters."
- Query Parameters:
  - department(str): Filter by department
  - status(str): Filter by status
  - position(str): Filter by position
  - managerId(str): Filter by manager
  - page(int): Page number
  - size(int): Page size
- Response Format:
  ```json
  {
    "agents": [
      {
        "agentId": "string",
        "name": "string",
        "position": "string",
        "department": "string",
        "status": "string"
      }
    ],
    "totalCount": "number",
    "page": "number",
    "size": "number"
  }
  ```

### API Export Agents: [GET] /agents/export
- Behavior: "Export agent data with filters to Excel/CSV."
- Query Parameters:
  - department(str): Filter by department
  - status(str): Filter by status
  - position(str): Filter by position
  - format(str): Export format (xlsx/csv)
- Response Format:
  ```
  File download with appropriate content-type header
  ```

## Document Management APIs

### API Upload Agent Document: [POST] /agents/{agentId}/documents
- Behavior: "Upload document for an agent."
- Properties:
  - agentId(str): Agent identifier
  - documentType(enum): Document type (profilePhoto/idCardFront/idCardBack/medicalCert/marriageCert)
  - file(file): Document file
- Response Format:
  ```json
  {
    "documentId": "number",
    "documentType": "string",
    "fileName": "string",
    "status": "Uploaded",
    "uploadedAt": "timestamp"
  }
  ```

### API Get Agent Documents: [GET] /agents/{agentId}/documents
- Behavior: "Retrieve all documents for an agent."
- Properties:
  - agentId(str): Agent identifier
- Response Format:
  ```json
  {
    "documents": [
      {
        "documentId": "number",
        "documentType": "string",
        "fileName": "string",
        "status": "Uploaded|Verified|Rejected",
        "uploadedAt": "timestamp"
      }
    ]
  }
  ```

### API Delete Agent Document: [DELETE] /agents/{agentId}/documents/{documentId}
- Behavior: "Delete a specific document."
- Properties:
  - agentId(str): Agent identifier
  - documentId(number): Document identifier
- Response Format:
  ```json
  {
    "message": "Document deleted successfully"
  }
  ```

## Emergency Contact APIs

### API Create Emergency Contact: [POST] /agents/{agentId}/emergency-contacts
- Behavior: "Create emergency contact for an agent."
- Properties:
  - agentId(str): Agent identifier
  - name(str): Contact name
  - relationship(str): Relationship to agent
  - phone(str): Contact phone number
  - address(str): Contact address
- Response Format:
  ```json
  {
    "contactId": "number",
    "name": "string",
    "relationship": "string",
    "phone": "string",
    "address": "string"
  }
  ```

### API Update Emergency Contact: [PUT] /agents/{agentId}/emergency-contacts/{contactId}
- Behavior: "Update emergency contact information."
- Properties:
  - agentId(str): Agent identifier
  - contactId(number): Contact identifier
  - name(str): Contact name
  - relationship(str): Relationship to agent
  - phone(str): Contact phone number
  - address(str): Contact address
- Response Format:
  ```json
  {
    "contactId": "number",
    "name": "string",
    "relationship": "string",
    "phone": "string",
    "address": "string"
  }
  ```

### API Get Emergency Contacts: [GET] /agents/{agentId}/emergency-contacts
- Behavior: "Retrieve all emergency contacts for an agent."
- Properties:
  - agentId(str): Agent identifier
- Response Format:
  ```json
  {
    "contacts": [
      {
        "contactId": "number",
        "name": "string",
        "relationship": "string",
        "phone": "string",
        "address": "string"
      }
    ]
  }
  ```

## Punch Clock APIs

### API Punch Clock: [POST] /agents/{agentId}/punch
- Behavior: "Agent punches in or out for work."
- Properties:
  - agentId(str): Agent identifier
  - punchType(enum): Punch type (IN/OUT)
  - location(str): Punch location
  - ipAddress(str): IP address
  - device(str): Device information
- Response Format:
  ```json
  {
    "recordId": "number",
    "agentId": "string",
    "date": "date",
    "punchType": "IN|OUT",
    "timestamp": "timestamp",
    "location": "string",
    "status": "Valid|Invalid",
    "workStatus": "Punched In|Punched Out",
    "message": "string"
  }
  ```

### API Get Punch Status: [GET] /agents/{agentId}/punch/status
- Behavior: "Get current punch status for an agent."
- Properties:
  - agentId(str): Agent identifier
- Response Format:
  ```json
  {
    "agentId": "string",
    "date": "date",
    "workStatus": "Not Punched|Punched In|Punched Out|On Leave",
    "lastPunch": {
      "punchType": "IN|OUT",
      "timestamp": "timestamp",
      "location": "string"
    },
    "message": "string"
  }
  ```

## Work Schedule APIs

### API Set Work Schedule: [POST] /agents/{agentId}/schedules
- Behavior: "Agent sets flexible work schedule for a specific date."
- Properties:
  - agentId(str): Agent identifier
  - date(date): Work date
  - workStart(time): Work start time
  - workEnd(time): Work end time
  - breakStart(time): Break start time
  - breakEnd(time): Break end time
  - totalHours(decimal): Total work hours
- Response Format:
  ```json
  {
    "scheduleId": "number",
    "agentId": "string",
    "date": "date",
    "workStart": "time",
    "workEnd": "time",
    "breakStart": "time",
    "breakEnd": "time",
    "totalHours": "decimal",
    "status": "Active"
  }
  ```

### API Update Work Schedule: [PUT] /agents/{agentId}/schedules/{date}
- Behavior: "Agent modifies work schedule before 10AM next day."
- Properties:
  - agentId(str): Agent identifier
  - date(date): Work date
  - workStart(time): Work start time
  - workEnd(time): Work end time
  - breakStart(time): Break start time
  - breakEnd(time): Break end time
  - totalHours(decimal): Total work hours
- Response Format:
  ```json
  {
    "scheduleId": "number",
    "agentId": "string",
    "date": "date",
    "workStart": "time",
    "workEnd": "time",
    "breakStart": "time",
    "breakEnd": "time",
    "totalHours": "decimal",
    "status": "Modified",
    "modifiedAt": "timestamp"
  }
  ```

### API Get Work Schedules: [GET] /agents/{agentId}/schedules
- Behavior: "Retrieve work schedules for an agent."
- Query Parameters:
  - startDate(date): Start date filter
  - endDate(date): End date filter
- Properties:
  - agentId(str): Agent identifier
- Response Format:
  ```json
  {
    "schedules": [
      {
        "scheduleId": "number",
        "date": "date",
        "workStart": "time",
        "workEnd": "time",
        "breakStart": "time",
        "breakEnd": "time",
        "totalHours": "decimal",
        "status": "Active|Modified|Cancelled"
      }
    ]
  }
  ```

## Attendance Management APIs

### API Get Attendance Records: [GET] /agents/{agentId}/attendance
- Behavior: "Retrieve attendance records for an agent."
- Query Parameters:
  - startDate(date): Start date filter
  - endDate(date): End date filter
- Properties:
  - agentId(str): Agent identifier
- Response Format:
  ```json
  {
    "attendanceRecords": [
      {
        "attendanceId": "number",
        "date": "date",
        "punchIn": "timestamp",
        "punchOut": "timestamp",
        "workHours": "decimal",
        "status": "Normal|Adjusted|OnLeave|Absent"
      }
    ]
  }
  ```

### API Adjust Attendance Record: [PUT] /attendance/{attendanceId}
- Behavior: "HR staff adjusts attendance record with audit trail."
- Properties:
  - attendanceId(number): Attendance record identifier
  - adjustments(object): Fields to adjust
  - reason(str): Adjustment reason
  - operatorId(str): Operator's ID
- Response Format:
  ```json
  {
    "attendanceId": "number",
    "agentId": "string",
    "date": "date",
    "punchIn": "timestamp",
    "punchOut": "timestamp",
    "workHours": "decimal",
    "status": "Adjusted",
    "auditLogId": "number"
  }
  ```

### API Export Attendance Data: [GET] /attendance/export
- Behavior: "Export attendance data with filters."
- Query Parameters:
  - startDate(date): Start date filter
  - endDate(date): End date filter
  - department(str): Department filter
  - status(str): Status filter
  - format(str): Export format (xlsx/csv)
- Response Format:
  ```
  File download with appropriate content-type header
  ```

### API Get Attendance Exceptions: [GET] /attendance/exceptions
- Behavior: "HR views attendance exception alerts."
- Query Parameters:
  - startDate(date): Start date filter
  - endDate(date): End date filter
  - isResolved(boolean): Filter by resolution status
- Response Format:
  ```json
  {
    "exceptions": [
      {
        "exceptionId": "number",
        "agentId": "string",
        "agentName": "string",
        "date": "date",
        "alertType": "string",
        "description": "string",
        "isResolved": "boolean",
        "createdAt": "timestamp"
      }
    ]
  }
  ```

## Leave Management APIs

### API Submit Leave Request: [POST] /agents/{agentId}/leave-requests
- Behavior: "Agent submits a leave request."
- Properties:
  - agentId(str): Agent identifier
  - leaveType(str): Leave type (特休/事假/病假/婚假/喪假/產假/陪產假/生理假)
  - startDate(date): Leave start date
  - endDate(date): Leave end date
  - startTime(time): Start time for partial day leave
  - endTime(time): End time for partial day leave
  - duration(decimal): Leave duration in hours/days
  - isFullDay(boolean): Whether it's full day leave
  - reason(str): Leave reason
  - medicalCert(file): Medical certificate for sick leave
- Response Format:
  ```json
  {
    "requestId": "string",
    "agentId": "string",
    "leaveType": "string",
    "startDate": "date",
    "endDate": "date",
    "duration": "decimal",
    "status": "Pending|Auto-approved",
    "submittedAt": "timestamp"
  }
  ```

### API Cancel Leave Request: [PUT] /leave-requests/{requestId}/cancel
- Behavior: "Agent cancels a pending leave request."
- Properties:
  - requestId(str): Leave request identifier
- Response Format:
  ```json
  {
    "requestId": "string",
    "status": "Cancelled",
    "cancelledAt": "timestamp",
    "message": "Leave request cancelled successfully"
  }
  ```

### API Review Leave Request: [PUT] /leave-requests/{requestId}/review
- Behavior: "Manager approves or rejects a leave request."
- Properties:
  - requestId(str): Leave request identifier
  - action(enum): Review action (Approve/Reject)
  - comments(str): Review comments
  - reviewerId(str): Reviewer's agent ID
- Response Format:
  ```json
  {
    "requestId": "string",
    "status": "Approved|Rejected",
    "reviewedBy": "string",
    "reviewedAt": "timestamp",
    "comments": "string"
  }
  ```

### API Get Leave Requests: [GET] /agents/{agentId}/leave-requests
- Behavior: "Retrieve leave requests for an agent."
- Query Parameters:
  - status(str): Filter by status
  - year(int): Filter by year
- Properties:
  - agentId(str): Agent identifier
- Response Format:
  ```json
  {
    "leaveRequests": [
      {
        "requestId": "string",
        "leaveType": "string",
        "startDate": "date",
        "endDate": "date",
        "duration": "decimal",
        "status": "Pending|Approved|Rejected|Cancelled",
        "submittedAt": "timestamp",
        "reviewedBy": "string",
        "reviewedAt": "timestamp",
        "comments": "string"
      }
    ]
  }
  ```

### API Get Leave Balances: [GET] /agents/{agentId}/leave-balances
- Behavior: "Retrieve leave balances for an agent."
- Query Parameters:
  - year(int): Filter by year
- Properties:
  - agentId(str): Agent identifier
- Response Format:
  ```json
  {
    "leaveBalances": [
      {
        "balanceId": "number",
        "leaveType": "string",
        "year": "number",
        "totalDays": "decimal",
        "usedDays": "decimal",
        "availableDays": "decimal",
        "expiryDate": "date"
      }
    ]
  }
  ```

### API Set Leave Rules for Position: [POST] /leave-rules/positions/{position}
- Behavior: "HR sets leave rules for a specific position."
- Properties:
  - position(str): Position name
  - rules(array): Array of leave rules
- Response Format:
  ```json
  {
    "position": "string",
    "positionType": "特務|專職",
    "rules": [
      {
        "ruleId": "number",
        "leaveType": "string",
        "annualQuota": "decimal",
        "canReject": "boolean"
      }
    ]
  }
  ```

### API Get Leave Rules for Position: [GET] /leave-rules/positions/{position}
- Behavior: "Retrieve leave rules for a specific position."
- Properties:
  - position(str): Position name
- Response Format:
  ```json
  {
    "position": "string",
    "positionType": "特務|專職",
    "rules": [
      {
        "ruleId": "number",
        "leaveType": "string",
        "annualQuota": "decimal",
        "canReject": "boolean"
      }
    ]
  }
  ```

## Notification APIs

### API Get Notifications: [GET] /agents/{agentId}/notifications
- Behavior: "Retrieve notifications for an agent."
- Query Parameters:
  - isRead(boolean): Filter by read status
  - type(str): Filter by notification type
- Properties:
  - agentId(str): Agent identifier
- Response Format:
  ```json
  {
    "notifications": [
      {
        "notificationId": "number",
        "type": "LeaveRequest|LeaveApproval|LeaveRejection|AttendanceAlert",
        "message": "string",
        "relatedId": "string",
        "isRead": "boolean",
        "createdAt": "timestamp"
      }
    ]
  }
  ```

### API Mark Notification as Read: [PUT] /notifications/{notificationId}/read
- Behavior: "Mark a notification as read."
- Properties:
  - notificationId(number): Notification identifier
- Response Format:
  ```json
  {
    "notificationId": "number",
    "isRead": true,
    "message": "Notification marked as read"
  }
  ```

## Audit Log APIs

### API Get Audit Logs: [GET] /agents/{agentId}/audit-logs
- Behavior: "Retrieve audit logs for an agent."
- Query Parameters:
  - tableName(str): Filter by table name
  - field(str): Filter by field name
  - startDate(date): Start date filter
  - endDate(date): End date filter
- Properties:
  - agentId(str): Agent identifier
- Response Format:
  ```json
  {
    "auditLogs": [
      {
        "logId": "number",
        "tableName": "string",
        "field": "string",
        "oldValue": "string",
        "newValue": "string",
        "operatorId": "string",
        "operatorName": "string",
        "modifiedAt": "timestamp"
      }
    ]
  }
  ```

## System APIs

### API Get System Status: [GET] /system/status
- Behavior: "Check system health and status."
- Response Format:
  ```json
  {
    "status": "healthy",
    "timestamp": "timestamp",
    "version": "string",
    "uptime": "string"
  }
  ```

## Error Response Format

All APIs may return error responses in the following format:

```json
{
  "error": {
    "code": "string",
    "message": "string",
    "details": "string",
    "timestamp": "timestamp"
  }
}
```

Common HTTP Status Codes:
- 200: Success
- 201: Created
- 400: Bad Request
- 401: Unauthorized
- 403: Forbidden
- 404: Not Found
- 409: Conflict
- 422: Unprocessable Entity
- 500: Internal Server Error 