# GetProtectionRunsStatusResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionRunsStatsList** | Pointer to [**[]ProtectionRunsStatsList**](ProtectionRunsStatsList.md) | Specifies a list of protection runs stats taken at different time. | [optional] 

## Methods

### NewGetProtectionRunsStatusResponseBody

`func NewGetProtectionRunsStatusResponseBody() *GetProtectionRunsStatusResponseBody`

NewGetProtectionRunsStatusResponseBody instantiates a new GetProtectionRunsStatusResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProtectionRunsStatusResponseBodyWithDefaults

`func NewGetProtectionRunsStatusResponseBodyWithDefaults() *GetProtectionRunsStatusResponseBody`

NewGetProtectionRunsStatusResponseBodyWithDefaults instantiates a new GetProtectionRunsStatusResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionRunsStatsList

`func (o *GetProtectionRunsStatusResponseBody) GetProtectionRunsStatsList() []ProtectionRunsStatsList`

GetProtectionRunsStatsList returns the ProtectionRunsStatsList field if non-nil, zero value otherwise.

### GetProtectionRunsStatsListOk

`func (o *GetProtectionRunsStatusResponseBody) GetProtectionRunsStatsListOk() (*[]ProtectionRunsStatsList, bool)`

GetProtectionRunsStatsListOk returns a tuple with the ProtectionRunsStatsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionRunsStatsList

`func (o *GetProtectionRunsStatusResponseBody) SetProtectionRunsStatsList(v []ProtectionRunsStatsList)`

SetProtectionRunsStatsList sets ProtectionRunsStatsList field to given value.

### HasProtectionRunsStatsList

`func (o *GetProtectionRunsStatusResponseBody) HasProtectionRunsStatsList() bool`

HasProtectionRunsStatsList returns a boolean if a field has been set.

### SetProtectionRunsStatsListNil

`func (o *GetProtectionRunsStatusResponseBody) SetProtectionRunsStatsListNil(b bool)`

 SetProtectionRunsStatsListNil sets the value for ProtectionRunsStatsList to be an explicit nil

### UnsetProtectionRunsStatsList
`func (o *GetProtectionRunsStatusResponseBody) UnsetProtectionRunsStatsList()`

UnsetProtectionRunsStatsList ensures that no value is present for ProtectionRunsStatsList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


