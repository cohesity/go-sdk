# BackupRunFilterParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunTypes** | Pointer to **[]string** | Specifies the run types to filter backup runs. | [optional] 
**ProtectionEnvironmentTypes** | Pointer to **[]string** | Specifies the protection environment types to filter backup runs. | [optional] 

## Methods

### NewBackupRunFilterParams

`func NewBackupRunFilterParams() *BackupRunFilterParams`

NewBackupRunFilterParams instantiates a new BackupRunFilterParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRunFilterParamsWithDefaults

`func NewBackupRunFilterParamsWithDefaults() *BackupRunFilterParams`

NewBackupRunFilterParamsWithDefaults instantiates a new BackupRunFilterParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunTypes

`func (o *BackupRunFilterParams) GetRunTypes() []string`

GetRunTypes returns the RunTypes field if non-nil, zero value otherwise.

### GetRunTypesOk

`func (o *BackupRunFilterParams) GetRunTypesOk() (*[]string, bool)`

GetRunTypesOk returns a tuple with the RunTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTypes

`func (o *BackupRunFilterParams) SetRunTypes(v []string)`

SetRunTypes sets RunTypes field to given value.

### HasRunTypes

`func (o *BackupRunFilterParams) HasRunTypes() bool`

HasRunTypes returns a boolean if a field has been set.

### SetRunTypesNil

`func (o *BackupRunFilterParams) SetRunTypesNil(b bool)`

 SetRunTypesNil sets the value for RunTypes to be an explicit nil

### UnsetRunTypes
`func (o *BackupRunFilterParams) UnsetRunTypes()`

UnsetRunTypes ensures that no value is present for RunTypes, not even an explicit nil
### GetProtectionEnvironmentTypes

`func (o *BackupRunFilterParams) GetProtectionEnvironmentTypes() []string`

GetProtectionEnvironmentTypes returns the ProtectionEnvironmentTypes field if non-nil, zero value otherwise.

### GetProtectionEnvironmentTypesOk

`func (o *BackupRunFilterParams) GetProtectionEnvironmentTypesOk() (*[]string, bool)`

GetProtectionEnvironmentTypesOk returns a tuple with the ProtectionEnvironmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionEnvironmentTypes

`func (o *BackupRunFilterParams) SetProtectionEnvironmentTypes(v []string)`

SetProtectionEnvironmentTypes sets ProtectionEnvironmentTypes field to given value.

### HasProtectionEnvironmentTypes

`func (o *BackupRunFilterParams) HasProtectionEnvironmentTypes() bool`

HasProtectionEnvironmentTypes returns a boolean if a field has been set.

### SetProtectionEnvironmentTypesNil

`func (o *BackupRunFilterParams) SetProtectionEnvironmentTypesNil(b bool)`

 SetProtectionEnvironmentTypesNil sets the value for ProtectionEnvironmentTypes to be an explicit nil

### UnsetProtectionEnvironmentTypes
`func (o *BackupRunFilterParams) UnsetProtectionEnvironmentTypes()`

UnsetProtectionEnvironmentTypes ensures that no value is present for ProtectionEnvironmentTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


