# CreateViewFailoverRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlannedFailoverParams** | Pointer to [**CreateViewFailoverRequestPlannedFailoverParams**](CreateViewFailoverRequestPlannedFailoverParams.md) |  | [optional] 
**Type** | **NullableString** | Specifies the failover type.&lt;br&gt; &#39;Planned&#39; indicates this is a planned failover.&lt;br&gt; &#39;Unplanned&#39; indicates this is an unplanned failover, which is used when source cluster is down. | 
**UnplannedFailoverParams** | Pointer to [**CreateViewFailoverRequestUnplannedFailoverParams**](CreateViewFailoverRequestUnplannedFailoverParams.md) |  | [optional] 

## Methods

### NewCreateViewFailoverRequest

`func NewCreateViewFailoverRequest(type_ NullableString, ) *CreateViewFailoverRequest`

NewCreateViewFailoverRequest instantiates a new CreateViewFailoverRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateViewFailoverRequestWithDefaults

`func NewCreateViewFailoverRequestWithDefaults() *CreateViewFailoverRequest`

NewCreateViewFailoverRequestWithDefaults instantiates a new CreateViewFailoverRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlannedFailoverParams

`func (o *CreateViewFailoverRequest) GetPlannedFailoverParams() CreateViewFailoverRequestPlannedFailoverParams`

GetPlannedFailoverParams returns the PlannedFailoverParams field if non-nil, zero value otherwise.

### GetPlannedFailoverParamsOk

`func (o *CreateViewFailoverRequest) GetPlannedFailoverParamsOk() (*CreateViewFailoverRequestPlannedFailoverParams, bool)`

GetPlannedFailoverParamsOk returns a tuple with the PlannedFailoverParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedFailoverParams

`func (o *CreateViewFailoverRequest) SetPlannedFailoverParams(v CreateViewFailoverRequestPlannedFailoverParams)`

SetPlannedFailoverParams sets PlannedFailoverParams field to given value.

### HasPlannedFailoverParams

`func (o *CreateViewFailoverRequest) HasPlannedFailoverParams() bool`

HasPlannedFailoverParams returns a boolean if a field has been set.

### GetType

`func (o *CreateViewFailoverRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateViewFailoverRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateViewFailoverRequest) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CreateViewFailoverRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CreateViewFailoverRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUnplannedFailoverParams

`func (o *CreateViewFailoverRequest) GetUnplannedFailoverParams() CreateViewFailoverRequestUnplannedFailoverParams`

GetUnplannedFailoverParams returns the UnplannedFailoverParams field if non-nil, zero value otherwise.

### GetUnplannedFailoverParamsOk

`func (o *CreateViewFailoverRequest) GetUnplannedFailoverParamsOk() (*CreateViewFailoverRequestUnplannedFailoverParams, bool)`

GetUnplannedFailoverParamsOk returns a tuple with the UnplannedFailoverParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnplannedFailoverParams

`func (o *CreateViewFailoverRequest) SetUnplannedFailoverParams(v CreateViewFailoverRequestUnplannedFailoverParams)`

SetUnplannedFailoverParams sets UnplannedFailoverParams field to given value.

### HasUnplannedFailoverParams

`func (o *CreateViewFailoverRequest) HasUnplannedFailoverParams() bool`

HasUnplannedFailoverParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


