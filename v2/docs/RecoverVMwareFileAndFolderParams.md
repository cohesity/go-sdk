# RecoverVMwareFileAndFolderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesAndFolders** | [**[]CommonRecoverFileAndFolderInfo**](CommonRecoverFileAndFolderInfo.md) | Specifies the info about the files and folders to be recovered. | 
**GlacierRetrievalType** | Pointer to **NullableString** | Specifies the glacier retrieval type when restoring or downloding files or folders from a Glacier-based cloud snapshot. | [optional] 
**ParentRecoveryId** | Pointer to **NullableString** | If current recovery is child task triggered through another parent recovery operation, then this field will specify the id of the parent recovery. | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**VmwareTargetParams** | Pointer to [**NullableRecoverVMwareFileAndFolderParamsVmwareTargetParams**](RecoverVMwareFileAndFolderParamsVmwareTargetParams.md) |  | [optional] 

## Methods

### NewRecoverVMwareFileAndFolderParams

`func NewRecoverVMwareFileAndFolderParams(filesAndFolders []CommonRecoverFileAndFolderInfo, targetEnvironment string, ) *RecoverVMwareFileAndFolderParams`

NewRecoverVMwareFileAndFolderParams instantiates a new RecoverVMwareFileAndFolderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVMwareFileAndFolderParamsWithDefaults

`func NewRecoverVMwareFileAndFolderParamsWithDefaults() *RecoverVMwareFileAndFolderParams`

NewRecoverVMwareFileAndFolderParamsWithDefaults instantiates a new RecoverVMwareFileAndFolderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesAndFolders

`func (o *RecoverVMwareFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *RecoverVMwareFileAndFolderParams) GetFilesAndFoldersOk() (*[]CommonRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *RecoverVMwareFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.


### SetFilesAndFoldersNil

`func (o *RecoverVMwareFileAndFolderParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *RecoverVMwareFileAndFolderParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetGlacierRetrievalType

`func (o *RecoverVMwareFileAndFolderParams) GetGlacierRetrievalType() string`

GetGlacierRetrievalType returns the GlacierRetrievalType field if non-nil, zero value otherwise.

### GetGlacierRetrievalTypeOk

`func (o *RecoverVMwareFileAndFolderParams) GetGlacierRetrievalTypeOk() (*string, bool)`

GetGlacierRetrievalTypeOk returns a tuple with the GlacierRetrievalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlacierRetrievalType

`func (o *RecoverVMwareFileAndFolderParams) SetGlacierRetrievalType(v string)`

SetGlacierRetrievalType sets GlacierRetrievalType field to given value.

### HasGlacierRetrievalType

`func (o *RecoverVMwareFileAndFolderParams) HasGlacierRetrievalType() bool`

HasGlacierRetrievalType returns a boolean if a field has been set.

### SetGlacierRetrievalTypeNil

`func (o *RecoverVMwareFileAndFolderParams) SetGlacierRetrievalTypeNil(b bool)`

 SetGlacierRetrievalTypeNil sets the value for GlacierRetrievalType to be an explicit nil

### UnsetGlacierRetrievalType
`func (o *RecoverVMwareFileAndFolderParams) UnsetGlacierRetrievalType()`

UnsetGlacierRetrievalType ensures that no value is present for GlacierRetrievalType, not even an explicit nil
### GetParentRecoveryId

`func (o *RecoverVMwareFileAndFolderParams) GetParentRecoveryId() string`

GetParentRecoveryId returns the ParentRecoveryId field if non-nil, zero value otherwise.

### GetParentRecoveryIdOk

`func (o *RecoverVMwareFileAndFolderParams) GetParentRecoveryIdOk() (*string, bool)`

GetParentRecoveryIdOk returns a tuple with the ParentRecoveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRecoveryId

`func (o *RecoverVMwareFileAndFolderParams) SetParentRecoveryId(v string)`

SetParentRecoveryId sets ParentRecoveryId field to given value.

### HasParentRecoveryId

`func (o *RecoverVMwareFileAndFolderParams) HasParentRecoveryId() bool`

HasParentRecoveryId returns a boolean if a field has been set.

### SetParentRecoveryIdNil

`func (o *RecoverVMwareFileAndFolderParams) SetParentRecoveryIdNil(b bool)`

 SetParentRecoveryIdNil sets the value for ParentRecoveryId to be an explicit nil

### UnsetParentRecoveryId
`func (o *RecoverVMwareFileAndFolderParams) UnsetParentRecoveryId()`

UnsetParentRecoveryId ensures that no value is present for ParentRecoveryId, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverVMwareFileAndFolderParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverVMwareFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverVMwareFileAndFolderParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetVmwareTargetParams

`func (o *RecoverVMwareFileAndFolderParams) GetVmwareTargetParams() RecoverVMwareFileAndFolderParamsVmwareTargetParams`

GetVmwareTargetParams returns the VmwareTargetParams field if non-nil, zero value otherwise.

### GetVmwareTargetParamsOk

`func (o *RecoverVMwareFileAndFolderParams) GetVmwareTargetParamsOk() (*RecoverVMwareFileAndFolderParamsVmwareTargetParams, bool)`

GetVmwareTargetParamsOk returns a tuple with the VmwareTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareTargetParams

`func (o *RecoverVMwareFileAndFolderParams) SetVmwareTargetParams(v RecoverVMwareFileAndFolderParamsVmwareTargetParams)`

SetVmwareTargetParams sets VmwareTargetParams field to given value.

### HasVmwareTargetParams

`func (o *RecoverVMwareFileAndFolderParams) HasVmwareTargetParams() bool`

HasVmwareTargetParams returns a boolean if a field has been set.

### SetVmwareTargetParamsNil

`func (o *RecoverVMwareFileAndFolderParams) SetVmwareTargetParamsNil(b bool)`

 SetVmwareTargetParamsNil sets the value for VmwareTargetParams to be an explicit nil

### UnsetVmwareTargetParams
`func (o *RecoverVMwareFileAndFolderParams) UnsetVmwareTargetParams()`

UnsetVmwareTargetParams ensures that no value is present for VmwareTargetParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


