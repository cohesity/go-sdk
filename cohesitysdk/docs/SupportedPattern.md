# SupportedPattern

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSystemDefined** | Pointer to **NullableBool** | Specifies whether the pattern has been defined by the system or the user. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Pattern. | [optional] 
**Pattern** | Pointer to **NullableString** | Specifies the value of the pattern(Regex). | [optional] 
**PatternType** | Pointer to **NullableString** | Specifies the Pattern type. &#39;REGULAR&#39; indicates that the pattern contains a regular expression. &#39;TEMPLATE&#39; indicates that the pattern has a pre defined input pattern such as date of the form &#39;DD-MM-YYYY&#39;. | [optional] 

## Methods

### NewSupportedPattern

`func NewSupportedPattern() *SupportedPattern`

NewSupportedPattern instantiates a new SupportedPattern object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedPatternWithDefaults

`func NewSupportedPatternWithDefaults() *SupportedPattern`

NewSupportedPatternWithDefaults instantiates a new SupportedPattern object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSystemDefined

`func (o *SupportedPattern) GetIsSystemDefined() bool`

GetIsSystemDefined returns the IsSystemDefined field if non-nil, zero value otherwise.

### GetIsSystemDefinedOk

`func (o *SupportedPattern) GetIsSystemDefinedOk() (*bool, bool)`

GetIsSystemDefinedOk returns a tuple with the IsSystemDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemDefined

`func (o *SupportedPattern) SetIsSystemDefined(v bool)`

SetIsSystemDefined sets IsSystemDefined field to given value.

### HasIsSystemDefined

`func (o *SupportedPattern) HasIsSystemDefined() bool`

HasIsSystemDefined returns a boolean if a field has been set.

### SetIsSystemDefinedNil

`func (o *SupportedPattern) SetIsSystemDefinedNil(b bool)`

 SetIsSystemDefinedNil sets the value for IsSystemDefined to be an explicit nil

### UnsetIsSystemDefined
`func (o *SupportedPattern) UnsetIsSystemDefined()`

UnsetIsSystemDefined ensures that no value is present for IsSystemDefined, not even an explicit nil
### GetName

`func (o *SupportedPattern) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SupportedPattern) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SupportedPattern) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SupportedPattern) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SupportedPattern) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SupportedPattern) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPattern

`func (o *SupportedPattern) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *SupportedPattern) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *SupportedPattern) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *SupportedPattern) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### SetPatternNil

`func (o *SupportedPattern) SetPatternNil(b bool)`

 SetPatternNil sets the value for Pattern to be an explicit nil

### UnsetPattern
`func (o *SupportedPattern) UnsetPattern()`

UnsetPattern ensures that no value is present for Pattern, not even an explicit nil
### GetPatternType

`func (o *SupportedPattern) GetPatternType() string`

GetPatternType returns the PatternType field if non-nil, zero value otherwise.

### GetPatternTypeOk

`func (o *SupportedPattern) GetPatternTypeOk() (*string, bool)`

GetPatternTypeOk returns a tuple with the PatternType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternType

`func (o *SupportedPattern) SetPatternType(v string)`

SetPatternType sets PatternType field to given value.

### HasPatternType

`func (o *SupportedPattern) HasPatternType() bool`

HasPatternType returns a boolean if a field has been set.

### SetPatternTypeNil

`func (o *SupportedPattern) SetPatternTypeNil(b bool)`

 SetPatternTypeNil sets the value for PatternType to be an explicit nil

### UnsetPatternType
`func (o *SupportedPattern) UnsetPatternType()`

UnsetPatternType ensures that no value is present for PatternType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


