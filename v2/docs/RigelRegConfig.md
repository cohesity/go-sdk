# RigelRegConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControlPlaneConnectionInfo** | Pointer to [**RigelConnectionInfo**](RigelConnectionInfo.md) |  | [optional] 
**DataPlaneConnectionInfo** | Pointer to [**RigelConnectionInfo**](RigelConnectionInfo.md) |  | [optional] 
**IsCertificateValid** | Pointer to **NullableBool** | Flag to indicate if certificate is valid. | [optional] 
**RegInfo** | Pointer to [**RigelClaimInfo**](RigelClaimInfo.md) |  | [optional] 

## Methods

### NewRigelRegConfig

`func NewRigelRegConfig() *RigelRegConfig`

NewRigelRegConfig instantiates a new RigelRegConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRigelRegConfigWithDefaults

`func NewRigelRegConfigWithDefaults() *RigelRegConfig`

NewRigelRegConfigWithDefaults instantiates a new RigelRegConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControlPlaneConnectionInfo

`func (o *RigelRegConfig) GetControlPlaneConnectionInfo() RigelConnectionInfo`

GetControlPlaneConnectionInfo returns the ControlPlaneConnectionInfo field if non-nil, zero value otherwise.

### GetControlPlaneConnectionInfoOk

`func (o *RigelRegConfig) GetControlPlaneConnectionInfoOk() (*RigelConnectionInfo, bool)`

GetControlPlaneConnectionInfoOk returns a tuple with the ControlPlaneConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneConnectionInfo

`func (o *RigelRegConfig) SetControlPlaneConnectionInfo(v RigelConnectionInfo)`

SetControlPlaneConnectionInfo sets ControlPlaneConnectionInfo field to given value.

### HasControlPlaneConnectionInfo

`func (o *RigelRegConfig) HasControlPlaneConnectionInfo() bool`

HasControlPlaneConnectionInfo returns a boolean if a field has been set.

### GetDataPlaneConnectionInfo

`func (o *RigelRegConfig) GetDataPlaneConnectionInfo() RigelConnectionInfo`

GetDataPlaneConnectionInfo returns the DataPlaneConnectionInfo field if non-nil, zero value otherwise.

### GetDataPlaneConnectionInfoOk

`func (o *RigelRegConfig) GetDataPlaneConnectionInfoOk() (*RigelConnectionInfo, bool)`

GetDataPlaneConnectionInfoOk returns a tuple with the DataPlaneConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPlaneConnectionInfo

`func (o *RigelRegConfig) SetDataPlaneConnectionInfo(v RigelConnectionInfo)`

SetDataPlaneConnectionInfo sets DataPlaneConnectionInfo field to given value.

### HasDataPlaneConnectionInfo

`func (o *RigelRegConfig) HasDataPlaneConnectionInfo() bool`

HasDataPlaneConnectionInfo returns a boolean if a field has been set.

### GetIsCertificateValid

`func (o *RigelRegConfig) GetIsCertificateValid() bool`

GetIsCertificateValid returns the IsCertificateValid field if non-nil, zero value otherwise.

### GetIsCertificateValidOk

`func (o *RigelRegConfig) GetIsCertificateValidOk() (*bool, bool)`

GetIsCertificateValidOk returns a tuple with the IsCertificateValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCertificateValid

`func (o *RigelRegConfig) SetIsCertificateValid(v bool)`

SetIsCertificateValid sets IsCertificateValid field to given value.

### HasIsCertificateValid

`func (o *RigelRegConfig) HasIsCertificateValid() bool`

HasIsCertificateValid returns a boolean if a field has been set.

### SetIsCertificateValidNil

`func (o *RigelRegConfig) SetIsCertificateValidNil(b bool)`

 SetIsCertificateValidNil sets the value for IsCertificateValid to be an explicit nil

### UnsetIsCertificateValid
`func (o *RigelRegConfig) UnsetIsCertificateValid()`

UnsetIsCertificateValid ensures that no value is present for IsCertificateValid, not even an explicit nil
### GetRegInfo

`func (o *RigelRegConfig) GetRegInfo() RigelClaimInfo`

GetRegInfo returns the RegInfo field if non-nil, zero value otherwise.

### GetRegInfoOk

`func (o *RigelRegConfig) GetRegInfoOk() (*RigelClaimInfo, bool)`

GetRegInfoOk returns a tuple with the RegInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegInfo

`func (o *RigelRegConfig) SetRegInfo(v RigelClaimInfo)`

SetRegInfo sets RegInfo field to given value.

### HasRegInfo

`func (o *RigelRegConfig) HasRegInfo() bool`

HasRegInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


