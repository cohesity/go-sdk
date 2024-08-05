# RecoverNetappFilesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElastifileTargetParams** | Pointer to [**NullableRecoverNetappFilesParamsElastifileTargetParams**](RecoverNetappFilesParamsElastifileTargetParams.md) |  | [optional] 
**FilesAndFolders** | [**[]NetappRecoverFileAndFolderInfo**](NetappRecoverFileAndFolderInfo.md) | Specifies the list of info about the netapp files and folders to be recovered. | 
**FlashbladeTargetParams** | Pointer to [**NullableRecoverNetappFilesParamsFlashbladeTargetParams**](RecoverNetappFilesParamsFlashbladeTargetParams.md) |  | [optional] 
**GenericNasTargetParams** | Pointer to [**NullableRecoverNetappFilesParamsGenericNasTargetParams**](RecoverNetappFilesParamsGenericNasTargetParams.md) |  | [optional] 
**GpfsTargetParams** | Pointer to [**NullableRecoverNetappFilesParamsGpfsTargetParams**](RecoverNetappFilesParamsGpfsTargetParams.md) |  | [optional] 
**IsFromSourceInitiatedProtection** | Pointer to **NullableBool** | Specifies if the snapshot trying to recover is from a source initiated protection. | [optional] 
**IsilonTargetParams** | Pointer to [**NullableRecoverNetappFilesParamsIsilonTargetParams**](RecoverNetappFilesParamsIsilonTargetParams.md) |  | [optional] 
**NetappTargetParams** | Pointer to [**NullableRecoverNetappFilesParamsNetappTargetParams**](RecoverNetappFilesParamsNetappTargetParams.md) |  | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverNetappFilesParams

`func NewRecoverNetappFilesParams(filesAndFolders []NetappRecoverFileAndFolderInfo, targetEnvironment string, ) *RecoverNetappFilesParams`

NewRecoverNetappFilesParams instantiates a new RecoverNetappFilesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverNetappFilesParamsWithDefaults

`func NewRecoverNetappFilesParamsWithDefaults() *RecoverNetappFilesParams`

NewRecoverNetappFilesParamsWithDefaults instantiates a new RecoverNetappFilesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElastifileTargetParams

`func (o *RecoverNetappFilesParams) GetElastifileTargetParams() RecoverNetappFilesParamsElastifileTargetParams`

GetElastifileTargetParams returns the ElastifileTargetParams field if non-nil, zero value otherwise.

### GetElastifileTargetParamsOk

`func (o *RecoverNetappFilesParams) GetElastifileTargetParamsOk() (*RecoverNetappFilesParamsElastifileTargetParams, bool)`

GetElastifileTargetParamsOk returns a tuple with the ElastifileTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileTargetParams

`func (o *RecoverNetappFilesParams) SetElastifileTargetParams(v RecoverNetappFilesParamsElastifileTargetParams)`

SetElastifileTargetParams sets ElastifileTargetParams field to given value.

### HasElastifileTargetParams

`func (o *RecoverNetappFilesParams) HasElastifileTargetParams() bool`

HasElastifileTargetParams returns a boolean if a field has been set.

### SetElastifileTargetParamsNil

`func (o *RecoverNetappFilesParams) SetElastifileTargetParamsNil(b bool)`

 SetElastifileTargetParamsNil sets the value for ElastifileTargetParams to be an explicit nil

### UnsetElastifileTargetParams
`func (o *RecoverNetappFilesParams) UnsetElastifileTargetParams()`

UnsetElastifileTargetParams ensures that no value is present for ElastifileTargetParams, not even an explicit nil
### GetFilesAndFolders

`func (o *RecoverNetappFilesParams) GetFilesAndFolders() []NetappRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *RecoverNetappFilesParams) GetFilesAndFoldersOk() (*[]NetappRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *RecoverNetappFilesParams) SetFilesAndFolders(v []NetappRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.


### SetFilesAndFoldersNil

`func (o *RecoverNetappFilesParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *RecoverNetappFilesParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetFlashbladeTargetParams

`func (o *RecoverNetappFilesParams) GetFlashbladeTargetParams() RecoverNetappFilesParamsFlashbladeTargetParams`

GetFlashbladeTargetParams returns the FlashbladeTargetParams field if non-nil, zero value otherwise.

### GetFlashbladeTargetParamsOk

`func (o *RecoverNetappFilesParams) GetFlashbladeTargetParamsOk() (*RecoverNetappFilesParamsFlashbladeTargetParams, bool)`

GetFlashbladeTargetParamsOk returns a tuple with the FlashbladeTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeTargetParams

`func (o *RecoverNetappFilesParams) SetFlashbladeTargetParams(v RecoverNetappFilesParamsFlashbladeTargetParams)`

SetFlashbladeTargetParams sets FlashbladeTargetParams field to given value.

### HasFlashbladeTargetParams

`func (o *RecoverNetappFilesParams) HasFlashbladeTargetParams() bool`

HasFlashbladeTargetParams returns a boolean if a field has been set.

### SetFlashbladeTargetParamsNil

`func (o *RecoverNetappFilesParams) SetFlashbladeTargetParamsNil(b bool)`

 SetFlashbladeTargetParamsNil sets the value for FlashbladeTargetParams to be an explicit nil

### UnsetFlashbladeTargetParams
`func (o *RecoverNetappFilesParams) UnsetFlashbladeTargetParams()`

UnsetFlashbladeTargetParams ensures that no value is present for FlashbladeTargetParams, not even an explicit nil
### GetGenericNasTargetParams

`func (o *RecoverNetappFilesParams) GetGenericNasTargetParams() RecoverNetappFilesParamsGenericNasTargetParams`

GetGenericNasTargetParams returns the GenericNasTargetParams field if non-nil, zero value otherwise.

### GetGenericNasTargetParamsOk

`func (o *RecoverNetappFilesParams) GetGenericNasTargetParamsOk() (*RecoverNetappFilesParamsGenericNasTargetParams, bool)`

GetGenericNasTargetParamsOk returns a tuple with the GenericNasTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasTargetParams

`func (o *RecoverNetappFilesParams) SetGenericNasTargetParams(v RecoverNetappFilesParamsGenericNasTargetParams)`

SetGenericNasTargetParams sets GenericNasTargetParams field to given value.

### HasGenericNasTargetParams

`func (o *RecoverNetappFilesParams) HasGenericNasTargetParams() bool`

HasGenericNasTargetParams returns a boolean if a field has been set.

### SetGenericNasTargetParamsNil

`func (o *RecoverNetappFilesParams) SetGenericNasTargetParamsNil(b bool)`

 SetGenericNasTargetParamsNil sets the value for GenericNasTargetParams to be an explicit nil

### UnsetGenericNasTargetParams
`func (o *RecoverNetappFilesParams) UnsetGenericNasTargetParams()`

UnsetGenericNasTargetParams ensures that no value is present for GenericNasTargetParams, not even an explicit nil
### GetGpfsTargetParams

`func (o *RecoverNetappFilesParams) GetGpfsTargetParams() RecoverNetappFilesParamsGpfsTargetParams`

GetGpfsTargetParams returns the GpfsTargetParams field if non-nil, zero value otherwise.

### GetGpfsTargetParamsOk

`func (o *RecoverNetappFilesParams) GetGpfsTargetParamsOk() (*RecoverNetappFilesParamsGpfsTargetParams, bool)`

GetGpfsTargetParamsOk returns a tuple with the GpfsTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsTargetParams

`func (o *RecoverNetappFilesParams) SetGpfsTargetParams(v RecoverNetappFilesParamsGpfsTargetParams)`

SetGpfsTargetParams sets GpfsTargetParams field to given value.

### HasGpfsTargetParams

`func (o *RecoverNetappFilesParams) HasGpfsTargetParams() bool`

HasGpfsTargetParams returns a boolean if a field has been set.

### SetGpfsTargetParamsNil

`func (o *RecoverNetappFilesParams) SetGpfsTargetParamsNil(b bool)`

 SetGpfsTargetParamsNil sets the value for GpfsTargetParams to be an explicit nil

### UnsetGpfsTargetParams
`func (o *RecoverNetappFilesParams) UnsetGpfsTargetParams()`

UnsetGpfsTargetParams ensures that no value is present for GpfsTargetParams, not even an explicit nil
### GetIsFromSourceInitiatedProtection

`func (o *RecoverNetappFilesParams) GetIsFromSourceInitiatedProtection() bool`

GetIsFromSourceInitiatedProtection returns the IsFromSourceInitiatedProtection field if non-nil, zero value otherwise.

### GetIsFromSourceInitiatedProtectionOk

`func (o *RecoverNetappFilesParams) GetIsFromSourceInitiatedProtectionOk() (*bool, bool)`

GetIsFromSourceInitiatedProtectionOk returns a tuple with the IsFromSourceInitiatedProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFromSourceInitiatedProtection

`func (o *RecoverNetappFilesParams) SetIsFromSourceInitiatedProtection(v bool)`

SetIsFromSourceInitiatedProtection sets IsFromSourceInitiatedProtection field to given value.

### HasIsFromSourceInitiatedProtection

`func (o *RecoverNetappFilesParams) HasIsFromSourceInitiatedProtection() bool`

HasIsFromSourceInitiatedProtection returns a boolean if a field has been set.

### SetIsFromSourceInitiatedProtectionNil

`func (o *RecoverNetappFilesParams) SetIsFromSourceInitiatedProtectionNil(b bool)`

 SetIsFromSourceInitiatedProtectionNil sets the value for IsFromSourceInitiatedProtection to be an explicit nil

### UnsetIsFromSourceInitiatedProtection
`func (o *RecoverNetappFilesParams) UnsetIsFromSourceInitiatedProtection()`

UnsetIsFromSourceInitiatedProtection ensures that no value is present for IsFromSourceInitiatedProtection, not even an explicit nil
### GetIsilonTargetParams

`func (o *RecoverNetappFilesParams) GetIsilonTargetParams() RecoverNetappFilesParamsIsilonTargetParams`

GetIsilonTargetParams returns the IsilonTargetParams field if non-nil, zero value otherwise.

### GetIsilonTargetParamsOk

`func (o *RecoverNetappFilesParams) GetIsilonTargetParamsOk() (*RecoverNetappFilesParamsIsilonTargetParams, bool)`

GetIsilonTargetParamsOk returns a tuple with the IsilonTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonTargetParams

`func (o *RecoverNetappFilesParams) SetIsilonTargetParams(v RecoverNetappFilesParamsIsilonTargetParams)`

SetIsilonTargetParams sets IsilonTargetParams field to given value.

### HasIsilonTargetParams

`func (o *RecoverNetappFilesParams) HasIsilonTargetParams() bool`

HasIsilonTargetParams returns a boolean if a field has been set.

### SetIsilonTargetParamsNil

`func (o *RecoverNetappFilesParams) SetIsilonTargetParamsNil(b bool)`

 SetIsilonTargetParamsNil sets the value for IsilonTargetParams to be an explicit nil

### UnsetIsilonTargetParams
`func (o *RecoverNetappFilesParams) UnsetIsilonTargetParams()`

UnsetIsilonTargetParams ensures that no value is present for IsilonTargetParams, not even an explicit nil
### GetNetappTargetParams

`func (o *RecoverNetappFilesParams) GetNetappTargetParams() RecoverNetappFilesParamsNetappTargetParams`

GetNetappTargetParams returns the NetappTargetParams field if non-nil, zero value otherwise.

### GetNetappTargetParamsOk

`func (o *RecoverNetappFilesParams) GetNetappTargetParamsOk() (*RecoverNetappFilesParamsNetappTargetParams, bool)`

GetNetappTargetParamsOk returns a tuple with the NetappTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappTargetParams

`func (o *RecoverNetappFilesParams) SetNetappTargetParams(v RecoverNetappFilesParamsNetappTargetParams)`

SetNetappTargetParams sets NetappTargetParams field to given value.

### HasNetappTargetParams

`func (o *RecoverNetappFilesParams) HasNetappTargetParams() bool`

HasNetappTargetParams returns a boolean if a field has been set.

### SetNetappTargetParamsNil

`func (o *RecoverNetappFilesParams) SetNetappTargetParamsNil(b bool)`

 SetNetappTargetParamsNil sets the value for NetappTargetParams to be an explicit nil

### UnsetNetappTargetParams
`func (o *RecoverNetappFilesParams) UnsetNetappTargetParams()`

UnsetNetappTargetParams ensures that no value is present for NetappTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverNetappFilesParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverNetappFilesParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverNetappFilesParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


