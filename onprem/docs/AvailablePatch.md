# AvailablePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | Specifies the name of the service. | [optional] 
**Component** | Pointer to **string** | Specifies the description of the service. | [optional] 
**Version** | Pointer to **string** | Specifies the version of the patch. | [optional] 
**Count** | Pointer to **int64** | Specifies the number of fixed issues. | [optional] 
**Dependencies** | Pointer to **[]string** | Specifies the services for which their patches must be applied together. | [optional] 
**FixedIssues** | Pointer to [**[]FixedIssue**](FixedIssue.md) | Specifies the details of the issues fixed in the patch. | [optional] 

## Methods

### NewAvailablePatch

`func NewAvailablePatch() *AvailablePatch`

NewAvailablePatch instantiates a new AvailablePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailablePatchWithDefaults

`func NewAvailablePatchWithDefaults() *AvailablePatch`

NewAvailablePatchWithDefaults instantiates a new AvailablePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *AvailablePatch) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AvailablePatch) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AvailablePatch) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *AvailablePatch) HasService() bool`

HasService returns a boolean if a field has been set.

### GetComponent

`func (o *AvailablePatch) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AvailablePatch) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AvailablePatch) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AvailablePatch) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetVersion

`func (o *AvailablePatch) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AvailablePatch) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AvailablePatch) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AvailablePatch) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCount

`func (o *AvailablePatch) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AvailablePatch) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AvailablePatch) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *AvailablePatch) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDependencies

`func (o *AvailablePatch) GetDependencies() []string`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *AvailablePatch) GetDependenciesOk() (*[]string, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *AvailablePatch) SetDependencies(v []string)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *AvailablePatch) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### SetDependenciesNil

`func (o *AvailablePatch) SetDependenciesNil(b bool)`

 SetDependenciesNil sets the value for Dependencies to be an explicit nil

### UnsetDependencies
`func (o *AvailablePatch) UnsetDependencies()`

UnsetDependencies ensures that no value is present for Dependencies, not even an explicit nil
### GetFixedIssues

`func (o *AvailablePatch) GetFixedIssues() []FixedIssue`

GetFixedIssues returns the FixedIssues field if non-nil, zero value otherwise.

### GetFixedIssuesOk

`func (o *AvailablePatch) GetFixedIssuesOk() (*[]FixedIssue, bool)`

GetFixedIssuesOk returns a tuple with the FixedIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIssues

`func (o *AvailablePatch) SetFixedIssues(v []FixedIssue)`

SetFixedIssues sets FixedIssues field to given value.

### HasFixedIssues

`func (o *AvailablePatch) HasFixedIssues() bool`

HasFixedIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


