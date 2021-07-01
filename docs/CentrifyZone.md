# CentrifyZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CentrifySchema** | Pointer to **NullableString** | Specifies the schema of this Centrify zone. The below list of schemas and their values are taken from the document Centrify Server Suite 2016 Windows API Programmer&#39;s Guide https://docs.centrify.com/en/css/suite2016/centrify-win-progguide.pdf &#39;kCentrifyDynamicSchema_1_0&#39; specifies dynamic schema, version 1.0. &#39;kCentrifyDynamicSchema_2_0&#39; specifies dynamic schema, version 2.0. &#39;kCentrifyDynamicSchema_3_0&#39; specifies dynamic schema, version 3.0. &#39;kCentrifyDynamicSchema_5_0&#39; specifies dynamic schema, version 5.0. &#39;kCentrifySfu_3_0&#39; specifies sfu schema, version 3.0. &#39;kCentrifySfu_3_0_V5&#39; specifies sfu schema, 3.0.5. &#39;kCentrifySfu_4_0&#39; specifies sfu schema, version 4.0. &#39;kCentrifyCdcRfc2307&#39; specifies cdcrfc2307 schema. &#39;kCentrifyCdcRfc2307_2&#39; specifies cdcrfc2307, version 2. &#39;kCentrifyCdcRfc2307_3&#39; specifies cdcrfc2307, version 3. | [optional] 
**Description** | Pointer to **NullableString** | Specifies a description of the Centrify zone. | [optional] 
**DistinguishedName** | Pointer to **NullableString** | Specifies the distinguished name of the Centrify zone. | [optional] 

## Methods

### NewCentrifyZone

`func NewCentrifyZone() *CentrifyZone`

NewCentrifyZone instantiates a new CentrifyZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCentrifyZoneWithDefaults

`func NewCentrifyZoneWithDefaults() *CentrifyZone`

NewCentrifyZoneWithDefaults instantiates a new CentrifyZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCentrifySchema

`func (o *CentrifyZone) GetCentrifySchema() string`

GetCentrifySchema returns the CentrifySchema field if non-nil, zero value otherwise.

### GetCentrifySchemaOk

`func (o *CentrifyZone) GetCentrifySchemaOk() (*string, bool)`

GetCentrifySchemaOk returns a tuple with the CentrifySchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentrifySchema

`func (o *CentrifyZone) SetCentrifySchema(v string)`

SetCentrifySchema sets CentrifySchema field to given value.

### HasCentrifySchema

`func (o *CentrifyZone) HasCentrifySchema() bool`

HasCentrifySchema returns a boolean if a field has been set.

### SetCentrifySchemaNil

`func (o *CentrifyZone) SetCentrifySchemaNil(b bool)`

 SetCentrifySchemaNil sets the value for CentrifySchema to be an explicit nil

### UnsetCentrifySchema
`func (o *CentrifyZone) UnsetCentrifySchema()`

UnsetCentrifySchema ensures that no value is present for CentrifySchema, not even an explicit nil
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

### HasDescription

`func (o *CentrifyZone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

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

### HasDistinguishedName

`func (o *CentrifyZone) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### SetDistinguishedNameNil

`func (o *CentrifyZone) SetDistinguishedNameNil(b bool)`

 SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil

### UnsetDistinguishedName
`func (o *CentrifyZone) UnsetDistinguishedName()`

UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


