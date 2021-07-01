# InputSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesSelector** | Pointer to [**InputSpecInputFilesSelector**](InputSpecInputFilesSelector.md) |  | [optional] 
**OnNfsFiles** | Pointer to **NullableBool** | Selects whether input is files inside vmdks or files on NFS. One of vm_selector or files_selector will be chosen based on this flag. | [optional] 
**VmSelector** | Pointer to [**InputSpecInputVMsSelector**](InputSpecInputVMsSelector.md) |  | [optional] 

## Methods

### NewInputSpec

`func NewInputSpec() *InputSpec`

NewInputSpec instantiates a new InputSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputSpecWithDefaults

`func NewInputSpecWithDefaults() *InputSpec`

NewInputSpecWithDefaults instantiates a new InputSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesSelector

`func (o *InputSpec) GetFilesSelector() InputSpecInputFilesSelector`

GetFilesSelector returns the FilesSelector field if non-nil, zero value otherwise.

### GetFilesSelectorOk

`func (o *InputSpec) GetFilesSelectorOk() (*InputSpecInputFilesSelector, bool)`

GetFilesSelectorOk returns a tuple with the FilesSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesSelector

`func (o *InputSpec) SetFilesSelector(v InputSpecInputFilesSelector)`

SetFilesSelector sets FilesSelector field to given value.

### HasFilesSelector

`func (o *InputSpec) HasFilesSelector() bool`

HasFilesSelector returns a boolean if a field has been set.

### GetOnNfsFiles

`func (o *InputSpec) GetOnNfsFiles() bool`

GetOnNfsFiles returns the OnNfsFiles field if non-nil, zero value otherwise.

### GetOnNfsFilesOk

`func (o *InputSpec) GetOnNfsFilesOk() (*bool, bool)`

GetOnNfsFilesOk returns a tuple with the OnNfsFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnNfsFiles

`func (o *InputSpec) SetOnNfsFiles(v bool)`

SetOnNfsFiles sets OnNfsFiles field to given value.

### HasOnNfsFiles

`func (o *InputSpec) HasOnNfsFiles() bool`

HasOnNfsFiles returns a boolean if a field has been set.

### SetOnNfsFilesNil

`func (o *InputSpec) SetOnNfsFilesNil(b bool)`

 SetOnNfsFilesNil sets the value for OnNfsFiles to be an explicit nil

### UnsetOnNfsFiles
`func (o *InputSpec) UnsetOnNfsFiles()`

UnsetOnNfsFiles ensures that no value is present for OnNfsFiles, not even an explicit nil
### GetVmSelector

`func (o *InputSpec) GetVmSelector() InputSpecInputVMsSelector`

GetVmSelector returns the VmSelector field if non-nil, zero value otherwise.

### GetVmSelectorOk

`func (o *InputSpec) GetVmSelectorOk() (*InputSpecInputVMsSelector, bool)`

GetVmSelectorOk returns a tuple with the VmSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSelector

`func (o *InputSpec) SetVmSelector(v InputSpecInputVMsSelector)`

SetVmSelector sets VmSelector field to given value.

### HasVmSelector

`func (o *InputSpec) HasVmSelector() bool`

HasVmSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


