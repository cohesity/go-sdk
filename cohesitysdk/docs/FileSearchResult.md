# FileSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdObjectMetaData** | Pointer to [**AdObjectMetaData**](AdObjectMetaData.md) |  | [optional] 
**DocumentType** | Pointer to **NullableString** | Specifies the inferred document type. | [optional] 
**EmailMetaData** | Pointer to [**EmailMetaData**](EmailMetaData.md) |  | [optional] 
**FileVersions** | Pointer to [**[]FileVersion**](FileVersion.md) | Array of File Versions.  Specifies the different snapshot versions of a file or folder that were captured at different times. | [optional] 
**Filename** | Pointer to **NullableString** | Specifies the name of the found file or folder. | [optional] 
**IsFolder** | Pointer to **NullableBool** | Specifies if the found item is a folder. If true, the found item is a folder. | [optional] 
**JobId** | Pointer to **NullableInt64** | Specifies the Job id for the Protection Job that is currently associated with object that contains the backed up file or folder. If the file or folder was backed up on current Cohesity Cluster, this field contains the id for the Job that captured the object that contains the file or folder. If the file or folder was backed up on a Primary Cluster and replicated to this Cohesity Cluster, a new Inactive Job is created, the object that contains the file or folder is now associated with new Inactive Job, and this field contains the id of the new Inactive Job. | [optional] 
**JobUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies the universal id of the Protection Job that backed up the object that contains the file or folder. | [optional] 
**OneDriveDocumentMetadata** | Pointer to [**OneDriveDocumentMetadata**](OneDriveDocumentMetadata.md) |  | [optional] 
**ProtectionSource** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 
**RegisteredSourceId** | Pointer to **NullableInt64** | Specifies the id of the top-level registered source (such as a vCenter Server) where the source object that contains the the file or folder is stored. | [optional] 
**SharepointDocumentMetadata** | Pointer to [**SharepointDocumentMetadata**](SharepointDocumentMetadata.md) |  | [optional] 
**SnapshotTags** | Pointer to **[]string** | Snapshot tags present on this document | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the source id of the object that contains the file or folder. | [optional] 
**Tags** | Pointer to **[]string** | Tags present on this document. | [optional] 
**TagsToSnapshotsMap** | Pointer to **map[string][]int64** | Mapping from snapshot tags to. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the file document such as KDirectory, kFile, etc. | [optional] 
**ViewBoxId** | Pointer to **NullableInt64** | Specifies the id of the Domain (View Box) where the source object that contains the file or folder is stored. | [optional] 

## Methods

### NewFileSearchResult

`func NewFileSearchResult() *FileSearchResult`

NewFileSearchResult instantiates a new FileSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSearchResultWithDefaults

`func NewFileSearchResultWithDefaults() *FileSearchResult`

NewFileSearchResultWithDefaults instantiates a new FileSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdObjectMetaData

`func (o *FileSearchResult) GetAdObjectMetaData() AdObjectMetaData`

GetAdObjectMetaData returns the AdObjectMetaData field if non-nil, zero value otherwise.

### GetAdObjectMetaDataOk

`func (o *FileSearchResult) GetAdObjectMetaDataOk() (*AdObjectMetaData, bool)`

GetAdObjectMetaDataOk returns a tuple with the AdObjectMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdObjectMetaData

`func (o *FileSearchResult) SetAdObjectMetaData(v AdObjectMetaData)`

SetAdObjectMetaData sets AdObjectMetaData field to given value.

### HasAdObjectMetaData

`func (o *FileSearchResult) HasAdObjectMetaData() bool`

HasAdObjectMetaData returns a boolean if a field has been set.

### GetDocumentType

`func (o *FileSearchResult) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *FileSearchResult) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *FileSearchResult) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *FileSearchResult) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### SetDocumentTypeNil

`func (o *FileSearchResult) SetDocumentTypeNil(b bool)`

 SetDocumentTypeNil sets the value for DocumentType to be an explicit nil

### UnsetDocumentType
`func (o *FileSearchResult) UnsetDocumentType()`

UnsetDocumentType ensures that no value is present for DocumentType, not even an explicit nil
### GetEmailMetaData

`func (o *FileSearchResult) GetEmailMetaData() EmailMetaData`

GetEmailMetaData returns the EmailMetaData field if non-nil, zero value otherwise.

### GetEmailMetaDataOk

`func (o *FileSearchResult) GetEmailMetaDataOk() (*EmailMetaData, bool)`

GetEmailMetaDataOk returns a tuple with the EmailMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMetaData

`func (o *FileSearchResult) SetEmailMetaData(v EmailMetaData)`

SetEmailMetaData sets EmailMetaData field to given value.

### HasEmailMetaData

`func (o *FileSearchResult) HasEmailMetaData() bool`

HasEmailMetaData returns a boolean if a field has been set.

### GetFileVersions

`func (o *FileSearchResult) GetFileVersions() []FileVersion`

GetFileVersions returns the FileVersions field if non-nil, zero value otherwise.

### GetFileVersionsOk

`func (o *FileSearchResult) GetFileVersionsOk() (*[]FileVersion, bool)`

GetFileVersionsOk returns a tuple with the FileVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileVersions

`func (o *FileSearchResult) SetFileVersions(v []FileVersion)`

SetFileVersions sets FileVersions field to given value.

### HasFileVersions

`func (o *FileSearchResult) HasFileVersions() bool`

HasFileVersions returns a boolean if a field has been set.

### SetFileVersionsNil

`func (o *FileSearchResult) SetFileVersionsNil(b bool)`

 SetFileVersionsNil sets the value for FileVersions to be an explicit nil

### UnsetFileVersions
`func (o *FileSearchResult) UnsetFileVersions()`

UnsetFileVersions ensures that no value is present for FileVersions, not even an explicit nil
### GetFilename

`func (o *FileSearchResult) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *FileSearchResult) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *FileSearchResult) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *FileSearchResult) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### SetFilenameNil

`func (o *FileSearchResult) SetFilenameNil(b bool)`

 SetFilenameNil sets the value for Filename to be an explicit nil

### UnsetFilename
`func (o *FileSearchResult) UnsetFilename()`

UnsetFilename ensures that no value is present for Filename, not even an explicit nil
### GetIsFolder

`func (o *FileSearchResult) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *FileSearchResult) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *FileSearchResult) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *FileSearchResult) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### SetIsFolderNil

`func (o *FileSearchResult) SetIsFolderNil(b bool)`

 SetIsFolderNil sets the value for IsFolder to be an explicit nil

### UnsetIsFolder
`func (o *FileSearchResult) UnsetIsFolder()`

UnsetIsFolder ensures that no value is present for IsFolder, not even an explicit nil
### GetJobId

`func (o *FileSearchResult) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *FileSearchResult) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *FileSearchResult) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *FileSearchResult) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *FileSearchResult) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *FileSearchResult) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobUid

`func (o *FileSearchResult) GetJobUid() UniversalId`

GetJobUid returns the JobUid field if non-nil, zero value otherwise.

### GetJobUidOk

`func (o *FileSearchResult) GetJobUidOk() (*UniversalId, bool)`

GetJobUidOk returns a tuple with the JobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobUid

`func (o *FileSearchResult) SetJobUid(v UniversalId)`

SetJobUid sets JobUid field to given value.

### HasJobUid

`func (o *FileSearchResult) HasJobUid() bool`

HasJobUid returns a boolean if a field has been set.

### SetJobUidNil

`func (o *FileSearchResult) SetJobUidNil(b bool)`

 SetJobUidNil sets the value for JobUid to be an explicit nil

### UnsetJobUid
`func (o *FileSearchResult) UnsetJobUid()`

UnsetJobUid ensures that no value is present for JobUid, not even an explicit nil
### GetOneDriveDocumentMetadata

`func (o *FileSearchResult) GetOneDriveDocumentMetadata() OneDriveDocumentMetadata`

GetOneDriveDocumentMetadata returns the OneDriveDocumentMetadata field if non-nil, zero value otherwise.

### GetOneDriveDocumentMetadataOk

`func (o *FileSearchResult) GetOneDriveDocumentMetadataOk() (*OneDriveDocumentMetadata, bool)`

GetOneDriveDocumentMetadataOk returns a tuple with the OneDriveDocumentMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneDriveDocumentMetadata

`func (o *FileSearchResult) SetOneDriveDocumentMetadata(v OneDriveDocumentMetadata)`

SetOneDriveDocumentMetadata sets OneDriveDocumentMetadata field to given value.

### HasOneDriveDocumentMetadata

`func (o *FileSearchResult) HasOneDriveDocumentMetadata() bool`

HasOneDriveDocumentMetadata returns a boolean if a field has been set.

### GetProtectionSource

`func (o *FileSearchResult) GetProtectionSource() ProtectionSource`

GetProtectionSource returns the ProtectionSource field if non-nil, zero value otherwise.

### GetProtectionSourceOk

`func (o *FileSearchResult) GetProtectionSourceOk() (*ProtectionSource, bool)`

GetProtectionSourceOk returns a tuple with the ProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSource

`func (o *FileSearchResult) SetProtectionSource(v ProtectionSource)`

SetProtectionSource sets ProtectionSource field to given value.

### HasProtectionSource

`func (o *FileSearchResult) HasProtectionSource() bool`

HasProtectionSource returns a boolean if a field has been set.

### GetRegisteredSourceId

`func (o *FileSearchResult) GetRegisteredSourceId() int64`

GetRegisteredSourceId returns the RegisteredSourceId field if non-nil, zero value otherwise.

### GetRegisteredSourceIdOk

`func (o *FileSearchResult) GetRegisteredSourceIdOk() (*int64, bool)`

GetRegisteredSourceIdOk returns a tuple with the RegisteredSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredSourceId

`func (o *FileSearchResult) SetRegisteredSourceId(v int64)`

SetRegisteredSourceId sets RegisteredSourceId field to given value.

### HasRegisteredSourceId

`func (o *FileSearchResult) HasRegisteredSourceId() bool`

HasRegisteredSourceId returns a boolean if a field has been set.

### SetRegisteredSourceIdNil

`func (o *FileSearchResult) SetRegisteredSourceIdNil(b bool)`

 SetRegisteredSourceIdNil sets the value for RegisteredSourceId to be an explicit nil

### UnsetRegisteredSourceId
`func (o *FileSearchResult) UnsetRegisteredSourceId()`

UnsetRegisteredSourceId ensures that no value is present for RegisteredSourceId, not even an explicit nil
### GetSharepointDocumentMetadata

`func (o *FileSearchResult) GetSharepointDocumentMetadata() SharepointDocumentMetadata`

GetSharepointDocumentMetadata returns the SharepointDocumentMetadata field if non-nil, zero value otherwise.

### GetSharepointDocumentMetadataOk

`func (o *FileSearchResult) GetSharepointDocumentMetadataOk() (*SharepointDocumentMetadata, bool)`

GetSharepointDocumentMetadataOk returns a tuple with the SharepointDocumentMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointDocumentMetadata

`func (o *FileSearchResult) SetSharepointDocumentMetadata(v SharepointDocumentMetadata)`

SetSharepointDocumentMetadata sets SharepointDocumentMetadata field to given value.

### HasSharepointDocumentMetadata

`func (o *FileSearchResult) HasSharepointDocumentMetadata() bool`

HasSharepointDocumentMetadata returns a boolean if a field has been set.

### GetSnapshotTags

`func (o *FileSearchResult) GetSnapshotTags() []string`

GetSnapshotTags returns the SnapshotTags field if non-nil, zero value otherwise.

### GetSnapshotTagsOk

`func (o *FileSearchResult) GetSnapshotTagsOk() (*[]string, bool)`

GetSnapshotTagsOk returns a tuple with the SnapshotTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTags

`func (o *FileSearchResult) SetSnapshotTags(v []string)`

SetSnapshotTags sets SnapshotTags field to given value.

### HasSnapshotTags

`func (o *FileSearchResult) HasSnapshotTags() bool`

HasSnapshotTags returns a boolean if a field has been set.

### SetSnapshotTagsNil

`func (o *FileSearchResult) SetSnapshotTagsNil(b bool)`

 SetSnapshotTagsNil sets the value for SnapshotTags to be an explicit nil

### UnsetSnapshotTags
`func (o *FileSearchResult) UnsetSnapshotTags()`

UnsetSnapshotTags ensures that no value is present for SnapshotTags, not even an explicit nil
### GetSourceId

`func (o *FileSearchResult) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *FileSearchResult) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *FileSearchResult) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *FileSearchResult) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *FileSearchResult) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *FileSearchResult) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetTags

`func (o *FileSearchResult) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FileSearchResult) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FileSearchResult) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FileSearchResult) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *FileSearchResult) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *FileSearchResult) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTagsToSnapshotsMap

`func (o *FileSearchResult) GetTagsToSnapshotsMap() map[string][]int64`

GetTagsToSnapshotsMap returns the TagsToSnapshotsMap field if non-nil, zero value otherwise.

### GetTagsToSnapshotsMapOk

`func (o *FileSearchResult) GetTagsToSnapshotsMapOk() (*map[string][]int64, bool)`

GetTagsToSnapshotsMapOk returns a tuple with the TagsToSnapshotsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsToSnapshotsMap

`func (o *FileSearchResult) SetTagsToSnapshotsMap(v map[string][]int64)`

SetTagsToSnapshotsMap sets TagsToSnapshotsMap field to given value.

### HasTagsToSnapshotsMap

`func (o *FileSearchResult) HasTagsToSnapshotsMap() bool`

HasTagsToSnapshotsMap returns a boolean if a field has been set.

### SetTagsToSnapshotsMapNil

`func (o *FileSearchResult) SetTagsToSnapshotsMapNil(b bool)`

 SetTagsToSnapshotsMapNil sets the value for TagsToSnapshotsMap to be an explicit nil

### UnsetTagsToSnapshotsMap
`func (o *FileSearchResult) UnsetTagsToSnapshotsMap()`

UnsetTagsToSnapshotsMap ensures that no value is present for TagsToSnapshotsMap, not even an explicit nil
### GetType

`func (o *FileSearchResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileSearchResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileSearchResult) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FileSearchResult) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *FileSearchResult) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FileSearchResult) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetViewBoxId

`func (o *FileSearchResult) GetViewBoxId() int64`

GetViewBoxId returns the ViewBoxId field if non-nil, zero value otherwise.

### GetViewBoxIdOk

`func (o *FileSearchResult) GetViewBoxIdOk() (*int64, bool)`

GetViewBoxIdOk returns a tuple with the ViewBoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxId

`func (o *FileSearchResult) SetViewBoxId(v int64)`

SetViewBoxId sets ViewBoxId field to given value.

### HasViewBoxId

`func (o *FileSearchResult) HasViewBoxId() bool`

HasViewBoxId returns a boolean if a field has been set.

### SetViewBoxIdNil

`func (o *FileSearchResult) SetViewBoxIdNil(b bool)`

 SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil

### UnsetViewBoxId
`func (o *FileSearchResult) UnsetViewBoxId()`

UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


