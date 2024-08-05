# ServiceGflags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gflags** | Pointer to [**[]Gflag**](Gflag.md) | Specifies a list of gflags for the service. | [optional] 
**ServiceName** | Pointer to **NullableString** | Specifies the service name. | [optional] 

## Methods

### NewServiceGflags

`func NewServiceGflags() *ServiceGflags`

NewServiceGflags instantiates a new ServiceGflags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceGflagsWithDefaults

`func NewServiceGflagsWithDefaults() *ServiceGflags`

NewServiceGflagsWithDefaults instantiates a new ServiceGflags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGflags

`func (o *ServiceGflags) GetGflags() []Gflag`

GetGflags returns the Gflags field if non-nil, zero value otherwise.

### GetGflagsOk

`func (o *ServiceGflags) GetGflagsOk() (*[]Gflag, bool)`

GetGflagsOk returns a tuple with the Gflags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGflags

`func (o *ServiceGflags) SetGflags(v []Gflag)`

SetGflags sets Gflags field to given value.

### HasGflags

`func (o *ServiceGflags) HasGflags() bool`

HasGflags returns a boolean if a field has been set.

### GetServiceName

`func (o *ServiceGflags) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServiceGflags) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServiceGflags) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ServiceGflags) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *ServiceGflags) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *ServiceGflags) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


