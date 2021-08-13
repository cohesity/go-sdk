# DownloadFilesAndFoldersRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the name of the recovery task. This field must be set and must be a unique name. | 
**Object** | [**CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) |  | 
**ParentRecoveryId** | Pointer to **NullableString** | If current recovery is child task triggered through another parent recovery operation, then this field will specify the id of the parent recovery. | [optional] 
**FilesAndFolders** | [**[]FilesAndFoldersObject**](FilesAndFoldersObject.md) | Specifies the list of files and folders to download. | 

## Methods

### NewDownloadFilesAndFoldersRequestParams

`func NewDownloadFilesAndFoldersRequestParams(name NullableString, object CommonRecoverObjectSnapshotParams, filesAndFolders []FilesAndFoldersObject, ) *DownloadFilesAndFoldersRequestParams`

NewDownloadFilesAndFoldersRequestParams instantiates a new DownloadFilesAndFoldersRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadFilesAndFoldersRequestParamsWithDefaults

`func NewDownloadFilesAndFoldersRequestParamsWithDefaults() *DownloadFilesAndFoldersRequestParams`

NewDownloadFilesAndFoldersRequestParamsWithDefaults instantiates a new DownloadFilesAndFoldersRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DownloadFilesAndFoldersRequestParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DownloadFilesAndFoldersRequestParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DownloadFilesAndFoldersRequestParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *DownloadFilesAndFoldersRequestParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DownloadFilesAndFoldersRequestParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetObject

`func (o *DownloadFilesAndFoldersRequestParams) GetObject() CommonRecoverObjectSnapshotParams`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *DownloadFilesAndFoldersRequestParams) GetObjectOk() (*CommonRecoverObjectSnapshotParams, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *DownloadFilesAndFoldersRequestParams) SetObject(v CommonRecoverObjectSnapshotParams)`

SetObject sets Object field to given value.


### GetParentRecoveryId

`func (o *DownloadFilesAndFoldersRequestParams) GetParentRecoveryId() string`

GetParentRecoveryId returns the ParentRecoveryId field if non-nil, zero value otherwise.

### GetParentRecoveryIdOk

`func (o *DownloadFilesAndFoldersRequestParams) GetParentRecoveryIdOk() (*string, bool)`

GetParentRecoveryIdOk returns a tuple with the ParentRecoveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRecoveryId

`func (o *DownloadFilesAndFoldersRequestParams) SetParentRecoveryId(v string)`

SetParentRecoveryId sets ParentRecoveryId field to given value.

### HasParentRecoveryId

`func (o *DownloadFilesAndFoldersRequestParams) HasParentRecoveryId() bool`

HasParentRecoveryId returns a boolean if a field has been set.

### SetParentRecoveryIdNil

`func (o *DownloadFilesAndFoldersRequestParams) SetParentRecoveryIdNil(b bool)`

 SetParentRecoveryIdNil sets the value for ParentRecoveryId to be an explicit nil

### UnsetParentRecoveryId
`func (o *DownloadFilesAndFoldersRequestParams) UnsetParentRecoveryId()`

UnsetParentRecoveryId ensures that no value is present for ParentRecoveryId, not even an explicit nil
### GetFilesAndFolders

`func (o *DownloadFilesAndFoldersRequestParams) GetFilesAndFolders() []FilesAndFoldersObject`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *DownloadFilesAndFoldersRequestParams) GetFilesAndFoldersOk() (*[]FilesAndFoldersObject, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *DownloadFilesAndFoldersRequestParams) SetFilesAndFolders(v []FilesAndFoldersObject)`

SetFilesAndFolders sets FilesAndFolders field to given value.


### SetFilesAndFoldersNil

`func (o *DownloadFilesAndFoldersRequestParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *DownloadFilesAndFoldersRequestParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


