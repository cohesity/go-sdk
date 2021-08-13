# AppliedPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | Specifies the name of the service. | [optional] 
**Component** | Pointer to **string** | Specifies the description of the service. | [optional] 
**Version** | Pointer to **string** | Specifies the version of the patch. | [optional] 
**VersionReplaced** | Pointer to **string** | Specifies the version it replaced. | [optional] 
**PatchLevel** | Pointer to **int64** | Specifies the number of patches applied so far for this service. | [optional] 
**AppliedTimeMsecs** | Pointer to **int64** | Specifies the time when the patch was applied in Unix epoch in milliseconds. | [optional] 
**Count** | Pointer to **int64** | Specifies the number of fixed issues. | [optional] 
**Dependencies** | Pointer to **[]string** | Specifies the services for which their patches were applied together. They will also be reverted together. | [optional] 
**FixedIssues** | Pointer to [**[]FixedIssue**](FixedIssue.md) | Specifies the details of the issues fixed in the patch. | [optional] 

## Methods

### NewAppliedPatch

`func NewAppliedPatch() *AppliedPatch`

NewAppliedPatch instantiates a new AppliedPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppliedPatchWithDefaults

`func NewAppliedPatchWithDefaults() *AppliedPatch`

NewAppliedPatchWithDefaults instantiates a new AppliedPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *AppliedPatch) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AppliedPatch) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AppliedPatch) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *AppliedPatch) HasService() bool`

HasService returns a boolean if a field has been set.

### GetComponent

`func (o *AppliedPatch) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AppliedPatch) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AppliedPatch) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AppliedPatch) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetVersion

`func (o *AppliedPatch) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AppliedPatch) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AppliedPatch) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AppliedPatch) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionReplaced

`func (o *AppliedPatch) GetVersionReplaced() string`

GetVersionReplaced returns the VersionReplaced field if non-nil, zero value otherwise.

### GetVersionReplacedOk

`func (o *AppliedPatch) GetVersionReplacedOk() (*string, bool)`

GetVersionReplacedOk returns a tuple with the VersionReplaced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionReplaced

`func (o *AppliedPatch) SetVersionReplaced(v string)`

SetVersionReplaced sets VersionReplaced field to given value.

### HasVersionReplaced

`func (o *AppliedPatch) HasVersionReplaced() bool`

HasVersionReplaced returns a boolean if a field has been set.

### GetPatchLevel

`func (o *AppliedPatch) GetPatchLevel() int64`

GetPatchLevel returns the PatchLevel field if non-nil, zero value otherwise.

### GetPatchLevelOk

`func (o *AppliedPatch) GetPatchLevelOk() (*int64, bool)`

GetPatchLevelOk returns a tuple with the PatchLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchLevel

`func (o *AppliedPatch) SetPatchLevel(v int64)`

SetPatchLevel sets PatchLevel field to given value.

### HasPatchLevel

`func (o *AppliedPatch) HasPatchLevel() bool`

HasPatchLevel returns a boolean if a field has been set.

### GetAppliedTimeMsecs

`func (o *AppliedPatch) GetAppliedTimeMsecs() int64`

GetAppliedTimeMsecs returns the AppliedTimeMsecs field if non-nil, zero value otherwise.

### GetAppliedTimeMsecsOk

`func (o *AppliedPatch) GetAppliedTimeMsecsOk() (*int64, bool)`

GetAppliedTimeMsecsOk returns a tuple with the AppliedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTimeMsecs

`func (o *AppliedPatch) SetAppliedTimeMsecs(v int64)`

SetAppliedTimeMsecs sets AppliedTimeMsecs field to given value.

### HasAppliedTimeMsecs

`func (o *AppliedPatch) HasAppliedTimeMsecs() bool`

HasAppliedTimeMsecs returns a boolean if a field has been set.

### GetCount

`func (o *AppliedPatch) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AppliedPatch) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AppliedPatch) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *AppliedPatch) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDependencies

`func (o *AppliedPatch) GetDependencies() []string`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *AppliedPatch) GetDependenciesOk() (*[]string, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *AppliedPatch) SetDependencies(v []string)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *AppliedPatch) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### SetDependenciesNil

`func (o *AppliedPatch) SetDependenciesNil(b bool)`

 SetDependenciesNil sets the value for Dependencies to be an explicit nil

### UnsetDependencies
`func (o *AppliedPatch) UnsetDependencies()`

UnsetDependencies ensures that no value is present for Dependencies, not even an explicit nil
### GetFixedIssues

`func (o *AppliedPatch) GetFixedIssues() []FixedIssue`

GetFixedIssues returns the FixedIssues field if non-nil, zero value otherwise.

### GetFixedIssuesOk

`func (o *AppliedPatch) GetFixedIssuesOk() (*[]FixedIssue, bool)`

GetFixedIssuesOk returns a tuple with the FixedIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIssues

`func (o *AppliedPatch) SetFixedIssues(v []FixedIssue)`

SetFixedIssues sets FixedIssues field to given value.

### HasFixedIssues

`func (o *AppliedPatch) HasFixedIssues() bool`

HasFixedIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


