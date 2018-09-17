package simplecqrs

// Events are just plain structs

// InventoryItemCreated event
type InventoryItemCreated struct {
	ID   string
	Name string
}

// InventoryItemRenamed event
type InventoryItemRenamed struct {
	ID      string
	NewName string
	Original int
}

// InventoryItemDeactivated event
type InventoryItemDeactivated struct {
	ID string
}

// ItemsRemovedFromInventory event
type ItemsRemovedFromInventory struct {
	ID    string
	Count int
}

// ItemsCheckedIntoInventory event
type ItemsCheckedIntoInventory struct {
	ID    string
	Count int
}
