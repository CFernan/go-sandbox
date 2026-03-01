package multimap

import "fmt"

// devuelve el índice de la primera ocurrencia de item en slice.
// Si no existe, devuelve −1
func index(item int, slice []int) int {
	for i, v := range slice {
		if item == v {
			return i
		}
	}

	return -1
}

// elimina el elemento en el índice indicado. Devuelve un error
// si el índice no existe; en otro caso devuelve el slice
// resultante
func remove(index int, slice []int) ([]int, error) {
	if index >= len(slice) || index < 0 {
		return slice, fmt.Errorf(
			"index out of bounds of slice (%v)", len(slice))
	}

	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1], nil
}

// inserta value en el contenedor indexado por key
func Set(key string, value int, multimap map[string][]int) (
	map[string][]int, error,
) {
	if multimap == nil {
		multimap = make(map[string][]int)
	}

	if _, ok := multimap[key]; !ok {
		multimap[key] = []int{}
	}

	multimap[key] = append(multimap[key], value)

	return multimap, nil
}

// devuelve todos los valores indexados por key y ok a true;
// si no hay valores indexados por key devuelve un slice nulo
// y ok=false
func Get(key string, multimap map[string][]int) (
	[]int, bool,
) {
	v, ok := multimap[key]
	if !ok {
		return []int{}, false
	}
	return v, true
}

// elimina value del contenedor indexado por key. Devuelve true
// si efectivamente ha sido eliminado , y falso en otro caso
func Remove(key string, index int, multimap map[string][]int) bool {
	values, ok := multimap[key]
	if !ok {
		return false
	}

	if index >= len(values) {
		return false
	}

	valuesR, err := remove(index, values)
	if err != nil {
		return false
	}

	if len(valuesR) == 0 {
		delete(multimap, key)
	} else {
		multimap[key] = valuesR
	}
	return true
}
