# RootPublicFolderParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoverObject** | [**CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the RootPublicFolder recover Object info. | 
**RecoverEntireRootPublicFolder** | Pointer to **NullableBool** | Specifies whether to recover the whole RootPublicFolder. | [optional] 
**RecoverFolders** | Pointer to [**[]PublicFolder**](PublicFolder.md) | Specifies a list of Public Folders to recover. This field is applicable only if &#39;recoverEntireRootPublicFolder&#39; is false. | [optional] 

## Methods

### NewRootPublicFolderParam

`func NewRootPublicFolderParam(recoverObject CommonRecoverObjectSnapshotParams, ) *RootPublicFolderParam`

NewRootPublicFolderParam instantiates a new RootPublicFolderParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootPublicFolderParamWithDefaults

`func NewRootPublicFolderParamWithDefaults() *RootPublicFolderParam`

NewRootPublicFolderParamWithDefaults instantiates a new RootPublicFolderParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoverObject

`func (o *RootPublicFolderParam) GetRecoverObject() CommonRecoverObjectSnapshotParams`

GetRecoverObject returns the RecoverObject field if non-nil, zero value otherwise.

### GetRecoverObjectOk

`func (o *RootPublicFolderParam) GetRecoverObjectOk() (*CommonRecoverObjectSnapshotParams, bool)`

GetRecoverObjectOk returns a tuple with the RecoverObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverObject

`func (o *RootPublicFolderParam) SetRecoverObject(v CommonRecoverObjectSnapshotParams)`

SetRecoverObject sets RecoverObject field to given value.


### GetRecoverEntireRootPublicFolder

`func (o *RootPublicFolderParam) GetRecoverEntireRootPublicFolder() bool`

GetRecoverEntireRootPublicFolder returns the RecoverEntireRootPublicFolder field if non-nil, zero value otherwise.

### GetRecoverEntireRootPublicFolderOk

`func (o *RootPublicFolderParam) GetRecoverEntireRootPublicFolderOk() (*bool, bool)`

GetRecoverEntireRootPublicFolderOk returns a tuple with the RecoverEntireRootPublicFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverEntireRootPublicFolder

`func (o *RootPublicFolderParam) SetRecoverEntireRootPublicFolder(v bool)`

SetRecoverEntireRootPublicFolder sets RecoverEntireRootPublicFolder field to given value.

### HasRecoverEntireRootPublicFolder

`func (o *RootPublicFolderParam) HasRecoverEntireRootPublicFolder() bool`

HasRecoverEntireRootPublicFolder returns a boolean if a field has been set.

### SetRecoverEntireRootPublicFolderNil

`func (o *RootPublicFolderParam) SetRecoverEntireRootPublicFolderNil(b bool)`

 SetRecoverEntireRootPublicFolderNil sets the value for RecoverEntireRootPublicFolder to be an explicit nil

### UnsetRecoverEntireRootPublicFolder
`func (o *RootPublicFolderParam) UnsetRecoverEntireRootPublicFolder()`

UnsetRecoverEntireRootPublicFolder ensures that no value is present for RecoverEntireRootPublicFolder, not even an explicit nil
### GetRecoverFolders

`func (o *RootPublicFolderParam) GetRecoverFolders() []PublicFolder`

GetRecoverFolders returns the RecoverFolders field if non-nil, zero value otherwise.

### GetRecoverFoldersOk

`func (o *RootPublicFolderParam) GetRecoverFoldersOk() (*[]PublicFolder, bool)`

GetRecoverFoldersOk returns a tuple with the RecoverFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverFolders

`func (o *RootPublicFolderParam) SetRecoverFolders(v []PublicFolder)`

SetRecoverFolders sets RecoverFolders field to given value.

### HasRecoverFolders

`func (o *RootPublicFolderParam) HasRecoverFolders() bool`

HasRecoverFolders returns a boolean if a field has been set.

### SetRecoverFoldersNil

`func (o *RootPublicFolderParam) SetRecoverFoldersNil(b bool)`

 SetRecoverFoldersNil sets the value for RecoverFolders to be an explicit nil

### UnsetRecoverFolders
`func (o *RootPublicFolderParam) UnsetRecoverFolders()`

UnsetRecoverFolders ensures that no value is present for RecoverFolders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


