# ListCentrifyZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CentrifySchema** | Pointer to **NullableString** | Specifies the schema of this Centrify zone. The below list of schemas and their values are taken from the document Centrify Server Suite 2016 Windows API Programmer&#39;s Guide https://docs.centrify.com/en/css/suite2016/centrify-win-progguide.pdf &#39;kCentrifyDynamicSchema_1_0&#39; specifies dynamic schema, version 1.0. &#39;kCentrifyDynamicSchema_2_0&#39; specifies dynamic schema, version 2.0. &#39;kCentrifyDynamicSchema_3_0&#39; specifies dynamic schema, version 3.0. &#39;kCentrifyDynamicSchema_5_0&#39; specifies dynamic schema, version 5.0. &#39;kCentrifySfu_3_0&#39; specifies sfu schema, version 3.0. &#39;kCentrifySfu_3_0_V5&#39; specifies sfu schema, 3.0.5. &#39;kCentrifySfu_4_0&#39; specifies sfu schema, version 4.0. &#39;kCentrifyCdcRfc2307&#39; specifies cdcrfc2307 schema. &#39;kCentrifyCdcRfc2307_2&#39; specifies cdcrfc2307, version 2. &#39;kCentrifyCdcRfc2307_3&#39; specifies cdcrfc2307, version 3. | [optional] 
**Description** | Pointer to **NullableString** | Specifies a description of the Centrify zone. | [optional] 
**DistinguishedName** | Pointer to **NullableString** | Specifies the distinguished name of the Centrify zone. | [optional] 
**ZoneName** | Pointer to **NullableString** | Specifies the zone name. | [optional] 

## Methods

### NewListCentrifyZone

`func NewListCentrifyZone() *ListCentrifyZone`

NewListCentrifyZone instantiates a new ListCentrifyZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCentrifyZoneWithDefaults

`func NewListCentrifyZoneWithDefaults() *ListCentrifyZone`

NewListCentrifyZoneWithDefaults instantiates a new ListCentrifyZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCentrifySchema

`func (o *ListCentrifyZone) GetCentrifySchema() string`

GetCentrifySchema returns the CentrifySchema field if non-nil, zero value otherwise.

### GetCentrifySchemaOk

`func (o *ListCentrifyZone) GetCentrifySchemaOk() (*string, bool)`

GetCentrifySchemaOk returns a tuple with the CentrifySchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentrifySchema

`func (o *ListCentrifyZone) SetCentrifySchema(v string)`

SetCentrifySchema sets CentrifySchema field to given value.

### HasCentrifySchema

`func (o *ListCentrifyZone) HasCentrifySchema() bool`

HasCentrifySchema returns a boolean if a field has been set.

### SetCentrifySchemaNil

`func (o *ListCentrifyZone) SetCentrifySchemaNil(b bool)`

 SetCentrifySchemaNil sets the value for CentrifySchema to be an explicit nil

### UnsetCentrifySchema
`func (o *ListCentrifyZone) UnsetCentrifySchema()`

UnsetCentrifySchema ensures that no value is present for CentrifySchema, not even an explicit nil
### GetDescription

`func (o *ListCentrifyZone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListCentrifyZone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListCentrifyZone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListCentrifyZone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListCentrifyZone) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListCentrifyZone) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDistinguishedName

`func (o *ListCentrifyZone) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *ListCentrifyZone) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *ListCentrifyZone) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *ListCentrifyZone) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### SetDistinguishedNameNil

`func (o *ListCentrifyZone) SetDistinguishedNameNil(b bool)`

 SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil

### UnsetDistinguishedName
`func (o *ListCentrifyZone) UnsetDistinguishedName()`

UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil
### GetZoneName

`func (o *ListCentrifyZone) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *ListCentrifyZone) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *ListCentrifyZone) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *ListCentrifyZone) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### SetZoneNameNil

`func (o *ListCentrifyZone) SetZoneNameNil(b bool)`

 SetZoneNameNil sets the value for ZoneName to be an explicit nil

### UnsetZoneName
`func (o *ListCentrifyZone) UnsetZoneName()`

UnsetZoneName ensures that no value is present for ZoneName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


