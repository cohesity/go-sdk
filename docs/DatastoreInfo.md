# DatastoreInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **NullableInt32** | Specifies the capacity of the datastore in bytes. | [optional] 
**FreeSpace** | Pointer to **NullableInt32** | Specifies the available space on the datastore in bytes. | [optional] 

## Methods

### NewDatastoreInfo

`func NewDatastoreInfo() *DatastoreInfo`

NewDatastoreInfo instantiates a new DatastoreInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatastoreInfoWithDefaults

`func NewDatastoreInfoWithDefaults() *DatastoreInfo`

NewDatastoreInfoWithDefaults instantiates a new DatastoreInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *DatastoreInfo) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *DatastoreInfo) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *DatastoreInfo) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *DatastoreInfo) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### SetCapacityNil

`func (o *DatastoreInfo) SetCapacityNil(b bool)`

 SetCapacityNil sets the value for Capacity to be an explicit nil

### UnsetCapacity
`func (o *DatastoreInfo) UnsetCapacity()`

UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil
### GetFreeSpace

`func (o *DatastoreInfo) GetFreeSpace() int32`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *DatastoreInfo) GetFreeSpaceOk() (*int32, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *DatastoreInfo) SetFreeSpace(v int32)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *DatastoreInfo) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### SetFreeSpaceNil

`func (o *DatastoreInfo) SetFreeSpaceNil(b bool)`

 SetFreeSpaceNil sets the value for FreeSpace to be an explicit nil

### UnsetFreeSpace
`func (o *DatastoreInfo) UnsetFreeSpace()`

UnsetFreeSpace ensures that no value is present for FreeSpace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


