# TatOpeEx

TatOpeEx is the Schema for the tatopeexes API

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_version** | **str** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**kind** | **str** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**metadata** | [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**spec** | [**TatOpeExSpec**](TatOpeExSpec.md) |  | [optional] 
**status** | [**TatOpeExStatus**](TatOpeExStatus.md) |  | [optional] 

## Example

```python
from tatope.models.tat_ope_ex import TatOpeEx

# TODO update the JSON string below
json = "{}"
# create an instance of TatOpeEx from a JSON string
tat_ope_ex_instance = TatOpeEx.from_json(json)
# print the JSON string representation of the object
print(TatOpeEx.to_json())

# convert the object into a dict
tat_ope_ex_dict = tat_ope_ex_instance.to_dict()
# create an instance of TatOpeEx from a dict
tat_ope_ex_from_dict = TatOpeEx.from_dict(tat_ope_ex_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


