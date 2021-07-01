# AlertResolution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertIdList** | Pointer to **[]string** | Specifies list of Alerts resolved by a Resolution, which are specified by Alert Ids. | [optional] 
**ResolutionDetails** | Pointer to [**AlertResolutionDetails**](AlertResolutionDetails.md) |  | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies unique tenantIds of the alert contained in this resolution. | [optional] 

## Methods

### NewAlertResolution

`func NewAlertResolution() *AlertResolution`

NewAlertResolution instantiates a new AlertResolution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertResolutionWithDefaults

`func NewAlertResolutionWithDefaults() *AlertResolution`

NewAlertResolutionWithDefaults instantiates a new AlertResolution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertIdList

`func (o *AlertResolution) GetAlertIdList() []string`

GetAlertIdList returns the AlertIdList field if non-nil, zero value otherwise.

### GetAlertIdListOk

`func (o *AlertResolution) GetAlertIdListOk() (*[]string, bool)`

GetAlertIdListOk returns a tuple with the AlertIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertIdList

`func (o *AlertResolution) SetAlertIdList(v []string)`

SetAlertIdList sets AlertIdList field to given value.

### HasAlertIdList

`func (o *AlertResolution) HasAlertIdList() bool`

HasAlertIdList returns a boolean if a field has been set.

### SetAlertIdListNil

`func (o *AlertResolution) SetAlertIdListNil(b bool)`

 SetAlertIdListNil sets the value for AlertIdList to be an explicit nil

### UnsetAlertIdList
`func (o *AlertResolution) UnsetAlertIdList()`

UnsetAlertIdList ensures that no value is present for AlertIdList, not even an explicit nil
### GetResolutionDetails

`func (o *AlertResolution) GetResolutionDetails() AlertResolutionDetails`

GetResolutionDetails returns the ResolutionDetails field if non-nil, zero value otherwise.

### GetResolutionDetailsOk

`func (o *AlertResolution) GetResolutionDetailsOk() (*AlertResolutionDetails, bool)`

GetResolutionDetailsOk returns a tuple with the ResolutionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionDetails

`func (o *AlertResolution) SetResolutionDetails(v AlertResolutionDetails)`

SetResolutionDetails sets ResolutionDetails field to given value.

### HasResolutionDetails

`func (o *AlertResolution) HasResolutionDetails() bool`

HasResolutionDetails returns a boolean if a field has been set.

### GetTenantIds

`func (o *AlertResolution) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *AlertResolution) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *AlertResolution) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *AlertResolution) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### SetTenantIdsNil

`func (o *AlertResolution) SetTenantIdsNil(b bool)`

 SetTenantIdsNil sets the value for TenantIds to be an explicit nil

### UnsetTenantIds
`func (o *AlertResolution) UnsetTenantIds()`

UnsetTenantIds ensures that no value is present for TenantIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


