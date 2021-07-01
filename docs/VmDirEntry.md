# VmDirEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FstatInfo** | Pointer to [**FileStatInfo**](FileStatInfo.md) |  | [optional] 
**FullPath** | Pointer to **NullableString** | FullPath is the full path of the file/directory. | [optional] 
**Name** | Pointer to **NullableString** | Name is the name of the file or folder. For /test/file.txt, name will be file.txt. | [optional] 
**Type** | Pointer to **NullableString** | DirEntryType is the type of entry i.e. file/folder. Specifies the type of directory entry.  &#39;kFile&#39; indicates that current entry is of file type. &#39;kDirectory&#39; indicates that current entry is of directory type. &#39;kSymlink&#39; indicates that current entry is of symbolic link. | [optional] 

## Methods

### NewVmDirEntry

`func NewVmDirEntry() *VmDirEntry`

NewVmDirEntry instantiates a new VmDirEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmDirEntryWithDefaults

`func NewVmDirEntryWithDefaults() *VmDirEntry`

NewVmDirEntryWithDefaults instantiates a new VmDirEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFstatInfo

`func (o *VmDirEntry) GetFstatInfo() FileStatInfo`

GetFstatInfo returns the FstatInfo field if non-nil, zero value otherwise.

### GetFstatInfoOk

`func (o *VmDirEntry) GetFstatInfoOk() (*FileStatInfo, bool)`

GetFstatInfoOk returns a tuple with the FstatInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFstatInfo

`func (o *VmDirEntry) SetFstatInfo(v FileStatInfo)`

SetFstatInfo sets FstatInfo field to given value.

### HasFstatInfo

`func (o *VmDirEntry) HasFstatInfo() bool`

HasFstatInfo returns a boolean if a field has been set.

### GetFullPath

`func (o *VmDirEntry) GetFullPath() string`

GetFullPath returns the FullPath field if non-nil, zero value otherwise.

### GetFullPathOk

`func (o *VmDirEntry) GetFullPathOk() (*string, bool)`

GetFullPathOk returns a tuple with the FullPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPath

`func (o *VmDirEntry) SetFullPath(v string)`

SetFullPath sets FullPath field to given value.

### HasFullPath

`func (o *VmDirEntry) HasFullPath() bool`

HasFullPath returns a boolean if a field has been set.

### SetFullPathNil

`func (o *VmDirEntry) SetFullPathNil(b bool)`

 SetFullPathNil sets the value for FullPath to be an explicit nil

### UnsetFullPath
`func (o *VmDirEntry) UnsetFullPath()`

UnsetFullPath ensures that no value is present for FullPath, not even an explicit nil
### GetName

`func (o *VmDirEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmDirEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmDirEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmDirEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VmDirEntry) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VmDirEntry) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *VmDirEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VmDirEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VmDirEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VmDirEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *VmDirEntry) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *VmDirEntry) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


