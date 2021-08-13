# DownloadFilesAndFoldersResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the id of the Recovery. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the recovery task. This field must be set and must be a unique name. | [optional] 
**Object** | Pointer to [**CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) |  | [optional] 
**FilesAndFolders** | Pointer to [**[]FilesAndFoldersObject**](FilesAndFoldersObject.md) | Specifies the list of files and folders to download. | [optional] 

## Methods

### NewDownloadFilesAndFoldersResponseParams

`func NewDownloadFilesAndFoldersResponseParams() *DownloadFilesAndFoldersResponseParams`

NewDownloadFilesAndFoldersResponseParams instantiates a new DownloadFilesAndFoldersResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadFilesAndFoldersResponseParamsWithDefaults

`func NewDownloadFilesAndFoldersResponseParamsWithDefaults() *DownloadFilesAndFoldersResponseParams`

NewDownloadFilesAndFoldersResponseParamsWithDefaults instantiates a new DownloadFilesAndFoldersResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DownloadFilesAndFoldersResponseParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DownloadFilesAndFoldersResponseParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DownloadFilesAndFoldersResponseParams) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DownloadFilesAndFoldersResponseParams) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DownloadFilesAndFoldersResponseParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DownloadFilesAndFoldersResponseParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *DownloadFilesAndFoldersResponseParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DownloadFilesAndFoldersResponseParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DownloadFilesAndFoldersResponseParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DownloadFilesAndFoldersResponseParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DownloadFilesAndFoldersResponseParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DownloadFilesAndFoldersResponseParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetObject

`func (o *DownloadFilesAndFoldersResponseParams) GetObject() CommonRecoverObjectSnapshotParams`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *DownloadFilesAndFoldersResponseParams) GetObjectOk() (*CommonRecoverObjectSnapshotParams, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *DownloadFilesAndFoldersResponseParams) SetObject(v CommonRecoverObjectSnapshotParams)`

SetObject sets Object field to given value.

### HasObject

`func (o *DownloadFilesAndFoldersResponseParams) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetFilesAndFolders

`func (o *DownloadFilesAndFoldersResponseParams) GetFilesAndFolders() []FilesAndFoldersObject`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *DownloadFilesAndFoldersResponseParams) GetFilesAndFoldersOk() (*[]FilesAndFoldersObject, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *DownloadFilesAndFoldersResponseParams) SetFilesAndFolders(v []FilesAndFoldersObject)`

SetFilesAndFolders sets FilesAndFolders field to given value.

### HasFilesAndFolders

`func (o *DownloadFilesAndFoldersResponseParams) HasFilesAndFolders() bool`

HasFilesAndFolders returns a boolean if a field has been set.

### SetFilesAndFoldersNil

`func (o *DownloadFilesAndFoldersResponseParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *DownloadFilesAndFoldersResponseParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


