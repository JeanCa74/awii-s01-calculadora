package cafeteria

import (
	"errors"
)

// errores
var (
	ErrClienteNoEncontrado  = errors.New("cliente no encontrado")
	ErrProductoNoEncontrado = errors.New("producto no encontrado")
	ErrStockInsuficiente    = errors.New("stock insuficiente")
	ErrSaldoInsuficiente    = errors.New("saldo insuficiente del cliente")
)

// tipos
type Categoria struct {
	ID     int
	Nombre string
}

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
	Categoria Categoria
}

type Pedido struct {
	ID       int
	Cliente  Cliente
	Producto Producto
	Cantidad int
	Total    float64
	Fecha    string
}

// INTERFAZ REPOSITORIO
type Repositorio interface {
	GuardarCliente(cliente Cliente) error
	ObtenerCliente(id int) (Cliente, error)
	ListarClientes() []Cliente
	GuardarProducto(producto Producto) error
	ObtenerProducto(id int) (Producto, error)
	ListarProductos() []Producto
}

// RepoMemoria
type RepoMemoria struct {
	clientes  []Cliente
	productos []Producto
	pedidos   []Pedido
}

// Constructor
func NewRepoMemoria() *RepoMemoria {
	return &RepoMemoria{
		clientes:  []Cliente{},
		productos: []Producto{},
		pedidos:   []Pedido{},
	}
}

// metodos para Clientes
func (r *RepoMemoria) GuardarCliente(c Cliente) error {
	r.clientes = append(r.clientes, c)
	return nil
}

func (r *RepoMemoria) ObtenerCliente(id int) (Cliente, error) {
	for _, c := range r.clientes {
		if c.ID == id {
			return c, nil
		}
	}
	return Cliente{}, ErrClienteNoEncontrado
}

func (r *RepoMemoria) ListarClientes() []Cliente {
	return r.clientes
}

//metodos para productos

func (r *RepoMemoria) GuardarProducto(p Producto) error {
	r.productos = append(r.productos, p)
	return nil
}

func (r *RepoMemoria) ObtenerProducto(id int) (Producto, error) {
	for _, p := range r.productos {
		if p.ID == id {
			return p, nil
		}
	}
	return Producto{}, ErrProductoNoEncontrado
}

func (r *RepoMemoria) ListarProductos() []Producto {
	return r.productos
}

// Verificación
var _ Repositorio = (*RepoMemoria)(nil)
