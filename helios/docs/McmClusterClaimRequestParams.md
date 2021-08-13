# McmClusterClaimRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **int64** | Specifies the cluster id. | 
**ClusterIncarnationId** | **int64** | Specifies the cluster incarnation id. | 
**ClusterName** | **string** | Specifies the cluster name. | 
**ClaimToken** | **string** | Claim token used for authentication. | 

## Methods

### NewMcmClusterClaimRequestParams

`func NewMcmClusterClaimRequestParams(clusterId int64, clusterIncarnationId int64, clusterName string, claimToken string, ) *McmClusterClaimRequestParams`

NewMcmClusterClaimRequestParams instantiates a new McmClusterClaimRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmClusterClaimRequestParamsWithDefaults

`func NewMcmClusterClaimRequestParamsWithDefaults() *McmClusterClaimRequestParams`

NewMcmClusterClaimRequestParamsWithDefaults instantiates a new McmClusterClaimRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *McmClusterClaimRequestParams) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *McmClusterClaimRequestParams) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *McmClusterClaimRequestParams) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.


### GetClusterIncarnationId

`func (o *McmClusterClaimRequestParams) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *McmClusterClaimRequestParams) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *McmClusterClaimRequestParams) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.


### GetClusterName

`func (o *McmClusterClaimRequestParams) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *McmClusterClaimRequestParams) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *McmClusterClaimRequestParams) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetClaimToken

`func (o *McmClusterClaimRequestParams) GetClaimToken() string`

GetClaimToken returns the ClaimToken field if non-nil, zero value otherwise.

### GetClaimTokenOk

`func (o *McmClusterClaimRequestParams) GetClaimTokenOk() (*string, bool)`

GetClaimTokenOk returns a tuple with the ClaimToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimToken

`func (o *McmClusterClaimRequestParams) SetClaimToken(v string)`

SetClaimToken sets ClaimToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


