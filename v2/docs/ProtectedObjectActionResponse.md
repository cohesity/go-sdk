# ProtectedObjectActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Specifies the action type to be performed on object getting protected. Based on selected action, provide the action params. | [optional] 
**Objects** | Pointer to [**[]ActionObjectLevelResponse**](ActionObjectLevelResponse.md) | Specifies the list of objects on which the provided action was performed. | [optional] 

## Methods

### NewProtectedObjectActionResponse

`func NewProtectedObjectActionResponse() *ProtectedObjectActionResponse`

NewProtectedObjectActionResponse instantiates a new ProtectedObjectActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedObjectActionResponseWithDefaults

`func NewProtectedObjectActionResponseWithDefaults() *ProtectedObjectActionResponse`

NewProtectedObjectActionResponseWithDefaults instantiates a new ProtectedObjectActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ProtectedObjectActionResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ProtectedObjectActionResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ProtectedObjectActionResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ProtectedObjectActionResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetObjects

`func (o *ProtectedObjectActionResponse) GetObjects() []ActionObjectLevelResponse`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ProtectedObjectActionResponse) GetObjectsOk() (*[]ActionObjectLevelResponse, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ProtectedObjectActionResponse) SetObjects(v []ActionObjectLevelResponse)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ProtectedObjectActionResponse) HasObjects() bool`

HasObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


