# ServiceProcessEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessIds** | Pointer to **[]int64** | Specifies the process Ids. | [optional] 
**ServiceName** | Pointer to **NullableString** | Specifies the service name. | [optional] 

## Methods

### NewServiceProcessEntry

`func NewServiceProcessEntry() *ServiceProcessEntry`

NewServiceProcessEntry instantiates a new ServiceProcessEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProcessEntryWithDefaults

`func NewServiceProcessEntryWithDefaults() *ServiceProcessEntry`

NewServiceProcessEntryWithDefaults instantiates a new ServiceProcessEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessIds

`func (o *ServiceProcessEntry) GetProcessIds() []int64`

GetProcessIds returns the ProcessIds field if non-nil, zero value otherwise.

### GetProcessIdsOk

`func (o *ServiceProcessEntry) GetProcessIdsOk() (*[]int64, bool)`

GetProcessIdsOk returns a tuple with the ProcessIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessIds

`func (o *ServiceProcessEntry) SetProcessIds(v []int64)`

SetProcessIds sets ProcessIds field to given value.

### HasProcessIds

`func (o *ServiceProcessEntry) HasProcessIds() bool`

HasProcessIds returns a boolean if a field has been set.

### SetProcessIdsNil

`func (o *ServiceProcessEntry) SetProcessIdsNil(b bool)`

 SetProcessIdsNil sets the value for ProcessIds to be an explicit nil

### UnsetProcessIds
`func (o *ServiceProcessEntry) UnsetProcessIds()`

UnsetProcessIds ensures that no value is present for ProcessIds, not even an explicit nil
### GetServiceName

`func (o *ServiceProcessEntry) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServiceProcessEntry) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServiceProcessEntry) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ServiceProcessEntry) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *ServiceProcessEntry) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *ServiceProcessEntry) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


