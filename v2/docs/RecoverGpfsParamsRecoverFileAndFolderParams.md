# RecoverGpfsParamsRecoverFileAndFolderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElastifileTargetParams** | Pointer to [**NullableRecoverFlashbladeFilesParamsElastifileTargetParams**](RecoverFlashbladeFilesParamsElastifileTargetParams.md) |  | [optional] 
**FilesAndFolders** | [**[]CommonRecoverFileAndFolderInfo**](CommonRecoverFileAndFolderInfo.md) | Specifies the info about the files and folders to be recovered. | 
**FlashbladeTargetParams** | Pointer to [**NullableRecoverGenericNasFilesParamsFlashbladeTargetParams**](RecoverGenericNasFilesParamsFlashbladeTargetParams.md) |  | [optional] 
**GenericNasTargetParams** | Pointer to [**NullableRecoverGpfsFilesParamsGenericNasTargetParams**](RecoverGpfsFilesParamsGenericNasTargetParams.md) |  | [optional] 
**GpfsTargetParams** | Pointer to [**NullableRecoverGpfsFilesParamsGpfsTargetParams**](RecoverGpfsFilesParamsGpfsTargetParams.md) |  | [optional] 
**IsilonTargetParams** | Pointer to [**NullableRecoverGpfsFilesParamsIsilonTargetParams**](RecoverGpfsFilesParamsIsilonTargetParams.md) |  | [optional] 
**NetappTargetParams** | Pointer to [**NullableRecoverElastifileFilesParamsNetappTargetParams**](RecoverElastifileFilesParamsNetappTargetParams.md) |  | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverGpfsParamsRecoverFileAndFolderParams

`func NewRecoverGpfsParamsRecoverFileAndFolderParams(filesAndFolders []CommonRecoverFileAndFolderInfo, targetEnvironment string, ) *RecoverGpfsParamsRecoverFileAndFolderParams`

NewRecoverGpfsParamsRecoverFileAndFolderParams instantiates a new RecoverGpfsParamsRecoverFileAndFolderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverGpfsParamsRecoverFileAndFolderParamsWithDefaults

`func NewRecoverGpfsParamsRecoverFileAndFolderParamsWithDefaults() *RecoverGpfsParamsRecoverFileAndFolderParams`

NewRecoverGpfsParamsRecoverFileAndFolderParamsWithDefaults instantiates a new RecoverGpfsParamsRecoverFileAndFolderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElastifileTargetParams

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) GetElastifileTargetParams() RecoverFlashbladeFilesParamsElastifileTargetParams`

GetElastifileTargetParams returns the ElastifileTargetParams field if non-nil, zero value otherwise.

### GetElastifileTargetParamsOk

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) GetElastifileTargetParamsOk() (*RecoverFlashbladeFilesParamsElastifileTargetParams, bool)`

GetElastifileTargetParamsOk returns a tuple with the ElastifileTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileTargetParams

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) SetElastifileTargetParams(v RecoverFlashbladeFilesParamsElastifileTargetParams)`

SetElastifileTargetParams sets ElastifileTargetParams field to given value.

### HasElastifileTargetParams

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) HasElastifileTargetParams() bool`

HasElastifileTargetParams returns a boolean if a field has been set.

### SetElastifileTargetParamsNil

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) SetElastifileTargetParamsNil(b bool)`

 SetElastifileTargetParamsNil sets the value for ElastifileTargetParams to be an explicit nil

### UnsetElastifileTargetParams
`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) UnsetElastifileTargetParams()`

UnsetElastifileTargetParams ensures that no value is present for ElastifileTargetParams, not even an explicit nil
### GetFilesAndFolders

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) GetFilesAndFoldersOk() (*[]CommonRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.


### SetFilesAndFoldersNil

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetFlashbladeTargetParams

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) GetFlashbladeTargetParams() RecoverGenericNasFilesParamsFlashbladeTargetParams`

GetFlashbladeTargetParams returns the FlashbladeTargetParams field if non-nil, zero value otherwise.

### GetFlashbladeTargetParamsOk

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) GetFlashbladeTargetParamsOk() (*RecoverGenericNasFilesParamsFlashbladeTargetParams, bool)`

GetFlashbladeTargetParamsOk returns a tuple with the FlashbladeTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeTargetParams

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) SetFlashbladeTargetParams(v RecoverGenericNasFilesParamsFlashbladeTargetParams)`

SetFlashbladeTargetParams sets FlashbladeTargetParams field to given value.

### HasFlashbladeTargetParams

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) HasFlashbladeTargetParams() bool`

HasFlashbladeTargetParams returns a boolean if a field has been set.

### SetFlashbladeTargetParamsNil

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) SetFlashbladeTargetParamsNil(b bool)`

 SetFlashbladeTargetParamsNil sets the value for FlashbladeTargetParams to be an explicit nil

### UnsetFlashbladeTargetParams
`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) UnsetFlashbladeTargetParams()`

UnsetFlashbladeTargetParams ensures that no value is present for FlashbladeTargetParams, not even an explicit nil
### GetGenericNasTargetParams

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) GetGenericNasTargetParams() RecoverGpfsFilesParamsGenericNasTargetParams`

GetGenericNasTargetParams returns the GenericNasTargetParams field if non-nil, zero value otherwise.

### GetGenericNasTargetParamsOk

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) GetGenericNasTargetParamsOk() (*RecoverGpfsFilesParamsGenericNasTargetParams, bool)`

GetGenericNasTargetParamsOk returns a tuple with the GenericNasTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasTargetParams

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) SetGenericNasTargetParams(v RecoverGpfsFilesParamsGenericNasTargetParams)`

SetGenericNasTargetParams sets GenericNasTargetParams field to given value.

### HasGenericNasTargetParams

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) HasGenericNasTargetParams() bool`

HasGenericNasTargetParams returns a boolean if a field has been set.

### SetGenericNasTargetParamsNil

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) SetGenericNasTargetParamsNil(b bool)`

 SetGenericNasTargetParamsNil sets the value for GenericNasTargetParams to be an explicit nil

### UnsetGenericNasTargetParams
`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) UnsetGenericNasTargetParams()`

UnsetGenericNasTargetParams ensures that no value is present for GenericNasTargetParams, not even an explicit nil
### GetGpfsTargetParams

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) GetGpfsTargetParams() RecoverGpfsFilesParamsGpfsTargetParams`

GetGpfsTargetParams returns the GpfsTargetParams field if non-nil, zero value otherwise.

### GetGpfsTargetParamsOk

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) GetGpfsTargetParamsOk() (*RecoverGpfsFilesParamsGpfsTargetParams, bool)`

GetGpfsTargetParamsOk returns a tuple with the GpfsTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsTargetParams

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) SetGpfsTargetParams(v RecoverGpfsFilesParamsGpfsTargetParams)`

SetGpfsTargetParams sets GpfsTargetParams field to given value.

### HasGpfsTargetParams

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) HasGpfsTargetParams() bool`

HasGpfsTargetParams returns a boolean if a field has been set.

### SetGpfsTargetParamsNil

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) SetGpfsTargetParamsNil(b bool)`

 SetGpfsTargetParamsNil sets the value for GpfsTargetParams to be an explicit nil

### UnsetGpfsTargetParams
`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) UnsetGpfsTargetParams()`

UnsetGpfsTargetParams ensures that no value is present for GpfsTargetParams, not even an explicit nil
### GetIsilonTargetParams

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) GetIsilonTargetParams() RecoverGpfsFilesParamsIsilonTargetParams`

GetIsilonTargetParams returns the IsilonTargetParams field if non-nil, zero value otherwise.

### GetIsilonTargetParamsOk

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) GetIsilonTargetParamsOk() (*RecoverGpfsFilesParamsIsilonTargetParams, bool)`

GetIsilonTargetParamsOk returns a tuple with the IsilonTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonTargetParams

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) SetIsilonTargetParams(v RecoverGpfsFilesParamsIsilonTargetParams)`

SetIsilonTargetParams sets IsilonTargetParams field to given value.

### HasIsilonTargetParams

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) HasIsilonTargetParams() bool`

HasIsilonTargetParams returns a boolean if a field has been set.

### SetIsilonTargetParamsNil

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) SetIsilonTargetParamsNil(b bool)`

 SetIsilonTargetParamsNil sets the value for IsilonTargetParams to be an explicit nil

### UnsetIsilonTargetParams
`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) UnsetIsilonTargetParams()`

UnsetIsilonTargetParams ensures that no value is present for IsilonTargetParams, not even an explicit nil
### GetNetappTargetParams

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) GetNetappTargetParams() RecoverElastifileFilesParamsNetappTargetParams`

GetNetappTargetParams returns the NetappTargetParams field if non-nil, zero value otherwise.

### GetNetappTargetParamsOk

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) GetNetappTargetParamsOk() (*RecoverElastifileFilesParamsNetappTargetParams, bool)`

GetNetappTargetParamsOk returns a tuple with the NetappTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappTargetParams

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) SetNetappTargetParams(v RecoverElastifileFilesParamsNetappTargetParams)`

SetNetappTargetParams sets NetappTargetParams field to given value.

### HasNetappTargetParams

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) HasNetappTargetParams() bool`

HasNetappTargetParams returns a boolean if a field has been set.

### SetNetappTargetParamsNil

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) SetNetappTargetParamsNil(b bool)`

 SetNetappTargetParamsNil sets the value for NetappTargetParams to be an explicit nil

### UnsetNetappTargetParams
`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) UnsetNetappTargetParams()`

UnsetNetappTargetParams ensures that no value is present for NetappTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverGpfsParamsRecoverFileAndFolderParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


