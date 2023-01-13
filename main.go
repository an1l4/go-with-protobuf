package main

import (
	"fmt"
	"reflect"

	// "github.com/an1l4/go-with-protobuf/proto"
	pb "github.com/an1l4/go-with-protobuf/proto"

	"google.golang.org/protobuf/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "A Name",
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}

}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{Id: 34, Name: "anila"},
		MultipleDummies: []*pb.Dummy{
			{Id: 56, Name: "anu"},
			{Id: 23, Name: "anil"},
			{Id: 12, Name: "john"},
		},
	}
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: 1, //pb.EyeColor_EYE_COLOR_GREEN,
	}
}

func doOneOf(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Errorf("message has unexpected type: %v", x)
	}

}

func doMaps() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myid1": {Id: 1},
			"myid2": {Id: 2},
			"myid3": {Id: 3},
		},
	}

}

func doFile(p proto.Message) {
	path := "simple.bin"

	writeToFile(path, p)
	message := &pb.Simple{}
	readFromFile(path, message)
	fmt.Println(message)
}

func doToJson(p proto.Message) string {
	jsonString:=toJson(p)
	fmt.Println(jsonString)
	return jsonString

}
 
func doFromJson(jsonString string, t reflect.Type) proto.Message {
	message:=reflect.New(t).Interface().(proto.Message)
	fromJson(jsonString,message)
	return message
}

func main() {
	// fmt.Println(doSimple())
	// fmt.Println(doComplex())
	//fmt.Println(doEnum())
	// fmt.Println("this should be an id")
	// doOneOf(&pb.Result_Id{Id: 42})
	// fmt.Println("this should be a message")
	// doOneOf(&pb.Result_Message{Message: "anila"})

	// fmt.Println(doMaps())
	// doFile(doSimple())
	jsonString:=doToJson(doSimple())
	message:=doFromJson(jsonString,reflect.TypeOf(pb.Simple{}))

	fmt.Println(jsonString)
	fmt.Println(message)

	fmt.Println(doFromJson(`{"id":42,"unknown":"test"}`,reflect.TypeOf(pb.Simple{})))


	// jsonString = doToJson(doComplex())
	// message = doFromJson(jsonString,reflect.TypeOf(pb.Complex{}))

	// fmt.Println(jsonString)
	// fmt.Println(message)

}
