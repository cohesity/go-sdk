# HostMappings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to [**[]HostEntry**](HostEntry.md) | Specifies the list of host entries | [optional] 

## Methods

### NewHostMappings

`func NewHostMappings() *HostMappings`

NewHostMappings instantiates a new HostMappings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostMappingsWithDefaults

`func NewHostMappingsWithDefaults() *HostMappings`

NewHostMappingsWithDefaults instantiates a new HostMappings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *HostMappings) GetHosts() []HostEntry`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *HostMappings) GetHostsOk() (*[]HostEntry, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *HostMappings) SetHosts(v []HostEntry)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *HostMappings) HasHosts() bool`

HasHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


