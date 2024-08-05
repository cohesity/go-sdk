# ProtectionRunSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludePublicStatuses** | Pointer to **[]string** | If this is empty, any run status will be qualified. Else, only the runs matching one of the status will be qualified. | [optional] 
**StartTimeInterval** | [**TimeInterval**](TimeInterval.md) |  | 

## Methods

### NewProtectionRunSelector

`func NewProtectionRunSelector(startTimeInterval TimeInterval, ) *ProtectionRunSelector`

NewProtectionRunSelector instantiates a new ProtectionRunSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionRunSelectorWithDefaults

`func NewProtectionRunSelectorWithDefaults() *ProtectionRunSelector`

NewProtectionRunSelectorWithDefaults instantiates a new ProtectionRunSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludePublicStatuses

`func (o *ProtectionRunSelector) GetIncludePublicStatuses() []string`

GetIncludePublicStatuses returns the IncludePublicStatuses field if non-nil, zero value otherwise.

### GetIncludePublicStatusesOk

`func (o *ProtectionRunSelector) GetIncludePublicStatusesOk() (*[]string, bool)`

GetIncludePublicStatusesOk returns a tuple with the IncludePublicStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePublicStatuses

`func (o *ProtectionRunSelector) SetIncludePublicStatuses(v []string)`

SetIncludePublicStatuses sets IncludePublicStatuses field to given value.

### HasIncludePublicStatuses

`func (o *ProtectionRunSelector) HasIncludePublicStatuses() bool`

HasIncludePublicStatuses returns a boolean if a field has been set.

### GetStartTimeInterval

`func (o *ProtectionRunSelector) GetStartTimeInterval() TimeInterval`

GetStartTimeInterval returns the StartTimeInterval field if non-nil, zero value otherwise.

### GetStartTimeIntervalOk

`func (o *ProtectionRunSelector) GetStartTimeIntervalOk() (*TimeInterval, bool)`

GetStartTimeIntervalOk returns a tuple with the StartTimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeInterval

`func (o *ProtectionRunSelector) SetStartTimeInterval(v TimeInterval)`

SetStartTimeInterval sets StartTimeInterval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


