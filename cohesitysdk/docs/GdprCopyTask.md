# GdprCopyTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **NullableInt64** | Specifies the job with which this copy task is tied to. Note: this is only used for internal aggregation. | [optional] 
**CloudTargetType** | Pointer to **NullableString** | Specifies the cloud deploy target type. For example &#39;kAzure&#39;,&#39;kAWS&#39;, &#39;kGCP&#39; | [optional] 
**ExpiryTimeUsecs** | Pointer to **NullableInt64** | Specifies the expiry of the copy task. | [optional] 
**TargetId** | Pointer to **NullableInt64** | Specifies the id for the target. | [optional] 
**TargetName** | Pointer to **NullableString** | Specifies the target of the replication or archival tasks. | [optional] 
**TotalSnapshots** | Pointer to **NullableInt64** | Specifies the total number of snapshots. | [optional] 
**Type** | Pointer to **NullableString** | Specifies details about the Copy Run of a Job Run. A Copy task copies snapshots resulted from a backup run to an external target which could be &#39;kLocal&#39;, &#39;kArchival&#39;, or &#39;kRemote&#39;. | [optional] 

## Methods

### NewGdprCopyTask

`func NewGdprCopyTask() *GdprCopyTask`

NewGdprCopyTask instantiates a new GdprCopyTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGdprCopyTaskWithDefaults

`func NewGdprCopyTaskWithDefaults() *GdprCopyTask`

NewGdprCopyTaskWithDefaults instantiates a new GdprCopyTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *GdprCopyTask) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *GdprCopyTask) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *GdprCopyTask) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *GdprCopyTask) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *GdprCopyTask) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *GdprCopyTask) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetCloudTargetType

`func (o *GdprCopyTask) GetCloudTargetType() string`

GetCloudTargetType returns the CloudTargetType field if non-nil, zero value otherwise.

### GetCloudTargetTypeOk

`func (o *GdprCopyTask) GetCloudTargetTypeOk() (*string, bool)`

GetCloudTargetTypeOk returns a tuple with the CloudTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudTargetType

`func (o *GdprCopyTask) SetCloudTargetType(v string)`

SetCloudTargetType sets CloudTargetType field to given value.

### HasCloudTargetType

`func (o *GdprCopyTask) HasCloudTargetType() bool`

HasCloudTargetType returns a boolean if a field has been set.

### SetCloudTargetTypeNil

`func (o *GdprCopyTask) SetCloudTargetTypeNil(b bool)`

 SetCloudTargetTypeNil sets the value for CloudTargetType to be an explicit nil

### UnsetCloudTargetType
`func (o *GdprCopyTask) UnsetCloudTargetType()`

UnsetCloudTargetType ensures that no value is present for CloudTargetType, not even an explicit nil
### GetExpiryTimeUsecs

`func (o *GdprCopyTask) GetExpiryTimeUsecs() int64`

GetExpiryTimeUsecs returns the ExpiryTimeUsecs field if non-nil, zero value otherwise.

### GetExpiryTimeUsecsOk

`func (o *GdprCopyTask) GetExpiryTimeUsecsOk() (*int64, bool)`

GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUsecs

`func (o *GdprCopyTask) SetExpiryTimeUsecs(v int64)`

SetExpiryTimeUsecs sets ExpiryTimeUsecs field to given value.

### HasExpiryTimeUsecs

`func (o *GdprCopyTask) HasExpiryTimeUsecs() bool`

HasExpiryTimeUsecs returns a boolean if a field has been set.

### SetExpiryTimeUsecsNil

`func (o *GdprCopyTask) SetExpiryTimeUsecsNil(b bool)`

 SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil

### UnsetExpiryTimeUsecs
`func (o *GdprCopyTask) UnsetExpiryTimeUsecs()`

UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil
### GetTargetId

`func (o *GdprCopyTask) GetTargetId() int64`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *GdprCopyTask) GetTargetIdOk() (*int64, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *GdprCopyTask) SetTargetId(v int64)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *GdprCopyTask) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *GdprCopyTask) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *GdprCopyTask) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetName

`func (o *GdprCopyTask) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GdprCopyTask) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GdprCopyTask) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GdprCopyTask) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### SetTargetNameNil

`func (o *GdprCopyTask) SetTargetNameNil(b bool)`

 SetTargetNameNil sets the value for TargetName to be an explicit nil

### UnsetTargetName
`func (o *GdprCopyTask) UnsetTargetName()`

UnsetTargetName ensures that no value is present for TargetName, not even an explicit nil
### GetTotalSnapshots

`func (o *GdprCopyTask) GetTotalSnapshots() int64`

GetTotalSnapshots returns the TotalSnapshots field if non-nil, zero value otherwise.

### GetTotalSnapshotsOk

`func (o *GdprCopyTask) GetTotalSnapshotsOk() (*int64, bool)`

GetTotalSnapshotsOk returns a tuple with the TotalSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSnapshots

`func (o *GdprCopyTask) SetTotalSnapshots(v int64)`

SetTotalSnapshots sets TotalSnapshots field to given value.

### HasTotalSnapshots

`func (o *GdprCopyTask) HasTotalSnapshots() bool`

HasTotalSnapshots returns a boolean if a field has been set.

### SetTotalSnapshotsNil

`func (o *GdprCopyTask) SetTotalSnapshotsNil(b bool)`

 SetTotalSnapshotsNil sets the value for TotalSnapshots to be an explicit nil

### UnsetTotalSnapshots
`func (o *GdprCopyTask) UnsetTotalSnapshots()`

UnsetTotalSnapshots ensures that no value is present for TotalSnapshots, not even an explicit nil
### GetType

`func (o *GdprCopyTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GdprCopyTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GdprCopyTask) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GdprCopyTask) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GdprCopyTask) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GdprCopyTask) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


