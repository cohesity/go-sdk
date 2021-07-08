# ObjectInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessibleUsers** | Pointer to **[]string** | Species the list of user who have access to this object. | [optional] 
**AuditLogs** | Pointer to [**[]ClusterAuditLog**](ClusterAuditLog.md) | Specifies the audit log information. | [optional] 
**CopyTaskInfo** | Pointer to [**[]GdprCopyTask**](GdprCopyTask.md) | Specifies the copy task information. | [optional] 
**IsProtected** | Pointer to **NullableBool** | Specifies the protection status of the object. | [optional] 
**Location** | Pointer to **NullableString** | Specifies the location of the parent source. | [optional] 
**ProtectionInfo** | Pointer to [**[]ProtectionInfo**](ProtectionInfo.md) | Specifies the data locations for the protected objects. | [optional] 
**RootNodeId** | Pointer to **NullableInt64** | Specifies the id of the root node. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the Protection Source. | [optional] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the object. | [optional] 

## Methods

### NewObjectInformation

`func NewObjectInformation() *ObjectInformation`

NewObjectInformation instantiates a new ObjectInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectInformationWithDefaults

`func NewObjectInformationWithDefaults() *ObjectInformation`

NewObjectInformationWithDefaults instantiates a new ObjectInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibleUsers

`func (o *ObjectInformation) GetAccessibleUsers() []string`

GetAccessibleUsers returns the AccessibleUsers field if non-nil, zero value otherwise.

### GetAccessibleUsersOk

`func (o *ObjectInformation) GetAccessibleUsersOk() (*[]string, bool)`

GetAccessibleUsersOk returns a tuple with the AccessibleUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibleUsers

`func (o *ObjectInformation) SetAccessibleUsers(v []string)`

SetAccessibleUsers sets AccessibleUsers field to given value.

### HasAccessibleUsers

`func (o *ObjectInformation) HasAccessibleUsers() bool`

HasAccessibleUsers returns a boolean if a field has been set.

### SetAccessibleUsersNil

`func (o *ObjectInformation) SetAccessibleUsersNil(b bool)`

 SetAccessibleUsersNil sets the value for AccessibleUsers to be an explicit nil

### UnsetAccessibleUsers
`func (o *ObjectInformation) UnsetAccessibleUsers()`

UnsetAccessibleUsers ensures that no value is present for AccessibleUsers, not even an explicit nil
### GetAuditLogs

`func (o *ObjectInformation) GetAuditLogs() []ClusterAuditLog`

GetAuditLogs returns the AuditLogs field if non-nil, zero value otherwise.

### GetAuditLogsOk

`func (o *ObjectInformation) GetAuditLogsOk() (*[]ClusterAuditLog, bool)`

GetAuditLogsOk returns a tuple with the AuditLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogs

`func (o *ObjectInformation) SetAuditLogs(v []ClusterAuditLog)`

SetAuditLogs sets AuditLogs field to given value.

### HasAuditLogs

`func (o *ObjectInformation) HasAuditLogs() bool`

HasAuditLogs returns a boolean if a field has been set.

### SetAuditLogsNil

`func (o *ObjectInformation) SetAuditLogsNil(b bool)`

 SetAuditLogsNil sets the value for AuditLogs to be an explicit nil

### UnsetAuditLogs
`func (o *ObjectInformation) UnsetAuditLogs()`

UnsetAuditLogs ensures that no value is present for AuditLogs, not even an explicit nil
### GetCopyTaskInfo

`func (o *ObjectInformation) GetCopyTaskInfo() []GdprCopyTask`

GetCopyTaskInfo returns the CopyTaskInfo field if non-nil, zero value otherwise.

### GetCopyTaskInfoOk

`func (o *ObjectInformation) GetCopyTaskInfoOk() (*[]GdprCopyTask, bool)`

GetCopyTaskInfoOk returns a tuple with the CopyTaskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTaskInfo

`func (o *ObjectInformation) SetCopyTaskInfo(v []GdprCopyTask)`

SetCopyTaskInfo sets CopyTaskInfo field to given value.

### HasCopyTaskInfo

`func (o *ObjectInformation) HasCopyTaskInfo() bool`

HasCopyTaskInfo returns a boolean if a field has been set.

### SetCopyTaskInfoNil

`func (o *ObjectInformation) SetCopyTaskInfoNil(b bool)`

 SetCopyTaskInfoNil sets the value for CopyTaskInfo to be an explicit nil

### UnsetCopyTaskInfo
`func (o *ObjectInformation) UnsetCopyTaskInfo()`

UnsetCopyTaskInfo ensures that no value is present for CopyTaskInfo, not even an explicit nil
### GetIsProtected

`func (o *ObjectInformation) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *ObjectInformation) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *ObjectInformation) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *ObjectInformation) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### SetIsProtectedNil

`func (o *ObjectInformation) SetIsProtectedNil(b bool)`

 SetIsProtectedNil sets the value for IsProtected to be an explicit nil

### UnsetIsProtected
`func (o *ObjectInformation) UnsetIsProtected()`

UnsetIsProtected ensures that no value is present for IsProtected, not even an explicit nil
### GetLocation

`func (o *ObjectInformation) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ObjectInformation) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ObjectInformation) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ObjectInformation) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *ObjectInformation) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *ObjectInformation) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetProtectionInfo

`func (o *ObjectInformation) GetProtectionInfo() []ProtectionInfo`

GetProtectionInfo returns the ProtectionInfo field if non-nil, zero value otherwise.

### GetProtectionInfoOk

`func (o *ObjectInformation) GetProtectionInfoOk() (*[]ProtectionInfo, bool)`

GetProtectionInfoOk returns a tuple with the ProtectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionInfo

`func (o *ObjectInformation) SetProtectionInfo(v []ProtectionInfo)`

SetProtectionInfo sets ProtectionInfo field to given value.

### HasProtectionInfo

`func (o *ObjectInformation) HasProtectionInfo() bool`

HasProtectionInfo returns a boolean if a field has been set.

### SetProtectionInfoNil

`func (o *ObjectInformation) SetProtectionInfoNil(b bool)`

 SetProtectionInfoNil sets the value for ProtectionInfo to be an explicit nil

### UnsetProtectionInfo
`func (o *ObjectInformation) UnsetProtectionInfo()`

UnsetProtectionInfo ensures that no value is present for ProtectionInfo, not even an explicit nil
### GetRootNodeId

`func (o *ObjectInformation) GetRootNodeId() int64`

GetRootNodeId returns the RootNodeId field if non-nil, zero value otherwise.

### GetRootNodeIdOk

`func (o *ObjectInformation) GetRootNodeIdOk() (*int64, bool)`

GetRootNodeIdOk returns a tuple with the RootNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootNodeId

`func (o *ObjectInformation) SetRootNodeId(v int64)`

SetRootNodeId sets RootNodeId field to given value.

### HasRootNodeId

`func (o *ObjectInformation) HasRootNodeId() bool`

HasRootNodeId returns a boolean if a field has been set.

### SetRootNodeIdNil

`func (o *ObjectInformation) SetRootNodeIdNil(b bool)`

 SetRootNodeIdNil sets the value for RootNodeId to be an explicit nil

### UnsetRootNodeId
`func (o *ObjectInformation) UnsetRootNodeId()`

UnsetRootNodeId ensures that no value is present for RootNodeId, not even an explicit nil
### GetSourceId

`func (o *ObjectInformation) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ObjectInformation) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ObjectInformation) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ObjectInformation) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *ObjectInformation) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *ObjectInformation) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *ObjectInformation) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *ObjectInformation) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *ObjectInformation) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *ObjectInformation) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *ObjectInformation) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *ObjectInformation) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


