# UserIdMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CentrifyZoneMapping** | Pointer to [**CentrifyZone**](CentrifyZone.md) |  | [optional] 
**CustomAttributesMapping** | Pointer to [**CustomUnixIdAttributes**](CustomUnixIdAttributes.md) |  | [optional] 
**FixedMapping** | Pointer to [**FixedUnixIdMapping**](FixedUnixIdMapping.md) |  | [optional] 
**Type** | Pointer to **NullableString** | Specifies the mapping type used. &#39;kRid&#39; indicates the kRid mapping type. &#39;kRfc2307&#39; indicates the kRfc2307 mapping type. &#39;kSfu30&#39; indicates the kSfu30 mapping type. &#39;kCentrify&#39; indicates the mapping type to refer to a centrify zone. &#39;kFixed&#39; indicates the mapping from all Active Directory users to a fixed Unix uid, and gid. &#39;kCustomAttributes&#39; indicates the mapping to derive from custom attributes defined in an AD domain. &#39;kLdapProvider&#39; indicates the Active Directory to LDAP provider mapping. | [optional] 

## Methods

### NewUserIdMapping

`func NewUserIdMapping() *UserIdMapping`

NewUserIdMapping instantiates a new UserIdMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdMappingWithDefaults

`func NewUserIdMappingWithDefaults() *UserIdMapping`

NewUserIdMappingWithDefaults instantiates a new UserIdMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCentrifyZoneMapping

`func (o *UserIdMapping) GetCentrifyZoneMapping() CentrifyZone`

GetCentrifyZoneMapping returns the CentrifyZoneMapping field if non-nil, zero value otherwise.

### GetCentrifyZoneMappingOk

`func (o *UserIdMapping) GetCentrifyZoneMappingOk() (*CentrifyZone, bool)`

GetCentrifyZoneMappingOk returns a tuple with the CentrifyZoneMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentrifyZoneMapping

`func (o *UserIdMapping) SetCentrifyZoneMapping(v CentrifyZone)`

SetCentrifyZoneMapping sets CentrifyZoneMapping field to given value.

### HasCentrifyZoneMapping

`func (o *UserIdMapping) HasCentrifyZoneMapping() bool`

HasCentrifyZoneMapping returns a boolean if a field has been set.

### GetCustomAttributesMapping

`func (o *UserIdMapping) GetCustomAttributesMapping() CustomUnixIdAttributes`

GetCustomAttributesMapping returns the CustomAttributesMapping field if non-nil, zero value otherwise.

### GetCustomAttributesMappingOk

`func (o *UserIdMapping) GetCustomAttributesMappingOk() (*CustomUnixIdAttributes, bool)`

GetCustomAttributesMappingOk returns a tuple with the CustomAttributesMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributesMapping

`func (o *UserIdMapping) SetCustomAttributesMapping(v CustomUnixIdAttributes)`

SetCustomAttributesMapping sets CustomAttributesMapping field to given value.

### HasCustomAttributesMapping

`func (o *UserIdMapping) HasCustomAttributesMapping() bool`

HasCustomAttributesMapping returns a boolean if a field has been set.

### GetFixedMapping

`func (o *UserIdMapping) GetFixedMapping() FixedUnixIdMapping`

GetFixedMapping returns the FixedMapping field if non-nil, zero value otherwise.

### GetFixedMappingOk

`func (o *UserIdMapping) GetFixedMappingOk() (*FixedUnixIdMapping, bool)`

GetFixedMappingOk returns a tuple with the FixedMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedMapping

`func (o *UserIdMapping) SetFixedMapping(v FixedUnixIdMapping)`

SetFixedMapping sets FixedMapping field to given value.

### HasFixedMapping

`func (o *UserIdMapping) HasFixedMapping() bool`

HasFixedMapping returns a boolean if a field has been set.

### GetType

`func (o *UserIdMapping) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserIdMapping) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserIdMapping) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserIdMapping) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *UserIdMapping) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *UserIdMapping) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


