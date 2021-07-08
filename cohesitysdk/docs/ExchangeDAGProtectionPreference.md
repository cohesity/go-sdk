# ExchangeDAGProtectionPreference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PassiveCopyPreferenceServerGuidList** | Pointer to **[]string** | Specifies the preference order of the exchange servers from which passive database copies should be protected. The preference order is descending which indicates that passive database copy in the first server in the list gets the highest preference. | [optional] 
**PassiveOnly** | Pointer to **NullableBool** | Specifies that only passive database copies should be protected if this is set to true. If this is set to false, both active and passive database copies can be protected. | [optional] 
**UseUserSpecifiedPassivePreferenceOrder** | Pointer to **NullableBool** | Specifies to use the user specified preference order of exchange servers from which the passive database copies should be protected if this is set to true.  Every copy of an Exchange database in a DAG is assigned an activation preference number. This number is used by the system as part of the passive database activation process. If this bool flag is set to false, the reverse order of activation is used while choosing between passive copies. | [optional] 

## Methods

### NewExchangeDAGProtectionPreference

`func NewExchangeDAGProtectionPreference() *ExchangeDAGProtectionPreference`

NewExchangeDAGProtectionPreference instantiates a new ExchangeDAGProtectionPreference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeDAGProtectionPreferenceWithDefaults

`func NewExchangeDAGProtectionPreferenceWithDefaults() *ExchangeDAGProtectionPreference`

NewExchangeDAGProtectionPreferenceWithDefaults instantiates a new ExchangeDAGProtectionPreference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassiveCopyPreferenceServerGuidList

`func (o *ExchangeDAGProtectionPreference) GetPassiveCopyPreferenceServerGuidList() []string`

GetPassiveCopyPreferenceServerGuidList returns the PassiveCopyPreferenceServerGuidList field if non-nil, zero value otherwise.

### GetPassiveCopyPreferenceServerGuidListOk

`func (o *ExchangeDAGProtectionPreference) GetPassiveCopyPreferenceServerGuidListOk() (*[]string, bool)`

GetPassiveCopyPreferenceServerGuidListOk returns a tuple with the PassiveCopyPreferenceServerGuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassiveCopyPreferenceServerGuidList

`func (o *ExchangeDAGProtectionPreference) SetPassiveCopyPreferenceServerGuidList(v []string)`

SetPassiveCopyPreferenceServerGuidList sets PassiveCopyPreferenceServerGuidList field to given value.

### HasPassiveCopyPreferenceServerGuidList

`func (o *ExchangeDAGProtectionPreference) HasPassiveCopyPreferenceServerGuidList() bool`

HasPassiveCopyPreferenceServerGuidList returns a boolean if a field has been set.

### SetPassiveCopyPreferenceServerGuidListNil

`func (o *ExchangeDAGProtectionPreference) SetPassiveCopyPreferenceServerGuidListNil(b bool)`

 SetPassiveCopyPreferenceServerGuidListNil sets the value for PassiveCopyPreferenceServerGuidList to be an explicit nil

### UnsetPassiveCopyPreferenceServerGuidList
`func (o *ExchangeDAGProtectionPreference) UnsetPassiveCopyPreferenceServerGuidList()`

UnsetPassiveCopyPreferenceServerGuidList ensures that no value is present for PassiveCopyPreferenceServerGuidList, not even an explicit nil
### GetPassiveOnly

`func (o *ExchangeDAGProtectionPreference) GetPassiveOnly() bool`

GetPassiveOnly returns the PassiveOnly field if non-nil, zero value otherwise.

### GetPassiveOnlyOk

`func (o *ExchangeDAGProtectionPreference) GetPassiveOnlyOk() (*bool, bool)`

GetPassiveOnlyOk returns a tuple with the PassiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassiveOnly

`func (o *ExchangeDAGProtectionPreference) SetPassiveOnly(v bool)`

SetPassiveOnly sets PassiveOnly field to given value.

### HasPassiveOnly

`func (o *ExchangeDAGProtectionPreference) HasPassiveOnly() bool`

HasPassiveOnly returns a boolean if a field has been set.

### SetPassiveOnlyNil

`func (o *ExchangeDAGProtectionPreference) SetPassiveOnlyNil(b bool)`

 SetPassiveOnlyNil sets the value for PassiveOnly to be an explicit nil

### UnsetPassiveOnly
`func (o *ExchangeDAGProtectionPreference) UnsetPassiveOnly()`

UnsetPassiveOnly ensures that no value is present for PassiveOnly, not even an explicit nil
### GetUseUserSpecifiedPassivePreferenceOrder

`func (o *ExchangeDAGProtectionPreference) GetUseUserSpecifiedPassivePreferenceOrder() bool`

GetUseUserSpecifiedPassivePreferenceOrder returns the UseUserSpecifiedPassivePreferenceOrder field if non-nil, zero value otherwise.

### GetUseUserSpecifiedPassivePreferenceOrderOk

`func (o *ExchangeDAGProtectionPreference) GetUseUserSpecifiedPassivePreferenceOrderOk() (*bool, bool)`

GetUseUserSpecifiedPassivePreferenceOrderOk returns a tuple with the UseUserSpecifiedPassivePreferenceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUserSpecifiedPassivePreferenceOrder

`func (o *ExchangeDAGProtectionPreference) SetUseUserSpecifiedPassivePreferenceOrder(v bool)`

SetUseUserSpecifiedPassivePreferenceOrder sets UseUserSpecifiedPassivePreferenceOrder field to given value.

### HasUseUserSpecifiedPassivePreferenceOrder

`func (o *ExchangeDAGProtectionPreference) HasUseUserSpecifiedPassivePreferenceOrder() bool`

HasUseUserSpecifiedPassivePreferenceOrder returns a boolean if a field has been set.

### SetUseUserSpecifiedPassivePreferenceOrderNil

`func (o *ExchangeDAGProtectionPreference) SetUseUserSpecifiedPassivePreferenceOrderNil(b bool)`

 SetUseUserSpecifiedPassivePreferenceOrderNil sets the value for UseUserSpecifiedPassivePreferenceOrder to be an explicit nil

### UnsetUseUserSpecifiedPassivePreferenceOrder
`func (o *ExchangeDAGProtectionPreference) UnsetUseUserSpecifiedPassivePreferenceOrder()`

UnsetUseUserSpecifiedPassivePreferenceOrder ensures that no value is present for UseUserSpecifiedPassivePreferenceOrder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


