# CentrifyZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **NullableString** | Specifies a description of the Centrify zone. | 
**DistinguishedName** | **NullableString** | Specifies the distinguished name of the Centrify zone. | 
**Schema** | **NullableString** | Specifies the schema of this Centrify zone. | 
**ZoneName** | Pointer to **NullableString** | Specifies the zone name of the Centrify zone. | [optional] [readonly] 
**ZoneDomain** | Pointer to **NullableString** | Specifies the zone domain of the Centrify zone. | [optional] [readonly] 

## Methods

### NewCentrifyZone

`func NewCentrifyZone(description NullableString, distinguishedName NullableString, schema NullableString, ) *CentrifyZone`

NewCentrifyZone instantiates a new CentrifyZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCentrifyZoneWithDefaults

`func NewCentrifyZoneWithDefaults() *CentrifyZone`

NewCentrifyZoneWithDefaults instantiates a new CentrifyZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CentrifyZone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CentrifyZone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CentrifyZone) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *CentrifyZone) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CentrifyZone) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDistinguishedName

`func (o *CentrifyZone) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *CentrifyZone) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *CentrifyZone) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.


### SetDistinguishedNameNil

`func (o *CentrifyZone) SetDistinguishedNameNil(b bool)`

 SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil

### UnsetDistinguishedName
`func (o *CentrifyZone) UnsetDistinguishedName()`

UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil
### GetSchema

`func (o *CentrifyZone) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *CentrifyZone) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *CentrifyZone) SetSchema(v string)`

SetSchema sets Schema field to given value.


### SetSchemaNil

`func (o *CentrifyZone) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *CentrifyZone) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetZoneName

`func (o *CentrifyZone) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *CentrifyZone) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *CentrifyZone) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *CentrifyZone) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### SetZoneNameNil

`func (o *CentrifyZone) SetZoneNameNil(b bool)`

 SetZoneNameNil sets the value for ZoneName to be an explicit nil

### UnsetZoneName
`func (o *CentrifyZone) UnsetZoneName()`

UnsetZoneName ensures that no value is present for ZoneName, not even an explicit nil
### GetZoneDomain

`func (o *CentrifyZone) GetZoneDomain() string`

GetZoneDomain returns the ZoneDomain field if non-nil, zero value otherwise.

### GetZoneDomainOk

`func (o *CentrifyZone) GetZoneDomainOk() (*string, bool)`

GetZoneDomainOk returns a tuple with the ZoneDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneDomain

`func (o *CentrifyZone) SetZoneDomain(v string)`

SetZoneDomain sets ZoneDomain field to given value.

### HasZoneDomain

`func (o *CentrifyZone) HasZoneDomain() bool`

HasZoneDomain returns a boolean if a field has been set.

### SetZoneDomainNil

`func (o *CentrifyZone) SetZoneDomainNil(b bool)`

 SetZoneDomainNil sets the value for ZoneDomain to be an explicit nil

### UnsetZoneDomain
`func (o *CentrifyZone) UnsetZoneDomain()`

UnsetZoneDomain ensures that no value is present for ZoneDomain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


