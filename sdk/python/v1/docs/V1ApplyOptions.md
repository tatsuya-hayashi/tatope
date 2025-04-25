# V1ApplyOptions

ApplyOptions may be provided when applying an API object. FieldManager is required for apply requests. ApplyOptions is equivalent to PatchOptions. It is provided as a convenience with documentation that speaks specifically to how the options fields relate to apply.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_version** | **str** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**dry_run** | **List[str]** | When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | [optional] 
**field_manager** | **str** | fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint. This field is required. | [default to '']
**force** | **bool** | Force is going to \&quot;force\&quot; Apply requests. It means user will re-acquire conflicting fields owned by other people. | [default to False]
**kind** | **str** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 

## Example

```python
from tatope.models.v1_apply_options import V1ApplyOptions

# TODO update the JSON string below
json = "{}"
# create an instance of V1ApplyOptions from a JSON string
v1_apply_options_instance = V1ApplyOptions.from_json(json)
# print the JSON string representation of the object
print(V1ApplyOptions.to_json())

# convert the object into a dict
v1_apply_options_dict = v1_apply_options_instance.to_dict()
# create an instance of V1ApplyOptions from a dict
v1_apply_options_from_dict = V1ApplyOptions.from_dict(v1_apply_options_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


