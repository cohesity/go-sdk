# ADObjectRestoreParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 
**GuidVec** | Pointer to **[]string** | Array of AD object guids to restore either from recycle bin or from AD snapshot. The guid should not exist in production AD. If it exits, then kDuplicate error is output in status. | [optional] 
**OptionFlags** | Pointer to **NullableInt32** | Restore option flags of type ADObjectRestoreOptionFlags. | [optional] 
**OuPath** | Pointer to **NullableString** | Distinguished name(DN) of the target Organization Unit (OU) to restore non-OU object. This can be empty, in which case objects are restored to their original OU. The &#39;credential&#39; specified must have privileges to add objects to this OU. Example: &#39;OU&#x3D;SJC,OU&#x3D;EngOU,DC&#x3D;contoso,DC&#x3D;com&#39;. This parameter is ignored for OU objects. | [optional] 
**SrcSysvolFolder** | Pointer to **NullableString** | When restoring a GPO, need to know the absolute path for SYSVOL folder. | [optional] 

## Methods

### NewADObjectRestoreParam

`func NewADObjectRestoreParam() *ADObjectRestoreParam`

NewADObjectRestoreParam instantiates a new ADObjectRestoreParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADObjectRestoreParamWithDefaults

`func NewADObjectRestoreParamWithDefaults() *ADObjectRestoreParam`

NewADObjectRestoreParamWithDefaults instantiates a new ADObjectRestoreParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ADObjectRestoreParam) GetCredentials() Credentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ADObjectRestoreParam) GetCredentialsOk() (*Credentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ADObjectRestoreParam) SetCredentials(v Credentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ADObjectRestoreParam) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetGuidVec

`func (o *ADObjectRestoreParam) GetGuidVec() []string`

GetGuidVec returns the GuidVec field if non-nil, zero value otherwise.

### GetGuidVecOk

`func (o *ADObjectRestoreParam) GetGuidVecOk() (*[]string, bool)`

GetGuidVecOk returns a tuple with the GuidVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidVec

`func (o *ADObjectRestoreParam) SetGuidVec(v []string)`

SetGuidVec sets GuidVec field to given value.

### HasGuidVec

`func (o *ADObjectRestoreParam) HasGuidVec() bool`

HasGuidVec returns a boolean if a field has been set.

### SetGuidVecNil

`func (o *ADObjectRestoreParam) SetGuidVecNil(b bool)`

 SetGuidVecNil sets the value for GuidVec to be an explicit nil

### UnsetGuidVec
`func (o *ADObjectRestoreParam) UnsetGuidVec()`

UnsetGuidVec ensures that no value is present for GuidVec, not even an explicit nil
### GetOptionFlags

`func (o *ADObjectRestoreParam) GetOptionFlags() int32`

GetOptionFlags returns the OptionFlags field if non-nil, zero value otherwise.

### GetOptionFlagsOk

`func (o *ADObjectRestoreParam) GetOptionFlagsOk() (*int32, bool)`

GetOptionFlagsOk returns a tuple with the OptionFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionFlags

`func (o *ADObjectRestoreParam) SetOptionFlags(v int32)`

SetOptionFlags sets OptionFlags field to given value.

### HasOptionFlags

`func (o *ADObjectRestoreParam) HasOptionFlags() bool`

HasOptionFlags returns a boolean if a field has been set.

### SetOptionFlagsNil

`func (o *ADObjectRestoreParam) SetOptionFlagsNil(b bool)`

 SetOptionFlagsNil sets the value for OptionFlags to be an explicit nil

### UnsetOptionFlags
`func (o *ADObjectRestoreParam) UnsetOptionFlags()`

UnsetOptionFlags ensures that no value is present for OptionFlags, not even an explicit nil
### GetOuPath

`func (o *ADObjectRestoreParam) GetOuPath() string`

GetOuPath returns the OuPath field if non-nil, zero value otherwise.

### GetOuPathOk

`func (o *ADObjectRestoreParam) GetOuPathOk() (*string, bool)`

GetOuPathOk returns a tuple with the OuPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuPath

`func (o *ADObjectRestoreParam) SetOuPath(v string)`

SetOuPath sets OuPath field to given value.

### HasOuPath

`func (o *ADObjectRestoreParam) HasOuPath() bool`

HasOuPath returns a boolean if a field has been set.

### SetOuPathNil

`func (o *ADObjectRestoreParam) SetOuPathNil(b bool)`

 SetOuPathNil sets the value for OuPath to be an explicit nil

### UnsetOuPath
`func (o *ADObjectRestoreParam) UnsetOuPath()`

UnsetOuPath ensures that no value is present for OuPath, not even an explicit nil
### GetSrcSysvolFolder

`func (o *ADObjectRestoreParam) GetSrcSysvolFolder() string`

GetSrcSysvolFolder returns the SrcSysvolFolder field if non-nil, zero value otherwise.

### GetSrcSysvolFolderOk

`func (o *ADObjectRestoreParam) GetSrcSysvolFolderOk() (*string, bool)`

GetSrcSysvolFolderOk returns a tuple with the SrcSysvolFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcSysvolFolder

`func (o *ADObjectRestoreParam) SetSrcSysvolFolder(v string)`

SetSrcSysvolFolder sets SrcSysvolFolder field to given value.

### HasSrcSysvolFolder

`func (o *ADObjectRestoreParam) HasSrcSysvolFolder() bool`

HasSrcSysvolFolder returns a boolean if a field has been set.

### SetSrcSysvolFolderNil

`func (o *ADObjectRestoreParam) SetSrcSysvolFolderNil(b bool)`

 SetSrcSysvolFolderNil sets the value for SrcSysvolFolder to be an explicit nil

### UnsetSrcSysvolFolder
`func (o *ADObjectRestoreParam) UnsetSrcSysvolFolder()`

UnsetSrcSysvolFolder ensures that no value is present for SrcSysvolFolder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


