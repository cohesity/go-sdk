# AdObjectAttributeParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdGuidPairs** | Pointer to [**[]AdGuidPair1**](AdGuidPair1.md) | Specifies the array of source and destination object guid pairs to restore attributes. | [optional] 
**ExcludeLdapProperties** | Pointer to **[]string** | Specifies the array of LDAP property names to excluded from &#39;property_vec&#39;. Excluded property name cannot contain wildcard character &#39;*&#39;.  Property names are case insensitive. | [optional] 
**LdapProperties** | Pointer to **[]string** | Specifies the array of LDAP property(attribute) names. The name can be standard or custom property defined in AD schema partition. The property can contain wildcard character &#39;*&#39;. If this array is empty, then &#39;*&#39; is assigned, means restore all properties except default system excluded properties. Wildcards will be expanded. If &#39;memberOf&#39; property is included, group membership of the objects specified in &#39;guid_vec&#39; will be restored. Property that does not exist for an object is ignored and no error info is returned for that property. Property names are case insensitive. | [optional] 
**MergeMultiValProperties** | Pointer to **NullableBool** | Specifies the Option to merge multi-valued values vs clearing and setting values from the AD snapshot. Defaults is to overwrite production AD values from the source AD snapshot. When updating group membership (using &#39;memberOf&#39; property), setting this option to true will result in group membership merge. This is useful in cases where production AD has later updates that should not be overridden while restoring an attribute. | [optional] 

## Methods

### NewAdObjectAttributeParameters

`func NewAdObjectAttributeParameters() *AdObjectAttributeParameters`

NewAdObjectAttributeParameters instantiates a new AdObjectAttributeParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdObjectAttributeParametersWithDefaults

`func NewAdObjectAttributeParametersWithDefaults() *AdObjectAttributeParameters`

NewAdObjectAttributeParametersWithDefaults instantiates a new AdObjectAttributeParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdGuidPairs

`func (o *AdObjectAttributeParameters) GetAdGuidPairs() []AdGuidPair1`

GetAdGuidPairs returns the AdGuidPairs field if non-nil, zero value otherwise.

### GetAdGuidPairsOk

`func (o *AdObjectAttributeParameters) GetAdGuidPairsOk() (*[]AdGuidPair1, bool)`

GetAdGuidPairsOk returns a tuple with the AdGuidPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdGuidPairs

`func (o *AdObjectAttributeParameters) SetAdGuidPairs(v []AdGuidPair1)`

SetAdGuidPairs sets AdGuidPairs field to given value.

### HasAdGuidPairs

`func (o *AdObjectAttributeParameters) HasAdGuidPairs() bool`

HasAdGuidPairs returns a boolean if a field has been set.

### SetAdGuidPairsNil

`func (o *AdObjectAttributeParameters) SetAdGuidPairsNil(b bool)`

 SetAdGuidPairsNil sets the value for AdGuidPairs to be an explicit nil

### UnsetAdGuidPairs
`func (o *AdObjectAttributeParameters) UnsetAdGuidPairs()`

UnsetAdGuidPairs ensures that no value is present for AdGuidPairs, not even an explicit nil
### GetExcludeLdapProperties

`func (o *AdObjectAttributeParameters) GetExcludeLdapProperties() []string`

GetExcludeLdapProperties returns the ExcludeLdapProperties field if non-nil, zero value otherwise.

### GetExcludeLdapPropertiesOk

`func (o *AdObjectAttributeParameters) GetExcludeLdapPropertiesOk() (*[]string, bool)`

GetExcludeLdapPropertiesOk returns a tuple with the ExcludeLdapProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLdapProperties

`func (o *AdObjectAttributeParameters) SetExcludeLdapProperties(v []string)`

SetExcludeLdapProperties sets ExcludeLdapProperties field to given value.

### HasExcludeLdapProperties

`func (o *AdObjectAttributeParameters) HasExcludeLdapProperties() bool`

HasExcludeLdapProperties returns a boolean if a field has been set.

### SetExcludeLdapPropertiesNil

`func (o *AdObjectAttributeParameters) SetExcludeLdapPropertiesNil(b bool)`

 SetExcludeLdapPropertiesNil sets the value for ExcludeLdapProperties to be an explicit nil

### UnsetExcludeLdapProperties
`func (o *AdObjectAttributeParameters) UnsetExcludeLdapProperties()`

UnsetExcludeLdapProperties ensures that no value is present for ExcludeLdapProperties, not even an explicit nil
### GetLdapProperties

`func (o *AdObjectAttributeParameters) GetLdapProperties() []string`

GetLdapProperties returns the LdapProperties field if non-nil, zero value otherwise.

### GetLdapPropertiesOk

`func (o *AdObjectAttributeParameters) GetLdapPropertiesOk() (*[]string, bool)`

GetLdapPropertiesOk returns a tuple with the LdapProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProperties

`func (o *AdObjectAttributeParameters) SetLdapProperties(v []string)`

SetLdapProperties sets LdapProperties field to given value.

### HasLdapProperties

`func (o *AdObjectAttributeParameters) HasLdapProperties() bool`

HasLdapProperties returns a boolean if a field has been set.

### SetLdapPropertiesNil

`func (o *AdObjectAttributeParameters) SetLdapPropertiesNil(b bool)`

 SetLdapPropertiesNil sets the value for LdapProperties to be an explicit nil

### UnsetLdapProperties
`func (o *AdObjectAttributeParameters) UnsetLdapProperties()`

UnsetLdapProperties ensures that no value is present for LdapProperties, not even an explicit nil
### GetMergeMultiValProperties

`func (o *AdObjectAttributeParameters) GetMergeMultiValProperties() bool`

GetMergeMultiValProperties returns the MergeMultiValProperties field if non-nil, zero value otherwise.

### GetMergeMultiValPropertiesOk

`func (o *AdObjectAttributeParameters) GetMergeMultiValPropertiesOk() (*bool, bool)`

GetMergeMultiValPropertiesOk returns a tuple with the MergeMultiValProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeMultiValProperties

`func (o *AdObjectAttributeParameters) SetMergeMultiValProperties(v bool)`

SetMergeMultiValProperties sets MergeMultiValProperties field to given value.

### HasMergeMultiValProperties

`func (o *AdObjectAttributeParameters) HasMergeMultiValProperties() bool`

HasMergeMultiValProperties returns a boolean if a field has been set.

### SetMergeMultiValPropertiesNil

`func (o *AdObjectAttributeParameters) SetMergeMultiValPropertiesNil(b bool)`

 SetMergeMultiValPropertiesNil sets the value for MergeMultiValProperties to be an explicit nil

### UnsetMergeMultiValProperties
`func (o *AdObjectAttributeParameters) UnsetMergeMultiValProperties()`

UnsetMergeMultiValProperties ensures that no value is present for MergeMultiValProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


