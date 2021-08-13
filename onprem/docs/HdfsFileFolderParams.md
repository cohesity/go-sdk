# HdfsFileFolderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the name. | [optional] 
**Path** | Pointer to **NullableString** | Specifies the path. | [optional] 
**Type** | Pointer to **string** | Specifies the type of the contents. | [optional] 
**LastModifiedTimeUsecs** | Pointer to **NullableInt64** | Specifies the last time file was modified in unix timestamp. | [optional] 
**SizeBytes** | Pointer to **NullableInt64** | Specifies the file size in bytes. | [optional] 

## Methods

### NewHdfsFileFolderParams

`func NewHdfsFileFolderParams() *HdfsFileFolderParams`

NewHdfsFileFolderParams instantiates a new HdfsFileFolderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHdfsFileFolderParamsWithDefaults

`func NewHdfsFileFolderParamsWithDefaults() *HdfsFileFolderParams`

NewHdfsFileFolderParamsWithDefaults instantiates a new HdfsFileFolderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HdfsFileFolderParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HdfsFileFolderParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HdfsFileFolderParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HdfsFileFolderParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HdfsFileFolderParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HdfsFileFolderParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *HdfsFileFolderParams) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HdfsFileFolderParams) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HdfsFileFolderParams) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HdfsFileFolderParams) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *HdfsFileFolderParams) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *HdfsFileFolderParams) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetType

`func (o *HdfsFileFolderParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HdfsFileFolderParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HdfsFileFolderParams) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HdfsFileFolderParams) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLastModifiedTimeUsecs

`func (o *HdfsFileFolderParams) GetLastModifiedTimeUsecs() int64`

GetLastModifiedTimeUsecs returns the LastModifiedTimeUsecs field if non-nil, zero value otherwise.

### GetLastModifiedTimeUsecsOk

`func (o *HdfsFileFolderParams) GetLastModifiedTimeUsecsOk() (*int64, bool)`

GetLastModifiedTimeUsecsOk returns a tuple with the LastModifiedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTimeUsecs

`func (o *HdfsFileFolderParams) SetLastModifiedTimeUsecs(v int64)`

SetLastModifiedTimeUsecs sets LastModifiedTimeUsecs field to given value.

### HasLastModifiedTimeUsecs

`func (o *HdfsFileFolderParams) HasLastModifiedTimeUsecs() bool`

HasLastModifiedTimeUsecs returns a boolean if a field has been set.

### SetLastModifiedTimeUsecsNil

`func (o *HdfsFileFolderParams) SetLastModifiedTimeUsecsNil(b bool)`

 SetLastModifiedTimeUsecsNil sets the value for LastModifiedTimeUsecs to be an explicit nil

### UnsetLastModifiedTimeUsecs
`func (o *HdfsFileFolderParams) UnsetLastModifiedTimeUsecs()`

UnsetLastModifiedTimeUsecs ensures that no value is present for LastModifiedTimeUsecs, not even an explicit nil
### GetSizeBytes

`func (o *HdfsFileFolderParams) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *HdfsFileFolderParams) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *HdfsFileFolderParams) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *HdfsFileFolderParams) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### SetSizeBytesNil

`func (o *HdfsFileFolderParams) SetSizeBytesNil(b bool)`

 SetSizeBytesNil sets the value for SizeBytes to be an explicit nil

### UnsetSizeBytes
`func (o *HdfsFileFolderParams) UnsetSizeBytes()`

UnsetSizeBytes ensures that no value is present for SizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


