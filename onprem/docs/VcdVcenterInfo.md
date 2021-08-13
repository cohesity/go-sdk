# VcdVcenterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the name of the vCenter. | [optional] 
**Endpoint** | Pointer to **NullableString** | Specifies the endpoint of the vCenter. | [optional] 

## Methods

### NewVcdVcenterInfo

`func NewVcdVcenterInfo() *VcdVcenterInfo`

NewVcdVcenterInfo instantiates a new VcdVcenterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcdVcenterInfoWithDefaults

`func NewVcdVcenterInfoWithDefaults() *VcdVcenterInfo`

NewVcdVcenterInfoWithDefaults instantiates a new VcdVcenterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcdVcenterInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcdVcenterInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcdVcenterInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VcdVcenterInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VcdVcenterInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VcdVcenterInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEndpoint

`func (o *VcdVcenterInfo) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *VcdVcenterInfo) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *VcdVcenterInfo) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *VcdVcenterInfo) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *VcdVcenterInfo) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *VcdVcenterInfo) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


