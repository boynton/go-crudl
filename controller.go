package main

import(
	"fmt"
	"sort"
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
	item := req.Entity
	if item == nil {
		return nil, &crudl.BadRequest{Message: "Bad Request - entity was not a valid Item"}
	}
	key := item.Id
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, ok := c.storage[key]; ok {
		return nil, &crudl.BadRequest{Message: fmt.Sprintf("Already exists: %v", key)}
	}
	item.Modified = &crudl.Timestamp{Time: time.Now()}
	c.storage[key] = item
	return &crudl.CreateItemResponse{
		Entity: item,
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
			Entity: item,
			Modified: item.Modified,
		}, nil
	}
	return nil, &crudl.NotFound{Message: fmt.Sprintf("Item not found: %s", req.Id)}
}

func (c *CrudlController)PutItem(req *crudl.PutItemRequest) (*crudl.PutItemResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
   if _, ok := c.storage[req.Id]; ok {
		item := req.Entity
		item.Modified = &crudl.Timestamp{Time: time.Now()}
		c.storage[req.Id] = item
		return &crudl.PutItemResponse{
			Entity: item,
		}, nil
	}
	return nil, &crudl.NotFound{Message: fmt.Sprintf("Item not found: %s", req.Id)}
}

func (c *CrudlController)DeleteItem(req *crudl.DeleteItemRequest) (*crudl.DeleteItemResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
   if _, ok := c.storage[req.Id]; ok {
		delete(c.storage, req.Id);
		return &crudl.DeleteItemResponse{}, nil
	}
	return nil, &crudl.NotFound{Message: fmt.Sprintf("Item not found: %s", req.Id)}
}

func (c *CrudlController)ListItems(req *crudl.ListItemsRequest) (*crudl.ListItemsResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	var next crudl.ItemId
	var count int32
	limit := req.Limit
	if limit == 0 { //smithy has no default values, to you get 0 here if the parameter is missing.
		limit = 10
	}
	skip := req.Skip
	lst := make([]*crudl.Item, 0)
	if len(c.storage) > 0 {
		keys := make([]string, 0, len(c.storage))
		for k, _ := range c.storage {
			keys = append(keys, string(k))
		}
		sort.Strings(keys)
		for _, ks := range keys {
			k := crudl.ItemId(ks)
			v := c.storage[k]
			if skip != "" {
				if skip != k {
					continue
				}
				skip = ""
			}
			count++
			if count > limit {
				next = k
				break
			}
			lst = append(lst, v)
		}
	}
	return &crudl.ListItemsResponse{
		Entity: &crudl.ItemListing{
			Items: lst,
			Next: next,
		},
	}, nil
}

