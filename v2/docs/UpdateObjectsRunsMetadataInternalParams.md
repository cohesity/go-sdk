# UpdateObjectsRunsMetadataInternalParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **string** | Specifies the environment. | [optional] 
**LegalHold** | Pointer to **NullableString** | Specifies whether to retain the snapshot for legal purpose. If set to &#39;enable&#39;, the snapshots cannot be deleted until the retention period. Note that using this option may cause the Cluster to run out of space. This field can be set only by a User having Data Security Role.  If set to &#39;release&#39;, the snapshots under legal hold will be released. | [optional] 
**Retention** | Pointer to [**Retention**](Retention.md) |  | [optional] 
**TargetObjectRuns** | Pointer to [**[]TargetObjectRun**](TargetObjectRun.md) | An array of objects. Each containing object id and the run start time that we want to target. | [optional] 
**TimeRange** | Pointer to [**TimeRangeUsecs**](TimeRangeUsecs.md) |  | [optional] 
**TenantId** | Pointer to **NullableString** | Tenant id associated with the incoming request | [optional] 

## Methods

### NewUpdateObjectsRunsMetadataInternalParams

`func NewUpdateObjectsRunsMetadataInternalParams() *UpdateObjectsRunsMetadataInternalParams`

NewUpdateObjectsRunsMetadataInternalParams instantiates a new UpdateObjectsRunsMetadataInternalParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateObjectsRunsMetadataInternalParamsWithDefaults

`func NewUpdateObjectsRunsMetadataInternalParamsWithDefaults() *UpdateObjectsRunsMetadataInternalParams`

NewUpdateObjectsRunsMetadataInternalParamsWithDefaults instantiates a new UpdateObjectsRunsMetadataInternalParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *UpdateObjectsRunsMetadataInternalParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *UpdateObjectsRunsMetadataInternalParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *UpdateObjectsRunsMetadataInternalParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *UpdateObjectsRunsMetadataInternalParams) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetLegalHold

`func (o *UpdateObjectsRunsMetadataInternalParams) GetLegalHold() string`

GetLegalHold returns the LegalHold field if non-nil, zero value otherwise.

### GetLegalHoldOk

`func (o *UpdateObjectsRunsMetadataInternalParams) GetLegalHoldOk() (*string, bool)`

GetLegalHoldOk returns a tuple with the LegalHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalHold

`func (o *UpdateObjectsRunsMetadataInternalParams) SetLegalHold(v string)`

SetLegalHold sets LegalHold field to given value.

### HasLegalHold

`func (o *UpdateObjectsRunsMetadataInternalParams) HasLegalHold() bool`

HasLegalHold returns a boolean if a field has been set.

### SetLegalHoldNil

`func (o *UpdateObjectsRunsMetadataInternalParams) SetLegalHoldNil(b bool)`

 SetLegalHoldNil sets the value for LegalHold to be an explicit nil

### UnsetLegalHold
`func (o *UpdateObjectsRunsMetadataInternalParams) UnsetLegalHold()`

UnsetLegalHold ensures that no value is present for LegalHold, not even an explicit nil
### GetRetention

`func (o *UpdateObjectsRunsMetadataInternalParams) GetRetention() Retention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *UpdateObjectsRunsMetadataInternalParams) GetRetentionOk() (*Retention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *UpdateObjectsRunsMetadataInternalParams) SetRetention(v Retention)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *UpdateObjectsRunsMetadataInternalParams) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetTargetObjectRuns

`func (o *UpdateObjectsRunsMetadataInternalParams) GetTargetObjectRuns() []TargetObjectRun`

GetTargetObjectRuns returns the TargetObjectRuns field if non-nil, zero value otherwise.

### GetTargetObjectRunsOk

`func (o *UpdateObjectsRunsMetadataInternalParams) GetTargetObjectRunsOk() (*[]TargetObjectRun, bool)`

GetTargetObjectRunsOk returns a tuple with the TargetObjectRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetObjectRuns

`func (o *UpdateObjectsRunsMetadataInternalParams) SetTargetObjectRuns(v []TargetObjectRun)`

SetTargetObjectRuns sets TargetObjectRuns field to given value.

### HasTargetObjectRuns

`func (o *UpdateObjectsRunsMetadataInternalParams) HasTargetObjectRuns() bool`

HasTargetObjectRuns returns a boolean if a field has been set.

### GetTimeRange

`func (o *UpdateObjectsRunsMetadataInternalParams) GetTimeRange() TimeRangeUsecs`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *UpdateObjectsRunsMetadataInternalParams) GetTimeRangeOk() (*TimeRangeUsecs, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *UpdateObjectsRunsMetadataInternalParams) SetTimeRange(v TimeRangeUsecs)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *UpdateObjectsRunsMetadataInternalParams) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.

### GetTenantId

`func (o *UpdateObjectsRunsMetadataInternalParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UpdateObjectsRunsMetadataInternalParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UpdateObjectsRunsMetadataInternalParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UpdateObjectsRunsMetadataInternalParams) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *UpdateObjectsRunsMetadataInternalParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *UpdateObjectsRunsMetadataInternalParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


