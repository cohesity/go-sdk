# UpdateObjectsRunsMetadataParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **string** | Specifies the environment. | [optional] 
**LegalHold** | Pointer to **NullableString** | Specifies whether to retain the snapshot for legal purpose. If set to &#39;enable&#39;, the snapshots cannot be deleted until the retention period. Note that using this option may cause the Cluster to run out of space. This field can be set only by a User having Data Security Role.  If set to &#39;release&#39;, the snapshots under legal hold will be released. | [optional] 
**Retention** | Pointer to [**Retention**](Retention.md) |  | [optional] 
**TargetObjectRuns** | Pointer to [**[]TargetObjectRun**](TargetObjectRun.md) | An array of objects. Each containing object id and the run start time that we want to target. | [optional] 
**TimeRange** | Pointer to [**TimeRangeUsecs**](TimeRangeUsecs.md) |  | [optional] 

## Methods

### NewUpdateObjectsRunsMetadataParams

`func NewUpdateObjectsRunsMetadataParams() *UpdateObjectsRunsMetadataParams`

NewUpdateObjectsRunsMetadataParams instantiates a new UpdateObjectsRunsMetadataParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateObjectsRunsMetadataParamsWithDefaults

`func NewUpdateObjectsRunsMetadataParamsWithDefaults() *UpdateObjectsRunsMetadataParams`

NewUpdateObjectsRunsMetadataParamsWithDefaults instantiates a new UpdateObjectsRunsMetadataParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *UpdateObjectsRunsMetadataParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *UpdateObjectsRunsMetadataParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *UpdateObjectsRunsMetadataParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *UpdateObjectsRunsMetadataParams) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetLegalHold

`func (o *UpdateObjectsRunsMetadataParams) GetLegalHold() string`

GetLegalHold returns the LegalHold field if non-nil, zero value otherwise.

### GetLegalHoldOk

`func (o *UpdateObjectsRunsMetadataParams) GetLegalHoldOk() (*string, bool)`

GetLegalHoldOk returns a tuple with the LegalHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalHold

`func (o *UpdateObjectsRunsMetadataParams) SetLegalHold(v string)`

SetLegalHold sets LegalHold field to given value.

### HasLegalHold

`func (o *UpdateObjectsRunsMetadataParams) HasLegalHold() bool`

HasLegalHold returns a boolean if a field has been set.

### SetLegalHoldNil

`func (o *UpdateObjectsRunsMetadataParams) SetLegalHoldNil(b bool)`

 SetLegalHoldNil sets the value for LegalHold to be an explicit nil

### UnsetLegalHold
`func (o *UpdateObjectsRunsMetadataParams) UnsetLegalHold()`

UnsetLegalHold ensures that no value is present for LegalHold, not even an explicit nil
### GetRetention

`func (o *UpdateObjectsRunsMetadataParams) GetRetention() Retention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *UpdateObjectsRunsMetadataParams) GetRetentionOk() (*Retention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *UpdateObjectsRunsMetadataParams) SetRetention(v Retention)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *UpdateObjectsRunsMetadataParams) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetTargetObjectRuns

`func (o *UpdateObjectsRunsMetadataParams) GetTargetObjectRuns() []TargetObjectRun`

GetTargetObjectRuns returns the TargetObjectRuns field if non-nil, zero value otherwise.

### GetTargetObjectRunsOk

`func (o *UpdateObjectsRunsMetadataParams) GetTargetObjectRunsOk() (*[]TargetObjectRun, bool)`

GetTargetObjectRunsOk returns a tuple with the TargetObjectRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetObjectRuns

`func (o *UpdateObjectsRunsMetadataParams) SetTargetObjectRuns(v []TargetObjectRun)`

SetTargetObjectRuns sets TargetObjectRuns field to given value.

### HasTargetObjectRuns

`func (o *UpdateObjectsRunsMetadataParams) HasTargetObjectRuns() bool`

HasTargetObjectRuns returns a boolean if a field has been set.

### GetTimeRange

`func (o *UpdateObjectsRunsMetadataParams) GetTimeRange() TimeRangeUsecs`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *UpdateObjectsRunsMetadataParams) GetTimeRangeOk() (*TimeRangeUsecs, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *UpdateObjectsRunsMetadataParams) SetTimeRange(v TimeRangeUsecs)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *UpdateObjectsRunsMetadataParams) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


