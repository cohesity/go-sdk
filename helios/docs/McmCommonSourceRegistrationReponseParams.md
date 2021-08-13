# McmCommonSourceRegistrationReponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster id. | [optional] [readonly] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the cluster incarnation id. | [optional] [readonly] 
**RegionId** | Pointer to **NullableString** | Specifies the region id. | [optional] [readonly] 
**Id** | Pointer to **NullableString** | Source Registration ID. This can be used to retrieve, edit or delete the source registration. | [optional] [readonly] 
**ConnectionId** | Pointer to **NullableInt64** | Specifies the id of the connection from where this source is reachable. | [optional] 
**SourceId** | Pointer to **NullableString** | ID of top level source object discovered after the registration. | [optional] [readonly] 
**Environment** | Pointer to **NullableString** | Specifies the environment type of the Protection Source. | [optional] 

## Methods

### NewMcmCommonSourceRegistrationReponseParams

`func NewMcmCommonSourceRegistrationReponseParams() *McmCommonSourceRegistrationReponseParams`

NewMcmCommonSourceRegistrationReponseParams instantiates a new McmCommonSourceRegistrationReponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmCommonSourceRegistrationReponseParamsWithDefaults

`func NewMcmCommonSourceRegistrationReponseParamsWithDefaults() *McmCommonSourceRegistrationReponseParams`

NewMcmCommonSourceRegistrationReponseParamsWithDefaults instantiates a new McmCommonSourceRegistrationReponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *McmCommonSourceRegistrationReponseParams) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *McmCommonSourceRegistrationReponseParams) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *McmCommonSourceRegistrationReponseParams) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *McmCommonSourceRegistrationReponseParams) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *McmCommonSourceRegistrationReponseParams) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *McmCommonSourceRegistrationReponseParams) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *McmCommonSourceRegistrationReponseParams) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *McmCommonSourceRegistrationReponseParams) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *McmCommonSourceRegistrationReponseParams) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *McmCommonSourceRegistrationReponseParams) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *McmCommonSourceRegistrationReponseParams) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *McmCommonSourceRegistrationReponseParams) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetRegionId

`func (o *McmCommonSourceRegistrationReponseParams) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *McmCommonSourceRegistrationReponseParams) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *McmCommonSourceRegistrationReponseParams) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *McmCommonSourceRegistrationReponseParams) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *McmCommonSourceRegistrationReponseParams) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *McmCommonSourceRegistrationReponseParams) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetId

`func (o *McmCommonSourceRegistrationReponseParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *McmCommonSourceRegistrationReponseParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *McmCommonSourceRegistrationReponseParams) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *McmCommonSourceRegistrationReponseParams) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *McmCommonSourceRegistrationReponseParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *McmCommonSourceRegistrationReponseParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetConnectionId

`func (o *McmCommonSourceRegistrationReponseParams) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *McmCommonSourceRegistrationReponseParams) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *McmCommonSourceRegistrationReponseParams) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *McmCommonSourceRegistrationReponseParams) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *McmCommonSourceRegistrationReponseParams) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *McmCommonSourceRegistrationReponseParams) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetSourceId

`func (o *McmCommonSourceRegistrationReponseParams) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *McmCommonSourceRegistrationReponseParams) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *McmCommonSourceRegistrationReponseParams) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *McmCommonSourceRegistrationReponseParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *McmCommonSourceRegistrationReponseParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *McmCommonSourceRegistrationReponseParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetEnvironment

`func (o *McmCommonSourceRegistrationReponseParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *McmCommonSourceRegistrationReponseParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *McmCommonSourceRegistrationReponseParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *McmCommonSourceRegistrationReponseParams) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *McmCommonSourceRegistrationReponseParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *McmCommonSourceRegistrationReponseParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


