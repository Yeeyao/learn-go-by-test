package main

type Dictionary map[string]string

// 把错误提取为变量，避免 magic error
// 使用 constant error，类似宏定义？

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// 添加元素
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	// 找不到，赋值
	case ErrNotFound:
		d[word] = definition
	// 已经存在了，提示错误
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil

	// d[word] = definition
	// return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
