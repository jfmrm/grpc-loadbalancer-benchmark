# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: helloworld.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x10helloworld.proto\"\x1c\n\x0cHelloRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"\x1d\n\nHelloReply\x12\x0f\n\x07message\x18\x01 \x01(\t28\n\x0cHelloService\x12(\n\x08SayHello\x12\r.HelloRequest\x1a\x0b.HelloReply\"\x00\x42\x39Z7github.com/jfmrm/grpc-loadbalancer-benchmark/helloworldb\x06proto3')



_HELLOREQUEST = DESCRIPTOR.message_types_by_name['HelloRequest']
_HELLOREPLY = DESCRIPTOR.message_types_by_name['HelloReply']
HelloRequest = _reflection.GeneratedProtocolMessageType('HelloRequest', (_message.Message,), {
  'DESCRIPTOR' : _HELLOREQUEST,
  '__module__' : 'helloworld_pb2'
  # @@protoc_insertion_point(class_scope:HelloRequest)
  })
_sym_db.RegisterMessage(HelloRequest)

HelloReply = _reflection.GeneratedProtocolMessageType('HelloReply', (_message.Message,), {
  'DESCRIPTOR' : _HELLOREPLY,
  '__module__' : 'helloworld_pb2'
  # @@protoc_insertion_point(class_scope:HelloReply)
  })
_sym_db.RegisterMessage(HelloReply)

_HELLOSERVICE = DESCRIPTOR.services_by_name['HelloService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z7github.com/jfmrm/grpc-loadbalancer-benchmark/helloworld'
  _HELLOREQUEST._serialized_start=20
  _HELLOREQUEST._serialized_end=48
  _HELLOREPLY._serialized_start=50
  _HELLOREPLY._serialized_end=79
  _HELLOSERVICE._serialized_start=81
  _HELLOSERVICE._serialized_end=137
# @@protoc_insertion_point(module_scope)
