# BaseosPatchLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Log** | Pointer to **string** | Baseos patch Log file. | [optional] 
**Status** | Pointer to **string** | Baseos patch application status | [optional] 

## Methods

### NewBaseosPatchLog

`func NewBaseosPatchLog() *BaseosPatchLog`

NewBaseosPatchLog instantiates a new BaseosPatchLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseosPatchLogWithDefaults

`func NewBaseosPatchLogWithDefaults() *BaseosPatchLog`

NewBaseosPatchLogWithDefaults instantiates a new BaseosPatchLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLog

`func (o *BaseosPatchLog) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *BaseosPatchLog) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *BaseosPatchLog) SetLog(v string)`

SetLog sets Log field to given value.

### HasLog

`func (o *BaseosPatchLog) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetStatus

`func (o *BaseosPatchLog) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseosPatchLog) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseosPatchLog) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseosPatchLog) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


