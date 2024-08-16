type HasName interface {
	GetName() string
}

function SayHello(hasName HasName){
	fmt.Println("Hello", hasName.GetName())
}