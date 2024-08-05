# RecoverGenericNasParamsRecoverFileAndFolderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElastifileTargetParams** | Pointer to [**NullableRecoverFlashbladeFilesParamsElastifileTargetParams**](RecoverFlashbladeFilesParamsElastifileTargetParams.md) |  | [optional] 
**FilesAndFolders** | [**[]CommonRecoverFileAndFolderInfo**](CommonRecoverFileAndFolderInfo.md) | Specifies the info about the files and folders to be recovered. | 
**FlashbladeTargetParams** | Pointer to [**NullableRecoverGenericNasFilesParamsFlashbladeTargetParams**](RecoverGenericNasFilesParamsFlashbladeTargetParams.md) |  | [optional] 
**GenericNasTargetParams** | Pointer to [**NullableRecoverGenericNasFilesParamsGenericNasTargetParams**](RecoverGenericNasFilesParamsGenericNasTargetParams.md) |  | [optional] 
**GpfsTargetParams** | Pointer to [**NullableRecoverElastifileFilesParamsGpfsTargetParams**](RecoverElastifileFilesParamsGpfsTargetParams.md) |  | [optional] 
**IsilonTargetParams** | Pointer to [**NullableRecoverElastifileFilesParamsIsilonTargetParams**](RecoverElastifileFilesParamsIsilonTargetParams.md) |  | [optional] 
**NetappTargetParams** | Pointer to [**NullableRecoverElastifileFilesParamsNetappTargetParams**](RecoverElastifileFilesParamsNetappTargetParams.md) |  | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverGenericNasParamsRecoverFileAndFolderParams

`func NewRecoverGenericNasParamsRecoverFileAndFolderParams(filesAndFolders []CommonRecoverFileAndFolderInfo, targetEnvironment string, ) *RecoverGenericNasParamsRecoverFileAndFolderParams`

NewRecoverGenericNasParamsRecoverFileAndFolderParams instantiates a new RecoverGenericNasParamsRecoverFileAndFolderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverGenericNasParamsRecoverFileAndFolderParamsWithDefaults

`func NewRecoverGenericNasParamsRecoverFileAndFolderParamsWithDefaults() *RecoverGenericNasParamsRecoverFileAndFolderParams`

NewRecoverGenericNasParamsRecoverFileAndFolderParamsWithDefaults instantiates a new RecoverGenericNasParamsRecoverFileAndFolderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElastifileTargetParams

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) GetElastifileTargetParams() RecoverFlashbladeFilesParamsElastifileTargetParams`

GetElastifileTargetParams returns the ElastifileTargetParams field if non-nil, zero value otherwise.

### GetElastifileTargetParamsOk

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) GetElastifileTargetParamsOk() (*RecoverFlashbladeFilesParamsElastifileTargetParams, bool)`

GetElastifileTargetParamsOk returns a tuple with the ElastifileTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileTargetParams

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) SetElastifileTargetParams(v RecoverFlashbladeFilesParamsElastifileTargetParams)`

SetElastifileTargetParams sets ElastifileTargetParams field to given value.

### HasElastifileTargetParams

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) HasElastifileTargetParams() bool`

HasElastifileTargetParams returns a boolean if a field has been set.

### SetElastifileTargetParamsNil

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) SetElastifileTargetParamsNil(b bool)`

 SetElastifileTargetParamsNil sets the value for ElastifileTargetParams to be an explicit nil

### UnsetElastifileTargetParams
`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) UnsetElastifileTargetParams()`

UnsetElastifileTargetParams ensures that no value is present for ElastifileTargetParams, not even an explicit nil
### GetFilesAndFolders

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) GetFilesAndFoldersOk() (*[]CommonRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.


### SetFilesAndFoldersNil

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetFlashbladeTargetParams

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) GetFlashbladeTargetParams() RecoverGenericNasFilesParamsFlashbladeTargetParams`

GetFlashbladeTargetParams returns the FlashbladeTargetParams field if non-nil, zero value otherwise.

### GetFlashbladeTargetParamsOk

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) GetFlashbladeTargetParamsOk() (*RecoverGenericNasFilesParamsFlashbladeTargetParams, bool)`

GetFlashbladeTargetParamsOk returns a tuple with the FlashbladeTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeTargetParams

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) SetFlashbladeTargetParams(v RecoverGenericNasFilesParamsFlashbladeTargetParams)`

SetFlashbladeTargetParams sets FlashbladeTargetParams field to given value.

### HasFlashbladeTargetParams

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) HasFlashbladeTargetParams() bool`

HasFlashbladeTargetParams returns a boolean if a field has been set.

### SetFlashbladeTargetParamsNil

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) SetFlashbladeTargetParamsNil(b bool)`

 SetFlashbladeTargetParamsNil sets the value for FlashbladeTargetParams to be an explicit nil

### UnsetFlashbladeTargetParams
`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) UnsetFlashbladeTargetParams()`

UnsetFlashbladeTargetParams ensures that no value is present for FlashbladeTargetParams, not even an explicit nil
### GetGenericNasTargetParams

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) GetGenericNasTargetParams() RecoverGenericNasFilesParamsGenericNasTargetParams`

GetGenericNasTargetParams returns the GenericNasTargetParams field if non-nil, zero value otherwise.

### GetGenericNasTargetParamsOk

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) GetGenericNasTargetParamsOk() (*RecoverGenericNasFilesParamsGenericNasTargetParams, bool)`

GetGenericNasTargetParamsOk returns a tuple with the GenericNasTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasTargetParams

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) SetGenericNasTargetParams(v RecoverGenericNasFilesParamsGenericNasTargetParams)`

SetGenericNasTargetParams sets GenericNasTargetParams field to given value.

### HasGenericNasTargetParams

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) HasGenericNasTargetParams() bool`

HasGenericNasTargetParams returns a boolean if a field has been set.

### SetGenericNasTargetParamsNil

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) SetGenericNasTargetParamsNil(b bool)`

 SetGenericNasTargetParamsNil sets the value for GenericNasTargetParams to be an explicit nil

### UnsetGenericNasTargetParams
`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) UnsetGenericNasTargetParams()`

UnsetGenericNasTargetParams ensures that no value is present for GenericNasTargetParams, not even an explicit nil
### GetGpfsTargetParams

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) GetGpfsTargetParams() RecoverElastifileFilesParamsGpfsTargetParams`

GetGpfsTargetParams returns the GpfsTargetParams field if non-nil, zero value otherwise.

### GetGpfsTargetParamsOk

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) GetGpfsTargetParamsOk() (*RecoverElastifileFilesParamsGpfsTargetParams, bool)`

GetGpfsTargetParamsOk returns a tuple with the GpfsTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsTargetParams

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) SetGpfsTargetParams(v RecoverElastifileFilesParamsGpfsTargetParams)`

SetGpfsTargetParams sets GpfsTargetParams field to given value.

### HasGpfsTargetParams

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) HasGpfsTargetParams() bool`

HasGpfsTargetParams returns a boolean if a field has been set.

### SetGpfsTargetParamsNil

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) SetGpfsTargetParamsNil(b bool)`

 SetGpfsTargetParamsNil sets the value for GpfsTargetParams to be an explicit nil

### UnsetGpfsTargetParams
`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) UnsetGpfsTargetParams()`

UnsetGpfsTargetParams ensures that no value is present for GpfsTargetParams, not even an explicit nil
### GetIsilonTargetParams

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) GetIsilonTargetParams() RecoverElastifileFilesParamsIsilonTargetParams`

GetIsilonTargetParams returns the IsilonTargetParams field if non-nil, zero value otherwise.

### GetIsilonTargetParamsOk

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) GetIsilonTargetParamsOk() (*RecoverElastifileFilesParamsIsilonTargetParams, bool)`

GetIsilonTargetParamsOk returns a tuple with the IsilonTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonTargetParams

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) SetIsilonTargetParams(v RecoverElastifileFilesParamsIsilonTargetParams)`

SetIsilonTargetParams sets IsilonTargetParams field to given value.

### HasIsilonTargetParams

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) HasIsilonTargetParams() bool`

HasIsilonTargetParams returns a boolean if a field has been set.

### SetIsilonTargetParamsNil

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) SetIsilonTargetParamsNil(b bool)`

 SetIsilonTargetParamsNil sets the value for IsilonTargetParams to be an explicit nil

### UnsetIsilonTargetParams
`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) UnsetIsilonTargetParams()`

UnsetIsilonTargetParams ensures that no value is present for IsilonTargetParams, not even an explicit nil
### GetNetappTargetParams

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) GetNetappTargetParams() RecoverElastifileFilesParamsNetappTargetParams`

GetNetappTargetParams returns the NetappTargetParams field if non-nil, zero value otherwise.

### GetNetappTargetParamsOk

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) GetNetappTargetParamsOk() (*RecoverElastifileFilesParamsNetappTargetParams, bool)`

GetNetappTargetParamsOk returns a tuple with the NetappTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappTargetParams

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) SetNetappTargetParams(v RecoverElastifileFilesParamsNetappTargetParams)`

SetNetappTargetParams sets NetappTargetParams field to given value.

### HasNetappTargetParams

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) HasNetappTargetParams() bool`

HasNetappTargetParams returns a boolean if a field has been set.

### SetNetappTargetParamsNil

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) SetNetappTargetParamsNil(b bool)`

 SetNetappTargetParamsNil sets the value for NetappTargetParams to be an explicit nil

### UnsetNetappTargetParams
`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) UnsetNetappTargetParams()`

UnsetNetappTargetParams ensures that no value is present for NetappTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverGenericNasParamsRecoverFileAndFolderParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


