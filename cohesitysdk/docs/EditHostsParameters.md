# EditHostsParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to [**[]HostEntry**](HostEntry.md) | Specifies the list of host entries to be edited. Each IP address listed in the list of host entries will have its corresponding domain names in the /etc/hosts file replaced with the domain names specified here. | [optional] 

## Methods

### NewEditHostsParameters

`func NewEditHostsParameters() *EditHostsParameters`

NewEditHostsParameters instantiates a new EditHostsParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditHostsParametersWithDefaults

`func NewEditHostsParametersWithDefaults() *EditHostsParameters`

NewEditHostsParametersWithDefaults instantiates a new EditHostsParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *EditHostsParameters) GetHosts() []HostEntry`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *EditHostsParameters) GetHostsOk() (*[]HostEntry, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *EditHostsParameters) SetHosts(v []HostEntry)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *EditHostsParameters) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *EditHostsParameters) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *EditHostsParameters) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


