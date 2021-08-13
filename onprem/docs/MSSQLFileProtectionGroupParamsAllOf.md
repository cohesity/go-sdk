# MSSQLFileProtectionGroupParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]MSSQLFileProtectionGroupObjectParams**](MSSQLFileProtectionGroupObjectParams.md) | Specifies the list of object params to be protected. | 
**PerformSourceSideDeduplication** | Pointer to **NullableBool** | Specifies whether or not to perform source side deduplication on this Protection Group. | [optional] 
**AdditionalHostParams** | Pointer to [**[]MSSQLFileProtectionGroupHostParams**](MSSQLFileProtectionGroupHostParams.md) | Specifies settings which are to be applied to specific host containers in this protection group. | [optional] 

## Methods

### NewMSSQLFileProtectionGroupParamsAllOf

`func NewMSSQLFileProtectionGroupParamsAllOf(objects []MSSQLFileProtectionGroupObjectParams, ) *MSSQLFileProtectionGroupParamsAllOf`

NewMSSQLFileProtectionGroupParamsAllOf instantiates a new MSSQLFileProtectionGroupParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMSSQLFileProtectionGroupParamsAllOfWithDefaults

`func NewMSSQLFileProtectionGroupParamsAllOfWithDefaults() *MSSQLFileProtectionGroupParamsAllOf`

NewMSSQLFileProtectionGroupParamsAllOfWithDefaults instantiates a new MSSQLFileProtectionGroupParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *MSSQLFileProtectionGroupParamsAllOf) GetObjects() []MSSQLFileProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *MSSQLFileProtectionGroupParamsAllOf) GetObjectsOk() (*[]MSSQLFileProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *MSSQLFileProtectionGroupParamsAllOf) SetObjects(v []MSSQLFileProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *MSSQLFileProtectionGroupParamsAllOf) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *MSSQLFileProtectionGroupParamsAllOf) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetPerformSourceSideDeduplication

`func (o *MSSQLFileProtectionGroupParamsAllOf) GetPerformSourceSideDeduplication() bool`

GetPerformSourceSideDeduplication returns the PerformSourceSideDeduplication field if non-nil, zero value otherwise.

### GetPerformSourceSideDeduplicationOk

`func (o *MSSQLFileProtectionGroupParamsAllOf) GetPerformSourceSideDeduplicationOk() (*bool, bool)`

GetPerformSourceSideDeduplicationOk returns a tuple with the PerformSourceSideDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformSourceSideDeduplication

`func (o *MSSQLFileProtectionGroupParamsAllOf) SetPerformSourceSideDeduplication(v bool)`

SetPerformSourceSideDeduplication sets PerformSourceSideDeduplication field to given value.

### HasPerformSourceSideDeduplication

`func (o *MSSQLFileProtectionGroupParamsAllOf) HasPerformSourceSideDeduplication() bool`

HasPerformSourceSideDeduplication returns a boolean if a field has been set.

### SetPerformSourceSideDeduplicationNil

`func (o *MSSQLFileProtectionGroupParamsAllOf) SetPerformSourceSideDeduplicationNil(b bool)`

 SetPerformSourceSideDeduplicationNil sets the value for PerformSourceSideDeduplication to be an explicit nil

### UnsetPerformSourceSideDeduplication
`func (o *MSSQLFileProtectionGroupParamsAllOf) UnsetPerformSourceSideDeduplication()`

UnsetPerformSourceSideDeduplication ensures that no value is present for PerformSourceSideDeduplication, not even an explicit nil
### GetAdditionalHostParams

`func (o *MSSQLFileProtectionGroupParamsAllOf) GetAdditionalHostParams() []MSSQLFileProtectionGroupHostParams`

GetAdditionalHostParams returns the AdditionalHostParams field if non-nil, zero value otherwise.

### GetAdditionalHostParamsOk

`func (o *MSSQLFileProtectionGroupParamsAllOf) GetAdditionalHostParamsOk() (*[]MSSQLFileProtectionGroupHostParams, bool)`

GetAdditionalHostParamsOk returns a tuple with the AdditionalHostParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalHostParams

`func (o *MSSQLFileProtectionGroupParamsAllOf) SetAdditionalHostParams(v []MSSQLFileProtectionGroupHostParams)`

SetAdditionalHostParams sets AdditionalHostParams field to given value.

### HasAdditionalHostParams

`func (o *MSSQLFileProtectionGroupParamsAllOf) HasAdditionalHostParams() bool`

HasAdditionalHostParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


