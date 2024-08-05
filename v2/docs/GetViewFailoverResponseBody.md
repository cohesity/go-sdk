# GetViewFailoverResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failovers** | Pointer to [**[]Failover**](Failover.md) | Specifies a list of failovers. | [optional] 

## Methods

### NewGetViewFailoverResponseBody

`func NewGetViewFailoverResponseBody() *GetViewFailoverResponseBody`

NewGetViewFailoverResponseBody instantiates a new GetViewFailoverResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetViewFailoverResponseBodyWithDefaults

`func NewGetViewFailoverResponseBodyWithDefaults() *GetViewFailoverResponseBody`

NewGetViewFailoverResponseBodyWithDefaults instantiates a new GetViewFailoverResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailovers

`func (o *GetViewFailoverResponseBody) GetFailovers() []Failover`

GetFailovers returns the Failovers field if non-nil, zero value otherwise.

### GetFailoversOk

`func (o *GetViewFailoverResponseBody) GetFailoversOk() (*[]Failover, bool)`

GetFailoversOk returns a tuple with the Failovers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailovers

`func (o *GetViewFailoverResponseBody) SetFailovers(v []Failover)`

SetFailovers sets Failovers field to given value.

### HasFailovers

`func (o *GetViewFailoverResponseBody) HasFailovers() bool`

HasFailovers returns a boolean if a field has been set.

### SetFailoversNil

`func (o *GetViewFailoverResponseBody) SetFailoversNil(b bool)`

 SetFailoversNil sets the value for Failovers to be an explicit nil

### UnsetFailovers
`func (o *GetViewFailoverResponseBody) UnsetFailovers()`

UnsetFailovers ensures that no value is present for Failovers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


