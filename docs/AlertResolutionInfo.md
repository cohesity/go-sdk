# AlertResolutionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResolutionDetails** | Pointer to **NullableString** | Specifies detailed notes about the Resolution. | [optional] 
**ResolutionSummary** | Pointer to **NullableString** | Specifies short description about the Resolution. | [optional] 

## Methods

### NewAlertResolutionInfo

`func NewAlertResolutionInfo() *AlertResolutionInfo`

NewAlertResolutionInfo instantiates a new AlertResolutionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertResolutionInfoWithDefaults

`func NewAlertResolutionInfoWithDefaults() *AlertResolutionInfo`

NewAlertResolutionInfoWithDefaults instantiates a new AlertResolutionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolutionDetails

`func (o *AlertResolutionInfo) GetResolutionDetails() string`

GetResolutionDetails returns the ResolutionDetails field if non-nil, zero value otherwise.

### GetResolutionDetailsOk

`func (o *AlertResolutionInfo) GetResolutionDetailsOk() (*string, bool)`

GetResolutionDetailsOk returns a tuple with the ResolutionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionDetails

`func (o *AlertResolutionInfo) SetResolutionDetails(v string)`

SetResolutionDetails sets ResolutionDetails field to given value.

### HasResolutionDetails

`func (o *AlertResolutionInfo) HasResolutionDetails() bool`

HasResolutionDetails returns a boolean if a field has been set.

### SetResolutionDetailsNil

`func (o *AlertResolutionInfo) SetResolutionDetailsNil(b bool)`

 SetResolutionDetailsNil sets the value for ResolutionDetails to be an explicit nil

### UnsetResolutionDetails
`func (o *AlertResolutionInfo) UnsetResolutionDetails()`

UnsetResolutionDetails ensures that no value is present for ResolutionDetails, not even an explicit nil
### GetResolutionSummary

`func (o *AlertResolutionInfo) GetResolutionSummary() string`

GetResolutionSummary returns the ResolutionSummary field if non-nil, zero value otherwise.

### GetResolutionSummaryOk

`func (o *AlertResolutionInfo) GetResolutionSummaryOk() (*string, bool)`

GetResolutionSummaryOk returns a tuple with the ResolutionSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionSummary

`func (o *AlertResolutionInfo) SetResolutionSummary(v string)`

SetResolutionSummary sets ResolutionSummary field to given value.

### HasResolutionSummary

`func (o *AlertResolutionInfo) HasResolutionSummary() bool`

HasResolutionSummary returns a boolean if a field has been set.

### SetResolutionSummaryNil

`func (o *AlertResolutionInfo) SetResolutionSummaryNil(b bool)`

 SetResolutionSummaryNil sets the value for ResolutionSummary to be an explicit nil

### UnsetResolutionSummary
`func (o *AlertResolutionInfo) UnsetResolutionSummary()`

UnsetResolutionSummary ensures that no value is present for ResolutionSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


