# ExportMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataFormatType** | Pointer to **string** | Denotes the format in which the metadata will be exported. | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**MetadataSelector** | Pointer to [**MetadataSelector**](MetadataSelector.md) |  | [optional] 

## Methods

### NewExportMetadataRequest

`func NewExportMetadataRequest() *ExportMetadataRequest`

NewExportMetadataRequest instantiates a new ExportMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportMetadataRequestWithDefaults

`func NewExportMetadataRequestWithDefaults() *ExportMetadataRequest`

NewExportMetadataRequestWithDefaults instantiates a new ExportMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataFormatType

`func (o *ExportMetadataRequest) GetDataFormatType() string`

GetDataFormatType returns the DataFormatType field if non-nil, zero value otherwise.

### GetDataFormatTypeOk

`func (o *ExportMetadataRequest) GetDataFormatTypeOk() (*string, bool)`

GetDataFormatTypeOk returns a tuple with the DataFormatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFormatType

`func (o *ExportMetadataRequest) SetDataFormatType(v string)`

SetDataFormatType sets DataFormatType field to given value.

### HasDataFormatType

`func (o *ExportMetadataRequest) HasDataFormatType() bool`

HasDataFormatType returns a boolean if a field has been set.

### GetLocation

`func (o *ExportMetadataRequest) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ExportMetadataRequest) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ExportMetadataRequest) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ExportMetadataRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMetadataSelector

`func (o *ExportMetadataRequest) GetMetadataSelector() MetadataSelector`

GetMetadataSelector returns the MetadataSelector field if non-nil, zero value otherwise.

### GetMetadataSelectorOk

`func (o *ExportMetadataRequest) GetMetadataSelectorOk() (*MetadataSelector, bool)`

GetMetadataSelectorOk returns a tuple with the MetadataSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSelector

`func (o *ExportMetadataRequest) SetMetadataSelector(v MetadataSelector)`

SetMetadataSelector sets MetadataSelector field to given value.

### HasMetadataSelector

`func (o *ExportMetadataRequest) HasMetadataSelector() bool`

HasMetadataSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


