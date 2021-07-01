# AppendHostsParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to [**[]HostEntry**](HostEntry.md) | Specifies the list of host entries to be added to the Cluster&#39;s etc/hosts file. | [optional] 

## Methods

### NewAppendHostsParameters

`func NewAppendHostsParameters() *AppendHostsParameters`

NewAppendHostsParameters instantiates a new AppendHostsParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppendHostsParametersWithDefaults

`func NewAppendHostsParametersWithDefaults() *AppendHostsParameters`

NewAppendHostsParametersWithDefaults instantiates a new AppendHostsParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *AppendHostsParameters) GetHosts() []HostEntry`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *AppendHostsParameters) GetHostsOk() (*[]HostEntry, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *AppendHostsParameters) SetHosts(v []HostEntry)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *AppendHostsParameters) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *AppendHostsParameters) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *AppendHostsParameters) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


