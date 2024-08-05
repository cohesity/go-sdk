# ClaimConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimToken** | **string** | Claim token used for authentication. | 
**ClusterId** | Pointer to **int64** | Cluster id of the connector. | [optional] 
**ClusterIncarnationId** | Pointer to **int64** | Incarnation id of the connector. | [optional] 
**ConnectorGuid** | **int64** | Unique id of the connector. | 
**ConnectorIp** | Pointer to **string** | IP address of the connector. | [optional] 
**ConnectorName** | Pointer to **string** | Name of the connector. | [optional] 
**SoftwareVersion** | Pointer to **string** | Software version of the connector. | [optional] 

## Methods

### NewClaimConnectorRequest

`func NewClaimConnectorRequest(claimToken string, connectorGuid int64, ) *ClaimConnectorRequest`

NewClaimConnectorRequest instantiates a new ClaimConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimConnectorRequestWithDefaults

`func NewClaimConnectorRequestWithDefaults() *ClaimConnectorRequest`

NewClaimConnectorRequestWithDefaults instantiates a new ClaimConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimToken

`func (o *ClaimConnectorRequest) GetClaimToken() string`

GetClaimToken returns the ClaimToken field if non-nil, zero value otherwise.

### GetClaimTokenOk

`func (o *ClaimConnectorRequest) GetClaimTokenOk() (*string, bool)`

GetClaimTokenOk returns a tuple with the ClaimToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimToken

`func (o *ClaimConnectorRequest) SetClaimToken(v string)`

SetClaimToken sets ClaimToken field to given value.


### GetClusterId

`func (o *ClaimConnectorRequest) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClaimConnectorRequest) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClaimConnectorRequest) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ClaimConnectorRequest) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetClusterIncarnationId

`func (o *ClaimConnectorRequest) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *ClaimConnectorRequest) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *ClaimConnectorRequest) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *ClaimConnectorRequest) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### GetConnectorGuid

`func (o *ClaimConnectorRequest) GetConnectorGuid() int64`

GetConnectorGuid returns the ConnectorGuid field if non-nil, zero value otherwise.

### GetConnectorGuidOk

`func (o *ClaimConnectorRequest) GetConnectorGuidOk() (*int64, bool)`

GetConnectorGuidOk returns a tuple with the ConnectorGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorGuid

`func (o *ClaimConnectorRequest) SetConnectorGuid(v int64)`

SetConnectorGuid sets ConnectorGuid field to given value.


### GetConnectorIp

`func (o *ClaimConnectorRequest) GetConnectorIp() string`

GetConnectorIp returns the ConnectorIp field if non-nil, zero value otherwise.

### GetConnectorIpOk

`func (o *ClaimConnectorRequest) GetConnectorIpOk() (*string, bool)`

GetConnectorIpOk returns a tuple with the ConnectorIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorIp

`func (o *ClaimConnectorRequest) SetConnectorIp(v string)`

SetConnectorIp sets ConnectorIp field to given value.

### HasConnectorIp

`func (o *ClaimConnectorRequest) HasConnectorIp() bool`

HasConnectorIp returns a boolean if a field has been set.

### GetConnectorName

`func (o *ClaimConnectorRequest) GetConnectorName() string`

GetConnectorName returns the ConnectorName field if non-nil, zero value otherwise.

### GetConnectorNameOk

`func (o *ClaimConnectorRequest) GetConnectorNameOk() (*string, bool)`

GetConnectorNameOk returns a tuple with the ConnectorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorName

`func (o *ClaimConnectorRequest) SetConnectorName(v string)`

SetConnectorName sets ConnectorName field to given value.

### HasConnectorName

`func (o *ClaimConnectorRequest) HasConnectorName() bool`

HasConnectorName returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *ClaimConnectorRequest) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *ClaimConnectorRequest) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *ClaimConnectorRequest) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *ClaimConnectorRequest) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


