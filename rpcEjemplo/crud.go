package main

import "fmt"

type Item struct {
	title string
	body  string
}

var database []Item

func GetByName(title string) Item {
	var getItem Item

	for _, val := range database {
		if val.title == title {
			getItem = val
		}
	}

	return getItem

}

func CreateItem(item Item) Item {
	database = append(database, item)
	return item
}

func AddItem(item Item) Item {
	database = append(database, item)
	return item
}

func EditItem(title string, edit Item) Item {
	var changed Item
	for idx, val := range database {
		if val.title == title {
			database[idx] = edit
			changed = edit
		}
	}
	return changed
}

func DeleteItem(item Item) Item {
	var del Item

	for idx, val := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}

	}

	return del
}

func main() {

	fmt.Println("Base de datos inicial: ", database)
	a := Item{"primero", "item de prueba"}
	b := Item{"segundo", "segundo item"}
	c := Item{"tercero", "tercer item"}

	AddItem(a)
	AddItem(b)
	AddItem(c)
	fmt.Println("segunda version base ", database)

	DeleteItem(b)
	fmt.Println("tercera version base ", database)

	EditItem("tercero", Item{"cuarto", "item nuevo"})
	fmt.Println("cuarta version base ", database)

	x := GetByName("cuarto")
	y := GetByName("primero")
	fmt.Println(x, y)

}
