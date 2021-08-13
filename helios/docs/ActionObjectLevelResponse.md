# ActionObjectLevelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the ID of the object. | 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] [readonly] 
**ResumeStatus** | Pointer to [**ResumeActionObjectLevelResponse**](ResumeActionObjectLevelResponse.md) |  | [optional] 
**PauseStatus** | Pointer to [**PauseActionObjectLevelResponse**](PauseActionObjectLevelResponse.md) |  | [optional] 
**RunNowStatus** | Pointer to [**RunNowActionObjectLevelResponse**](RunNowActionObjectLevelResponse.md) |  | [optional] 
**UnProtectStatus** | Pointer to [**UnprotectActionObjectLevelResponse**](UnprotectActionObjectLevelResponse.md) |  | [optional] 

## Methods

### NewActionObjectLevelResponse

`func NewActionObjectLevelResponse(id NullableInt64, ) *ActionObjectLevelResponse`

NewActionObjectLevelResponse instantiates a new ActionObjectLevelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionObjectLevelResponseWithDefaults

`func NewActionObjectLevelResponseWithDefaults() *ActionObjectLevelResponse`

NewActionObjectLevelResponseWithDefaults instantiates a new ActionObjectLevelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionObjectLevelResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionObjectLevelResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionObjectLevelResponse) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ActionObjectLevelResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ActionObjectLevelResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ActionObjectLevelResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionObjectLevelResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionObjectLevelResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActionObjectLevelResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ActionObjectLevelResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ActionObjectLevelResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetResumeStatus

`func (o *ActionObjectLevelResponse) GetResumeStatus() ResumeActionObjectLevelResponse`

GetResumeStatus returns the ResumeStatus field if non-nil, zero value otherwise.

### GetResumeStatusOk

`func (o *ActionObjectLevelResponse) GetResumeStatusOk() (*ResumeActionObjectLevelResponse, bool)`

GetResumeStatusOk returns a tuple with the ResumeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeStatus

`func (o *ActionObjectLevelResponse) SetResumeStatus(v ResumeActionObjectLevelResponse)`

SetResumeStatus sets ResumeStatus field to given value.

### HasResumeStatus

`func (o *ActionObjectLevelResponse) HasResumeStatus() bool`

HasResumeStatus returns a boolean if a field has been set.

### GetPauseStatus

`func (o *ActionObjectLevelResponse) GetPauseStatus() PauseActionObjectLevelResponse`

GetPauseStatus returns the PauseStatus field if non-nil, zero value otherwise.

### GetPauseStatusOk

`func (o *ActionObjectLevelResponse) GetPauseStatusOk() (*PauseActionObjectLevelResponse, bool)`

GetPauseStatusOk returns a tuple with the PauseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseStatus

`func (o *ActionObjectLevelResponse) SetPauseStatus(v PauseActionObjectLevelResponse)`

SetPauseStatus sets PauseStatus field to given value.

### HasPauseStatus

`func (o *ActionObjectLevelResponse) HasPauseStatus() bool`

HasPauseStatus returns a boolean if a field has been set.

### GetRunNowStatus

`func (o *ActionObjectLevelResponse) GetRunNowStatus() RunNowActionObjectLevelResponse`

GetRunNowStatus returns the RunNowStatus field if non-nil, zero value otherwise.

### GetRunNowStatusOk

`func (o *ActionObjectLevelResponse) GetRunNowStatusOk() (*RunNowActionObjectLevelResponse, bool)`

GetRunNowStatusOk returns a tuple with the RunNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunNowStatus

`func (o *ActionObjectLevelResponse) SetRunNowStatus(v RunNowActionObjectLevelResponse)`

SetRunNowStatus sets RunNowStatus field to given value.

### HasRunNowStatus

`func (o *ActionObjectLevelResponse) HasRunNowStatus() bool`

HasRunNowStatus returns a boolean if a field has been set.

### GetUnProtectStatus

`func (o *ActionObjectLevelResponse) GetUnProtectStatus() UnprotectActionObjectLevelResponse`

GetUnProtectStatus returns the UnProtectStatus field if non-nil, zero value otherwise.

### GetUnProtectStatusOk

`func (o *ActionObjectLevelResponse) GetUnProtectStatusOk() (*UnprotectActionObjectLevelResponse, bool)`

GetUnProtectStatusOk returns a tuple with the UnProtectStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnProtectStatus

`func (o *ActionObjectLevelResponse) SetUnProtectStatus(v UnprotectActionObjectLevelResponse)`

SetUnProtectStatus sets UnProtectStatus field to given value.

### HasUnProtectStatus

`func (o *ActionObjectLevelResponse) HasUnProtectStatus() bool`

HasUnProtectStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


