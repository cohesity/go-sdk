# RecoverAzureSqlParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureTargetParams** | Pointer to [**NullableRecoverAzureSqlParamsAzureTargetParams**](RecoverAzureSqlParamsAzureTargetParams.md) |  | [optional] 
**Snapshots** | [**[]RecoverAzureSqlSnapshotParams**](RecoverAzureSqlSnapshotParams.md) | Specifies the details of the azure sql objects to be recovered. | 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverAzureSqlParams

`func NewRecoverAzureSqlParams(snapshots []RecoverAzureSqlSnapshotParams, targetEnvironment string, ) *RecoverAzureSqlParams`

NewRecoverAzureSqlParams instantiates a new RecoverAzureSqlParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAzureSqlParamsWithDefaults

`func NewRecoverAzureSqlParamsWithDefaults() *RecoverAzureSqlParams`

NewRecoverAzureSqlParamsWithDefaults instantiates a new RecoverAzureSqlParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureTargetParams

`func (o *RecoverAzureSqlParams) GetAzureTargetParams() RecoverAzureSqlParamsAzureTargetParams`

GetAzureTargetParams returns the AzureTargetParams field if non-nil, zero value otherwise.

### GetAzureTargetParamsOk

`func (o *RecoverAzureSqlParams) GetAzureTargetParamsOk() (*RecoverAzureSqlParamsAzureTargetParams, bool)`

GetAzureTargetParamsOk returns a tuple with the AzureTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTargetParams

`func (o *RecoverAzureSqlParams) SetAzureTargetParams(v RecoverAzureSqlParamsAzureTargetParams)`

SetAzureTargetParams sets AzureTargetParams field to given value.

### HasAzureTargetParams

`func (o *RecoverAzureSqlParams) HasAzureTargetParams() bool`

HasAzureTargetParams returns a boolean if a field has been set.

### SetAzureTargetParamsNil

`func (o *RecoverAzureSqlParams) SetAzureTargetParamsNil(b bool)`

 SetAzureTargetParamsNil sets the value for AzureTargetParams to be an explicit nil

### UnsetAzureTargetParams
`func (o *RecoverAzureSqlParams) UnsetAzureTargetParams()`

UnsetAzureTargetParams ensures that no value is present for AzureTargetParams, not even an explicit nil
### GetSnapshots

`func (o *RecoverAzureSqlParams) GetSnapshots() []RecoverAzureSqlSnapshotParams`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *RecoverAzureSqlParams) GetSnapshotsOk() (*[]RecoverAzureSqlSnapshotParams, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *RecoverAzureSqlParams) SetSnapshots(v []RecoverAzureSqlSnapshotParams)`

SetSnapshots sets Snapshots field to given value.


### SetSnapshotsNil

`func (o *RecoverAzureSqlParams) SetSnapshotsNil(b bool)`

 SetSnapshotsNil sets the value for Snapshots to be an explicit nil

### UnsetSnapshots
`func (o *RecoverAzureSqlParams) UnsetSnapshots()`

UnsetSnapshots ensures that no value is present for Snapshots, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverAzureSqlParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverAzureSqlParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverAzureSqlParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


