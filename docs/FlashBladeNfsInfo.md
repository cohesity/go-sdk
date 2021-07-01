# FlashBladeNfsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportRules** | Pointer to **NullableString** | Specifies NFS protocol export rules. Rules are in the form host(options). host represents one of the following categories:  IP address in the form ddd.ddd.ddd.ddd for IPv4, or xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx for IPv6.  Netmask in the form ddd.ddd.ddd.ddd/dd for IPv4, or xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx for IPv6.  Wildcard in the form * to represent all clients  options in parenthesis represents a comma-separated list of NFS export options. Valid export options are rw, ro, root_squash, no_root_squash, and fileid_32bit. | [optional] 

## Methods

### NewFlashBladeNfsInfo

`func NewFlashBladeNfsInfo() *FlashBladeNfsInfo`

NewFlashBladeNfsInfo instantiates a new FlashBladeNfsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlashBladeNfsInfoWithDefaults

`func NewFlashBladeNfsInfoWithDefaults() *FlashBladeNfsInfo`

NewFlashBladeNfsInfoWithDefaults instantiates a new FlashBladeNfsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportRules

`func (o *FlashBladeNfsInfo) GetExportRules() string`

GetExportRules returns the ExportRules field if non-nil, zero value otherwise.

### GetExportRulesOk

`func (o *FlashBladeNfsInfo) GetExportRulesOk() (*string, bool)`

GetExportRulesOk returns a tuple with the ExportRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportRules

`func (o *FlashBladeNfsInfo) SetExportRules(v string)`

SetExportRules sets ExportRules field to given value.

### HasExportRules

`func (o *FlashBladeNfsInfo) HasExportRules() bool`

HasExportRules returns a boolean if a field has been set.

### SetExportRulesNil

`func (o *FlashBladeNfsInfo) SetExportRulesNil(b bool)`

 SetExportRulesNil sets the value for ExportRules to be an explicit nil

### UnsetExportRules
`func (o *FlashBladeNfsInfo) UnsetExportRules()`

UnsetExportRules ensures that no value is present for ExportRules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


