# DownloadFilesAndFoldersResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documents** | Pointer to [**[]DocumentObject**](DocumentObject.md) | Specifies the list of documents to download using item ids. Only one of filesAndFolders or documents should be used. Currently only files are supported by documents. | [optional] 
**FilesAndFolders** | Pointer to [**[]FilesAndFoldersObject**](FilesAndFoldersObject.md) | Specifies the list of files and folders to download. Only one of filesAndFolders or documents should be used. | [optional] 
**GlacierRetrievalType** | Pointer to **NullableString** | Specifies the glacier retrieval type when restoring or downloding files or folders from a Glacier-based cloud snapshot. | [optional] 
**Id** | Pointer to **NullableString** | Specifies the id of the Recovery. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the recovery task. This field must be set and must be a unique name. | [optional] 
**Object** | Pointer to [**CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) |  | [optional] 

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

### GetDocuments

`func (o *DownloadFilesAndFoldersResponseParams) GetDocuments() []DocumentObject`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *DownloadFilesAndFoldersResponseParams) GetDocumentsOk() (*[]DocumentObject, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *DownloadFilesAndFoldersResponseParams) SetDocuments(v []DocumentObject)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *DownloadFilesAndFoldersResponseParams) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### SetDocumentsNil

`func (o *DownloadFilesAndFoldersResponseParams) SetDocumentsNil(b bool)`

 SetDocumentsNil sets the value for Documents to be an explicit nil

### UnsetDocuments
`func (o *DownloadFilesAndFoldersResponseParams) UnsetDocuments()`

UnsetDocuments ensures that no value is present for Documents, not even an explicit nil
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
### GetGlacierRetrievalType

`func (o *DownloadFilesAndFoldersResponseParams) GetGlacierRetrievalType() string`

GetGlacierRetrievalType returns the GlacierRetrievalType field if non-nil, zero value otherwise.

### GetGlacierRetrievalTypeOk

`func (o *DownloadFilesAndFoldersResponseParams) GetGlacierRetrievalTypeOk() (*string, bool)`

GetGlacierRetrievalTypeOk returns a tuple with the GlacierRetrievalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlacierRetrievalType

`func (o *DownloadFilesAndFoldersResponseParams) SetGlacierRetrievalType(v string)`

SetGlacierRetrievalType sets GlacierRetrievalType field to given value.

### HasGlacierRetrievalType

`func (o *DownloadFilesAndFoldersResponseParams) HasGlacierRetrievalType() bool`

HasGlacierRetrievalType returns a boolean if a field has been set.

### SetGlacierRetrievalTypeNil

`func (o *DownloadFilesAndFoldersResponseParams) SetGlacierRetrievalTypeNil(b bool)`

 SetGlacierRetrievalTypeNil sets the value for GlacierRetrievalType to be an explicit nil

### UnsetGlacierRetrievalType
`func (o *DownloadFilesAndFoldersResponseParams) UnsetGlacierRetrievalType()`

UnsetGlacierRetrievalType ensures that no value is present for GlacierRetrievalType, not even an explicit nil
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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


