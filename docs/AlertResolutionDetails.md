# AlertResolutionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResolutionDetails** | Pointer to **NullableString** | Specifies detailed notes about the Resolution. | [optional] 
**ResolutionId** | Pointer to **NullableInt64** | Specifies Unique id assigned by the Cohesity Cluster for this Resolution. | [optional] 
**ResolutionSummary** | Pointer to **NullableString** | Specifies short description about the Resolution. | [optional] 
**TimestampUsecs** | Pointer to **NullableInt64** | Specifies unix epoch timestamp (in microseconds) when the Alerts were resolved. | [optional] 
**UserName** | Pointer to **NullableString** | Specifies name of the Cohesity Cluster user who resolved the Alerts. | [optional] 

## Methods

### NewAlertResolutionDetails

`func NewAlertResolutionDetails() *AlertResolutionDetails`

NewAlertResolutionDetails instantiates a new AlertResolutionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertResolutionDetailsWithDefaults

`func NewAlertResolutionDetailsWithDefaults() *AlertResolutionDetails`

NewAlertResolutionDetailsWithDefaults instantiates a new AlertResolutionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolutionDetails

`func (o *AlertResolutionDetails) GetResolutionDetails() string`

GetResolutionDetails returns the ResolutionDetails field if non-nil, zero value otherwise.

### GetResolutionDetailsOk

`func (o *AlertResolutionDetails) GetResolutionDetailsOk() (*string, bool)`

GetResolutionDetailsOk returns a tuple with the ResolutionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionDetails

`func (o *AlertResolutionDetails) SetResolutionDetails(v string)`

SetResolutionDetails sets ResolutionDetails field to given value.

### HasResolutionDetails

`func (o *AlertResolutionDetails) HasResolutionDetails() bool`

HasResolutionDetails returns a boolean if a field has been set.

### SetResolutionDetailsNil

`func (o *AlertResolutionDetails) SetResolutionDetailsNil(b bool)`

 SetResolutionDetailsNil sets the value for ResolutionDetails to be an explicit nil

### UnsetResolutionDetails
`func (o *AlertResolutionDetails) UnsetResolutionDetails()`

UnsetResolutionDetails ensures that no value is present for ResolutionDetails, not even an explicit nil
### GetResolutionId

`func (o *AlertResolutionDetails) GetResolutionId() int64`

GetResolutionId returns the ResolutionId field if non-nil, zero value otherwise.

### GetResolutionIdOk

`func (o *AlertResolutionDetails) GetResolutionIdOk() (*int64, bool)`

GetResolutionIdOk returns a tuple with the ResolutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionId

`func (o *AlertResolutionDetails) SetResolutionId(v int64)`

SetResolutionId sets ResolutionId field to given value.

### HasResolutionId

`func (o *AlertResolutionDetails) HasResolutionId() bool`

HasResolutionId returns a boolean if a field has been set.

### SetResolutionIdNil

`func (o *AlertResolutionDetails) SetResolutionIdNil(b bool)`

 SetResolutionIdNil sets the value for ResolutionId to be an explicit nil

### UnsetResolutionId
`func (o *AlertResolutionDetails) UnsetResolutionId()`

UnsetResolutionId ensures that no value is present for ResolutionId, not even an explicit nil
### GetResolutionSummary

`func (o *AlertResolutionDetails) GetResolutionSummary() string`

GetResolutionSummary returns the ResolutionSummary field if non-nil, zero value otherwise.

### GetResolutionSummaryOk

`func (o *AlertResolutionDetails) GetResolutionSummaryOk() (*string, bool)`

GetResolutionSummaryOk returns a tuple with the ResolutionSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionSummary

`func (o *AlertResolutionDetails) SetResolutionSummary(v string)`

SetResolutionSummary sets ResolutionSummary field to given value.

### HasResolutionSummary

`func (o *AlertResolutionDetails) HasResolutionSummary() bool`

HasResolutionSummary returns a boolean if a field has been set.

### SetResolutionSummaryNil

`func (o *AlertResolutionDetails) SetResolutionSummaryNil(b bool)`

 SetResolutionSummaryNil sets the value for ResolutionSummary to be an explicit nil

### UnsetResolutionSummary
`func (o *AlertResolutionDetails) UnsetResolutionSummary()`

UnsetResolutionSummary ensures that no value is present for ResolutionSummary, not even an explicit nil
### GetTimestampUsecs

`func (o *AlertResolutionDetails) GetTimestampUsecs() int64`

GetTimestampUsecs returns the TimestampUsecs field if non-nil, zero value otherwise.

### GetTimestampUsecsOk

`func (o *AlertResolutionDetails) GetTimestampUsecsOk() (*int64, bool)`

GetTimestampUsecsOk returns a tuple with the TimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampUsecs

`func (o *AlertResolutionDetails) SetTimestampUsecs(v int64)`

SetTimestampUsecs sets TimestampUsecs field to given value.

### HasTimestampUsecs

`func (o *AlertResolutionDetails) HasTimestampUsecs() bool`

HasTimestampUsecs returns a boolean if a field has been set.

### SetTimestampUsecsNil

`func (o *AlertResolutionDetails) SetTimestampUsecsNil(b bool)`

 SetTimestampUsecsNil sets the value for TimestampUsecs to be an explicit nil

### UnsetTimestampUsecs
`func (o *AlertResolutionDetails) UnsetTimestampUsecs()`

UnsetTimestampUsecs ensures that no value is present for TimestampUsecs, not even an explicit nil
### GetUserName

`func (o *AlertResolutionDetails) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AlertResolutionDetails) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AlertResolutionDetails) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *AlertResolutionDetails) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *AlertResolutionDetails) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *AlertResolutionDetails) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


