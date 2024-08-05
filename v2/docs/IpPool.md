# IpPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ips** | **[]string** | IP addresses. | 
**Name** | **string** | Name of the IP pool. | 

## Methods

### NewIpPool

`func NewIpPool(ips []string, name string, ) *IpPool`

NewIpPool instantiates a new IpPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpPoolWithDefaults

`func NewIpPoolWithDefaults() *IpPool`

NewIpPoolWithDefaults instantiates a new IpPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIps

`func (o *IpPool) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *IpPool) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *IpPool) SetIps(v []string)`

SetIps sets Ips field to given value.


### GetName

`func (o *IpPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpPool) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


