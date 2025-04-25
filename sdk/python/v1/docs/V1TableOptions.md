# V1TableOptions

TableOptions are used when a Table is requested by the caller.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_version** | **str** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**include_object** | **str** | includeObject decides whether to include each object along with its columnar information. Specifying \&quot;None\&quot; will return no object, specifying \&quot;Object\&quot; will return the full object contents, and specifying \&quot;Metadata\&quot; (the default) will return the object&#39;s metadata in the PartialObjectMetadata kind in version v1beta1 of the meta.k8s.io API group. | [optional] 
**kind** | **str** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 

## Example

```python
from tatope.models.v1_table_options import V1TableOptions

# TODO update the JSON string below
json = "{}"
# create an instance of V1TableOptions from a JSON string
v1_table_options_instance = V1TableOptions.from_json(json)
# print the JSON string representation of the object
print(V1TableOptions.to_json())

# convert the object into a dict
v1_table_options_dict = v1_table_options_instance.to_dict()
# create an instance of V1TableOptions from a dict
v1_table_options_from_dict = V1TableOptions.from_dict(v1_table_options_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


