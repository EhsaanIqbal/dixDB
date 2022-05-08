package lsmt

import "testing"

func TestBinaryTree(t *testing.T) {
	items := []KVPair{
		{Key: "1", Value: "a"},
		{Key: "2", Value: "b"},
		{Key: "3", Value: "c"},
	}

	bt := NewBinaryTree(items)

	// Check tree size
	if bt.Size != len(items) {
		t.Errorf("got tree size %d; expected %d", bt.Size, len(items))
	}

	// Test absent item
	item, err := Search(bt, "99")
	if err == nil {
		t.Errorf("got %v; expected non existant kv pair", item)
	}

	// Test Search
	for _, tcase := range items {
		item, err = Search(bt, tcase.Key)
		if err != nil {
			t.Errorf("search key %s; error: %v", tcase.Key, err)
		}

		if item.Key != tcase.Key || item.Value != tcase.Value {
			t.Errorf("got kv %v; expected %v", item, tcase)
		}
	}

	// Test Insert
	newItem := KVPair{Key: "45", Value: "abc"}
	Insert(&bt, newItem)

	item, err = Search(bt, "45")
	if err != nil {
		t.Errorf("got %v; expected 45, abc", item)
	}

}
