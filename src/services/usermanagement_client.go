package protobuf_example_go

import (
	pb "github.com/jayatwork/protobuf-example-go/src/services"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect %v:", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_users = make(map[string]int32)
	new_users["Ericka"] = 34
	new_users["Jason"] = 42

	for name, age := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("Could not create new user: %v", err)
		}
		log.Printf(`User Detail: 
		NAME: %s
		AGE:	%d
		ID:		%d`, r.GetName(), r.GetAge(), r.GetId())
	}

}
