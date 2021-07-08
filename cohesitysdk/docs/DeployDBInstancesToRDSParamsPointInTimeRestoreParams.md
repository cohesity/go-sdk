# DeployDBInstancesToRDSParamsPointInTimeRestoreParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimestampMsecs** | Pointer to **NullableInt64** | Time in milliseconds since epoch that the DB needs to be restored to. | [optional] 

## Methods

### NewDeployDBInstancesToRDSParamsPointInTimeRestoreParams

`func NewDeployDBInstancesToRDSParamsPointInTimeRestoreParams() *DeployDBInstancesToRDSParamsPointInTimeRestoreParams`

NewDeployDBInstancesToRDSParamsPointInTimeRestoreParams instantiates a new DeployDBInstancesToRDSParamsPointInTimeRestoreParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployDBInstancesToRDSParamsPointInTimeRestoreParamsWithDefaults

`func NewDeployDBInstancesToRDSParamsPointInTimeRestoreParamsWithDefaults() *DeployDBInstancesToRDSParamsPointInTimeRestoreParams`

NewDeployDBInstancesToRDSParamsPointInTimeRestoreParamsWithDefaults instantiates a new DeployDBInstancesToRDSParamsPointInTimeRestoreParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestampMsecs

`func (o *DeployDBInstancesToRDSParamsPointInTimeRestoreParams) GetTimestampMsecs() int64`

GetTimestampMsecs returns the TimestampMsecs field if non-nil, zero value otherwise.

### GetTimestampMsecsOk

`func (o *DeployDBInstancesToRDSParamsPointInTimeRestoreParams) GetTimestampMsecsOk() (*int64, bool)`

GetTimestampMsecsOk returns a tuple with the TimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampMsecs

`func (o *DeployDBInstancesToRDSParamsPointInTimeRestoreParams) SetTimestampMsecs(v int64)`

SetTimestampMsecs sets TimestampMsecs field to given value.

### HasTimestampMsecs

`func (o *DeployDBInstancesToRDSParamsPointInTimeRestoreParams) HasTimestampMsecs() bool`

HasTimestampMsecs returns a boolean if a field has been set.

### SetTimestampMsecsNil

`func (o *DeployDBInstancesToRDSParamsPointInTimeRestoreParams) SetTimestampMsecsNil(b bool)`

 SetTimestampMsecsNil sets the value for TimestampMsecs to be an explicit nil

### UnsetTimestampMsecs
`func (o *DeployDBInstancesToRDSParamsPointInTimeRestoreParams) UnsetTimestampMsecs()`

UnsetTimestampMsecs ensures that no value is present for TimestampMsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


