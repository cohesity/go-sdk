# ProtectedObjectRunNowActionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]RunNowActionObjectLevelParams**](RunNowActionObjectLevelParams.md) | Specifies the list of objects to perform an action. If provided object id is not explicitly protected by object protection, then given action will not be performed on that. | [optional] 

## Methods

### NewProtectedObjectRunNowActionParams

`func NewProtectedObjectRunNowActionParams() *ProtectedObjectRunNowActionParams`

NewProtectedObjectRunNowActionParams instantiates a new ProtectedObjectRunNowActionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedObjectRunNowActionParamsWithDefaults

`func NewProtectedObjectRunNowActionParamsWithDefaults() *ProtectedObjectRunNowActionParams`

NewProtectedObjectRunNowActionParamsWithDefaults instantiates a new ProtectedObjectRunNowActionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *ProtectedObjectRunNowActionParams) GetObjects() []RunNowActionObjectLevelParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ProtectedObjectRunNowActionParams) GetObjectsOk() (*[]RunNowActionObjectLevelParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ProtectedObjectRunNowActionParams) SetObjects(v []RunNowActionObjectLevelParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ProtectedObjectRunNowActionParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


