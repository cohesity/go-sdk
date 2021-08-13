# PatchOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | Specifies the name of the service. | [optional] 
**Component** | Pointer to **string** | Specifies the description of the service. | [optional] 
**Version** | Pointer to **string** | Specifies the version of the patch. | [optional] 
**VersionReplaced** | Pointer to **string** | Specifies the version it replaced. | [optional] 
**Operation** | Pointer to **string** | Specifies what patch management operation was performed | [optional] 
**OperationTimeMsecs** | Pointer to **int64** | Specifies the time when the patch operation was done in Unix epoch in milliseconds. | [optional] 
**User** | Pointer to **string** | Specifies the user who performed the operation. | [optional] 
**Domain** | Pointer to **string** | Specifies the domain of the user. | [optional] 

## Methods

### NewPatchOperation

`func NewPatchOperation() *PatchOperation`

NewPatchOperation instantiates a new PatchOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchOperationWithDefaults

`func NewPatchOperationWithDefaults() *PatchOperation`

NewPatchOperationWithDefaults instantiates a new PatchOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *PatchOperation) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *PatchOperation) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *PatchOperation) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *PatchOperation) HasService() bool`

HasService returns a boolean if a field has been set.

### GetComponent

`func (o *PatchOperation) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *PatchOperation) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *PatchOperation) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *PatchOperation) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetVersion

`func (o *PatchOperation) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PatchOperation) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PatchOperation) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PatchOperation) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionReplaced

`func (o *PatchOperation) GetVersionReplaced() string`

GetVersionReplaced returns the VersionReplaced field if non-nil, zero value otherwise.

### GetVersionReplacedOk

`func (o *PatchOperation) GetVersionReplacedOk() (*string, bool)`

GetVersionReplacedOk returns a tuple with the VersionReplaced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionReplaced

`func (o *PatchOperation) SetVersionReplaced(v string)`

SetVersionReplaced sets VersionReplaced field to given value.

### HasVersionReplaced

`func (o *PatchOperation) HasVersionReplaced() bool`

HasVersionReplaced returns a boolean if a field has been set.

### GetOperation

`func (o *PatchOperation) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *PatchOperation) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *PatchOperation) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *PatchOperation) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetOperationTimeMsecs

`func (o *PatchOperation) GetOperationTimeMsecs() int64`

GetOperationTimeMsecs returns the OperationTimeMsecs field if non-nil, zero value otherwise.

### GetOperationTimeMsecsOk

`func (o *PatchOperation) GetOperationTimeMsecsOk() (*int64, bool)`

GetOperationTimeMsecsOk returns a tuple with the OperationTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationTimeMsecs

`func (o *PatchOperation) SetOperationTimeMsecs(v int64)`

SetOperationTimeMsecs sets OperationTimeMsecs field to given value.

### HasOperationTimeMsecs

`func (o *PatchOperation) HasOperationTimeMsecs() bool`

HasOperationTimeMsecs returns a boolean if a field has been set.

### GetUser

`func (o *PatchOperation) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchOperation) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchOperation) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchOperation) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetDomain

`func (o *PatchOperation) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PatchOperation) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PatchOperation) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PatchOperation) HasDomain() bool`

HasDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


