// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example.proto

package example

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ExampleMessage) Validate() error {
	if !(this.HelloId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("HelloId", fmt.Errorf(`value '%v' must be greater than '0'`, this.HelloId))
	}
	return nil
}
func (this *Empty) Validate() error {
	return nil
}
