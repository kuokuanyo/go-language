type tree struct {
	value int
	//S不可宣告型別同為S的欄位，但可宣告*S指標型別的欄位
	left, right *tree
}

//原址排序
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

//依序插入元素
//回傳結果slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		//等同回傳&tree
		//new()建立一個型別，回傳型別的位址
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}