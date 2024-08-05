# TdmObjectTimelineEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]TdmObjectTimelineEvent**](TdmObjectTimelineEvent.md) | Specifies the collection of the timeline events, filtered by the specified criteria. | [optional] 

## Methods

### NewTdmObjectTimelineEvents

`func NewTdmObjectTimelineEvents() *TdmObjectTimelineEvents`

NewTdmObjectTimelineEvents instantiates a new TdmObjectTimelineEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTdmObjectTimelineEventsWithDefaults

`func NewTdmObjectTimelineEventsWithDefaults() *TdmObjectTimelineEvents`

NewTdmObjectTimelineEventsWithDefaults instantiates a new TdmObjectTimelineEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *TdmObjectTimelineEvents) GetEvents() []TdmObjectTimelineEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TdmObjectTimelineEvents) GetEventsOk() (*[]TdmObjectTimelineEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TdmObjectTimelineEvents) SetEvents(v []TdmObjectTimelineEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *TdmObjectTimelineEvents) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *TdmObjectTimelineEvents) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *TdmObjectTimelineEvents) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


