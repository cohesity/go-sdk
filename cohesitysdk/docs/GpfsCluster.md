# GpfsCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CesAddresses** | Pointer to **[]string** | Specifies a list of CES(Cluster Export Services) IP addresses of a GPFS Cluster. | [optional] 
**Id** | Pointer to **NullableInt32** | Specifies a globally unique id of a GPFS Cluster. | [optional] 
**PrimaryServer** | Pointer to **NullableString** | Specifies a primary server of a GPFS Cluster. | [optional] 

## Methods

### NewGpfsCluster

`func NewGpfsCluster() *GpfsCluster`

NewGpfsCluster instantiates a new GpfsCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpfsClusterWithDefaults

`func NewGpfsClusterWithDefaults() *GpfsCluster`

NewGpfsClusterWithDefaults instantiates a new GpfsCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCesAddresses

`func (o *GpfsCluster) GetCesAddresses() []string`

GetCesAddresses returns the CesAddresses field if non-nil, zero value otherwise.

### GetCesAddressesOk

`func (o *GpfsCluster) GetCesAddressesOk() (*[]string, bool)`

GetCesAddressesOk returns a tuple with the CesAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCesAddresses

`func (o *GpfsCluster) SetCesAddresses(v []string)`

SetCesAddresses sets CesAddresses field to given value.

### HasCesAddresses

`func (o *GpfsCluster) HasCesAddresses() bool`

HasCesAddresses returns a boolean if a field has been set.

### SetCesAddressesNil

`func (o *GpfsCluster) SetCesAddressesNil(b bool)`

 SetCesAddressesNil sets the value for CesAddresses to be an explicit nil

### UnsetCesAddresses
`func (o *GpfsCluster) UnsetCesAddresses()`

UnsetCesAddresses ensures that no value is present for CesAddresses, not even an explicit nil
### GetId

`func (o *GpfsCluster) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GpfsCluster) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GpfsCluster) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GpfsCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GpfsCluster) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GpfsCluster) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPrimaryServer

`func (o *GpfsCluster) GetPrimaryServer() string`

GetPrimaryServer returns the PrimaryServer field if non-nil, zero value otherwise.

### GetPrimaryServerOk

`func (o *GpfsCluster) GetPrimaryServerOk() (*string, bool)`

GetPrimaryServerOk returns a tuple with the PrimaryServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryServer

`func (o *GpfsCluster) SetPrimaryServer(v string)`

SetPrimaryServer sets PrimaryServer field to given value.

### HasPrimaryServer

`func (o *GpfsCluster) HasPrimaryServer() bool`

HasPrimaryServer returns a boolean if a field has been set.

### SetPrimaryServerNil

`func (o *GpfsCluster) SetPrimaryServerNil(b bool)`

 SetPrimaryServerNil sets the value for PrimaryServer to be an explicit nil

### UnsetPrimaryServer
`func (o *GpfsCluster) UnsetPrimaryServer()`

UnsetPrimaryServer ensures that no value is present for PrimaryServer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


