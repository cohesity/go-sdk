# NetworkingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceVec** | Pointer to [**[]ClusterNetworkingResourceInformation**](ClusterNetworkingResourceInformation.md) | The list of resources on the system that are accessible by an IP address. | [optional] 

## Methods

### NewNetworkingInformation

`func NewNetworkingInformation() *NetworkingInformation`

NewNetworkingInformation instantiates a new NetworkingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingInformationWithDefaults

`func NewNetworkingInformationWithDefaults() *NetworkingInformation`

NewNetworkingInformationWithDefaults instantiates a new NetworkingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceVec

`func (o *NetworkingInformation) GetResourceVec() []ClusterNetworkingResourceInformation`

GetResourceVec returns the ResourceVec field if non-nil, zero value otherwise.

### GetResourceVecOk

`func (o *NetworkingInformation) GetResourceVecOk() (*[]ClusterNetworkingResourceInformation, bool)`

GetResourceVecOk returns a tuple with the ResourceVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVec

`func (o *NetworkingInformation) SetResourceVec(v []ClusterNetworkingResourceInformation)`

SetResourceVec sets ResourceVec field to given value.

### HasResourceVec

`func (o *NetworkingInformation) HasResourceVec() bool`

HasResourceVec returns a boolean if a field has been set.

### SetResourceVecNil

`func (o *NetworkingInformation) SetResourceVecNil(b bool)`

 SetResourceVecNil sets the value for ResourceVec to be an explicit nil

### UnsetResourceVec
`func (o *NetworkingInformation) UnsetResourceVec()`

UnsetResourceVec ensures that no value is present for ResourceVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


