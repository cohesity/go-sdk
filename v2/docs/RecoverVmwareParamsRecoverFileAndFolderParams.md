# RecoverVmwareParamsRecoverFileAndFolderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesAndFolders** | [**[]CommonRecoverFileAndFolderInfo**](CommonRecoverFileAndFolderInfo.md) | Specifies the info about the files and folders to be recovered. | 
**GlacierRetrievalType** | Pointer to **NullableString** | Specifies the glacier retrieval type when restoring or downloding files or folders from a Glacier-based cloud snapshot. | [optional] 
**ParentRecoveryId** | Pointer to **NullableString** | If current recovery is child task triggered through another parent recovery operation, then this field will specify the id of the parent recovery. | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**VmwareTargetParams** | Pointer to [**NullableRecoverVMwareFileAndFolderParamsVmwareTargetParams**](RecoverVMwareFileAndFolderParamsVmwareTargetParams.md) |  | [optional] 

## Methods

### NewRecoverVmwareParamsRecoverFileAndFolderParams

`func NewRecoverVmwareParamsRecoverFileAndFolderParams(filesAndFolders []CommonRecoverFileAndFolderInfo, targetEnvironment string, ) *RecoverVmwareParamsRecoverFileAndFolderParams`

NewRecoverVmwareParamsRecoverFileAndFolderParams instantiates a new RecoverVmwareParamsRecoverFileAndFolderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareParamsRecoverFileAndFolderParamsWithDefaults

`func NewRecoverVmwareParamsRecoverFileAndFolderParamsWithDefaults() *RecoverVmwareParamsRecoverFileAndFolderParams`

NewRecoverVmwareParamsRecoverFileAndFolderParamsWithDefaults instantiates a new RecoverVmwareParamsRecoverFileAndFolderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesAndFolders

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) GetFilesAndFoldersOk() (*[]CommonRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.


### SetFilesAndFoldersNil

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetGlacierRetrievalType

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) GetGlacierRetrievalType() string`

GetGlacierRetrievalType returns the GlacierRetrievalType field if non-nil, zero value otherwise.

### GetGlacierRetrievalTypeOk

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) GetGlacierRetrievalTypeOk() (*string, bool)`

GetGlacierRetrievalTypeOk returns a tuple with the GlacierRetrievalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlacierRetrievalType

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) SetGlacierRetrievalType(v string)`

SetGlacierRetrievalType sets GlacierRetrievalType field to given value.

### HasGlacierRetrievalType

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) HasGlacierRetrievalType() bool`

HasGlacierRetrievalType returns a boolean if a field has been set.

### SetGlacierRetrievalTypeNil

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) SetGlacierRetrievalTypeNil(b bool)`

 SetGlacierRetrievalTypeNil sets the value for GlacierRetrievalType to be an explicit nil

### UnsetGlacierRetrievalType
`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) UnsetGlacierRetrievalType()`

UnsetGlacierRetrievalType ensures that no value is present for GlacierRetrievalType, not even an explicit nil
### GetParentRecoveryId

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) GetParentRecoveryId() string`

GetParentRecoveryId returns the ParentRecoveryId field if non-nil, zero value otherwise.

### GetParentRecoveryIdOk

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) GetParentRecoveryIdOk() (*string, bool)`

GetParentRecoveryIdOk returns a tuple with the ParentRecoveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRecoveryId

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) SetParentRecoveryId(v string)`

SetParentRecoveryId sets ParentRecoveryId field to given value.

### HasParentRecoveryId

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) HasParentRecoveryId() bool`

HasParentRecoveryId returns a boolean if a field has been set.

### SetParentRecoveryIdNil

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) SetParentRecoveryIdNil(b bool)`

 SetParentRecoveryIdNil sets the value for ParentRecoveryId to be an explicit nil

### UnsetParentRecoveryId
`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) UnsetParentRecoveryId()`

UnsetParentRecoveryId ensures that no value is present for ParentRecoveryId, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetVmwareTargetParams

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) GetVmwareTargetParams() RecoverVMwareFileAndFolderParamsVmwareTargetParams`

GetVmwareTargetParams returns the VmwareTargetParams field if non-nil, zero value otherwise.

### GetVmwareTargetParamsOk

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) GetVmwareTargetParamsOk() (*RecoverVMwareFileAndFolderParamsVmwareTargetParams, bool)`

GetVmwareTargetParamsOk returns a tuple with the VmwareTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareTargetParams

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) SetVmwareTargetParams(v RecoverVMwareFileAndFolderParamsVmwareTargetParams)`

SetVmwareTargetParams sets VmwareTargetParams field to given value.

### HasVmwareTargetParams

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) HasVmwareTargetParams() bool`

HasVmwareTargetParams returns a boolean if a field has been set.

### SetVmwareTargetParamsNil

`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) SetVmwareTargetParamsNil(b bool)`

 SetVmwareTargetParamsNil sets the value for VmwareTargetParams to be an explicit nil

### UnsetVmwareTargetParams
`func (o *RecoverVmwareParamsRecoverFileAndFolderParams) UnsetVmwareTargetParams()`

UnsetVmwareTargetParams ensures that no value is present for VmwareTargetParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


