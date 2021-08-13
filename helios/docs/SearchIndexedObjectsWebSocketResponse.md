# SearchIndexedObjectsWebSocketResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **NullableInt64** | ID of the message, to which the response belongs. This ID is returned as sent by the client. | 
**Status** | Pointer to **NullableString** | Denotes if the search operation is complete, errored or in-progress and if it is safe to close the websocket by the client in case there aren&#39;t any additional search requests to be sent by the client. | [optional] 
**Body** | [**HeliosSearchIndexedObjectsResponseBody**](HeliosSearchIndexedObjectsResponseBody.md) |  | 
**Error** | Pointer to **NullableString** | Error message, this will be populated if status is Error | [optional] 

## Methods

### NewSearchIndexedObjectsWebSocketResponse

`func NewSearchIndexedObjectsWebSocketResponse(messageId NullableInt64, body HeliosSearchIndexedObjectsResponseBody, ) *SearchIndexedObjectsWebSocketResponse`

NewSearchIndexedObjectsWebSocketResponse instantiates a new SearchIndexedObjectsWebSocketResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIndexedObjectsWebSocketResponseWithDefaults

`func NewSearchIndexedObjectsWebSocketResponseWithDefaults() *SearchIndexedObjectsWebSocketResponse`

NewSearchIndexedObjectsWebSocketResponseWithDefaults instantiates a new SearchIndexedObjectsWebSocketResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *SearchIndexedObjectsWebSocketResponse) GetMessageId() int64`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *SearchIndexedObjectsWebSocketResponse) GetMessageIdOk() (*int64, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *SearchIndexedObjectsWebSocketResponse) SetMessageId(v int64)`

SetMessageId sets MessageId field to given value.


### SetMessageIdNil

`func (o *SearchIndexedObjectsWebSocketResponse) SetMessageIdNil(b bool)`

 SetMessageIdNil sets the value for MessageId to be an explicit nil

### UnsetMessageId
`func (o *SearchIndexedObjectsWebSocketResponse) UnsetMessageId()`

UnsetMessageId ensures that no value is present for MessageId, not even an explicit nil
### GetStatus

`func (o *SearchIndexedObjectsWebSocketResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchIndexedObjectsWebSocketResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchIndexedObjectsWebSocketResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SearchIndexedObjectsWebSocketResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SearchIndexedObjectsWebSocketResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SearchIndexedObjectsWebSocketResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetBody

`func (o *SearchIndexedObjectsWebSocketResponse) GetBody() HeliosSearchIndexedObjectsResponseBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SearchIndexedObjectsWebSocketResponse) GetBodyOk() (*HeliosSearchIndexedObjectsResponseBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SearchIndexedObjectsWebSocketResponse) SetBody(v HeliosSearchIndexedObjectsResponseBody)`

SetBody sets Body field to given value.


### GetError

`func (o *SearchIndexedObjectsWebSocketResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SearchIndexedObjectsWebSocketResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SearchIndexedObjectsWebSocketResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SearchIndexedObjectsWebSocketResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *SearchIndexedObjectsWebSocketResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *SearchIndexedObjectsWebSocketResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


