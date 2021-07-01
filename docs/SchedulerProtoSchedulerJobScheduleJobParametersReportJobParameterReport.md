# SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the report name. | [optional] 
**OutputFormat** | Pointer to **NullableString** | Specifies the output format of the report. | [optional] 
**Parameters** | Pointer to [**SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters**](SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters.md) |  | [optional] 
**SubjectLine** | Pointer to **NullableString** | Specifies the subject line for report. | [optional] 
**Type** | Pointer to **NullableInt32** | Specifies the report type. | [optional] 

## Methods

### NewSchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport

`func NewSchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport() *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport`

NewSchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport instantiates a new SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportWithDefaults

`func NewSchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportWithDefaults() *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport`

NewSchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportWithDefaults instantiates a new SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOutputFormat

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### SetOutputFormatNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) SetOutputFormatNil(b bool)`

 SetOutputFormatNil sets the value for OutputFormat to be an explicit nil

### UnsetOutputFormat
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) UnsetOutputFormat()`

UnsetOutputFormat ensures that no value is present for OutputFormat, not even an explicit nil
### GetParameters

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) GetParameters() SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) GetParametersOk() (*SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) SetParameters(v SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReportParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetSubjectLine

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) GetSubjectLine() string`

GetSubjectLine returns the SubjectLine field if non-nil, zero value otherwise.

### GetSubjectLineOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) GetSubjectLineOk() (*string, bool)`

GetSubjectLineOk returns a tuple with the SubjectLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectLine

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) SetSubjectLine(v string)`

SetSubjectLine sets SubjectLine field to given value.

### HasSubjectLine

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) HasSubjectLine() bool`

HasSubjectLine returns a boolean if a field has been set.

### SetSubjectLineNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) SetSubjectLineNil(b bool)`

 SetSubjectLineNil sets the value for SubjectLine to be an explicit nil

### UnsetSubjectLine
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) UnsetSubjectLine()`

UnsetSubjectLine ensures that no value is present for SubjectLine, not even an explicit nil
### GetType

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SchedulerProtoSchedulerJobScheduleJobParametersReportJobParameterReport) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


