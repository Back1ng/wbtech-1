package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.

// Laptop У ноутбука нет места для USB флешки, но есть свободный Type-C.
// Для этого нам пригодится адаптер USB - Type-C
type Laptop struct {
}

func (l Laptop) ReadData(connector TypeCConnector) {
	fmt.Println(connector.Read())
}

type TypeCConnector interface {
	Read() string
}

type TypeCFlashDrive struct {
	Data string
}

func (c TypeCFlashDrive) Read() string {
	return c.Data
}

type usbToTypeCConvertor struct {
	flash UsbFlashDrive
}

func (c usbToTypeCConvertor) Read() string {
	return c.flash.Data
}

type UsbFlashDrive struct {
	Data string // Some data on flash storage
}

func main() {
	laptop := Laptop{}
	usb := UsbFlashDrive{
		Data: "some useful data",
	}
	typeC := TypeCFlashDrive{
		Data: "another useful data",
	}

	convertor := usbToTypeCConvertor{
		flash: usb,
	}

	laptop.ReadData(typeC)
	laptop.ReadData(convertor)
}
