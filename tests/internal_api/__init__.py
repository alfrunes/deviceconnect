# coding: utf-8

# flake8: noqa

"""
    Device Connect

    Internal API for managing persistent device connections. Intended for use by the web GUI.   # noqa: E501

    The version of the OpenAPI document: 1
    Contact: support@mender.io
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

__version__ = "1.0.0"

# import apis into sdk package
from internal_api.api.internal_api_client import InternalAPIClient

# import ApiClient
from internal_api.api_client import ApiClient
from internal_api.configuration import Configuration
from internal_api.exceptions import OpenApiException
from internal_api.exceptions import ApiTypeError
from internal_api.exceptions import ApiValueError
from internal_api.exceptions import ApiKeyError
from internal_api.exceptions import ApiException
# import models into sdk package
from internal_api.models.device import Device
from internal_api.models.error import Error
from internal_api.models.new_tenant import NewTenant

