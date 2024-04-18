package main

/*
import "fmt"

// Estructura para representar un producto
type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

// Función para añadir un nuevo producto al inventario
func agregarProducto(inventario map[string]Producto, nombre string, precio float64, cantidad int) {
	// Crear un nuevo producto
	nuevoProducto := Producto{Nombre: nombre, Precio: precio, Cantidad: cantidad}

	// Agregar el nuevo producto al inventario
	inventario[nombre] = nuevoProducto

	fmt.Printf("Producto '%s' agregado al inventario.\n", nombre)
}

// Función para actualizar la cantidad disponible de un producto
func actualizarCantidad(inventario map[string]Producto, nombre string, cantidad int) {
	// Verificar si el producto existe en el inventario
	producto, existe := inventario[nombre]
	if existe {
		// Actualizar la cantidad del producto
		producto.Cantidad = cantidad
		inventario[nombre] = producto
		fmt.Printf("Cantidad del producto '%s' actualizada a %d.\n", nombre, cantidad)
	} else {
		fmt.Printf("El producto '%s' no existe en el inventario.\n", nombre)
	}
}

// Función para eliminar un producto del inventario
func eliminarProducto(inventario map[string]Producto, nombre string) {
	// Verificar si el producto existe en el inventario
	_, existe := inventario[nombre]
	if existe {
		// Eliminar el producto del inventario
		delete(inventario, nombre)
		fmt.Printf("Producto '%s' eliminado del inventario.\n", nombre)
	} else {
		fmt.Printf("El producto '%s' no existe en el inventario.\n", nombre)
	}
}

// Función para mostrar el inventario completo
func mostrarInventario(inventario map[string]Producto) {
	fmt.Println("Inventario:")
	for nombre, producto := range inventario {
		fmt.Printf("- Nombre: %s, Precio: %.2f, Cantidad: %d\n", nombre, producto.Precio, producto.Cantidad)
	}
}

func main() {
	inventario := make(map[string]Producto)

	// Agregar productos al inventario
	agregarProducto(inventario, "Camisa", 29.99, 50)
	agregarProducto(inventario, "Pantalón", 39.99, 30)

	// Mostrar el inventario
	mostrarInventario(inventario)

	// Actualizar la cantidad de un producto
	actualizarCantidad(inventario, "Camisa", 40)

	// Mostrar el inventario actualizado
	mostrarInventario(inventario)

	// Eliminar un producto del inventario
	eliminarProducto(inventario, "Pantalón")

	// Mostrar el inventario después de eliminar un producto
	mostrarInventario(inventario)
}*/

//ejercicio 2

import "fmt"

// Estructura para representar una tarea
type Tarea struct {
	Nombre      string
	Descripcion string
	Responsable string
	Estado      string // Estado de la tarea: "pendiente", "en progreso" o "completada"
}

// Función para crear una nueva tarea
func crearTarea(nombre, descripcion, responsable string) *Tarea {
	nuevaTarea := &Tarea{
		Nombre:      nombre,
		Descripcion: descripcion,
		Responsable: responsable,
		Estado:      "pendiente", // Por defecto, la tarea se crea como "pendiente"
	}
	return nuevaTarea
}

// Función para actualizar el estado de una tarea
func actualizarEstado(tarea *Tarea, nuevoEstado string) {
	tarea.Estado = nuevoEstado
}

// Función para mostrar todas las tareas pendientes
func mostrarTareasPendientes(tareas []*Tarea) {
	fmt.Println("Tareas Pendientes:")
	for _, tarea := range tareas {
		if tarea.Estado == "pendiente" {
			fmt.Printf("- Nombre: %s, Descripción: %s, Responsable: %s\n", tarea.Nombre, tarea.Descripcion, tarea.Responsable)
		}
	}
}

func main() {
	// Crear algunas tareas de ejemplo
	tarea1 := crearTarea("Implementar funcionalidad de inicio de sesión", "Implementar la funcionalidad de inicio de sesión en la aplicación web", "Juan")
	tarea2 := crearTarea("Diseñar interfaz de usuario", "Diseñar la interfaz de usuario para la página de perfil del usuario", "María")
	tarea3 := crearTarea("Depurar errores en la base de datos", "Investigar y solucionar errores en la base de datos de producción", "Pedro")

	// Mostrar las tareas pendientes
	tareas := []*Tarea{tarea1, tarea2, tarea3}
	mostrarTareasPendientes(tareas)

	// Actualizar el estado de una tarea
	actualizarEstado(tarea1, "en progreso")
	fmt.Println("\nDespués de actualizar el estado de una tarea:")
	mostrarTareasPendientes(tareas)
}

//ejercicio 3

/*package main

import "fmt"

// Estructura para representar un libro
type Libro struct {
	Titulo     string
	Autor      string
	Genero     string
	Disponible bool // Estado del libro: true si está disponible, false si está prestado
}

// Función para añadir un nuevo libro a la colección
func agregarLibro(coleccion map[string]Libro, titulo, autor, genero string) {
	nuevoLibro := Libro{
		Titulo:     titulo,
		Autor:      autor,
		Genero:     genero,
		Disponible: true, // Por defecto, el libro se añade como disponible
	}
	coleccion[titulo] = nuevoLibro
	fmt.Printf("Libro '%s' agregado a la colección.\n", titulo)
}

// Función para buscar un libro por título
func buscarPorTitulo(coleccion map[string]Libro, titulo string) (Libro, bool) {
	libro, encontrado := coleccion[titulo]
	if !encontrado {
		fmt.Printf("Libro '%s' no encontrado.\n", titulo)
	}
	return libro, encontrado
}

// Función para buscar libros por autor
func buscarPorAutor(coleccion map[string]Libro, autor string) []Libro {
	var librosPorAutor []Libro
	for _, libro := range coleccion {
		if libro.Autor == autor {
			librosPorAutor = append(librosPorAutor, libro)
		}
	}
	return librosPorAutor
}

// Función para actualizar el estado de un libro
func actualizarEstado(coleccion map[string]Libro, titulo string, disponible bool) {
	if libro, encontrado := coleccion[titulo]; encontrado {
		libro.Disponible = disponible
		coleccion[titulo] = libro
		estado := "disponible"
		if !disponible {
			estado = "prestado"
		}
		fmt.Printf("Estado del libro '%s' actualizado a %s.\n", titulo, estado)
	} else {
		fmt.Printf("Libro '%s' no encontrado.\n", titulo)
	}
}

// Función para eliminar un libro de la colección
func eliminarLibro(coleccion map[string]Libro, titulo string) {
	if _, encontrado := coleccion[titulo]; encontrado {
		delete(coleccion, titulo)
		fmt.Printf("Libro '%s' eliminado de la colección.\n", titulo)
	} else {
		fmt.Printf("Libro '%s' no encontrado.\n", titulo)
	}
}

// Función para mostrar la colección completa de libros
func mostrarColeccion(coleccion map[string]Libro) {
	fmt.Println("Colección de Libros:")
	for _, libro := range coleccion {
		estado := "disponible"
		if !libro.Disponible {
			estado = "prestado"
		}
		fmt.Printf("- Título: %s, Autor: %s, Género: %s, Estado: %s\n", libro.Titulo, libro.Autor, libro.Genero, estado)
	}
}

func main() {
	coleccion := make(map[string]Libro)

	// Agregar algunos libros a la colección
	agregarLibro(coleccion, "El principito", "Antoine de Saint-Exupéry", "Fábula")
	agregarLibro(coleccion, "Cien años de soledad", "Gabriel García Márquez", "Realismo mágico")
	agregarLibro(coleccion, "1984", "George Orwell", "Distopía")

	// Mostrar la colección completa
	mostrarColeccion(coleccion)

	// Actualizar el estado de un libro
	actualizarEstado(coleccion, "El principito", false)

	// Mostrar la colección después de actualizar el estado de un libro
	mostrarColeccion(coleccion)

	// Eliminar un libro de la colección
	eliminarLibro(coleccion, "1984")

	// Mostrar la colección después de eliminar un libro
	mostrarColeccion(coleccion)
}*/

//ejercicio 4

/*package main

import "fmt"

// Estructura para representar un contacto
type Contacto struct {
	Nombre    string
	Telefono  string
	Correo    string
	Direccion string
}

// Función para añadir un nuevo contacto
func agregarContacto(agenda map[string]Contacto, nombre, telefono, correo, direccion string) {
	nuevoContacto := Contacto{
		Nombre:    nombre,
		Telefono:  telefono,
		Correo:    correo,
		Direccion: direccion,
	}
	agenda[nombre] = nuevoContacto
	fmt.Printf("Contacto '%s' agregado a la agenda.\n", nombre)
}

// Función para buscar un contacto por nombre
func buscarPorNombre(agenda map[string]Contacto, nombre string) (Contacto, bool) {
	contacto, encontrado := agenda[nombre]
	if !encontrado {
		fmt.Printf("Contacto '%s' no encontrado.\n", nombre)
	}
	return contacto, encontrado
}

// Función para buscar un contacto por número de teléfono
func buscarPorTelefono(agenda map[string]Contacto, telefono string) []Contacto {
	var contactosPorTelefono []Contacto
	for _, contacto := range agenda {
		if contacto.Telefono == telefono {
			contactosPorTelefono = append(contactosPorTelefono, contacto)
		}
	}
	return contactosPorTelefono
}

// Función para actualizar la información de un contacto
func actualizarContacto(agenda map[string]Contacto, nombre, telefono, correo, direccion string) {
	if _, encontrado := agenda[nombre]; encontrado {
		contactoActualizado := Contacto{
			Nombre:    nombre,
			Telefono:  telefono,
			Correo:    correo,
			Direccion: direccion,
		}
		agenda[nombre] = contactoActualizado
		fmt.Printf("Contacto '%s' actualizado en la agenda.\n", nombre)
	} else {
		fmt.Printf("Contacto '%s' no encontrado.\n", nombre)
	}
}

// Función para eliminar un contacto de la agenda
func eliminarContacto(agenda map[string]Contacto, nombre string) {
	if _, encontrado := agenda[nombre]; encontrado {
		delete(agenda, nombre)
		fmt.Printf("Contacto '%s' eliminado de la agenda.\n", nombre)
	} else {
		fmt.Printf("Contacto '%s' no encontrado.\n", nombre)
	}
}

// Función para mostrar la agenda completa de contactos
func mostrarAgenda(agenda map[string]Contacto) {
	fmt.Println("Agenda de Contactos:")
	for _, contacto := range agenda {
		fmt.Printf("- Nombre: %s, Teléfono: %s, Correo: %s, Dirección: %s\n", contacto.Nombre, contacto.Telefono, contacto.Correo, contacto.Direccion)
	}
}

func main() {
	agenda := make(map[string]Contacto)

	// Agregar algunos contactos a la agenda
	agregarContacto(agenda, "Juan Pérez", "123456789", "juan@example.com", "Calle 123")
	agregarContacto(agenda, "María López", "987654321", "maria@example.com", "Avenida 456")

	// Mostrar la agenda completa
	mostrarAgenda(agenda)

	// Actualizar la información de un contacto
	actualizarContacto(agenda, "Juan Pérez", "123456789", "juan@gmail.com", "Calle 1234")

	// Mostrar la agenda después de actualizar un contacto
	mostrarAgenda(agenda)

	// Eliminar un contacto de la agenda
	eliminarContacto(agenda, "María López")

	// Mostrar la agenda después de eliminar un contacto
	mostrarAgenda(agenda)
}*/
