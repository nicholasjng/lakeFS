# coding: utf-8

"""
    lakeFS API

    lakeFS HTTP API

    The version of the OpenAPI document: 0.1.0
    Contact: services@treeverse.io
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from datetime import datetime
from typing import Optional
from pydantic import BaseModel, Field, StrictStr, validator

class ActionRun(BaseModel):
    """
    ActionRun
    """
    run_id: StrictStr = Field(...)
    branch: StrictStr = Field(...)
    start_time: datetime = Field(...)
    end_time: Optional[datetime] = None
    event_type: StrictStr = Field(...)
    status: StrictStr = Field(...)
    commit_id: StrictStr = Field(...)
    __properties = ["run_id", "branch", "start_time", "end_time", "event_type", "status", "commit_id"]

    @validator('status')
    def status_validate_enum(cls, value):
        """Validates the enum"""
        if value not in ('failed', 'completed'):
            raise ValueError("must be one of enum values ('failed', 'completed')")
        return value

    class Config:
        """Pydantic configuration"""
        allow_population_by_field_name = True
        validate_assignment = True

    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.dict(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> ActionRun:
        """Create an instance of ActionRun from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> ActionRun:
        """Create an instance of ActionRun from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return ActionRun.parse_obj(obj)

        _obj = ActionRun.parse_obj({
            "run_id": obj.get("run_id"),
            "branch": obj.get("branch"),
            "start_time": obj.get("start_time"),
            "end_time": obj.get("end_time"),
            "event_type": obj.get("event_type"),
            "status": obj.get("status"),
            "commit_id": obj.get("commit_id")
        })
        return _obj

