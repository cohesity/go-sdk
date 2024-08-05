# DnsDelegationZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsZoneResolvedVips** | Pointer to **[]string** | Dns zone resolved VIPs. | [optional] 
**DnsZoneVips** | Pointer to **[]string** | VIPs part of dns zone. | [optional] 
**Name** | **string** | Name of dns zone. | 

## Methods

### NewDnsDelegationZone

`func NewDnsDelegationZone(name string, ) *DnsDelegationZone`

NewDnsDelegationZone instantiates a new DnsDelegationZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsDelegationZoneWithDefaults

`func NewDnsDelegationZoneWithDefaults() *DnsDelegationZone`

NewDnsDelegationZoneWithDefaults instantiates a new DnsDelegationZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsZoneResolvedVips

`func (o *DnsDelegationZone) GetDnsZoneResolvedVips() []string`

GetDnsZoneResolvedVips returns the DnsZoneResolvedVips field if non-nil, zero value otherwise.

### GetDnsZoneResolvedVipsOk

`func (o *DnsDelegationZone) GetDnsZoneResolvedVipsOk() (*[]string, bool)`

GetDnsZoneResolvedVipsOk returns a tuple with the DnsZoneResolvedVips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZoneResolvedVips

`func (o *DnsDelegationZone) SetDnsZoneResolvedVips(v []string)`

SetDnsZoneResolvedVips sets DnsZoneResolvedVips field to given value.

### HasDnsZoneResolvedVips

`func (o *DnsDelegationZone) HasDnsZoneResolvedVips() bool`

HasDnsZoneResolvedVips returns a boolean if a field has been set.

### SetDnsZoneResolvedVipsNil

`func (o *DnsDelegationZone) SetDnsZoneResolvedVipsNil(b bool)`

 SetDnsZoneResolvedVipsNil sets the value for DnsZoneResolvedVips to be an explicit nil

### UnsetDnsZoneResolvedVips
`func (o *DnsDelegationZone) UnsetDnsZoneResolvedVips()`

UnsetDnsZoneResolvedVips ensures that no value is present for DnsZoneResolvedVips, not even an explicit nil
### GetDnsZoneVips

`func (o *DnsDelegationZone) GetDnsZoneVips() []string`

GetDnsZoneVips returns the DnsZoneVips field if non-nil, zero value otherwise.

### GetDnsZoneVipsOk

`func (o *DnsDelegationZone) GetDnsZoneVipsOk() (*[]string, bool)`

GetDnsZoneVipsOk returns a tuple with the DnsZoneVips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZoneVips

`func (o *DnsDelegationZone) SetDnsZoneVips(v []string)`

SetDnsZoneVips sets DnsZoneVips field to given value.

### HasDnsZoneVips

`func (o *DnsDelegationZone) HasDnsZoneVips() bool`

HasDnsZoneVips returns a boolean if a field has been set.

### SetDnsZoneVipsNil

`func (o *DnsDelegationZone) SetDnsZoneVipsNil(b bool)`

 SetDnsZoneVipsNil sets the value for DnsZoneVips to be an explicit nil

### UnsetDnsZoneVips
`func (o *DnsDelegationZone) UnsetDnsZoneVips()`

UnsetDnsZoneVips ensures that no value is present for DnsZoneVips, not even an explicit nil
### GetName

`func (o *DnsDelegationZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsDelegationZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsDelegationZone) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


