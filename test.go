package main

import(
	"fmt"
	"log"

	"example/crudl"
)

func main() {
	client, err := crudl.NewClient("http://localhost:8000")
	if err != nil {
		log.Fatalf("Cannot create client: %v\n", err)
	}
	for i := 0; i<10; i++ {
		_, err := client.CreateItem(&crudl.CreateItemRequest{
			Entity: &crudl.Item{
				Id: crudl.ItemId(fmt.Sprintf("item%02d", i)),
				Data: fmt.Sprintf("Item %d!", i),
			},
		})
		if err != nil {
			log.Fatalf("Cannot create item: %v\n", crudl.Pretty(err))
		}
	}
	//    GetItem(req *GetItemRequest) (*GetItemResponse, error)
	//    PutItem(req *PutItemRequest) (*PutItemResponse, error)
	//    DeleteItem(req *DeleteItemRequest) (*DeleteItemResponse, error)
	//	res, err := client.ListItems(&crudl.ListItemsRequest{Limit:5,Skip:"item3"})
	res, err := client.ListItems(&crudl.ListItemsRequest{Limit:5})
	if err != nil {
		log.Fatalf("client got an error: %v\n", err)
	}
	log.Printf("Empty listing: %s", crudl.Pretty(res))
}
