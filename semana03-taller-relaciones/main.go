package main

import (
	"errors"
	"fmt"
	"semana03-taller-relaciones/internal/cafeteria"
)

func main() {
	//repositorio por interfaz
	var repo cafeteria.Repositorio = cafeteria.NewRepoMemoria()

	//guardar 2 clientes y 3 productos
	repo.GuardarCliente(cafeteria.Cliente{ID: 1, Nombre: "Ana Torres", Carrera: "TI", Saldo: 15.00})
	repo.GuardarCliente(cafeteria.Cliente{ID: 2, Nombre: "Luis Vera", Carrera: "Civil", Saldo: 8.50})

	catBebidas := cafeteria.Categoria{ID: 1, Nombre: "Bebidas"}
	catSnacks := cafeteria.Categoria{ID: 2, Nombre: "Snacks"}

	repo.GuardarProducto(cafeteria.Producto{ID: 1, Nombre: "Café", Precio: 1.25, Stock: 20, Categoria: catBebidas})
	repo.GuardarProducto(cafeteria.Producto{ID: 2, Nombre: "Empanada", Precio: 0.75, Stock: 15, Categoria: catSnacks})
	repo.GuardarProducto(cafeteria.Producto{ID: 3, Nombre: "Jugo natural", Precio: 1.00, Stock: 12, Categoria: catBebidas})

	//imprimir un cliente que existe y uno que no
	fmt.Println("--- BUSCANDO CLIENTES ---")
	
	// Cliente que si
	c1, err := repo.ObtenerCliente(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Cliente encontrado: %s (Saldo: $%.2f)\n", c1.Nombre, c1.Saldo)
	}

	// Cliente que no
	_, errNoExiste := repo.ObtenerCliente(99)
	if errNoExiste != nil {
		if errors.Is(errNoExiste, cafeteria.ErrClienteNoEncontrado) {
			fmt.Println("Error esperado: El cliente 99 no fue encontrado en el sistema.")
		} else {
			fmt.Println("Error inesperado:", errNoExiste)
		}
	}

	//lista de productos
	fmt.Println("\n--- LISTA DE PRODUCTOS ---")
	for _, p := range repo.ListarProductos() {
		fmt.Printf("[%d] %s - $%.2f (Categoría: %s)\n", p.ID, p.Nombre, p.Precio, p.Categoria.Nombre)
	}

	fmt.Println("\n--- CREANDO UN PEDIDO ---")
	
	// simulacion de pedido
	p1, _ := repo.ObtenerProducto(1)

	pedido := cafeteria.Pedido{
		ID:       1,
		Cliente:  c1, 
		Producto: p1,
		Cantidad: 2,
		Total:    p1.Precio * 2,
		Fecha:    "2026-04-23",
	}

	//imprimir
	fmt.Printf("Pedido #%d: %s compró %d x %s. Total a pagar: $%.2f\n",
		pedido.ID, pedido.Cliente.Nombre, pedido.Cantidad, pedido.Producto.Nombre, pedido.Total)
}