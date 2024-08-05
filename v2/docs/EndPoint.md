# EndPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** | Index of the interface as given by &#39;ip a&#39; command. | [optional] 
**IpAddresses** | Pointer to **[]string** | IP addresses on the interface | [optional] 
**Name** | Pointer to **string** | Name of the interface like bond0. | [optional] 

## Methods

### NewEndPoint

`func NewEndPoint() *EndPoint`

NewEndPoint instantiates a new EndPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndPointWithDefaults

`func NewEndPointWithDefaults() *EndPoint`

NewEndPointWithDefaults instantiates a new EndPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *EndPoint) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *EndPoint) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *EndPoint) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *EndPoint) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetIpAddresses

`func (o *EndPoint) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *EndPoint) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *EndPoint) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *EndPoint) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetName

`func (o *EndPoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndPoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndPoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EndPoint) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


