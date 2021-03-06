// Code generated by protoc-gen-gogo.
// source: sizeunderscore.proto
// DO NOT EDIT!

/*
Package sizeunderscore is a generated protocol buffer package.

It is generated from these files:
	sizeunderscore.proto

It has these top-level messages:
	SizeMessage
*/
package sizeunderscore

import testing "testing"
import math_rand "math/rand"
import time "time"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

func TestSizeMessageProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedSizeMessage(popr, false)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &SizeMessage{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	littlefuzz := make([]byte, len(data))
	copy(littlefuzz, data)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_gogo_protobuf_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestSizeMessageMarshalTo(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedSizeMessage(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &SizeMessage{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestSizeMessageJSON(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedSizeMessage(popr, true)
	marshaler := github_com_gogo_protobuf_jsonpb.Marshaller{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatal(err)
	}
	msg := &SizeMessage{}
	err = github_com_gogo_protobuf_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatal(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestSizeMessageProtoText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedSizeMessage(popr, true)
	data := github_com_gogo_protobuf_proto.MarshalTextString(p)
	msg := &SizeMessage{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestSizeMessageProtoCompactText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedSizeMessage(popr, true)
	data := github_com_gogo_protobuf_proto.CompactTextString(p)
	msg := &SizeMessage{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestSizeMessageSize(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedSizeMessage(popr, true)
	size2 := github_com_gogo_protobuf_proto.Size(p)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	size := p.Size()
	if len(data) != size {
		t.Errorf("size %v != marshalled size %v", size, len(data))
	}
	if size2 != size {
		t.Errorf("size %v != before marshal proto.Size %v", size, size2)
	}
	size3 := github_com_gogo_protobuf_proto.Size(p)
	if size3 != size {
		t.Errorf("size %v != after marshal proto.Size %v", size, size3)
	}
}

//These tests are generated by github.com/gogo/protobuf/plugin/testgen
