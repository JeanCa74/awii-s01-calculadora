package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cliente struct {
	ID      int
	Nombre  string
	Carrera string
	Saldo   float64
}

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

type Pedido struct {
	ID         int
	ClienteID  int
	ProductoID int
	Cantidad   int
	Total      float64
	Fecha      string
}

func main() {
	clientes := []Cliente{
		{1, "Ana Vera", "Agropecuaria", 100.0},
		{2, "Mayra Zamora", "TI", 50.0},
		{3, "Liliana Lopez", "Educacion", 300.0},
	}
	productos := []Producto{
		{1, "cafe", 1.00, 150, "Bebida"},
		{2, "empanada", 0.50, 50, "Bocadito"},
		{3, "sanduche", 1.50, 50, "Refrigerio"},
	}
	var pedidos []Pedido

	lector := bufio.NewReader(os.Stdin)

	for {
		mostrarMenu()
		opcion := leerEntero(lector, "Elige una opción: ")

		switch opcion {
		case 1:
			ListarClientes(clientes)
		case 2:
			ListarProductos(productos)
		case 3:
			fmt.Println("\n--- NUEVO CLIENTE ---")
			id := leerEntero(lector, "ID: ")
			fmt.Print("Nombre: ")
			nombre := leerLinea(lector)
			fmt.Print("Carrera: ")
			carrera := leerLinea(lector)
			saldo := leerFloat(lector, "Saldo $: ")

			clientes = AgregarCliente(clientes, Cliente{id, nombre, carrera, saldo})
			fmt.Println("¡Cliente agregado!")

		case 4:
			fmt.Println("\n--- NUEVO PRODUCTO ---")
			id := leerEntero(lector, "ID: ")
			fmt.Print("Nombre: ")
			nombre := leerLinea(lector)
			precio := leerFloat(lector, "Precio $: ")
			stock := leerEntero(lector, "Stock inicial: ")
			fmt.Print("Categoría: ")
			categoria := leerLinea(lector)

			productos = AgregarProducto(productos, Producto{id, nombre, precio, stock, categoria})
			fmt.Println("¡Producto agregado!")

		case 5:
			fmt.Println("\n--- REGISTRAR PEDIDO ---")
			cID := leerEntero(lector, "ID del Cliente: ")
			pID := leerEntero(lector, "ID del Producto: ")
			cant := leerEntero(lector, "Cantidad: ")
			fmt.Print("Fecha (Ej. 2026-04-16): ")
			fecha := leerLinea(lector)

			nuevosPedidos, err := RegistrarPedido(clientes, productos, pedidos, cID, pID, cant, fecha)
			if err != nil {
				fmt.Println("Error al procesar:", err)
			} else {
				pedidos = nuevosPedidos
				fmt.Println("¡Pedido registrado con éxito! Saldos y stock actualizados.")
			}

		case 6:
			fmt.Println("\n--- REPORTE DE CLIENTE ---")
			cID := leerEntero(lector, "Ingrese el ID del Cliente: ")
			PedidosDeCliente(pedidos, clientes, productos, cID)

		case 0:
			fmt.Println("Cerrando el sistema... ¡Hasta luego!")
			return

		default:
			fmt.Println("Opción no válida. Por favor, escriba un número del 0 al 6.")
		}
	}
}

func ListarClientes(clientes []Cliente) {
	if len(clientes) == 0 {
		fmt.Println("No hay clientes Registrados")
		return
	}

	fmt.Println("\n---------> LISTA DE CLIENTES <----------")
	fmt.Printf("%-5s | %-20s | %-15s | %s\n", "ID", "NOMBRE", "CARRERA", "SALDO")
	fmt.Println("----------------------------------------------------------------")

	for _, c := range clientes {
		fmt.Printf("[%-3d] %-20s | %-15s | $%.2f\n", c.ID, c.Nombre, c.Carrera, c.Saldo)
	}

}

func BuscarClientesPorID(clientes []Cliente, id int) int {
	for i, c := range clientes {
		if c.ID == id {
			return i
		}
	}
	return -1
}

func AgregarCliente(clientes []Cliente, nuevo Cliente) []Cliente {
	return append(clientes, nuevo)
}

func EliminarCliente(clientes []Cliente, id int) []Cliente {
	idx := BuscarClientesPorID(clientes, id)
	if idx == -1 {
		fmt.Println("Error. Cliente ingresado, no encontrado")
	}
	return append(clientes[:idx], clientes[idx+1:]...)
}

func BuscarProductoPorID(productos []Producto, id int) int {
	for i, p := range productos {
		if p.ID == id {
			return i
		}
	}
	return -1
}

func AgregarProducto(productos []Producto, nuevo Producto) []Producto {
	return append(productos, nuevo)
}

func EliminarProducto(productos []Producto, id int) []Producto {
	idx := BuscarProductoPorID(productos, id)
	if idx == -1 {
		fmt.Println("No se encontro el producto que desea eliminar")
		return productos
	}
	return append(productos[:idx], productos[idx+1:]...)
}

func ListarProductos(productos []Producto) {
	if len(productos) == 0 {
		fmt.Println("No hay productos registrados")
		return
	}
	fmt.Println("\n---------> LISTA DE PRODUCTOS <----------")
	fmt.Printf("%-5s | %-20s | %-10s | %-6s | %s\n", "ID", "NOMBRE", "CATEGORIA", "STOCK", "PRECIO")
	fmt.Println("----------------------------------------------------------------")

	for _, p := range productos {
		fmt.Printf("[%-3d] %-20s | %-10s | %-6d | $%.2f\n", p.ID, p.Nombre, p.Categoria, p.Stock, p.Precio)
	}
}

func DescontarSaldo(cliente *Cliente, monto float64) error {
	if monto <= 0 {
		return errors.New("el monto a descontar debe ser mayor a cero")
	}
	if cliente.Saldo < monto {
		return errors.New("saldo insuficiente")
	}

	cliente.Saldo = cliente.Saldo - monto
	return nil
}

func DescontarStock(producto *Producto, cantidad int) error {
	if cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}
	if producto.Stock < cantidad {
		return errors.New("stock insuficiente")
	}
	producto.Stock = producto.Stock - cantidad
	return nil
}

func RegistrarPedido(clientes []Cliente, productos []Producto, pedidos []Pedido, clienteID int, productoID int, cantidad int, fecha string) ([]Pedido, error) {

	idxC := BuscarClientesPorID(clientes, clienteID)
	if idxC == -1 {
		return pedidos, errors.New("cliente no encontrado")
	}

	idxP := BuscarProductoPorID(productos, productoID)
	if idxP == -1 {
		return pedidos, errors.New("producto no encontrado")
	}

	total := productos[idxP].Precio * float64(cantidad)

	err := DescontarStock(&productos[idxP], cantidad)
	if err != nil {
		return pedidos, err
	}

	err = DescontarSaldo(&clientes[idxC], total)
	if err != nil {
		productos[idxP].Stock += cantidad
		return pedidos, err
	}
	nuevoPedido := Pedido{
		ID:         len(pedidos) + 1,
		ClienteID:  clienteID,
		ProductoID: productoID,
		Cantidad:   cantidad,
		Total:      total,
		Fecha:      fecha,
	}

	pedidos = append(pedidos, nuevoPedido)
	return pedidos, nil
}

func leerLinea(lector *bufio.Reader) string {
	linea, _ := lector.ReadString('\n')
	return strings.TrimSpace(linea)
}

func leerEntero(lector *bufio.Reader, prompt string) int {
	fmt.Print(prompt)
	texto := leerLinea(lector)
	n, err := strconv.Atoi(texto)
	if err != nil {
		return -1
	}
	return n
}

func leerFloat(lector *bufio.Reader, prompt string) float64 {
	fmt.Print(prompt)
	texto := leerLinea(lector)
	f, err := strconv.ParseFloat(texto, 64)
	if err != nil {
		return -1.0
	}
	return f
}

func mostrarMenu() {
	fmt.Println("\n=================================")
	fmt.Println("    MINI-CAFETERÍA ULEAM     ")
	fmt.Println("=================================")
	fmt.Println("1. Listar clientes")
	fmt.Println("2. Listar productos")
	fmt.Println("3. Agregar cliente")
	fmt.Println("4. Agregar producto")
	fmt.Println("5. Registrar pedido")
	fmt.Println("6. Ver historial de cliente (Reporte)")
	fmt.Println("0. Salir")
	fmt.Println("=================================")
}

func PedidosDeCliente(pedidos []Pedido, clientes []Cliente, productos []Producto, clienteID int) {

	idxC := BuscarClientesPorID(clientes, clienteID)
	if idxC == -1 {
		fmt.Println("Error: Cliente no encontrado.")
		return
	}

	cliente := clientes[idxC]
	fmt.Printf("\n=== HISTORIAL DE PEDIDOS: %s ===\n", cliente.Nombre)

	tienePedidos := false
	var totalGastado float64

	for _, ped := range pedidos {
		if ped.ClienteID == clienteID {
			tienePedidos = true

			idxP := BuscarProductoPorID(productos, ped.ProductoID)
			nombreProd := "Producto Borrado"
			if idxP != -1 {
				nombreProd = productos[idxP].Nombre
			}

			fmt.Printf("Pedido #%-2d | %-15s | Cant: %-2d | Fecha: %s | Total: $%.2f\n", ped.ID, nombreProd, ped.Cantidad, ped.Fecha, ped.Total)
			totalGastado += ped.Total
		}
	}

	if !tienePedidos {
		fmt.Println("Este cliente aún no ha realizado compras.")
	} else {
		fmt.Println("-------------------------------------------------------------------")
		fmt.Printf("TOTAL ACUMULADO GASTADO: $%.2f\n", totalGastado)
	}
}
