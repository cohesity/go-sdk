# FilerLifecycleRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgingPolicy** | Pointer to [**FilerLifecycleRuleAgingPolicy**](FilerLifecycleRuleAgingPolicy.md) |  | [optional] 
**FileFilter** | Pointer to [**FilerLifecycleRuleFileFilter**](FilerLifecycleRuleFileFilter.md) |  | [optional] 
**Name** | **NullableString** | Specifies the Unique identifier for the rule. No 2 rules in a policy should have the same name. The value cannot be longer than 255 characters. | 
**Status** | **NullableString** | Specifies if the rule is currently being applied. | 
**Type** | **NullableString** | Specifies if the rule is Allow or Deny type. | 

## Methods

### NewFilerLifecycleRule

`func NewFilerLifecycleRule(name NullableString, status NullableString, type_ NullableString, ) *FilerLifecycleRule`

NewFilerLifecycleRule instantiates a new FilerLifecycleRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilerLifecycleRuleWithDefaults

`func NewFilerLifecycleRuleWithDefaults() *FilerLifecycleRule`

NewFilerLifecycleRuleWithDefaults instantiates a new FilerLifecycleRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgingPolicy

`func (o *FilerLifecycleRule) GetAgingPolicy() FilerLifecycleRuleAgingPolicy`

GetAgingPolicy returns the AgingPolicy field if non-nil, zero value otherwise.

### GetAgingPolicyOk

`func (o *FilerLifecycleRule) GetAgingPolicyOk() (*FilerLifecycleRuleAgingPolicy, bool)`

GetAgingPolicyOk returns a tuple with the AgingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgingPolicy

`func (o *FilerLifecycleRule) SetAgingPolicy(v FilerLifecycleRuleAgingPolicy)`

SetAgingPolicy sets AgingPolicy field to given value.

### HasAgingPolicy

`func (o *FilerLifecycleRule) HasAgingPolicy() bool`

HasAgingPolicy returns a boolean if a field has been set.

### GetFileFilter

`func (o *FilerLifecycleRule) GetFileFilter() FilerLifecycleRuleFileFilter`

GetFileFilter returns the FileFilter field if non-nil, zero value otherwise.

### GetFileFilterOk

`func (o *FilerLifecycleRule) GetFileFilterOk() (*FilerLifecycleRuleFileFilter, bool)`

GetFileFilterOk returns a tuple with the FileFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFilter

`func (o *FilerLifecycleRule) SetFileFilter(v FilerLifecycleRuleFileFilter)`

SetFileFilter sets FileFilter field to given value.

### HasFileFilter

`func (o *FilerLifecycleRule) HasFileFilter() bool`

HasFileFilter returns a boolean if a field has been set.

### GetName

`func (o *FilerLifecycleRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilerLifecycleRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilerLifecycleRule) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *FilerLifecycleRule) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FilerLifecycleRule) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *FilerLifecycleRule) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FilerLifecycleRule) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FilerLifecycleRule) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *FilerLifecycleRule) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *FilerLifecycleRule) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetType

`func (o *FilerLifecycleRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FilerLifecycleRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FilerLifecycleRule) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *FilerLifecycleRule) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FilerLifecycleRule) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


