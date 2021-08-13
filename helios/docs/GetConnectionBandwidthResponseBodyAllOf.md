# GetConnectionBandwidthResponseBodyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **NullableInt64** | Connection Id for which bandwidth settings are to be returned | [optional] 

## Methods

### NewGetConnectionBandwidthResponseBodyAllOf

`func NewGetConnectionBandwidthResponseBodyAllOf() *GetConnectionBandwidthResponseBodyAllOf`

NewGetConnectionBandwidthResponseBodyAllOf instantiates a new GetConnectionBandwidthResponseBodyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionBandwidthResponseBodyAllOfWithDefaults

`func NewGetConnectionBandwidthResponseBodyAllOfWithDefaults() *GetConnectionBandwidthResponseBodyAllOf`

NewGetConnectionBandwidthResponseBodyAllOfWithDefaults instantiates a new GetConnectionBandwidthResponseBodyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *GetConnectionBandwidthResponseBodyAllOf) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *GetConnectionBandwidthResponseBodyAllOf) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *GetConnectionBandwidthResponseBodyAllOf) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *GetConnectionBandwidthResponseBodyAllOf) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *GetConnectionBandwidthResponseBodyAllOf) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *GetConnectionBandwidthResponseBodyAllOf) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


