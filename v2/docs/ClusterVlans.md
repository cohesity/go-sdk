# ClusterVlans

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vlans** | Pointer to [**[]ClusterVlanParams**](ClusterVlanParams.md) | List of vlans. | [optional] 

## Methods

### NewClusterVlans

`func NewClusterVlans() *ClusterVlans`

NewClusterVlans instantiates a new ClusterVlans object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterVlansWithDefaults

`func NewClusterVlansWithDefaults() *ClusterVlans`

NewClusterVlansWithDefaults instantiates a new ClusterVlans object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlans

`func (o *ClusterVlans) GetVlans() []ClusterVlanParams`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *ClusterVlans) GetVlansOk() (*[]ClusterVlanParams, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *ClusterVlans) SetVlans(v []ClusterVlanParams)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *ClusterVlans) HasVlans() bool`

HasVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


