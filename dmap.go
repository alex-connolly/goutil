package goutil

// A DMap is a deterministic map
type DMap struct {
	keyMap map[string]interface{}
	idxMap map[string]int
	array  []interface{}
	index  int
}

// Next ...
func (d *DMap) Next() interface{} {
	val := d.array[d.index]
	d.index++
	return val
}

// Get ...
func (d *DMap) Get(key string) interface{} {
	if d.keyMap == nil {
		return nil
	}
	return d.keyMap[key]
}

// Add ...
func (d *DMap) Add(key string, value interface{}) {
	if d.keyMap == nil {
		d.keyMap = make(map[string]interface{})
		d.idxMap = make(map[string]int)
		d.array = make([]interface{}, 0)
	}

	_, ok := d.keyMap[key]
	if ok {
		// already in map
		d.Remove(key)
		d.Add(key, value)
	} else {
		d.keyMap[key] = value
		d.idxMap[key] = len(d.array)
		d.array = append(d.array, value)
	}
}

// Remove ...
func (d *DMap) Remove(key string) {
	idx, ok := d.idxMap[key]
	if ok {
		d.array = append(d.array[:idx], d.array[idx+1:]...)
		delete(d.keyMap, key)
		delete(d.idxMap, key)
	}
}

// Length ...
func (d *DMap) Length() int {
	return len(d.array)
}

// Map ...
func (d *DMap) Map() map[string]interface{} {
	return d.keyMap
}

// Array ...
func (d *DMap) Array() []interface{} {
	return d.array
}
