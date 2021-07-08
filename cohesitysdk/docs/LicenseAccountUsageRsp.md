# LicenseAccountUsageRsp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureOverusage** | Pointer to [**[]Overusage**](Overusage.md) | Holds information about consumption usage of overused features | [optional] 
**FreeSetupMode** | Pointer to **NullableBool** | Free Setup Mode | [optional] 
**IsTrial** | Pointer to **NullableBool** | Check if trial license. | [optional] 
**Last12MonthsAvgEntitlement** | Pointer to [**map[string][]MonthlyEntitlement**](array.md) | Holds monthly avg usage values of feature | [optional] 
**Last12MonthsAvgUsage** | Pointer to [**map[string][]MonthlyUsage**](array.md) | Holds monthly avg usage values of feature | [optional] 
**Last30DaysEntitlement** | Pointer to [**map[string][]DailyEntitlement**](array.md) | Holds daily entitled capacity values of feature | [optional] 
**Last30DaysUsage** | Pointer to [**map[string][]DailyUsage**](array.md) | Holds daily usage values of feature | [optional] 
**LastUpdateTime** | Pointer to **NullableInt64** | Last time, this report was updated. | [optional] 
**LicensedUsage** | Pointer to [**[]LicensedUsage**](LicensedUsage.md) | LicenseFeatureUsages holds information about each feature from license orders. | [optional] 
**TrialExpiration** | Pointer to **NullableInt64** | Trial expiration period. | [optional] 
**Usage** | Pointer to [**map[string][]FeatureUsage**](array.md) | Creating a map of cluster id and feature usage to make it consistent display usage UI for the helios server license page UI. | [optional] 

## Methods

### NewLicenseAccountUsageRsp

`func NewLicenseAccountUsageRsp() *LicenseAccountUsageRsp`

NewLicenseAccountUsageRsp instantiates a new LicenseAccountUsageRsp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseAccountUsageRspWithDefaults

`func NewLicenseAccountUsageRspWithDefaults() *LicenseAccountUsageRsp`

NewLicenseAccountUsageRspWithDefaults instantiates a new LicenseAccountUsageRsp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureOverusage

`func (o *LicenseAccountUsageRsp) GetFeatureOverusage() []Overusage`

GetFeatureOverusage returns the FeatureOverusage field if non-nil, zero value otherwise.

### GetFeatureOverusageOk

`func (o *LicenseAccountUsageRsp) GetFeatureOverusageOk() (*[]Overusage, bool)`

GetFeatureOverusageOk returns a tuple with the FeatureOverusage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureOverusage

`func (o *LicenseAccountUsageRsp) SetFeatureOverusage(v []Overusage)`

SetFeatureOverusage sets FeatureOverusage field to given value.

### HasFeatureOverusage

`func (o *LicenseAccountUsageRsp) HasFeatureOverusage() bool`

HasFeatureOverusage returns a boolean if a field has been set.

### SetFeatureOverusageNil

`func (o *LicenseAccountUsageRsp) SetFeatureOverusageNil(b bool)`

 SetFeatureOverusageNil sets the value for FeatureOverusage to be an explicit nil

### UnsetFeatureOverusage
`func (o *LicenseAccountUsageRsp) UnsetFeatureOverusage()`

UnsetFeatureOverusage ensures that no value is present for FeatureOverusage, not even an explicit nil
### GetFreeSetupMode

`func (o *LicenseAccountUsageRsp) GetFreeSetupMode() bool`

GetFreeSetupMode returns the FreeSetupMode field if non-nil, zero value otherwise.

### GetFreeSetupModeOk

`func (o *LicenseAccountUsageRsp) GetFreeSetupModeOk() (*bool, bool)`

GetFreeSetupModeOk returns a tuple with the FreeSetupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSetupMode

`func (o *LicenseAccountUsageRsp) SetFreeSetupMode(v bool)`

SetFreeSetupMode sets FreeSetupMode field to given value.

### HasFreeSetupMode

`func (o *LicenseAccountUsageRsp) HasFreeSetupMode() bool`

HasFreeSetupMode returns a boolean if a field has been set.

### SetFreeSetupModeNil

`func (o *LicenseAccountUsageRsp) SetFreeSetupModeNil(b bool)`

 SetFreeSetupModeNil sets the value for FreeSetupMode to be an explicit nil

### UnsetFreeSetupMode
`func (o *LicenseAccountUsageRsp) UnsetFreeSetupMode()`

UnsetFreeSetupMode ensures that no value is present for FreeSetupMode, not even an explicit nil
### GetIsTrial

`func (o *LicenseAccountUsageRsp) GetIsTrial() bool`

GetIsTrial returns the IsTrial field if non-nil, zero value otherwise.

### GetIsTrialOk

`func (o *LicenseAccountUsageRsp) GetIsTrialOk() (*bool, bool)`

GetIsTrialOk returns a tuple with the IsTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrial

`func (o *LicenseAccountUsageRsp) SetIsTrial(v bool)`

SetIsTrial sets IsTrial field to given value.

### HasIsTrial

`func (o *LicenseAccountUsageRsp) HasIsTrial() bool`

HasIsTrial returns a boolean if a field has been set.

### SetIsTrialNil

`func (o *LicenseAccountUsageRsp) SetIsTrialNil(b bool)`

 SetIsTrialNil sets the value for IsTrial to be an explicit nil

### UnsetIsTrial
`func (o *LicenseAccountUsageRsp) UnsetIsTrial()`

UnsetIsTrial ensures that no value is present for IsTrial, not even an explicit nil
### GetLast12MonthsAvgEntitlement

`func (o *LicenseAccountUsageRsp) GetLast12MonthsAvgEntitlement() map[string][]MonthlyEntitlement`

GetLast12MonthsAvgEntitlement returns the Last12MonthsAvgEntitlement field if non-nil, zero value otherwise.

### GetLast12MonthsAvgEntitlementOk

`func (o *LicenseAccountUsageRsp) GetLast12MonthsAvgEntitlementOk() (*map[string][]MonthlyEntitlement, bool)`

GetLast12MonthsAvgEntitlementOk returns a tuple with the Last12MonthsAvgEntitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast12MonthsAvgEntitlement

`func (o *LicenseAccountUsageRsp) SetLast12MonthsAvgEntitlement(v map[string][]MonthlyEntitlement)`

SetLast12MonthsAvgEntitlement sets Last12MonthsAvgEntitlement field to given value.

### HasLast12MonthsAvgEntitlement

`func (o *LicenseAccountUsageRsp) HasLast12MonthsAvgEntitlement() bool`

HasLast12MonthsAvgEntitlement returns a boolean if a field has been set.

### SetLast12MonthsAvgEntitlementNil

`func (o *LicenseAccountUsageRsp) SetLast12MonthsAvgEntitlementNil(b bool)`

 SetLast12MonthsAvgEntitlementNil sets the value for Last12MonthsAvgEntitlement to be an explicit nil

### UnsetLast12MonthsAvgEntitlement
`func (o *LicenseAccountUsageRsp) UnsetLast12MonthsAvgEntitlement()`

UnsetLast12MonthsAvgEntitlement ensures that no value is present for Last12MonthsAvgEntitlement, not even an explicit nil
### GetLast12MonthsAvgUsage

`func (o *LicenseAccountUsageRsp) GetLast12MonthsAvgUsage() map[string][]MonthlyUsage`

GetLast12MonthsAvgUsage returns the Last12MonthsAvgUsage field if non-nil, zero value otherwise.

### GetLast12MonthsAvgUsageOk

`func (o *LicenseAccountUsageRsp) GetLast12MonthsAvgUsageOk() (*map[string][]MonthlyUsage, bool)`

GetLast12MonthsAvgUsageOk returns a tuple with the Last12MonthsAvgUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast12MonthsAvgUsage

`func (o *LicenseAccountUsageRsp) SetLast12MonthsAvgUsage(v map[string][]MonthlyUsage)`

SetLast12MonthsAvgUsage sets Last12MonthsAvgUsage field to given value.

### HasLast12MonthsAvgUsage

`func (o *LicenseAccountUsageRsp) HasLast12MonthsAvgUsage() bool`

HasLast12MonthsAvgUsage returns a boolean if a field has been set.

### SetLast12MonthsAvgUsageNil

`func (o *LicenseAccountUsageRsp) SetLast12MonthsAvgUsageNil(b bool)`

 SetLast12MonthsAvgUsageNil sets the value for Last12MonthsAvgUsage to be an explicit nil

### UnsetLast12MonthsAvgUsage
`func (o *LicenseAccountUsageRsp) UnsetLast12MonthsAvgUsage()`

UnsetLast12MonthsAvgUsage ensures that no value is present for Last12MonthsAvgUsage, not even an explicit nil
### GetLast30DaysEntitlement

`func (o *LicenseAccountUsageRsp) GetLast30DaysEntitlement() map[string][]DailyEntitlement`

GetLast30DaysEntitlement returns the Last30DaysEntitlement field if non-nil, zero value otherwise.

### GetLast30DaysEntitlementOk

`func (o *LicenseAccountUsageRsp) GetLast30DaysEntitlementOk() (*map[string][]DailyEntitlement, bool)`

GetLast30DaysEntitlementOk returns a tuple with the Last30DaysEntitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast30DaysEntitlement

`func (o *LicenseAccountUsageRsp) SetLast30DaysEntitlement(v map[string][]DailyEntitlement)`

SetLast30DaysEntitlement sets Last30DaysEntitlement field to given value.

### HasLast30DaysEntitlement

`func (o *LicenseAccountUsageRsp) HasLast30DaysEntitlement() bool`

HasLast30DaysEntitlement returns a boolean if a field has been set.

### SetLast30DaysEntitlementNil

`func (o *LicenseAccountUsageRsp) SetLast30DaysEntitlementNil(b bool)`

 SetLast30DaysEntitlementNil sets the value for Last30DaysEntitlement to be an explicit nil

### UnsetLast30DaysEntitlement
`func (o *LicenseAccountUsageRsp) UnsetLast30DaysEntitlement()`

UnsetLast30DaysEntitlement ensures that no value is present for Last30DaysEntitlement, not even an explicit nil
### GetLast30DaysUsage

`func (o *LicenseAccountUsageRsp) GetLast30DaysUsage() map[string][]DailyUsage`

GetLast30DaysUsage returns the Last30DaysUsage field if non-nil, zero value otherwise.

### GetLast30DaysUsageOk

`func (o *LicenseAccountUsageRsp) GetLast30DaysUsageOk() (*map[string][]DailyUsage, bool)`

GetLast30DaysUsageOk returns a tuple with the Last30DaysUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast30DaysUsage

`func (o *LicenseAccountUsageRsp) SetLast30DaysUsage(v map[string][]DailyUsage)`

SetLast30DaysUsage sets Last30DaysUsage field to given value.

### HasLast30DaysUsage

`func (o *LicenseAccountUsageRsp) HasLast30DaysUsage() bool`

HasLast30DaysUsage returns a boolean if a field has been set.

### SetLast30DaysUsageNil

`func (o *LicenseAccountUsageRsp) SetLast30DaysUsageNil(b bool)`

 SetLast30DaysUsageNil sets the value for Last30DaysUsage to be an explicit nil

### UnsetLast30DaysUsage
`func (o *LicenseAccountUsageRsp) UnsetLast30DaysUsage()`

UnsetLast30DaysUsage ensures that no value is present for Last30DaysUsage, not even an explicit nil
### GetLastUpdateTime

`func (o *LicenseAccountUsageRsp) GetLastUpdateTime() int64`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *LicenseAccountUsageRsp) GetLastUpdateTimeOk() (*int64, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *LicenseAccountUsageRsp) SetLastUpdateTime(v int64)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *LicenseAccountUsageRsp) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### SetLastUpdateTimeNil

`func (o *LicenseAccountUsageRsp) SetLastUpdateTimeNil(b bool)`

 SetLastUpdateTimeNil sets the value for LastUpdateTime to be an explicit nil

### UnsetLastUpdateTime
`func (o *LicenseAccountUsageRsp) UnsetLastUpdateTime()`

UnsetLastUpdateTime ensures that no value is present for LastUpdateTime, not even an explicit nil
### GetLicensedUsage

`func (o *LicenseAccountUsageRsp) GetLicensedUsage() []LicensedUsage`

GetLicensedUsage returns the LicensedUsage field if non-nil, zero value otherwise.

### GetLicensedUsageOk

`func (o *LicenseAccountUsageRsp) GetLicensedUsageOk() (*[]LicensedUsage, bool)`

GetLicensedUsageOk returns a tuple with the LicensedUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedUsage

`func (o *LicenseAccountUsageRsp) SetLicensedUsage(v []LicensedUsage)`

SetLicensedUsage sets LicensedUsage field to given value.

### HasLicensedUsage

`func (o *LicenseAccountUsageRsp) HasLicensedUsage() bool`

HasLicensedUsage returns a boolean if a field has been set.

### SetLicensedUsageNil

`func (o *LicenseAccountUsageRsp) SetLicensedUsageNil(b bool)`

 SetLicensedUsageNil sets the value for LicensedUsage to be an explicit nil

### UnsetLicensedUsage
`func (o *LicenseAccountUsageRsp) UnsetLicensedUsage()`

UnsetLicensedUsage ensures that no value is present for LicensedUsage, not even an explicit nil
### GetTrialExpiration

`func (o *LicenseAccountUsageRsp) GetTrialExpiration() int64`

GetTrialExpiration returns the TrialExpiration field if non-nil, zero value otherwise.

### GetTrialExpirationOk

`func (o *LicenseAccountUsageRsp) GetTrialExpirationOk() (*int64, bool)`

GetTrialExpirationOk returns a tuple with the TrialExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialExpiration

`func (o *LicenseAccountUsageRsp) SetTrialExpiration(v int64)`

SetTrialExpiration sets TrialExpiration field to given value.

### HasTrialExpiration

`func (o *LicenseAccountUsageRsp) HasTrialExpiration() bool`

HasTrialExpiration returns a boolean if a field has been set.

### SetTrialExpirationNil

`func (o *LicenseAccountUsageRsp) SetTrialExpirationNil(b bool)`

 SetTrialExpirationNil sets the value for TrialExpiration to be an explicit nil

### UnsetTrialExpiration
`func (o *LicenseAccountUsageRsp) UnsetTrialExpiration()`

UnsetTrialExpiration ensures that no value is present for TrialExpiration, not even an explicit nil
### GetUsage

`func (o *LicenseAccountUsageRsp) GetUsage() map[string][]FeatureUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *LicenseAccountUsageRsp) GetUsageOk() (*map[string][]FeatureUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *LicenseAccountUsageRsp) SetUsage(v map[string][]FeatureUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *LicenseAccountUsageRsp) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *LicenseAccountUsageRsp) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *LicenseAccountUsageRsp) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


