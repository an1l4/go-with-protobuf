package main

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func toJson(pb proto.Message) string {
	option := protojson.MarshalOptions{
		Multiline: true,
	}
	out, err := option.Marshal(pb)

	if err != nil {
		log.Fatalln("can't convert to json", err)
		return ""
	}

	return string(out)

}

func fromJson(in string, pb proto.Message) {
	option := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	if err := option.Unmarshal([]byte(in), pb); err != nil {
		log.Fatalln("cant unmarshal from json", err)
	}
}
