# ClusterCreateRigelParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimToken** | **NullableString** | Specifies the token which will be used to claim this Cluster to Helios. | 
**Nodes** | [**[]RigelClusterNode**](RigelClusterNode.md) |  | 

## Methods

### NewClusterCreateRigelParams

`func NewClusterCreateRigelParams(claimToken NullableString, nodes []RigelClusterNode, ) *ClusterCreateRigelParams`

NewClusterCreateRigelParams instantiates a new ClusterCreateRigelParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreateRigelParamsWithDefaults

`func NewClusterCreateRigelParamsWithDefaults() *ClusterCreateRigelParams`

NewClusterCreateRigelParamsWithDefaults instantiates a new ClusterCreateRigelParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimToken

`func (o *ClusterCreateRigelParams) GetClaimToken() string`

GetClaimToken returns the ClaimToken field if non-nil, zero value otherwise.

### GetClaimTokenOk

`func (o *ClusterCreateRigelParams) GetClaimTokenOk() (*string, bool)`

GetClaimTokenOk returns a tuple with the ClaimToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimToken

`func (o *ClusterCreateRigelParams) SetClaimToken(v string)`

SetClaimToken sets ClaimToken field to given value.


### SetClaimTokenNil

`func (o *ClusterCreateRigelParams) SetClaimTokenNil(b bool)`

 SetClaimTokenNil sets the value for ClaimToken to be an explicit nil

### UnsetClaimToken
`func (o *ClusterCreateRigelParams) UnsetClaimToken()`

UnsetClaimToken ensures that no value is present for ClaimToken, not even an explicit nil
### GetNodes

`func (o *ClusterCreateRigelParams) GetNodes() []RigelClusterNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ClusterCreateRigelParams) GetNodesOk() (*[]RigelClusterNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ClusterCreateRigelParams) SetNodes(v []RigelClusterNode)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


