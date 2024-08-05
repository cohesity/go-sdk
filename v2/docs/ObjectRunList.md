# ObjectRunList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **string** | Specifies an error (if any) corresponding to this objectId and run. | [optional] 
**ObjectId** | **int64** | Specifies the id of the object. | 
**RunStartTimeUsecs** | **int64** | Specifies the start time of the run that was updated. | 

## Methods

### NewObjectRunList

`func NewObjectRunList(objectId int64, runStartTimeUsecs int64, ) *ObjectRunList`

NewObjectRunList instantiates a new ObjectRunList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectRunListWithDefaults

`func NewObjectRunListWithDefaults() *ObjectRunList`

NewObjectRunListWithDefaults instantiates a new ObjectRunList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *ObjectRunList) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ObjectRunList) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ObjectRunList) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ObjectRunList) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetObjectId

`func (o *ObjectRunList) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ObjectRunList) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ObjectRunList) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetRunStartTimeUsecs

`func (o *ObjectRunList) GetRunStartTimeUsecs() int64`

GetRunStartTimeUsecs returns the RunStartTimeUsecs field if non-nil, zero value otherwise.

### GetRunStartTimeUsecsOk

`func (o *ObjectRunList) GetRunStartTimeUsecsOk() (*int64, bool)`

GetRunStartTimeUsecsOk returns a tuple with the RunStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStartTimeUsecs

`func (o *ObjectRunList) SetRunStartTimeUsecs(v int64)`

SetRunStartTimeUsecs sets RunStartTimeUsecs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


