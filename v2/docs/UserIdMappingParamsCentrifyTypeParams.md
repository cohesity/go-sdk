# UserIdMappingParamsCentrifyTypeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **NullableString** | Specifies a description of the Centrify zone. | 
**DistinguishedName** | **NullableString** | Specifies the distinguished name of the Centrify zone. | 
**Schema** | **NullableString** | Specifies the schema of this Centrify zone. | 
**ZoneDomain** | Pointer to **NullableString** | Specifies the zone domain of the Centrify zone. | [optional] [readonly] 
**ZoneName** | Pointer to **NullableString** | Specifies the zone name of the Centrify zone. | [optional] [readonly] 
**FallbackOption** | **map[string]interface{}** | Specifies a fallback user id mapping param in case the primary config does not work. | 

## Methods

### NewUserIdMappingParamsCentrifyTypeParams

`func NewUserIdMappingParamsCentrifyTypeParams(description NullableString, distinguishedName NullableString, schema NullableString, fallbackOption map[string]interface{}, ) *UserIdMappingParamsCentrifyTypeParams`

NewUserIdMappingParamsCentrifyTypeParams instantiates a new UserIdMappingParamsCentrifyTypeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdMappingParamsCentrifyTypeParamsWithDefaults

`func NewUserIdMappingParamsCentrifyTypeParamsWithDefaults() *UserIdMappingParamsCentrifyTypeParams`

NewUserIdMappingParamsCentrifyTypeParamsWithDefaults instantiates a new UserIdMappingParamsCentrifyTypeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UserIdMappingParamsCentrifyTypeParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserIdMappingParamsCentrifyTypeParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserIdMappingParamsCentrifyTypeParams) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *UserIdMappingParamsCentrifyTypeParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UserIdMappingParamsCentrifyTypeParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDistinguishedName

`func (o *UserIdMappingParamsCentrifyTypeParams) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *UserIdMappingParamsCentrifyTypeParams) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *UserIdMappingParamsCentrifyTypeParams) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.


### SetDistinguishedNameNil

`func (o *UserIdMappingParamsCentrifyTypeParams) SetDistinguishedNameNil(b bool)`

 SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil

### UnsetDistinguishedName
`func (o *UserIdMappingParamsCentrifyTypeParams) UnsetDistinguishedName()`

UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil
### GetSchema

`func (o *UserIdMappingParamsCentrifyTypeParams) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *UserIdMappingParamsCentrifyTypeParams) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *UserIdMappingParamsCentrifyTypeParams) SetSchema(v string)`

SetSchema sets Schema field to given value.


### SetSchemaNil

`func (o *UserIdMappingParamsCentrifyTypeParams) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *UserIdMappingParamsCentrifyTypeParams) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetZoneDomain

`func (o *UserIdMappingParamsCentrifyTypeParams) GetZoneDomain() string`

GetZoneDomain returns the ZoneDomain field if non-nil, zero value otherwise.

### GetZoneDomainOk

`func (o *UserIdMappingParamsCentrifyTypeParams) GetZoneDomainOk() (*string, bool)`

GetZoneDomainOk returns a tuple with the ZoneDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneDomain

`func (o *UserIdMappingParamsCentrifyTypeParams) SetZoneDomain(v string)`

SetZoneDomain sets ZoneDomain field to given value.

### HasZoneDomain

`func (o *UserIdMappingParamsCentrifyTypeParams) HasZoneDomain() bool`

HasZoneDomain returns a boolean if a field has been set.

### SetZoneDomainNil

`func (o *UserIdMappingParamsCentrifyTypeParams) SetZoneDomainNil(b bool)`

 SetZoneDomainNil sets the value for ZoneDomain to be an explicit nil

### UnsetZoneDomain
`func (o *UserIdMappingParamsCentrifyTypeParams) UnsetZoneDomain()`

UnsetZoneDomain ensures that no value is present for ZoneDomain, not even an explicit nil
### GetZoneName

`func (o *UserIdMappingParamsCentrifyTypeParams) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *UserIdMappingParamsCentrifyTypeParams) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *UserIdMappingParamsCentrifyTypeParams) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *UserIdMappingParamsCentrifyTypeParams) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### SetZoneNameNil

`func (o *UserIdMappingParamsCentrifyTypeParams) SetZoneNameNil(b bool)`

 SetZoneNameNil sets the value for ZoneName to be an explicit nil

### UnsetZoneName
`func (o *UserIdMappingParamsCentrifyTypeParams) UnsetZoneName()`

UnsetZoneName ensures that no value is present for ZoneName, not even an explicit nil
### GetFallbackOption

`func (o *UserIdMappingParamsCentrifyTypeParams) GetFallbackOption() map[string]interface{}`

GetFallbackOption returns the FallbackOption field if non-nil, zero value otherwise.

### GetFallbackOptionOk

`func (o *UserIdMappingParamsCentrifyTypeParams) GetFallbackOptionOk() (*map[string]interface{}, bool)`

GetFallbackOptionOk returns a tuple with the FallbackOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackOption

`func (o *UserIdMappingParamsCentrifyTypeParams) SetFallbackOption(v map[string]interface{})`

SetFallbackOption sets FallbackOption field to given value.


### SetFallbackOptionNil

`func (o *UserIdMappingParamsCentrifyTypeParams) SetFallbackOptionNil(b bool)`

 SetFallbackOptionNil sets the value for FallbackOption to be an explicit nil

### UnsetFallbackOption
`func (o *UserIdMappingParamsCentrifyTypeParams) UnsetFallbackOption()`

UnsetFallbackOption ensures that no value is present for FallbackOption, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


