# AuditLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditLogs** | Pointer to [**[]AuditLog**](AuditLog.md) | Specifies a list of audit logs. | [optional] 
**Count** | Pointer to **NullableInt64** | Specifies the total number of audit logs that match the filter and search criteria. Use this value to determine how many additional requests are required to get the full result. | [optional] 

## Methods

### NewAuditLogs

`func NewAuditLogs() *AuditLogs`

NewAuditLogs instantiates a new AuditLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogsWithDefaults

`func NewAuditLogsWithDefaults() *AuditLogs`

NewAuditLogsWithDefaults instantiates a new AuditLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditLogs

`func (o *AuditLogs) GetAuditLogs() []AuditLog`

GetAuditLogs returns the AuditLogs field if non-nil, zero value otherwise.

### GetAuditLogsOk

`func (o *AuditLogs) GetAuditLogsOk() (*[]AuditLog, bool)`

GetAuditLogsOk returns a tuple with the AuditLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogs

`func (o *AuditLogs) SetAuditLogs(v []AuditLog)`

SetAuditLogs sets AuditLogs field to given value.

### HasAuditLogs

`func (o *AuditLogs) HasAuditLogs() bool`

HasAuditLogs returns a boolean if a field has been set.

### SetAuditLogsNil

`func (o *AuditLogs) SetAuditLogsNil(b bool)`

 SetAuditLogsNil sets the value for AuditLogs to be an explicit nil

### UnsetAuditLogs
`func (o *AuditLogs) UnsetAuditLogs()`

UnsetAuditLogs ensures that no value is present for AuditLogs, not even an explicit nil
### GetCount

`func (o *AuditLogs) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AuditLogs) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AuditLogs) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *AuditLogs) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *AuditLogs) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *AuditLogs) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


