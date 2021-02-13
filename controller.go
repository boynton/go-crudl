package main

import(
	"fmt"
	"sync"
	"time"

	"example/crudl"
)

type CrudlController struct {
	storage map[crudl.ItemId]*crudl.Item
	mutex sync.Mutex
}

func NewController() *CrudlController {
	return &CrudlController{
		storage: make(map[crudl.ItemId]*crudl.Item, 0),
	}
}

func (c *CrudlController) CreateItem(req *crudl.CreateItemRequest) (*crudl.CreateItemResponse, error) {
	item := req.Item
	if item == nil {
		return nil, &crudl.BadRequest{Message: "Bad Request - entity was not a valid Item"}
	}
	key := item.Id
	c.mutex.Lock()
	if _, ok := c.storage[key]; ok {
		return nil, &crudl.BadRequest{Message: fmt.Sprintf("Already exists: %v", key)}
	}
	item.Modified = &crudl.Timestamp{Time: time.Now()}
	c.storage[key] = item
	c.mutex.Unlock()
	return &crudl.CreateItemResponse{
		Item: item,
	}, nil
}

func (c *CrudlController) GetItem(req *crudl.GetItemRequest) (*crudl.GetItemResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if item, ok := c.storage[req.Id]; ok {
		if req.IfNewer != nil {
			if (item.Modified.Time).After(req.IfNewer.Time) {
				return nil, &crudl.NotModified{Message: "Item not modified"}
			}
		}
		return &crudl.GetItemResponse{
			Item: item,
			Modified: item.Modified,
		}, nil
	}
	return nil, &crudl.NotFound{Message: fmt.Sprintf("Item not found: %s", req.Id)}
}

func (c *CrudlController)PutItem(req *crudl.PutItemRequest) (*crudl.PutItemResponse, error) {
	return nil, nil
}

func (c *CrudlController)DeleteItem(req *crudl.DeleteItemRequest) (*crudl.DeleteItemResponse, error) {
	return nil, nil
}

func (c *CrudlController)ListItems(req *crudl.ListItemsRequest) (*crudl.ListItemsResponse, error) {
	return nil, nil
}

