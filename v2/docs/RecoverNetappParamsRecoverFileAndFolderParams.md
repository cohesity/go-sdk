# RecoverNetappParamsRecoverFileAndFolderParams

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

### NewRecoverNetappParamsRecoverFileAndFolderParams

`func NewRecoverNetappParamsRecoverFileAndFolderParams(filesAndFolders []NetappRecoverFileAndFolderInfo, targetEnvironment string, ) *RecoverNetappParamsRecoverFileAndFolderParams`

NewRecoverNetappParamsRecoverFileAndFolderParams instantiates a new RecoverNetappParamsRecoverFileAndFolderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverNetappParamsRecoverFileAndFolderParamsWithDefaults

`func NewRecoverNetappParamsRecoverFileAndFolderParamsWithDefaults() *RecoverNetappParamsRecoverFileAndFolderParams`

NewRecoverNetappParamsRecoverFileAndFolderParamsWithDefaults instantiates a new RecoverNetappParamsRecoverFileAndFolderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElastifileTargetParams

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetElastifileTargetParams() RecoverNetappFilesParamsElastifileTargetParams`

GetElastifileTargetParams returns the ElastifileTargetParams field if non-nil, zero value otherwise.

### GetElastifileTargetParamsOk

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetElastifileTargetParamsOk() (*RecoverNetappFilesParamsElastifileTargetParams, bool)`

GetElastifileTargetParamsOk returns a tuple with the ElastifileTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileTargetParams

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetElastifileTargetParams(v RecoverNetappFilesParamsElastifileTargetParams)`

SetElastifileTargetParams sets ElastifileTargetParams field to given value.

### HasElastifileTargetParams

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) HasElastifileTargetParams() bool`

HasElastifileTargetParams returns a boolean if a field has been set.

### SetElastifileTargetParamsNil

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetElastifileTargetParamsNil(b bool)`

 SetElastifileTargetParamsNil sets the value for ElastifileTargetParams to be an explicit nil

### UnsetElastifileTargetParams
`func (o *RecoverNetappParamsRecoverFileAndFolderParams) UnsetElastifileTargetParams()`

UnsetElastifileTargetParams ensures that no value is present for ElastifileTargetParams, not even an explicit nil
### GetFilesAndFolders

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetFilesAndFolders() []NetappRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetFilesAndFoldersOk() (*[]NetappRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetFilesAndFolders(v []NetappRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.


### SetFilesAndFoldersNil

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *RecoverNetappParamsRecoverFileAndFolderParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetFlashbladeTargetParams

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetFlashbladeTargetParams() RecoverNetappFilesParamsFlashbladeTargetParams`

GetFlashbladeTargetParams returns the FlashbladeTargetParams field if non-nil, zero value otherwise.

### GetFlashbladeTargetParamsOk

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetFlashbladeTargetParamsOk() (*RecoverNetappFilesParamsFlashbladeTargetParams, bool)`

GetFlashbladeTargetParamsOk returns a tuple with the FlashbladeTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeTargetParams

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetFlashbladeTargetParams(v RecoverNetappFilesParamsFlashbladeTargetParams)`

SetFlashbladeTargetParams sets FlashbladeTargetParams field to given value.

### HasFlashbladeTargetParams

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) HasFlashbladeTargetParams() bool`

HasFlashbladeTargetParams returns a boolean if a field has been set.

### SetFlashbladeTargetParamsNil

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetFlashbladeTargetParamsNil(b bool)`

 SetFlashbladeTargetParamsNil sets the value for FlashbladeTargetParams to be an explicit nil

### UnsetFlashbladeTargetParams
`func (o *RecoverNetappParamsRecoverFileAndFolderParams) UnsetFlashbladeTargetParams()`

UnsetFlashbladeTargetParams ensures that no value is present for FlashbladeTargetParams, not even an explicit nil
### GetGenericNasTargetParams

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetGenericNasTargetParams() RecoverNetappFilesParamsGenericNasTargetParams`

GetGenericNasTargetParams returns the GenericNasTargetParams field if non-nil, zero value otherwise.

### GetGenericNasTargetParamsOk

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetGenericNasTargetParamsOk() (*RecoverNetappFilesParamsGenericNasTargetParams, bool)`

GetGenericNasTargetParamsOk returns a tuple with the GenericNasTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasTargetParams

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetGenericNasTargetParams(v RecoverNetappFilesParamsGenericNasTargetParams)`

SetGenericNasTargetParams sets GenericNasTargetParams field to given value.

### HasGenericNasTargetParams

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) HasGenericNasTargetParams() bool`

HasGenericNasTargetParams returns a boolean if a field has been set.

### SetGenericNasTargetParamsNil

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetGenericNasTargetParamsNil(b bool)`

 SetGenericNasTargetParamsNil sets the value for GenericNasTargetParams to be an explicit nil

### UnsetGenericNasTargetParams
`func (o *RecoverNetappParamsRecoverFileAndFolderParams) UnsetGenericNasTargetParams()`

UnsetGenericNasTargetParams ensures that no value is present for GenericNasTargetParams, not even an explicit nil
### GetGpfsTargetParams

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetGpfsTargetParams() RecoverNetappFilesParamsGpfsTargetParams`

GetGpfsTargetParams returns the GpfsTargetParams field if non-nil, zero value otherwise.

### GetGpfsTargetParamsOk

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetGpfsTargetParamsOk() (*RecoverNetappFilesParamsGpfsTargetParams, bool)`

GetGpfsTargetParamsOk returns a tuple with the GpfsTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsTargetParams

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetGpfsTargetParams(v RecoverNetappFilesParamsGpfsTargetParams)`

SetGpfsTargetParams sets GpfsTargetParams field to given value.

### HasGpfsTargetParams

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) HasGpfsTargetParams() bool`

HasGpfsTargetParams returns a boolean if a field has been set.

### SetGpfsTargetParamsNil

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetGpfsTargetParamsNil(b bool)`

 SetGpfsTargetParamsNil sets the value for GpfsTargetParams to be an explicit nil

### UnsetGpfsTargetParams
`func (o *RecoverNetappParamsRecoverFileAndFolderParams) UnsetGpfsTargetParams()`

UnsetGpfsTargetParams ensures that no value is present for GpfsTargetParams, not even an explicit nil
### GetIsFromSourceInitiatedProtection

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetIsFromSourceInitiatedProtection() bool`

GetIsFromSourceInitiatedProtection returns the IsFromSourceInitiatedProtection field if non-nil, zero value otherwise.

### GetIsFromSourceInitiatedProtectionOk

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetIsFromSourceInitiatedProtectionOk() (*bool, bool)`

GetIsFromSourceInitiatedProtectionOk returns a tuple with the IsFromSourceInitiatedProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFromSourceInitiatedProtection

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetIsFromSourceInitiatedProtection(v bool)`

SetIsFromSourceInitiatedProtection sets IsFromSourceInitiatedProtection field to given value.

### HasIsFromSourceInitiatedProtection

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) HasIsFromSourceInitiatedProtection() bool`

HasIsFromSourceInitiatedProtection returns a boolean if a field has been set.

### SetIsFromSourceInitiatedProtectionNil

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetIsFromSourceInitiatedProtectionNil(b bool)`

 SetIsFromSourceInitiatedProtectionNil sets the value for IsFromSourceInitiatedProtection to be an explicit nil

### UnsetIsFromSourceInitiatedProtection
`func (o *RecoverNetappParamsRecoverFileAndFolderParams) UnsetIsFromSourceInitiatedProtection()`

UnsetIsFromSourceInitiatedProtection ensures that no value is present for IsFromSourceInitiatedProtection, not even an explicit nil
### GetIsilonTargetParams

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetIsilonTargetParams() RecoverNetappFilesParamsIsilonTargetParams`

GetIsilonTargetParams returns the IsilonTargetParams field if non-nil, zero value otherwise.

### GetIsilonTargetParamsOk

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetIsilonTargetParamsOk() (*RecoverNetappFilesParamsIsilonTargetParams, bool)`

GetIsilonTargetParamsOk returns a tuple with the IsilonTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonTargetParams

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetIsilonTargetParams(v RecoverNetappFilesParamsIsilonTargetParams)`

SetIsilonTargetParams sets IsilonTargetParams field to given value.

### HasIsilonTargetParams

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) HasIsilonTargetParams() bool`

HasIsilonTargetParams returns a boolean if a field has been set.

### SetIsilonTargetParamsNil

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetIsilonTargetParamsNil(b bool)`

 SetIsilonTargetParamsNil sets the value for IsilonTargetParams to be an explicit nil

### UnsetIsilonTargetParams
`func (o *RecoverNetappParamsRecoverFileAndFolderParams) UnsetIsilonTargetParams()`

UnsetIsilonTargetParams ensures that no value is present for IsilonTargetParams, not even an explicit nil
### GetNetappTargetParams

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetNetappTargetParams() RecoverNetappFilesParamsNetappTargetParams`

GetNetappTargetParams returns the NetappTargetParams field if non-nil, zero value otherwise.

### GetNetappTargetParamsOk

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetNetappTargetParamsOk() (*RecoverNetappFilesParamsNetappTargetParams, bool)`

GetNetappTargetParamsOk returns a tuple with the NetappTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappTargetParams

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetNetappTargetParams(v RecoverNetappFilesParamsNetappTargetParams)`

SetNetappTargetParams sets NetappTargetParams field to given value.

### HasNetappTargetParams

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) HasNetappTargetParams() bool`

HasNetappTargetParams returns a boolean if a field has been set.

### SetNetappTargetParamsNil

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetNetappTargetParamsNil(b bool)`

 SetNetappTargetParamsNil sets the value for NetappTargetParams to be an explicit nil

### UnsetNetappTargetParams
`func (o *RecoverNetappParamsRecoverFileAndFolderParams) UnsetNetappTargetParams()`

UnsetNetappTargetParams ensures that no value is present for NetappTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


