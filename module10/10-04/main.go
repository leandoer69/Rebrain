package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type GoodForSale struct {
	Name  string
	Price int
	Qnty  int
}

type Lead struct {
	ID          int
	Name        string
	ContactName string
	Budget      int
	Items       []GoodForSale
	DateTime    string
	Status      int
	Win         bool
}

func main() {
	lead := &Lead{
		ID:          41,
		Name:        "Phone",
		ContactName: "John Smith",
		Budget:      30000,
		Items: []GoodForSale{
			{
				Name:  "Iphone 6",
				Price: 100,
				Qnty:  3,
			},
			{
				Name:  "Xiaomi Redmi 8",
				Price: 100,
				Qnty:  2,
			},
			{
				Name:  "Samsung Galaxy",
				Price: 200,
				Qnty:  4,
			},
		},
		DateTime: "12.08.2020 11:23:10",
		Status:   150,
		Win:      false,
	}

	//pt := (*[unsafe.Sizeof(Lead{})]byte)(unsafe.Pointer(&lead))
	//fmt.Println(pt)

	idPt := uintptr(unsafe.Pointer(&lead.ID))

	bugetPt := (*int)(unsafe.Pointer(idPt + unsafe.Offsetof(lead.Budget)))
	*bugetPt = 10

	fmt.Println("Change Budget", lead)

	/*
		or:
		idSz := unsafe.Sizeof(lead.ID)
		nameSz := unsafe.Sizeof(lead.Name)
		contactNameSz := unsafe.Sizeof(lead.ContactName)
		bugetPt := (*int)(unsafe.Pointer(idPt + idSz + nameSz + contactNameSz))
	*/

	itemsPt := uintptr(unsafe.Pointer(&lead.Items[0]))
	sliceHeader := &reflect.SliceHeader{
		Data: itemsPt,
		Len:  len(lead.Items),
		Cap:  cap(lead.Items),
	}
	newSlice := *(*[]GoodForSale)(unsafe.Pointer(sliceHeader))
	newSlice = append(newSlice,
		GoodForSale{
			Name:  "Iphone 12",
			Price: 60000,
			Qnty:  2},
		GoodForSale{
			Name:  "IPad 15 Pro",
			Price: 120000,
			Qnty:  3})
	fmt.Println("Changed new slice:\n", newSlice)

	lead.Items = newSlice
	fmt.Println("Change Items", lead)

}
