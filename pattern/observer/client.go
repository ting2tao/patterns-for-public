package observer

type ObInterface interface {
	ObserverPattern()
}

func ObserverClient() {
	shirtItem := NewItem("Nike Shirt")

	observerFirst := &Customer{id: "sct@gmail.com"}
	observerSecond := &Customer{id: "hwt@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}
