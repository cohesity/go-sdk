# PhysicalBackupEnvParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableIncrementalBackupAfterRestart** | Pointer to **NullableBool** | If this is set to true, then incremental backup will be performed after the server restarts, otherwise a full-backup will be done. NOTE: This is applicable to windows host environments. | [optional] 
**FilteringPolicy** | Pointer to [**FilteringPolicyProto**](FilteringPolicyProto.md) |  | [optional] 

## Methods

### NewPhysicalBackupEnvParams

`func NewPhysicalBackupEnvParams() *PhysicalBackupEnvParams`

NewPhysicalBackupEnvParams instantiates a new PhysicalBackupEnvParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalBackupEnvParamsWithDefaults

`func NewPhysicalBackupEnvParamsWithDefaults() *PhysicalBackupEnvParams`

NewPhysicalBackupEnvParamsWithDefaults instantiates a new PhysicalBackupEnvParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableIncrementalBackupAfterRestart

`func (o *PhysicalBackupEnvParams) GetEnableIncrementalBackupAfterRestart() bool`

GetEnableIncrementalBackupAfterRestart returns the EnableIncrementalBackupAfterRestart field if non-nil, zero value otherwise.

### GetEnableIncrementalBackupAfterRestartOk

`func (o *PhysicalBackupEnvParams) GetEnableIncrementalBackupAfterRestartOk() (*bool, bool)`

GetEnableIncrementalBackupAfterRestartOk returns a tuple with the EnableIncrementalBackupAfterRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIncrementalBackupAfterRestart

`func (o *PhysicalBackupEnvParams) SetEnableIncrementalBackupAfterRestart(v bool)`

SetEnableIncrementalBackupAfterRestart sets EnableIncrementalBackupAfterRestart field to given value.

### HasEnableIncrementalBackupAfterRestart

`func (o *PhysicalBackupEnvParams) HasEnableIncrementalBackupAfterRestart() bool`

HasEnableIncrementalBackupAfterRestart returns a boolean if a field has been set.

### SetEnableIncrementalBackupAfterRestartNil

`func (o *PhysicalBackupEnvParams) SetEnableIncrementalBackupAfterRestartNil(b bool)`

 SetEnableIncrementalBackupAfterRestartNil sets the value for EnableIncrementalBackupAfterRestart to be an explicit nil

### UnsetEnableIncrementalBackupAfterRestart
`func (o *PhysicalBackupEnvParams) UnsetEnableIncrementalBackupAfterRestart()`

UnsetEnableIncrementalBackupAfterRestart ensures that no value is present for EnableIncrementalBackupAfterRestart, not even an explicit nil
### GetFilteringPolicy

`func (o *PhysicalBackupEnvParams) GetFilteringPolicy() FilteringPolicyProto`

GetFilteringPolicy returns the FilteringPolicy field if non-nil, zero value otherwise.

### GetFilteringPolicyOk

`func (o *PhysicalBackupEnvParams) GetFilteringPolicyOk() (*FilteringPolicyProto, bool)`

GetFilteringPolicyOk returns a tuple with the FilteringPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteringPolicy

`func (o *PhysicalBackupEnvParams) SetFilteringPolicy(v FilteringPolicyProto)`

SetFilteringPolicy sets FilteringPolicy field to given value.

### HasFilteringPolicy

`func (o *PhysicalBackupEnvParams) HasFilteringPolicy() bool`

HasFilteringPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


