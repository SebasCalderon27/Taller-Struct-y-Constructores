package main

import "fmt"

type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func NuevoProducto(id int, name string, prec float64, cant int) *Producto {
	return &Producto{
		ID:       id,
		Nombre:   name,
		Precio:   prec,
		Cantidad: cant,
	}

}

func (p *Producto) ActualizarPrecio(NuevoPrecio float64) {
	p.Precio = NuevoPrecio
}

func (p *Producto) ActualizarCantidad(NuevaCantidad int) {
	p.Cantidad = NuevaCantidad
}

func (p Producto) MostrarInfo() {
	fmt.Println("El ID de el producto es: ", p.ID)

	fmt.Println("El NOMBRE de el producto es: ", p.Nombre)

	fmt.Println("El PRECIO de el producto es: ", p.Precio)

	fmt.Println("El CANTIDAD de el producto es: ", p.Cantidad)
}

type Inventario struct {
	Productos map[int]Producto
}

func NuevoInventario() Inventario {
	return Inventario{
		Productos: map[int]Producto{},
	}
}

func (inve *Inventario) AgregarProducto(p Producto) {
	p.ID += 1
	inve.Productos[p.ID] = p
}

func (inve *Inventario) EliminarProducto(p Producto) {
	p, existencia := inve.Productos[p.ID]
	if existencia == true {
		delete(inve.Productos, p.ID)
	} else {
		fmt.Println("El producto no se encontro")
	}
}

func (inve *Inventario) BuscarProducto(p Producto) {
	p, existencia := inve.Productos[p.ID]
	if existencia == true {
		fmt.Println(inve.Productos[p.ID])

	} else {
		fmt.Println("El producto no se encontro")
	}
}

