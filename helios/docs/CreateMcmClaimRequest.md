# CreateMcmClaimRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | **string** | Specfies the type of entity. | 
**RigelParams** | Pointer to [**McmRigelClaimRequestParams**](McmRigelClaimRequestParams.md) |  | [optional] 
**ClusterParams** | Pointer to [**McmClusterClaimRequestParams**](McmClusterClaimRequestParams.md) |  | [optional] 

## Methods

### NewCreateMcmClaimRequest

`func NewCreateMcmClaimRequest(entityType string, ) *CreateMcmClaimRequest`

NewCreateMcmClaimRequest instantiates a new CreateMcmClaimRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMcmClaimRequestWithDefaults

`func NewCreateMcmClaimRequestWithDefaults() *CreateMcmClaimRequest`

NewCreateMcmClaimRequestWithDefaults instantiates a new CreateMcmClaimRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *CreateMcmClaimRequest) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CreateMcmClaimRequest) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CreateMcmClaimRequest) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetRigelParams

`func (o *CreateMcmClaimRequest) GetRigelParams() McmRigelClaimRequestParams`

GetRigelParams returns the RigelParams field if non-nil, zero value otherwise.

### GetRigelParamsOk

`func (o *CreateMcmClaimRequest) GetRigelParamsOk() (*McmRigelClaimRequestParams, bool)`

GetRigelParamsOk returns a tuple with the RigelParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRigelParams

`func (o *CreateMcmClaimRequest) SetRigelParams(v McmRigelClaimRequestParams)`

SetRigelParams sets RigelParams field to given value.

### HasRigelParams

`func (o *CreateMcmClaimRequest) HasRigelParams() bool`

HasRigelParams returns a boolean if a field has been set.

### GetClusterParams

`func (o *CreateMcmClaimRequest) GetClusterParams() McmClusterClaimRequestParams`

GetClusterParams returns the ClusterParams field if non-nil, zero value otherwise.

### GetClusterParamsOk

`func (o *CreateMcmClaimRequest) GetClusterParamsOk() (*McmClusterClaimRequestParams, bool)`

GetClusterParamsOk returns a tuple with the ClusterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterParams

`func (o *CreateMcmClaimRequest) SetClusterParams(v McmClusterClaimRequestParams)`

SetClusterParams sets ClusterParams field to given value.

### HasClusterParams

`func (o *CreateMcmClaimRequest) HasClusterParams() bool`

HasClusterParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


