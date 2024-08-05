# DownloadFilesAndFoldersRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documents** | Pointer to [**[]DocumentObject**](DocumentObject.md) | Specifies the list of documents to download using item ids. Only one of filesAndFolders or documents should be used. Currently only files are supported by documents. | [optional] 
**FilesAndFolders** | Pointer to [**[]FilesAndFoldersObject**](FilesAndFoldersObject.md) | Specifies the list of files and folders to download. Only one of filesAndFolders or documents should be used. | [optional] 
**GlacierRetrievalType** | Pointer to **NullableString** | Specifies the glacier retrieval type when restoring or downloding files or folders from a Glacier-based cloud snapshot. | [optional] 
**Name** | **NullableString** | Specifies the name of the recovery task. This field must be set and must be a unique name. | 
**Object** | [**CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) |  | 
**ParentRecoveryId** | Pointer to **NullableString** | If current recovery is child task triggered through another parent recovery operation, then this field will specify the id of the parent recovery. | [optional] 

## Methods

### NewDownloadFilesAndFoldersRequestParams

`func NewDownloadFilesAndFoldersRequestParams(name NullableString, object CommonRecoverObjectSnapshotParams, ) *DownloadFilesAndFoldersRequestParams`

NewDownloadFilesAndFoldersRequestParams instantiates a new DownloadFilesAndFoldersRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadFilesAndFoldersRequestParamsWithDefaults

`func NewDownloadFilesAndFoldersRequestParamsWithDefaults() *DownloadFilesAndFoldersRequestParams`

NewDownloadFilesAndFoldersRequestParamsWithDefaults instantiates a new DownloadFilesAndFoldersRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocuments

`func (o *DownloadFilesAndFoldersRequestParams) GetDocuments() []DocumentObject`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *DownloadFilesAndFoldersRequestParams) GetDocumentsOk() (*[]DocumentObject, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *DownloadFilesAndFoldersRequestParams) SetDocuments(v []DocumentObject)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *DownloadFilesAndFoldersRequestParams) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### SetDocumentsNil

`func (o *DownloadFilesAndFoldersRequestParams) SetDocumentsNil(b bool)`

 SetDocumentsNil sets the value for Documents to be an explicit nil

### UnsetDocuments
`func (o *DownloadFilesAndFoldersRequestParams) UnsetDocuments()`

UnsetDocuments ensures that no value is present for Documents, not even an explicit nil
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

### HasFilesAndFolders

`func (o *DownloadFilesAndFoldersRequestParams) HasFilesAndFolders() bool`

HasFilesAndFolders returns a boolean if a field has been set.

### SetFilesAndFoldersNil

`func (o *DownloadFilesAndFoldersRequestParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *DownloadFilesAndFoldersRequestParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetGlacierRetrievalType

`func (o *DownloadFilesAndFoldersRequestParams) GetGlacierRetrievalType() string`

GetGlacierRetrievalType returns the GlacierRetrievalType field if non-nil, zero value otherwise.

### GetGlacierRetrievalTypeOk

`func (o *DownloadFilesAndFoldersRequestParams) GetGlacierRetrievalTypeOk() (*string, bool)`

GetGlacierRetrievalTypeOk returns a tuple with the GlacierRetrievalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlacierRetrievalType

`func (o *DownloadFilesAndFoldersRequestParams) SetGlacierRetrievalType(v string)`

SetGlacierRetrievalType sets GlacierRetrievalType field to given value.

### HasGlacierRetrievalType

`func (o *DownloadFilesAndFoldersRequestParams) HasGlacierRetrievalType() bool`

HasGlacierRetrievalType returns a boolean if a field has been set.

### SetGlacierRetrievalTypeNil

`func (o *DownloadFilesAndFoldersRequestParams) SetGlacierRetrievalTypeNil(b bool)`

 SetGlacierRetrievalTypeNil sets the value for GlacierRetrievalType to be an explicit nil

### UnsetGlacierRetrievalType
`func (o *DownloadFilesAndFoldersRequestParams) UnsetGlacierRetrievalType()`

UnsetGlacierRetrievalType ensures that no value is present for GlacierRetrievalType, not even an explicit nil
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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


