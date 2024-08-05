# PatchDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **NullableString** | Specifies the user friendly name of the service. | [optional] 
**ImportVersion** | Pointer to **NullableString** | Specifies the version of the imported service patch. | [optional] 
**Service** | Pointer to **NullableString** | Specifies the name of the service. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the status of the patch whether it is accepted or rejected. A patch is rejected if it is older than the version available or applied on the cluster. | [optional] 
**Version** | Pointer to **NullableString** | Specifies the existing version of the service. This is the available service patch version if exists. If there is no patch available, then it is the applied patch version if applied. If both don&#39;t exist, it is the base version of the service. | [optional] 

## Methods

### NewPatchDetail

`func NewPatchDetail() *PatchDetail`

NewPatchDetail instantiates a new PatchDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchDetailWithDefaults

`func NewPatchDetailWithDefaults() *PatchDetail`

NewPatchDetailWithDefaults instantiates a new PatchDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *PatchDetail) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *PatchDetail) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *PatchDetail) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *PatchDetail) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### SetComponentNil

`func (o *PatchDetail) SetComponentNil(b bool)`

 SetComponentNil sets the value for Component to be an explicit nil

### UnsetComponent
`func (o *PatchDetail) UnsetComponent()`

UnsetComponent ensures that no value is present for Component, not even an explicit nil
### GetImportVersion

`func (o *PatchDetail) GetImportVersion() string`

GetImportVersion returns the ImportVersion field if non-nil, zero value otherwise.

### GetImportVersionOk

`func (o *PatchDetail) GetImportVersionOk() (*string, bool)`

GetImportVersionOk returns a tuple with the ImportVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportVersion

`func (o *PatchDetail) SetImportVersion(v string)`

SetImportVersion sets ImportVersion field to given value.

### HasImportVersion

`func (o *PatchDetail) HasImportVersion() bool`

HasImportVersion returns a boolean if a field has been set.

### SetImportVersionNil

`func (o *PatchDetail) SetImportVersionNil(b bool)`

 SetImportVersionNil sets the value for ImportVersion to be an explicit nil

### UnsetImportVersion
`func (o *PatchDetail) UnsetImportVersion()`

UnsetImportVersion ensures that no value is present for ImportVersion, not even an explicit nil
### GetService

`func (o *PatchDetail) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *PatchDetail) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *PatchDetail) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *PatchDetail) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *PatchDetail) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *PatchDetail) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetStatus

`func (o *PatchDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PatchDetail) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PatchDetail) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetVersion

`func (o *PatchDetail) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PatchDetail) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PatchDetail) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PatchDetail) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *PatchDetail) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *PatchDetail) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


