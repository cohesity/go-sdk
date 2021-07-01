# DownloadFilesAndFoldersParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesAndFoldersInfo** | Pointer to [**[]FilesAndFoldersInfo**](FilesAndFoldersInfo.md) | Specifies the absolute paths for list of files and folders to download. | [optional] 
**Name** | **NullableString** | Specifies the name of the Download Task. This field must be set and must be a unique name. | 
**SourceObjectInfo** | Pointer to [**RestoreObjectDetails**](RestoreObjectDetails.md) |  | [optional] 

## Methods

### NewDownloadFilesAndFoldersParams

`func NewDownloadFilesAndFoldersParams(name NullableString, ) *DownloadFilesAndFoldersParams`

NewDownloadFilesAndFoldersParams instantiates a new DownloadFilesAndFoldersParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadFilesAndFoldersParamsWithDefaults

`func NewDownloadFilesAndFoldersParamsWithDefaults() *DownloadFilesAndFoldersParams`

NewDownloadFilesAndFoldersParamsWithDefaults instantiates a new DownloadFilesAndFoldersParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesAndFoldersInfo

`func (o *DownloadFilesAndFoldersParams) GetFilesAndFoldersInfo() []FilesAndFoldersInfo`

GetFilesAndFoldersInfo returns the FilesAndFoldersInfo field if non-nil, zero value otherwise.

### GetFilesAndFoldersInfoOk

`func (o *DownloadFilesAndFoldersParams) GetFilesAndFoldersInfoOk() (*[]FilesAndFoldersInfo, bool)`

GetFilesAndFoldersInfoOk returns a tuple with the FilesAndFoldersInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFoldersInfo

`func (o *DownloadFilesAndFoldersParams) SetFilesAndFoldersInfo(v []FilesAndFoldersInfo)`

SetFilesAndFoldersInfo sets FilesAndFoldersInfo field to given value.

### HasFilesAndFoldersInfo

`func (o *DownloadFilesAndFoldersParams) HasFilesAndFoldersInfo() bool`

HasFilesAndFoldersInfo returns a boolean if a field has been set.

### SetFilesAndFoldersInfoNil

`func (o *DownloadFilesAndFoldersParams) SetFilesAndFoldersInfoNil(b bool)`

 SetFilesAndFoldersInfoNil sets the value for FilesAndFoldersInfo to be an explicit nil

### UnsetFilesAndFoldersInfo
`func (o *DownloadFilesAndFoldersParams) UnsetFilesAndFoldersInfo()`

UnsetFilesAndFoldersInfo ensures that no value is present for FilesAndFoldersInfo, not even an explicit nil
### GetName

`func (o *DownloadFilesAndFoldersParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DownloadFilesAndFoldersParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DownloadFilesAndFoldersParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *DownloadFilesAndFoldersParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DownloadFilesAndFoldersParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSourceObjectInfo

`func (o *DownloadFilesAndFoldersParams) GetSourceObjectInfo() RestoreObjectDetails`

GetSourceObjectInfo returns the SourceObjectInfo field if non-nil, zero value otherwise.

### GetSourceObjectInfoOk

`func (o *DownloadFilesAndFoldersParams) GetSourceObjectInfoOk() (*RestoreObjectDetails, bool)`

GetSourceObjectInfoOk returns a tuple with the SourceObjectInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObjectInfo

`func (o *DownloadFilesAndFoldersParams) SetSourceObjectInfo(v RestoreObjectDetails)`

SetSourceObjectInfo sets SourceObjectInfo field to given value.

### HasSourceObjectInfo

`func (o *DownloadFilesAndFoldersParams) HasSourceObjectInfo() bool`

HasSourceObjectInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


