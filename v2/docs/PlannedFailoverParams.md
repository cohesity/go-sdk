# PlannedFailoverParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreparePlannedFailverParams** | Pointer to [**PlannedFailoverParamsPreparePlannedFailverParams**](PlannedFailoverParamsPreparePlannedFailverParams.md) |  | [optional] 
**Type** | **NullableString** | Spcifies the planned failover type.&lt;br&gt; &#39;Prepare&#39; indicates this is a preparation for failover.&lt;br&gt; &#39;Finalize&#39; indicates this is finalization of failover. After this is done, the view can be used as source view. | 

## Methods

### NewPlannedFailoverParams

`func NewPlannedFailoverParams(type_ NullableString, ) *PlannedFailoverParams`

NewPlannedFailoverParams instantiates a new PlannedFailoverParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannedFailoverParamsWithDefaults

`func NewPlannedFailoverParamsWithDefaults() *PlannedFailoverParams`

NewPlannedFailoverParamsWithDefaults instantiates a new PlannedFailoverParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreparePlannedFailverParams

`func (o *PlannedFailoverParams) GetPreparePlannedFailverParams() PlannedFailoverParamsPreparePlannedFailverParams`

GetPreparePlannedFailverParams returns the PreparePlannedFailverParams field if non-nil, zero value otherwise.

### GetPreparePlannedFailverParamsOk

`func (o *PlannedFailoverParams) GetPreparePlannedFailverParamsOk() (*PlannedFailoverParamsPreparePlannedFailverParams, bool)`

GetPreparePlannedFailverParamsOk returns a tuple with the PreparePlannedFailverParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparePlannedFailverParams

`func (o *PlannedFailoverParams) SetPreparePlannedFailverParams(v PlannedFailoverParamsPreparePlannedFailverParams)`

SetPreparePlannedFailverParams sets PreparePlannedFailverParams field to given value.

### HasPreparePlannedFailverParams

`func (o *PlannedFailoverParams) HasPreparePlannedFailverParams() bool`

HasPreparePlannedFailverParams returns a boolean if a field has been set.

### GetType

`func (o *PlannedFailoverParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlannedFailoverParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlannedFailoverParams) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PlannedFailoverParams) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PlannedFailoverParams) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


