# coding: utf-8

"""
    tatoperator

    Python SDK for Tat-Operator

    The version of the OpenAPI document: v1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from tatope.models.k8s_io_apimachinery_pkg_runtime_type_meta import K8sIoApimachineryPkgRuntimeTypeMeta

class TestK8sIoApimachineryPkgRuntimeTypeMeta(unittest.TestCase):
    """K8sIoApimachineryPkgRuntimeTypeMeta unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> K8sIoApimachineryPkgRuntimeTypeMeta:
        """Test K8sIoApimachineryPkgRuntimeTypeMeta
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `K8sIoApimachineryPkgRuntimeTypeMeta`
        """
        model = K8sIoApimachineryPkgRuntimeTypeMeta()
        if include_optional:
            return K8sIoApimachineryPkgRuntimeTypeMeta(
                api_version = '',
                kind = ''
            )
        else:
            return K8sIoApimachineryPkgRuntimeTypeMeta(
        )
        """

    def testK8sIoApimachineryPkgRuntimeTypeMeta(self):
        """Test K8sIoApimachineryPkgRuntimeTypeMeta"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
