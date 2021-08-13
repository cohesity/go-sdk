# McmObjectActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the unique id of the activity event. | [optional] 
**Object** | Pointer to [**ObjectSummary**](ObjectSummary.md) |  | [optional] 
**SourceInfo** | Pointer to [**ObjectSummary**](ObjectSummary.md) |  | [optional] 
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster id. | [optional] [readonly] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the cluster incarnation id. | [optional] [readonly] 
**RegionId** | Pointer to **NullableString** | Specifies the region id. Applicable only in case of DMaaS. | [optional] [readonly] 
**TimestampUsecs** | Pointer to **NullableInt64** | Specifies the timestamp in Unix timestamp epoch in microseconds at which this activity occured. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of activity event. | [optional] 
**BackupRunParams** | Pointer to [**McmObjectBackupRunActivityParams**](McmObjectBackupRunActivityParams.md) |  | [optional] 
**RecoveryParams** | Pointer to [**McmObjectRecoverActivityParams**](McmObjectRecoverActivityParams.md) |  | [optional] 
**ArchivalRunParams** | Pointer to [**McmObjectArchivalRunActivityParams**](McmObjectArchivalRunActivityParams.md) |  | [optional] 

## Methods

### NewMcmObjectActivity

`func NewMcmObjectActivity() *McmObjectActivity`

NewMcmObjectActivity instantiates a new McmObjectActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmObjectActivityWithDefaults

`func NewMcmObjectActivityWithDefaults() *McmObjectActivity`

NewMcmObjectActivityWithDefaults instantiates a new McmObjectActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *McmObjectActivity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *McmObjectActivity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *McmObjectActivity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *McmObjectActivity) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *McmObjectActivity) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *McmObjectActivity) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObject

`func (o *McmObjectActivity) GetObject() ObjectSummary`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *McmObjectActivity) GetObjectOk() (*ObjectSummary, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *McmObjectActivity) SetObject(v ObjectSummary)`

SetObject sets Object field to given value.

### HasObject

`func (o *McmObjectActivity) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetSourceInfo

`func (o *McmObjectActivity) GetSourceInfo() ObjectSummary`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *McmObjectActivity) GetSourceInfoOk() (*ObjectSummary, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *McmObjectActivity) SetSourceInfo(v ObjectSummary)`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *McmObjectActivity) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.

### GetClusterId

`func (o *McmObjectActivity) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *McmObjectActivity) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *McmObjectActivity) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *McmObjectActivity) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *McmObjectActivity) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *McmObjectActivity) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *McmObjectActivity) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *McmObjectActivity) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *McmObjectActivity) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *McmObjectActivity) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *McmObjectActivity) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *McmObjectActivity) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetRegionId

`func (o *McmObjectActivity) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *McmObjectActivity) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *McmObjectActivity) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *McmObjectActivity) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *McmObjectActivity) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *McmObjectActivity) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetTimestampUsecs

`func (o *McmObjectActivity) GetTimestampUsecs() int64`

GetTimestampUsecs returns the TimestampUsecs field if non-nil, zero value otherwise.

### GetTimestampUsecsOk

`func (o *McmObjectActivity) GetTimestampUsecsOk() (*int64, bool)`

GetTimestampUsecsOk returns a tuple with the TimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampUsecs

`func (o *McmObjectActivity) SetTimestampUsecs(v int64)`

SetTimestampUsecs sets TimestampUsecs field to given value.

### HasTimestampUsecs

`func (o *McmObjectActivity) HasTimestampUsecs() bool`

HasTimestampUsecs returns a boolean if a field has been set.

### SetTimestampUsecsNil

`func (o *McmObjectActivity) SetTimestampUsecsNil(b bool)`

 SetTimestampUsecsNil sets the value for TimestampUsecs to be an explicit nil

### UnsetTimestampUsecs
`func (o *McmObjectActivity) UnsetTimestampUsecs()`

UnsetTimestampUsecs ensures that no value is present for TimestampUsecs, not even an explicit nil
### GetType

`func (o *McmObjectActivity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *McmObjectActivity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *McmObjectActivity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *McmObjectActivity) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *McmObjectActivity) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *McmObjectActivity) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetBackupRunParams

`func (o *McmObjectActivity) GetBackupRunParams() McmObjectBackupRunActivityParams`

GetBackupRunParams returns the BackupRunParams field if non-nil, zero value otherwise.

### GetBackupRunParamsOk

`func (o *McmObjectActivity) GetBackupRunParamsOk() (*McmObjectBackupRunActivityParams, bool)`

GetBackupRunParamsOk returns a tuple with the BackupRunParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRunParams

`func (o *McmObjectActivity) SetBackupRunParams(v McmObjectBackupRunActivityParams)`

SetBackupRunParams sets BackupRunParams field to given value.

### HasBackupRunParams

`func (o *McmObjectActivity) HasBackupRunParams() bool`

HasBackupRunParams returns a boolean if a field has been set.

### GetRecoveryParams

`func (o *McmObjectActivity) GetRecoveryParams() McmObjectRecoverActivityParams`

GetRecoveryParams returns the RecoveryParams field if non-nil, zero value otherwise.

### GetRecoveryParamsOk

`func (o *McmObjectActivity) GetRecoveryParamsOk() (*McmObjectRecoverActivityParams, bool)`

GetRecoveryParamsOk returns a tuple with the RecoveryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryParams

`func (o *McmObjectActivity) SetRecoveryParams(v McmObjectRecoverActivityParams)`

SetRecoveryParams sets RecoveryParams field to given value.

### HasRecoveryParams

`func (o *McmObjectActivity) HasRecoveryParams() bool`

HasRecoveryParams returns a boolean if a field has been set.

### GetArchivalRunParams

`func (o *McmObjectActivity) GetArchivalRunParams() McmObjectArchivalRunActivityParams`

GetArchivalRunParams returns the ArchivalRunParams field if non-nil, zero value otherwise.

### GetArchivalRunParamsOk

`func (o *McmObjectActivity) GetArchivalRunParamsOk() (*McmObjectArchivalRunActivityParams, bool)`

GetArchivalRunParamsOk returns a tuple with the ArchivalRunParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalRunParams

`func (o *McmObjectActivity) SetArchivalRunParams(v McmObjectArchivalRunActivityParams)`

SetArchivalRunParams sets ArchivalRunParams field to given value.

### HasArchivalRunParams

`func (o *McmObjectActivity) HasArchivalRunParams() bool`

HasArchivalRunParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


