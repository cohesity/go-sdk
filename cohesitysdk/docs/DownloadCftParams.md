# DownloadCftParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **NullableString** | Specifies the file name of the cloud formation template. | [optional] 
**FilePath** | Pointer to **NullableString** | Specifies the file path of the template. If passed null, \&quot;/home/cohesity/bin\&quot; will be considered as file path. | [optional] 

## Methods

### NewDownloadCftParams

`func NewDownloadCftParams() *DownloadCftParams`

NewDownloadCftParams instantiates a new DownloadCftParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadCftParamsWithDefaults

`func NewDownloadCftParamsWithDefaults() *DownloadCftParams`

NewDownloadCftParamsWithDefaults instantiates a new DownloadCftParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *DownloadCftParams) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DownloadCftParams) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DownloadCftParams) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *DownloadCftParams) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *DownloadCftParams) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *DownloadCftParams) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetFilePath

`func (o *DownloadCftParams) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *DownloadCftParams) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *DownloadCftParams) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *DownloadCftParams) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *DownloadCftParams) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *DownloadCftParams) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


