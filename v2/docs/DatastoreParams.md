# DatastoreParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the Id of the datastore. | [optional] 
**MaxConcurrentStreams** | Pointer to **NullableInt32** | Specifies the max concurrent stream per datastore. | [optional] 

## Methods

### NewDatastoreParams

`func NewDatastoreParams() *DatastoreParams`

NewDatastoreParams instantiates a new DatastoreParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatastoreParamsWithDefaults

`func NewDatastoreParamsWithDefaults() *DatastoreParams`

NewDatastoreParamsWithDefaults instantiates a new DatastoreParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatastoreParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatastoreParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatastoreParams) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DatastoreParams) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DatastoreParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DatastoreParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMaxConcurrentStreams

`func (o *DatastoreParams) GetMaxConcurrentStreams() int32`

GetMaxConcurrentStreams returns the MaxConcurrentStreams field if non-nil, zero value otherwise.

### GetMaxConcurrentStreamsOk

`func (o *DatastoreParams) GetMaxConcurrentStreamsOk() (*int32, bool)`

GetMaxConcurrentStreamsOk returns a tuple with the MaxConcurrentStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentStreams

`func (o *DatastoreParams) SetMaxConcurrentStreams(v int32)`

SetMaxConcurrentStreams sets MaxConcurrentStreams field to given value.

### HasMaxConcurrentStreams

`func (o *DatastoreParams) HasMaxConcurrentStreams() bool`

HasMaxConcurrentStreams returns a boolean if a field has been set.

### SetMaxConcurrentStreamsNil

`func (o *DatastoreParams) SetMaxConcurrentStreamsNil(b bool)`

 SetMaxConcurrentStreamsNil sets the value for MaxConcurrentStreams to be an explicit nil

### UnsetMaxConcurrentStreams
`func (o *DatastoreParams) UnsetMaxConcurrentStreams()`

UnsetMaxConcurrentStreams ensures that no value is present for MaxConcurrentStreams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


