# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: document.proto
# Protobuf Python Version: 5.27.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    2,
    '',
    'document.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0e\x64ocument.proto\x12\ndocservice\"3\n\rUpdateRequest\x12\x0f\n\x07\x63ontent\x18\x01 \x01(\t\x12\x11\n\tclient_id\x18\x02 \x01(\t\"\\\n\x0eUpdateResponse\x12\x0f\n\x07\x63ontent\x18\x01 \x01(\t\x12\x11\n\tclient_id\x18\x02 \x01(\t\x12\x0f\n\x07success\x18\x03 \x01(\x08\x12\x15\n\rerror_message\x18\x04 \x01(\t\"\x14\n\x12GetDocumentRequest\"&\n\x13GetDocumentResponse\x12\x0f\n\x07\x63ontent\x18\x01 \x01(\t\"1\n\x0bLogResponse\x12\x0f\n\x07message\x18\x01 \x01(\t\x12\x11\n\ttimestamp\x18\x02 \x01(\t2\xf7\x01\n\x0f\x44ocumentService\x12L\n\rStreamUpdates\x12\x19.docservice.UpdateRequest\x1a\x1a.docservice.UpdateResponse\"\x00(\x01\x30\x01\x12P\n\x0bGetDocument\x12\x1e.docservice.GetDocumentRequest\x1a\x1f.docservice.GetDocumentResponse\"\x00\x12\x44\n\nLogUpdates\x12\x19.docservice.UpdateRequest\x1a\x17.docservice.LogResponse\"\x00\x30\x01\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'document_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_UPDATEREQUEST']._serialized_start=30
  _globals['_UPDATEREQUEST']._serialized_end=81
  _globals['_UPDATERESPONSE']._serialized_start=83
  _globals['_UPDATERESPONSE']._serialized_end=175
  _globals['_GETDOCUMENTREQUEST']._serialized_start=177
  _globals['_GETDOCUMENTREQUEST']._serialized_end=197
  _globals['_GETDOCUMENTRESPONSE']._serialized_start=199
  _globals['_GETDOCUMENTRESPONSE']._serialized_end=237
  _globals['_LOGRESPONSE']._serialized_start=239
  _globals['_LOGRESPONSE']._serialized_end=288
  _globals['_DOCUMENTSERVICE']._serialized_start=291
  _globals['_DOCUMENTSERVICE']._serialized_end=538
# @@protoc_insertion_point(module_scope)
