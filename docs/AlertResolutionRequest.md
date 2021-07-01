# AlertResolutionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertIdList** | Pointer to **[]string** | Specifies list of alerts resolved by a Resolution, which are specified by Alert Ids. | [optional] 
**ResolutionDetails** | Pointer to [**AlertResolutionInfo**](AlertResolutionInfo.md) |  | [optional] 

## Methods

### NewAlertResolutionRequest

`func NewAlertResolutionRequest() *AlertResolutionRequest`

NewAlertResolutionRequest instantiates a new AlertResolutionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertResolutionRequestWithDefaults

`func NewAlertResolutionRequestWithDefaults() *AlertResolutionRequest`

NewAlertResolutionRequestWithDefaults instantiates a new AlertResolutionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertIdList

`func (o *AlertResolutionRequest) GetAlertIdList() []string`

GetAlertIdList returns the AlertIdList field if non-nil, zero value otherwise.

### GetAlertIdListOk

`func (o *AlertResolutionRequest) GetAlertIdListOk() (*[]string, bool)`

GetAlertIdListOk returns a tuple with the AlertIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertIdList

`func (o *AlertResolutionRequest) SetAlertIdList(v []string)`

SetAlertIdList sets AlertIdList field to given value.

### HasAlertIdList

`func (o *AlertResolutionRequest) HasAlertIdList() bool`

HasAlertIdList returns a boolean if a field has been set.

### SetAlertIdListNil

`func (o *AlertResolutionRequest) SetAlertIdListNil(b bool)`

 SetAlertIdListNil sets the value for AlertIdList to be an explicit nil

### UnsetAlertIdList
`func (o *AlertResolutionRequest) UnsetAlertIdList()`

UnsetAlertIdList ensures that no value is present for AlertIdList, not even an explicit nil
### GetResolutionDetails

`func (o *AlertResolutionRequest) GetResolutionDetails() AlertResolutionInfo`

GetResolutionDetails returns the ResolutionDetails field if non-nil, zero value otherwise.

### GetResolutionDetailsOk

`func (o *AlertResolutionRequest) GetResolutionDetailsOk() (*AlertResolutionInfo, bool)`

GetResolutionDetailsOk returns a tuple with the ResolutionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionDetails

`func (o *AlertResolutionRequest) SetResolutionDetails(v AlertResolutionInfo)`

SetResolutionDetails sets ResolutionDetails field to given value.

### HasResolutionDetails

`func (o *AlertResolutionRequest) HasResolutionDetails() bool`

HasResolutionDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


