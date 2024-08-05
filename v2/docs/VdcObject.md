# VdcObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalogs** | Pointer to [**[]VdcCatalog**](VdcCatalog.md) | Specifies a list of VDC Catalogs. | [optional] 
**OrgNetworks** | Pointer to [**[]OrgVDCNetwork**](OrgVDCNetwork.md) | Specifies a list of Organization VDC Networks. | [optional] 

## Methods

### NewVdcObject

`func NewVdcObject() *VdcObject`

NewVdcObject instantiates a new VdcObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVdcObjectWithDefaults

`func NewVdcObjectWithDefaults() *VdcObject`

NewVdcObjectWithDefaults instantiates a new VdcObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogs

`func (o *VdcObject) GetCatalogs() []VdcCatalog`

GetCatalogs returns the Catalogs field if non-nil, zero value otherwise.

### GetCatalogsOk

`func (o *VdcObject) GetCatalogsOk() (*[]VdcCatalog, bool)`

GetCatalogsOk returns a tuple with the Catalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogs

`func (o *VdcObject) SetCatalogs(v []VdcCatalog)`

SetCatalogs sets Catalogs field to given value.

### HasCatalogs

`func (o *VdcObject) HasCatalogs() bool`

HasCatalogs returns a boolean if a field has been set.

### GetOrgNetworks

`func (o *VdcObject) GetOrgNetworks() []OrgVDCNetwork`

GetOrgNetworks returns the OrgNetworks field if non-nil, zero value otherwise.

### GetOrgNetworksOk

`func (o *VdcObject) GetOrgNetworksOk() (*[]OrgVDCNetwork, bool)`

GetOrgNetworksOk returns a tuple with the OrgNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgNetworks

`func (o *VdcObject) SetOrgNetworks(v []OrgVDCNetwork)`

SetOrgNetworks sets OrgNetworks field to given value.

### HasOrgNetworks

`func (o *VdcObject) HasOrgNetworks() bool`

HasOrgNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


