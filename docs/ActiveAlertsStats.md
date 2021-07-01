# ActiveAlertsStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumCriticalAlerts** | Pointer to **NullableInt64** | Specifies the count of active critical Alerts excluding alerts that belong to other bucket. | [optional] 
**NumCriticalAlertsCategories** | Pointer to **NullableInt64** | Specifies the count of active critical alerts categories. | [optional] 
**NumHardwareAlerts** | Pointer to **NullableInt64** | Specifies the count of active hardware Alerts. | [optional] 
**NumHardwareCriticalAlerts** | Pointer to **NullableInt64** | Specifies the count of active hardware critical Alerts. | [optional] 
**NumHardwareInfoAlerts** | Pointer to **NullableInt64** | Specifies the count of active hardware info Alerts. | [optional] 
**NumHardwareWarningAlerts** | Pointer to **NullableInt64** | Specifies the count of active hardware warning Alerts. | [optional] 
**NumInfoAlerts** | Pointer to **NullableInt64** | Specifies the count of active info Alerts excluding alerts that belong to other bucket. | [optional] 
**NumInfoAlertsCategories** | Pointer to **NullableInt64** | Specifies the count of active info alerts categories. | [optional] 
**NumOtherAlerts** | Pointer to **NullableInt64** | Specifies the count of active Alerts of other bucket | [optional] 
**NumOtherCriticalAlerts** | Pointer to **NullableInt64** | Specifies the count of active other critical Alerts. | [optional] 
**NumOtherInfoAlerts** | Pointer to **NullableInt64** | Specifies the count of active other info Alerts. | [optional] 
**NumOtherWarningAlerts** | Pointer to **NullableInt64** | Specifies the count of active other warning Alerts. | [optional] 
**NumServiceAlerts** | Pointer to **NullableInt64** | Specifies the count of active service Alerts. | [optional] 
**NumServiceCriticalAlerts** | Pointer to **NullableInt64** | Specifies the count of active service critical Alerts. | [optional] 
**NumServiceInfoAlerts** | Pointer to **NullableInt64** | Specifies the count of active service info Alerts. | [optional] 
**NumServiceWarningAlerts** | Pointer to **NullableInt64** | Specifies the count of active service warning Alerts. | [optional] 
**NumSoftwareAlerts** | Pointer to **NullableInt64** | Specifies the count of active software Alerts. | [optional] 
**NumSoftwareCriticalAlerts** | Pointer to **NullableInt64** | Specifies the count of active software critical Alerts. | [optional] 
**NumSoftwareInfoAlerts** | Pointer to **NullableInt64** | Specifies the count of active software info Alerts. | [optional] 
**NumSoftwareWarningAlerts** | Pointer to **NullableInt64** | Specifies the count of active software warning Alerts. | [optional] 
**NumWarningAlerts** | Pointer to **NullableInt64** | Specifies the count of active warning Alerts excluding alerts that belong to other bucket. | [optional] 
**NumWarningAlertsCategories** | Pointer to **NullableInt64** | Specifies the count of active warning alerts categories. | [optional] 

## Methods

### NewActiveAlertsStats

`func NewActiveAlertsStats() *ActiveAlertsStats`

NewActiveAlertsStats instantiates a new ActiveAlertsStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveAlertsStatsWithDefaults

`func NewActiveAlertsStatsWithDefaults() *ActiveAlertsStats`

NewActiveAlertsStatsWithDefaults instantiates a new ActiveAlertsStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumCriticalAlerts

`func (o *ActiveAlertsStats) GetNumCriticalAlerts() int64`

GetNumCriticalAlerts returns the NumCriticalAlerts field if non-nil, zero value otherwise.

### GetNumCriticalAlertsOk

`func (o *ActiveAlertsStats) GetNumCriticalAlertsOk() (*int64, bool)`

GetNumCriticalAlertsOk returns a tuple with the NumCriticalAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCriticalAlerts

`func (o *ActiveAlertsStats) SetNumCriticalAlerts(v int64)`

SetNumCriticalAlerts sets NumCriticalAlerts field to given value.

### HasNumCriticalAlerts

`func (o *ActiveAlertsStats) HasNumCriticalAlerts() bool`

HasNumCriticalAlerts returns a boolean if a field has been set.

### SetNumCriticalAlertsNil

`func (o *ActiveAlertsStats) SetNumCriticalAlertsNil(b bool)`

 SetNumCriticalAlertsNil sets the value for NumCriticalAlerts to be an explicit nil

### UnsetNumCriticalAlerts
`func (o *ActiveAlertsStats) UnsetNumCriticalAlerts()`

UnsetNumCriticalAlerts ensures that no value is present for NumCriticalAlerts, not even an explicit nil
### GetNumCriticalAlertsCategories

`func (o *ActiveAlertsStats) GetNumCriticalAlertsCategories() int64`

GetNumCriticalAlertsCategories returns the NumCriticalAlertsCategories field if non-nil, zero value otherwise.

### GetNumCriticalAlertsCategoriesOk

`func (o *ActiveAlertsStats) GetNumCriticalAlertsCategoriesOk() (*int64, bool)`

GetNumCriticalAlertsCategoriesOk returns a tuple with the NumCriticalAlertsCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCriticalAlertsCategories

`func (o *ActiveAlertsStats) SetNumCriticalAlertsCategories(v int64)`

SetNumCriticalAlertsCategories sets NumCriticalAlertsCategories field to given value.

### HasNumCriticalAlertsCategories

`func (o *ActiveAlertsStats) HasNumCriticalAlertsCategories() bool`

HasNumCriticalAlertsCategories returns a boolean if a field has been set.

### SetNumCriticalAlertsCategoriesNil

`func (o *ActiveAlertsStats) SetNumCriticalAlertsCategoriesNil(b bool)`

 SetNumCriticalAlertsCategoriesNil sets the value for NumCriticalAlertsCategories to be an explicit nil

### UnsetNumCriticalAlertsCategories
`func (o *ActiveAlertsStats) UnsetNumCriticalAlertsCategories()`

UnsetNumCriticalAlertsCategories ensures that no value is present for NumCriticalAlertsCategories, not even an explicit nil
### GetNumHardwareAlerts

`func (o *ActiveAlertsStats) GetNumHardwareAlerts() int64`

GetNumHardwareAlerts returns the NumHardwareAlerts field if non-nil, zero value otherwise.

### GetNumHardwareAlertsOk

`func (o *ActiveAlertsStats) GetNumHardwareAlertsOk() (*int64, bool)`

GetNumHardwareAlertsOk returns a tuple with the NumHardwareAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumHardwareAlerts

`func (o *ActiveAlertsStats) SetNumHardwareAlerts(v int64)`

SetNumHardwareAlerts sets NumHardwareAlerts field to given value.

### HasNumHardwareAlerts

`func (o *ActiveAlertsStats) HasNumHardwareAlerts() bool`

HasNumHardwareAlerts returns a boolean if a field has been set.

### SetNumHardwareAlertsNil

`func (o *ActiveAlertsStats) SetNumHardwareAlertsNil(b bool)`

 SetNumHardwareAlertsNil sets the value for NumHardwareAlerts to be an explicit nil

### UnsetNumHardwareAlerts
`func (o *ActiveAlertsStats) UnsetNumHardwareAlerts()`

UnsetNumHardwareAlerts ensures that no value is present for NumHardwareAlerts, not even an explicit nil
### GetNumHardwareCriticalAlerts

`func (o *ActiveAlertsStats) GetNumHardwareCriticalAlerts() int64`

GetNumHardwareCriticalAlerts returns the NumHardwareCriticalAlerts field if non-nil, zero value otherwise.

### GetNumHardwareCriticalAlertsOk

`func (o *ActiveAlertsStats) GetNumHardwareCriticalAlertsOk() (*int64, bool)`

GetNumHardwareCriticalAlertsOk returns a tuple with the NumHardwareCriticalAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumHardwareCriticalAlerts

`func (o *ActiveAlertsStats) SetNumHardwareCriticalAlerts(v int64)`

SetNumHardwareCriticalAlerts sets NumHardwareCriticalAlerts field to given value.

### HasNumHardwareCriticalAlerts

`func (o *ActiveAlertsStats) HasNumHardwareCriticalAlerts() bool`

HasNumHardwareCriticalAlerts returns a boolean if a field has been set.

### SetNumHardwareCriticalAlertsNil

`func (o *ActiveAlertsStats) SetNumHardwareCriticalAlertsNil(b bool)`

 SetNumHardwareCriticalAlertsNil sets the value for NumHardwareCriticalAlerts to be an explicit nil

### UnsetNumHardwareCriticalAlerts
`func (o *ActiveAlertsStats) UnsetNumHardwareCriticalAlerts()`

UnsetNumHardwareCriticalAlerts ensures that no value is present for NumHardwareCriticalAlerts, not even an explicit nil
### GetNumHardwareInfoAlerts

`func (o *ActiveAlertsStats) GetNumHardwareInfoAlerts() int64`

GetNumHardwareInfoAlerts returns the NumHardwareInfoAlerts field if non-nil, zero value otherwise.

### GetNumHardwareInfoAlertsOk

`func (o *ActiveAlertsStats) GetNumHardwareInfoAlertsOk() (*int64, bool)`

GetNumHardwareInfoAlertsOk returns a tuple with the NumHardwareInfoAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumHardwareInfoAlerts

`func (o *ActiveAlertsStats) SetNumHardwareInfoAlerts(v int64)`

SetNumHardwareInfoAlerts sets NumHardwareInfoAlerts field to given value.

### HasNumHardwareInfoAlerts

`func (o *ActiveAlertsStats) HasNumHardwareInfoAlerts() bool`

HasNumHardwareInfoAlerts returns a boolean if a field has been set.

### SetNumHardwareInfoAlertsNil

`func (o *ActiveAlertsStats) SetNumHardwareInfoAlertsNil(b bool)`

 SetNumHardwareInfoAlertsNil sets the value for NumHardwareInfoAlerts to be an explicit nil

### UnsetNumHardwareInfoAlerts
`func (o *ActiveAlertsStats) UnsetNumHardwareInfoAlerts()`

UnsetNumHardwareInfoAlerts ensures that no value is present for NumHardwareInfoAlerts, not even an explicit nil
### GetNumHardwareWarningAlerts

`func (o *ActiveAlertsStats) GetNumHardwareWarningAlerts() int64`

GetNumHardwareWarningAlerts returns the NumHardwareWarningAlerts field if non-nil, zero value otherwise.

### GetNumHardwareWarningAlertsOk

`func (o *ActiveAlertsStats) GetNumHardwareWarningAlertsOk() (*int64, bool)`

GetNumHardwareWarningAlertsOk returns a tuple with the NumHardwareWarningAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumHardwareWarningAlerts

`func (o *ActiveAlertsStats) SetNumHardwareWarningAlerts(v int64)`

SetNumHardwareWarningAlerts sets NumHardwareWarningAlerts field to given value.

### HasNumHardwareWarningAlerts

`func (o *ActiveAlertsStats) HasNumHardwareWarningAlerts() bool`

HasNumHardwareWarningAlerts returns a boolean if a field has been set.

### SetNumHardwareWarningAlertsNil

`func (o *ActiveAlertsStats) SetNumHardwareWarningAlertsNil(b bool)`

 SetNumHardwareWarningAlertsNil sets the value for NumHardwareWarningAlerts to be an explicit nil

### UnsetNumHardwareWarningAlerts
`func (o *ActiveAlertsStats) UnsetNumHardwareWarningAlerts()`

UnsetNumHardwareWarningAlerts ensures that no value is present for NumHardwareWarningAlerts, not even an explicit nil
### GetNumInfoAlerts

`func (o *ActiveAlertsStats) GetNumInfoAlerts() int64`

GetNumInfoAlerts returns the NumInfoAlerts field if non-nil, zero value otherwise.

### GetNumInfoAlertsOk

`func (o *ActiveAlertsStats) GetNumInfoAlertsOk() (*int64, bool)`

GetNumInfoAlertsOk returns a tuple with the NumInfoAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInfoAlerts

`func (o *ActiveAlertsStats) SetNumInfoAlerts(v int64)`

SetNumInfoAlerts sets NumInfoAlerts field to given value.

### HasNumInfoAlerts

`func (o *ActiveAlertsStats) HasNumInfoAlerts() bool`

HasNumInfoAlerts returns a boolean if a field has been set.

### SetNumInfoAlertsNil

`func (o *ActiveAlertsStats) SetNumInfoAlertsNil(b bool)`

 SetNumInfoAlertsNil sets the value for NumInfoAlerts to be an explicit nil

### UnsetNumInfoAlerts
`func (o *ActiveAlertsStats) UnsetNumInfoAlerts()`

UnsetNumInfoAlerts ensures that no value is present for NumInfoAlerts, not even an explicit nil
### GetNumInfoAlertsCategories

`func (o *ActiveAlertsStats) GetNumInfoAlertsCategories() int64`

GetNumInfoAlertsCategories returns the NumInfoAlertsCategories field if non-nil, zero value otherwise.

### GetNumInfoAlertsCategoriesOk

`func (o *ActiveAlertsStats) GetNumInfoAlertsCategoriesOk() (*int64, bool)`

GetNumInfoAlertsCategoriesOk returns a tuple with the NumInfoAlertsCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInfoAlertsCategories

`func (o *ActiveAlertsStats) SetNumInfoAlertsCategories(v int64)`

SetNumInfoAlertsCategories sets NumInfoAlertsCategories field to given value.

### HasNumInfoAlertsCategories

`func (o *ActiveAlertsStats) HasNumInfoAlertsCategories() bool`

HasNumInfoAlertsCategories returns a boolean if a field has been set.

### SetNumInfoAlertsCategoriesNil

`func (o *ActiveAlertsStats) SetNumInfoAlertsCategoriesNil(b bool)`

 SetNumInfoAlertsCategoriesNil sets the value for NumInfoAlertsCategories to be an explicit nil

### UnsetNumInfoAlertsCategories
`func (o *ActiveAlertsStats) UnsetNumInfoAlertsCategories()`

UnsetNumInfoAlertsCategories ensures that no value is present for NumInfoAlertsCategories, not even an explicit nil
### GetNumOtherAlerts

`func (o *ActiveAlertsStats) GetNumOtherAlerts() int64`

GetNumOtherAlerts returns the NumOtherAlerts field if non-nil, zero value otherwise.

### GetNumOtherAlertsOk

`func (o *ActiveAlertsStats) GetNumOtherAlertsOk() (*int64, bool)`

GetNumOtherAlertsOk returns a tuple with the NumOtherAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOtherAlerts

`func (o *ActiveAlertsStats) SetNumOtherAlerts(v int64)`

SetNumOtherAlerts sets NumOtherAlerts field to given value.

### HasNumOtherAlerts

`func (o *ActiveAlertsStats) HasNumOtherAlerts() bool`

HasNumOtherAlerts returns a boolean if a field has been set.

### SetNumOtherAlertsNil

`func (o *ActiveAlertsStats) SetNumOtherAlertsNil(b bool)`

 SetNumOtherAlertsNil sets the value for NumOtherAlerts to be an explicit nil

### UnsetNumOtherAlerts
`func (o *ActiveAlertsStats) UnsetNumOtherAlerts()`

UnsetNumOtherAlerts ensures that no value is present for NumOtherAlerts, not even an explicit nil
### GetNumOtherCriticalAlerts

`func (o *ActiveAlertsStats) GetNumOtherCriticalAlerts() int64`

GetNumOtherCriticalAlerts returns the NumOtherCriticalAlerts field if non-nil, zero value otherwise.

### GetNumOtherCriticalAlertsOk

`func (o *ActiveAlertsStats) GetNumOtherCriticalAlertsOk() (*int64, bool)`

GetNumOtherCriticalAlertsOk returns a tuple with the NumOtherCriticalAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOtherCriticalAlerts

`func (o *ActiveAlertsStats) SetNumOtherCriticalAlerts(v int64)`

SetNumOtherCriticalAlerts sets NumOtherCriticalAlerts field to given value.

### HasNumOtherCriticalAlerts

`func (o *ActiveAlertsStats) HasNumOtherCriticalAlerts() bool`

HasNumOtherCriticalAlerts returns a boolean if a field has been set.

### SetNumOtherCriticalAlertsNil

`func (o *ActiveAlertsStats) SetNumOtherCriticalAlertsNil(b bool)`

 SetNumOtherCriticalAlertsNil sets the value for NumOtherCriticalAlerts to be an explicit nil

### UnsetNumOtherCriticalAlerts
`func (o *ActiveAlertsStats) UnsetNumOtherCriticalAlerts()`

UnsetNumOtherCriticalAlerts ensures that no value is present for NumOtherCriticalAlerts, not even an explicit nil
### GetNumOtherInfoAlerts

`func (o *ActiveAlertsStats) GetNumOtherInfoAlerts() int64`

GetNumOtherInfoAlerts returns the NumOtherInfoAlerts field if non-nil, zero value otherwise.

### GetNumOtherInfoAlertsOk

`func (o *ActiveAlertsStats) GetNumOtherInfoAlertsOk() (*int64, bool)`

GetNumOtherInfoAlertsOk returns a tuple with the NumOtherInfoAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOtherInfoAlerts

`func (o *ActiveAlertsStats) SetNumOtherInfoAlerts(v int64)`

SetNumOtherInfoAlerts sets NumOtherInfoAlerts field to given value.

### HasNumOtherInfoAlerts

`func (o *ActiveAlertsStats) HasNumOtherInfoAlerts() bool`

HasNumOtherInfoAlerts returns a boolean if a field has been set.

### SetNumOtherInfoAlertsNil

`func (o *ActiveAlertsStats) SetNumOtherInfoAlertsNil(b bool)`

 SetNumOtherInfoAlertsNil sets the value for NumOtherInfoAlerts to be an explicit nil

### UnsetNumOtherInfoAlerts
`func (o *ActiveAlertsStats) UnsetNumOtherInfoAlerts()`

UnsetNumOtherInfoAlerts ensures that no value is present for NumOtherInfoAlerts, not even an explicit nil
### GetNumOtherWarningAlerts

`func (o *ActiveAlertsStats) GetNumOtherWarningAlerts() int64`

GetNumOtherWarningAlerts returns the NumOtherWarningAlerts field if non-nil, zero value otherwise.

### GetNumOtherWarningAlertsOk

`func (o *ActiveAlertsStats) GetNumOtherWarningAlertsOk() (*int64, bool)`

GetNumOtherWarningAlertsOk returns a tuple with the NumOtherWarningAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOtherWarningAlerts

`func (o *ActiveAlertsStats) SetNumOtherWarningAlerts(v int64)`

SetNumOtherWarningAlerts sets NumOtherWarningAlerts field to given value.

### HasNumOtherWarningAlerts

`func (o *ActiveAlertsStats) HasNumOtherWarningAlerts() bool`

HasNumOtherWarningAlerts returns a boolean if a field has been set.

### SetNumOtherWarningAlertsNil

`func (o *ActiveAlertsStats) SetNumOtherWarningAlertsNil(b bool)`

 SetNumOtherWarningAlertsNil sets the value for NumOtherWarningAlerts to be an explicit nil

### UnsetNumOtherWarningAlerts
`func (o *ActiveAlertsStats) UnsetNumOtherWarningAlerts()`

UnsetNumOtherWarningAlerts ensures that no value is present for NumOtherWarningAlerts, not even an explicit nil
### GetNumServiceAlerts

`func (o *ActiveAlertsStats) GetNumServiceAlerts() int64`

GetNumServiceAlerts returns the NumServiceAlerts field if non-nil, zero value otherwise.

### GetNumServiceAlertsOk

`func (o *ActiveAlertsStats) GetNumServiceAlertsOk() (*int64, bool)`

GetNumServiceAlertsOk returns a tuple with the NumServiceAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumServiceAlerts

`func (o *ActiveAlertsStats) SetNumServiceAlerts(v int64)`

SetNumServiceAlerts sets NumServiceAlerts field to given value.

### HasNumServiceAlerts

`func (o *ActiveAlertsStats) HasNumServiceAlerts() bool`

HasNumServiceAlerts returns a boolean if a field has been set.

### SetNumServiceAlertsNil

`func (o *ActiveAlertsStats) SetNumServiceAlertsNil(b bool)`

 SetNumServiceAlertsNil sets the value for NumServiceAlerts to be an explicit nil

### UnsetNumServiceAlerts
`func (o *ActiveAlertsStats) UnsetNumServiceAlerts()`

UnsetNumServiceAlerts ensures that no value is present for NumServiceAlerts, not even an explicit nil
### GetNumServiceCriticalAlerts

`func (o *ActiveAlertsStats) GetNumServiceCriticalAlerts() int64`

GetNumServiceCriticalAlerts returns the NumServiceCriticalAlerts field if non-nil, zero value otherwise.

### GetNumServiceCriticalAlertsOk

`func (o *ActiveAlertsStats) GetNumServiceCriticalAlertsOk() (*int64, bool)`

GetNumServiceCriticalAlertsOk returns a tuple with the NumServiceCriticalAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumServiceCriticalAlerts

`func (o *ActiveAlertsStats) SetNumServiceCriticalAlerts(v int64)`

SetNumServiceCriticalAlerts sets NumServiceCriticalAlerts field to given value.

### HasNumServiceCriticalAlerts

`func (o *ActiveAlertsStats) HasNumServiceCriticalAlerts() bool`

HasNumServiceCriticalAlerts returns a boolean if a field has been set.

### SetNumServiceCriticalAlertsNil

`func (o *ActiveAlertsStats) SetNumServiceCriticalAlertsNil(b bool)`

 SetNumServiceCriticalAlertsNil sets the value for NumServiceCriticalAlerts to be an explicit nil

### UnsetNumServiceCriticalAlerts
`func (o *ActiveAlertsStats) UnsetNumServiceCriticalAlerts()`

UnsetNumServiceCriticalAlerts ensures that no value is present for NumServiceCriticalAlerts, not even an explicit nil
### GetNumServiceInfoAlerts

`func (o *ActiveAlertsStats) GetNumServiceInfoAlerts() int64`

GetNumServiceInfoAlerts returns the NumServiceInfoAlerts field if non-nil, zero value otherwise.

### GetNumServiceInfoAlertsOk

`func (o *ActiveAlertsStats) GetNumServiceInfoAlertsOk() (*int64, bool)`

GetNumServiceInfoAlertsOk returns a tuple with the NumServiceInfoAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumServiceInfoAlerts

`func (o *ActiveAlertsStats) SetNumServiceInfoAlerts(v int64)`

SetNumServiceInfoAlerts sets NumServiceInfoAlerts field to given value.

### HasNumServiceInfoAlerts

`func (o *ActiveAlertsStats) HasNumServiceInfoAlerts() bool`

HasNumServiceInfoAlerts returns a boolean if a field has been set.

### SetNumServiceInfoAlertsNil

`func (o *ActiveAlertsStats) SetNumServiceInfoAlertsNil(b bool)`

 SetNumServiceInfoAlertsNil sets the value for NumServiceInfoAlerts to be an explicit nil

### UnsetNumServiceInfoAlerts
`func (o *ActiveAlertsStats) UnsetNumServiceInfoAlerts()`

UnsetNumServiceInfoAlerts ensures that no value is present for NumServiceInfoAlerts, not even an explicit nil
### GetNumServiceWarningAlerts

`func (o *ActiveAlertsStats) GetNumServiceWarningAlerts() int64`

GetNumServiceWarningAlerts returns the NumServiceWarningAlerts field if non-nil, zero value otherwise.

### GetNumServiceWarningAlertsOk

`func (o *ActiveAlertsStats) GetNumServiceWarningAlertsOk() (*int64, bool)`

GetNumServiceWarningAlertsOk returns a tuple with the NumServiceWarningAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumServiceWarningAlerts

`func (o *ActiveAlertsStats) SetNumServiceWarningAlerts(v int64)`

SetNumServiceWarningAlerts sets NumServiceWarningAlerts field to given value.

### HasNumServiceWarningAlerts

`func (o *ActiveAlertsStats) HasNumServiceWarningAlerts() bool`

HasNumServiceWarningAlerts returns a boolean if a field has been set.

### SetNumServiceWarningAlertsNil

`func (o *ActiveAlertsStats) SetNumServiceWarningAlertsNil(b bool)`

 SetNumServiceWarningAlertsNil sets the value for NumServiceWarningAlerts to be an explicit nil

### UnsetNumServiceWarningAlerts
`func (o *ActiveAlertsStats) UnsetNumServiceWarningAlerts()`

UnsetNumServiceWarningAlerts ensures that no value is present for NumServiceWarningAlerts, not even an explicit nil
### GetNumSoftwareAlerts

`func (o *ActiveAlertsStats) GetNumSoftwareAlerts() int64`

GetNumSoftwareAlerts returns the NumSoftwareAlerts field if non-nil, zero value otherwise.

### GetNumSoftwareAlertsOk

`func (o *ActiveAlertsStats) GetNumSoftwareAlertsOk() (*int64, bool)`

GetNumSoftwareAlertsOk returns a tuple with the NumSoftwareAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSoftwareAlerts

`func (o *ActiveAlertsStats) SetNumSoftwareAlerts(v int64)`

SetNumSoftwareAlerts sets NumSoftwareAlerts field to given value.

### HasNumSoftwareAlerts

`func (o *ActiveAlertsStats) HasNumSoftwareAlerts() bool`

HasNumSoftwareAlerts returns a boolean if a field has been set.

### SetNumSoftwareAlertsNil

`func (o *ActiveAlertsStats) SetNumSoftwareAlertsNil(b bool)`

 SetNumSoftwareAlertsNil sets the value for NumSoftwareAlerts to be an explicit nil

### UnsetNumSoftwareAlerts
`func (o *ActiveAlertsStats) UnsetNumSoftwareAlerts()`

UnsetNumSoftwareAlerts ensures that no value is present for NumSoftwareAlerts, not even an explicit nil
### GetNumSoftwareCriticalAlerts

`func (o *ActiveAlertsStats) GetNumSoftwareCriticalAlerts() int64`

GetNumSoftwareCriticalAlerts returns the NumSoftwareCriticalAlerts field if non-nil, zero value otherwise.

### GetNumSoftwareCriticalAlertsOk

`func (o *ActiveAlertsStats) GetNumSoftwareCriticalAlertsOk() (*int64, bool)`

GetNumSoftwareCriticalAlertsOk returns a tuple with the NumSoftwareCriticalAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSoftwareCriticalAlerts

`func (o *ActiveAlertsStats) SetNumSoftwareCriticalAlerts(v int64)`

SetNumSoftwareCriticalAlerts sets NumSoftwareCriticalAlerts field to given value.

### HasNumSoftwareCriticalAlerts

`func (o *ActiveAlertsStats) HasNumSoftwareCriticalAlerts() bool`

HasNumSoftwareCriticalAlerts returns a boolean if a field has been set.

### SetNumSoftwareCriticalAlertsNil

`func (o *ActiveAlertsStats) SetNumSoftwareCriticalAlertsNil(b bool)`

 SetNumSoftwareCriticalAlertsNil sets the value for NumSoftwareCriticalAlerts to be an explicit nil

### UnsetNumSoftwareCriticalAlerts
`func (o *ActiveAlertsStats) UnsetNumSoftwareCriticalAlerts()`

UnsetNumSoftwareCriticalAlerts ensures that no value is present for NumSoftwareCriticalAlerts, not even an explicit nil
### GetNumSoftwareInfoAlerts

`func (o *ActiveAlertsStats) GetNumSoftwareInfoAlerts() int64`

GetNumSoftwareInfoAlerts returns the NumSoftwareInfoAlerts field if non-nil, zero value otherwise.

### GetNumSoftwareInfoAlertsOk

`func (o *ActiveAlertsStats) GetNumSoftwareInfoAlertsOk() (*int64, bool)`

GetNumSoftwareInfoAlertsOk returns a tuple with the NumSoftwareInfoAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSoftwareInfoAlerts

`func (o *ActiveAlertsStats) SetNumSoftwareInfoAlerts(v int64)`

SetNumSoftwareInfoAlerts sets NumSoftwareInfoAlerts field to given value.

### HasNumSoftwareInfoAlerts

`func (o *ActiveAlertsStats) HasNumSoftwareInfoAlerts() bool`

HasNumSoftwareInfoAlerts returns a boolean if a field has been set.

### SetNumSoftwareInfoAlertsNil

`func (o *ActiveAlertsStats) SetNumSoftwareInfoAlertsNil(b bool)`

 SetNumSoftwareInfoAlertsNil sets the value for NumSoftwareInfoAlerts to be an explicit nil

### UnsetNumSoftwareInfoAlerts
`func (o *ActiveAlertsStats) UnsetNumSoftwareInfoAlerts()`

UnsetNumSoftwareInfoAlerts ensures that no value is present for NumSoftwareInfoAlerts, not even an explicit nil
### GetNumSoftwareWarningAlerts

`func (o *ActiveAlertsStats) GetNumSoftwareWarningAlerts() int64`

GetNumSoftwareWarningAlerts returns the NumSoftwareWarningAlerts field if non-nil, zero value otherwise.

### GetNumSoftwareWarningAlertsOk

`func (o *ActiveAlertsStats) GetNumSoftwareWarningAlertsOk() (*int64, bool)`

GetNumSoftwareWarningAlertsOk returns a tuple with the NumSoftwareWarningAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSoftwareWarningAlerts

`func (o *ActiveAlertsStats) SetNumSoftwareWarningAlerts(v int64)`

SetNumSoftwareWarningAlerts sets NumSoftwareWarningAlerts field to given value.

### HasNumSoftwareWarningAlerts

`func (o *ActiveAlertsStats) HasNumSoftwareWarningAlerts() bool`

HasNumSoftwareWarningAlerts returns a boolean if a field has been set.

### SetNumSoftwareWarningAlertsNil

`func (o *ActiveAlertsStats) SetNumSoftwareWarningAlertsNil(b bool)`

 SetNumSoftwareWarningAlertsNil sets the value for NumSoftwareWarningAlerts to be an explicit nil

### UnsetNumSoftwareWarningAlerts
`func (o *ActiveAlertsStats) UnsetNumSoftwareWarningAlerts()`

UnsetNumSoftwareWarningAlerts ensures that no value is present for NumSoftwareWarningAlerts, not even an explicit nil
### GetNumWarningAlerts

`func (o *ActiveAlertsStats) GetNumWarningAlerts() int64`

GetNumWarningAlerts returns the NumWarningAlerts field if non-nil, zero value otherwise.

### GetNumWarningAlertsOk

`func (o *ActiveAlertsStats) GetNumWarningAlertsOk() (*int64, bool)`

GetNumWarningAlertsOk returns a tuple with the NumWarningAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWarningAlerts

`func (o *ActiveAlertsStats) SetNumWarningAlerts(v int64)`

SetNumWarningAlerts sets NumWarningAlerts field to given value.

### HasNumWarningAlerts

`func (o *ActiveAlertsStats) HasNumWarningAlerts() bool`

HasNumWarningAlerts returns a boolean if a field has been set.

### SetNumWarningAlertsNil

`func (o *ActiveAlertsStats) SetNumWarningAlertsNil(b bool)`

 SetNumWarningAlertsNil sets the value for NumWarningAlerts to be an explicit nil

### UnsetNumWarningAlerts
`func (o *ActiveAlertsStats) UnsetNumWarningAlerts()`

UnsetNumWarningAlerts ensures that no value is present for NumWarningAlerts, not even an explicit nil
### GetNumWarningAlertsCategories

`func (o *ActiveAlertsStats) GetNumWarningAlertsCategories() int64`

GetNumWarningAlertsCategories returns the NumWarningAlertsCategories field if non-nil, zero value otherwise.

### GetNumWarningAlertsCategoriesOk

`func (o *ActiveAlertsStats) GetNumWarningAlertsCategoriesOk() (*int64, bool)`

GetNumWarningAlertsCategoriesOk returns a tuple with the NumWarningAlertsCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWarningAlertsCategories

`func (o *ActiveAlertsStats) SetNumWarningAlertsCategories(v int64)`

SetNumWarningAlertsCategories sets NumWarningAlertsCategories field to given value.

### HasNumWarningAlertsCategories

`func (o *ActiveAlertsStats) HasNumWarningAlertsCategories() bool`

HasNumWarningAlertsCategories returns a boolean if a field has been set.

### SetNumWarningAlertsCategoriesNil

`func (o *ActiveAlertsStats) SetNumWarningAlertsCategoriesNil(b bool)`

 SetNumWarningAlertsCategoriesNil sets the value for NumWarningAlertsCategories to be an explicit nil

### UnsetNumWarningAlertsCategories
`func (o *ActiveAlertsStats) UnsetNumWarningAlertsCategories()`

UnsetNumWarningAlertsCategories ensures that no value is present for NumWarningAlertsCategories, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


