# CommonRecoveryRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the name of the Recovery. | 
**SnapshotEnvironment** | **string** | Specifies the type of environment of snapshots for which the Recovery has to be performed. | 

## Methods

### NewCommonRecoveryRequestParams

`func NewCommonRecoveryRequestParams(name NullableString, snapshotEnvironment string, ) *CommonRecoveryRequestParams`

NewCommonRecoveryRequestParams instantiates a new CommonRecoveryRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonRecoveryRequestParamsWithDefaults

`func NewCommonRecoveryRequestParamsWithDefaults() *CommonRecoveryRequestParams`

NewCommonRecoveryRequestParamsWithDefaults instantiates a new CommonRecoveryRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CommonRecoveryRequestParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonRecoveryRequestParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonRecoveryRequestParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CommonRecoveryRequestParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonRecoveryRequestParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSnapshotEnvironment

`func (o *CommonRecoveryRequestParams) GetSnapshotEnvironment() string`

GetSnapshotEnvironment returns the SnapshotEnvironment field if non-nil, zero value otherwise.

### GetSnapshotEnvironmentOk

`func (o *CommonRecoveryRequestParams) GetSnapshotEnvironmentOk() (*string, bool)`

GetSnapshotEnvironmentOk returns a tuple with the SnapshotEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotEnvironment

`func (o *CommonRecoveryRequestParams) SetSnapshotEnvironment(v string)`

SetSnapshotEnvironment sets SnapshotEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


