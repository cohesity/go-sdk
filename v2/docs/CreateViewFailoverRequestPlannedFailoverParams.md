# CreateViewFailoverRequestPlannedFailoverParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreparePlannedFailverParams** | Pointer to [**PlannedFailoverParamsPreparePlannedFailverParams**](PlannedFailoverParamsPreparePlannedFailverParams.md) |  | [optional] 
**Type** | **NullableString** | Spcifies the planned failover type.&lt;br&gt; &#39;Prepare&#39; indicates this is a preparation for failover.&lt;br&gt; &#39;Finalize&#39; indicates this is finalization of failover. After this is done, the view can be used as source view. | 

## Methods

### NewCreateViewFailoverRequestPlannedFailoverParams

`func NewCreateViewFailoverRequestPlannedFailoverParams(type_ NullableString, ) *CreateViewFailoverRequestPlannedFailoverParams`

NewCreateViewFailoverRequestPlannedFailoverParams instantiates a new CreateViewFailoverRequestPlannedFailoverParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateViewFailoverRequestPlannedFailoverParamsWithDefaults

`func NewCreateViewFailoverRequestPlannedFailoverParamsWithDefaults() *CreateViewFailoverRequestPlannedFailoverParams`

NewCreateViewFailoverRequestPlannedFailoverParamsWithDefaults instantiates a new CreateViewFailoverRequestPlannedFailoverParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreparePlannedFailverParams

`func (o *CreateViewFailoverRequestPlannedFailoverParams) GetPreparePlannedFailverParams() PlannedFailoverParamsPreparePlannedFailverParams`

GetPreparePlannedFailverParams returns the PreparePlannedFailverParams field if non-nil, zero value otherwise.

### GetPreparePlannedFailverParamsOk

`func (o *CreateViewFailoverRequestPlannedFailoverParams) GetPreparePlannedFailverParamsOk() (*PlannedFailoverParamsPreparePlannedFailverParams, bool)`

GetPreparePlannedFailverParamsOk returns a tuple with the PreparePlannedFailverParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparePlannedFailverParams

`func (o *CreateViewFailoverRequestPlannedFailoverParams) SetPreparePlannedFailverParams(v PlannedFailoverParamsPreparePlannedFailverParams)`

SetPreparePlannedFailverParams sets PreparePlannedFailverParams field to given value.

### HasPreparePlannedFailverParams

`func (o *CreateViewFailoverRequestPlannedFailoverParams) HasPreparePlannedFailverParams() bool`

HasPreparePlannedFailverParams returns a boolean if a field has been set.

### GetType

`func (o *CreateViewFailoverRequestPlannedFailoverParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateViewFailoverRequestPlannedFailoverParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateViewFailoverRequestPlannedFailoverParams) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CreateViewFailoverRequestPlannedFailoverParams) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CreateViewFailoverRequestPlannedFailoverParams) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


