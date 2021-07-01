# NetappClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactInfo** | Pointer to **NullableString** | Specifies information about the contact for the NetApp cluster such as a name, phone number, and email address. | [optional] 
**Location** | Pointer to **NullableString** | Specifies where this NetApp cluster is located. This location identification string is configured by the NetApp system administrator. This field does not contain the NetApp cluster hostname or IP address. | [optional] 
**SerialNumber** | Pointer to **NullableString** | Specifies the serial number of the NetApp cluster in the format: x-xx-xxxxxx. | [optional] 

## Methods

### NewNetappClusterInfo

`func NewNetappClusterInfo() *NetappClusterInfo`

NewNetappClusterInfo instantiates a new NetappClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetappClusterInfoWithDefaults

`func NewNetappClusterInfoWithDefaults() *NetappClusterInfo`

NewNetappClusterInfoWithDefaults instantiates a new NetappClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactInfo

`func (o *NetappClusterInfo) GetContactInfo() string`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *NetappClusterInfo) GetContactInfoOk() (*string, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *NetappClusterInfo) SetContactInfo(v string)`

SetContactInfo sets ContactInfo field to given value.

### HasContactInfo

`func (o *NetappClusterInfo) HasContactInfo() bool`

HasContactInfo returns a boolean if a field has been set.

### SetContactInfoNil

`func (o *NetappClusterInfo) SetContactInfoNil(b bool)`

 SetContactInfoNil sets the value for ContactInfo to be an explicit nil

### UnsetContactInfo
`func (o *NetappClusterInfo) UnsetContactInfo()`

UnsetContactInfo ensures that no value is present for ContactInfo, not even an explicit nil
### GetLocation

`func (o *NetappClusterInfo) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *NetappClusterInfo) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *NetappClusterInfo) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *NetappClusterInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *NetappClusterInfo) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *NetappClusterInfo) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetSerialNumber

`func (o *NetappClusterInfo) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NetappClusterInfo) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NetappClusterInfo) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *NetappClusterInfo) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *NetappClusterInfo) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *NetappClusterInfo) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


