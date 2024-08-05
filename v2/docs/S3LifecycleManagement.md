# S3LifecycleManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]LifecycleRule**](LifecycleRule.md) | Specifies Lifecycle configuration rules for an Amazon S3 bucket. A maximum of 1000 rules can be specified. | [optional] 
**VersionId** | Pointer to **NullableInt64** | Specifies a unique monotonically increasing version for the lifecycle configuration. | [optional] 

## Methods

### NewS3LifecycleManagement

`func NewS3LifecycleManagement() *S3LifecycleManagement`

NewS3LifecycleManagement instantiates a new S3LifecycleManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LifecycleManagementWithDefaults

`func NewS3LifecycleManagementWithDefaults() *S3LifecycleManagement`

NewS3LifecycleManagementWithDefaults instantiates a new S3LifecycleManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *S3LifecycleManagement) GetRules() []LifecycleRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *S3LifecycleManagement) GetRulesOk() (*[]LifecycleRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *S3LifecycleManagement) SetRules(v []LifecycleRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *S3LifecycleManagement) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *S3LifecycleManagement) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *S3LifecycleManagement) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil
### GetVersionId

`func (o *S3LifecycleManagement) GetVersionId() int64`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *S3LifecycleManagement) GetVersionIdOk() (*int64, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *S3LifecycleManagement) SetVersionId(v int64)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *S3LifecycleManagement) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### SetVersionIdNil

`func (o *S3LifecycleManagement) SetVersionIdNil(b bool)`

 SetVersionIdNil sets the value for VersionId to be an explicit nil

### UnsetVersionId
`func (o *S3LifecycleManagement) UnsetVersionId()`

UnsetVersionId ensures that no value is present for VersionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


