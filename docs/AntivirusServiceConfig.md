# AntivirusServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Specifies the description of the Antivirus service. This could be any additional information admin might associate with the Antivirus service. | [optional] 
**IcapUri** | **NullableString** | Specifies the ICAP uri for this Antivirus service. It is of the form icap://&lt;ip-address&gt;[:&lt;port&gt;]/&lt;service&gt; | 
**Tag** | Pointer to **NullableString** | Specifies the tag of antivirus service. This is service-specific \&quot;cookie\&quot; sent from Antivirus server to clients that represents a service&#39;s current state. This tag validates that previous Antivirus server responses can still be considered fresh by an Antivirus client that may be caching them. If a change on the AV server invalidates previous responses, the AV server can invalidate portions of the Antivirus client&#39;s cache by changing its service tag. | [optional] 
**TagId** | Pointer to **NullableInt64** | Specifies the tag Id of antivirus service. | [optional] 

## Methods

### NewAntivirusServiceConfig

`func NewAntivirusServiceConfig(icapUri NullableString, ) *AntivirusServiceConfig`

NewAntivirusServiceConfig instantiates a new AntivirusServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntivirusServiceConfigWithDefaults

`func NewAntivirusServiceConfigWithDefaults() *AntivirusServiceConfig`

NewAntivirusServiceConfigWithDefaults instantiates a new AntivirusServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AntivirusServiceConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AntivirusServiceConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AntivirusServiceConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AntivirusServiceConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AntivirusServiceConfig) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AntivirusServiceConfig) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIcapUri

`func (o *AntivirusServiceConfig) GetIcapUri() string`

GetIcapUri returns the IcapUri field if non-nil, zero value otherwise.

### GetIcapUriOk

`func (o *AntivirusServiceConfig) GetIcapUriOk() (*string, bool)`

GetIcapUriOk returns a tuple with the IcapUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcapUri

`func (o *AntivirusServiceConfig) SetIcapUri(v string)`

SetIcapUri sets IcapUri field to given value.


### SetIcapUriNil

`func (o *AntivirusServiceConfig) SetIcapUriNil(b bool)`

 SetIcapUriNil sets the value for IcapUri to be an explicit nil

### UnsetIcapUri
`func (o *AntivirusServiceConfig) UnsetIcapUri()`

UnsetIcapUri ensures that no value is present for IcapUri, not even an explicit nil
### GetTag

`func (o *AntivirusServiceConfig) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AntivirusServiceConfig) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AntivirusServiceConfig) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AntivirusServiceConfig) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *AntivirusServiceConfig) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *AntivirusServiceConfig) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetTagId

`func (o *AntivirusServiceConfig) GetTagId() int64`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *AntivirusServiceConfig) GetTagIdOk() (*int64, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *AntivirusServiceConfig) SetTagId(v int64)`

SetTagId sets TagId field to given value.

### HasTagId

`func (o *AntivirusServiceConfig) HasTagId() bool`

HasTagId returns a boolean if a field has been set.

### SetTagIdNil

`func (o *AntivirusServiceConfig) SetTagIdNil(b bool)`

 SetTagIdNil sets the value for TagId to be an explicit nil

### UnsetTagId
`func (o *AntivirusServiceConfig) UnsetTagId()`

UnsetTagId ensures that no value is present for TagId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


