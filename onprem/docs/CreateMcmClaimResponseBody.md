# CreateMcmClaimResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | Pointer to **NullableString** | Specfies the type of entity. | [optional] 
**RigelParams** | Pointer to [**McmRigelClaimResponseParams**](McmRigelClaimResponseParams.md) |  | [optional] 
**ClusterParams** | Pointer to [**McmClusterClaimResponseParams**](McmClusterClaimResponseParams.md) |  | [optional] 

## Methods

### NewCreateMcmClaimResponseBody

`func NewCreateMcmClaimResponseBody() *CreateMcmClaimResponseBody`

NewCreateMcmClaimResponseBody instantiates a new CreateMcmClaimResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMcmClaimResponseBodyWithDefaults

`func NewCreateMcmClaimResponseBodyWithDefaults() *CreateMcmClaimResponseBody`

NewCreateMcmClaimResponseBodyWithDefaults instantiates a new CreateMcmClaimResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *CreateMcmClaimResponseBody) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CreateMcmClaimResponseBody) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CreateMcmClaimResponseBody) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *CreateMcmClaimResponseBody) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### SetEntityTypeNil

`func (o *CreateMcmClaimResponseBody) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *CreateMcmClaimResponseBody) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetRigelParams

`func (o *CreateMcmClaimResponseBody) GetRigelParams() McmRigelClaimResponseParams`

GetRigelParams returns the RigelParams field if non-nil, zero value otherwise.

### GetRigelParamsOk

`func (o *CreateMcmClaimResponseBody) GetRigelParamsOk() (*McmRigelClaimResponseParams, bool)`

GetRigelParamsOk returns a tuple with the RigelParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRigelParams

`func (o *CreateMcmClaimResponseBody) SetRigelParams(v McmRigelClaimResponseParams)`

SetRigelParams sets RigelParams field to given value.

### HasRigelParams

`func (o *CreateMcmClaimResponseBody) HasRigelParams() bool`

HasRigelParams returns a boolean if a field has been set.

### GetClusterParams

`func (o *CreateMcmClaimResponseBody) GetClusterParams() McmClusterClaimResponseParams`

GetClusterParams returns the ClusterParams field if non-nil, zero value otherwise.

### GetClusterParamsOk

`func (o *CreateMcmClaimResponseBody) GetClusterParamsOk() (*McmClusterClaimResponseParams, bool)`

GetClusterParamsOk returns a tuple with the ClusterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterParams

`func (o *CreateMcmClaimResponseBody) SetClusterParams(v McmClusterClaimResponseParams)`

SetClusterParams sets ClusterParams field to given value.

### HasClusterParams

`func (o *CreateMcmClaimResponseBody) HasClusterParams() bool`

HasClusterParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


