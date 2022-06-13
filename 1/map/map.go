package main

func main() {
	// Create a map
	//var m map[string]int									// -> bad way
	//m := map[string]int{}
	//m := make(map[string]int)
	//m := map[string]int{
	//	"key1": 1,
	//	"key2": 2,
	//}
	//fmt.Printf("Map: %v\n", m)

	//////////////////////////////

	// Add a new key/value pair
	//m["newKey"] = 10
	//fmt.Printf("Map: %v\n", m)

	//////////////////////////////

	// Get value by key
	//key := "newKey"
	//val := m[key]
	//fmt.Printf("Value for '%v': %v\n", key, val)

	//key = "notExistingKey"
	//val = m[key]
	//fmt.Printf("Value for '%v': %v\n", key, val)

	//////////////////////////////

	// Check that value exists
	//fmt.Printf("Map: %v\n", m)

	//key := "key1"
	//val, exists := m[key]
	//fmt.Printf("key: %v -> val: %v. exists: %v\n", key, val, exists)

	//key = "notExistingKey"
	//val, exists = m[key]
	//fmt.Printf("key: %v -> val: %v. exists: %v\n", key, val, exists)

	//////////////////////////////

	// Update value by key
	//m["newKey"] = -5
	//fmt.Printf("Map: %v\n", m)

	//////////////////////////////

	// Delete a key/value pair
	//delete(m, "newKey")
	//fmt.Printf("Map: %v\n", m)

	//////////////////////////////

	// Go through all elements
	//for k, v := range m {
	//	fmt.Printf("key: %v -> value: %v\n", k, v)
	//}

	//////////////////////////////

	// Get length of the map
	//length := len(m)
	//fmt.Printf("Length of the map: %v\n", length)

	//////////////////////////////

	// Using in functions
	//m = map[string]int{
	//	"key1": 1,
	//	"key2": 2,
	//	"key3": 3,
	//	"key4": 4,
	//	"key5": 5,
	//}
	//fmt.Printf("Map before:\n%v\n", m)
	//update(m)
	//fmt.Printf("Map after:\n%v\n", m)
}

func update(m map[string]int) {
	m["key1"] = -12
	m["newKey"] = 0
}
