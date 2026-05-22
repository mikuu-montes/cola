package cola

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

// Pre condicion: -
// Post condicion: Retorna un puntero a un nodo con el dato enviado y apunta como "siguiente nodo" a nil
func nodoCrear[T any](dato T) *nodo[T] {
	return &nodo[T]{dato: dato, siguiente: nil}
}

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

// Pre condicion: -
// Post condicion: Retorna una cola enlazada.
func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{
		primero: nil,
		ultimo:  nil,
	}
}

// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil && c.ultimo == nil
}

// VerPrimero obtiene el valor del primero de la cola. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}

	return c.primero.dato
}

// Encolar agrega un nuevo elemento a la cola, al final de la misma.
func (c *colaEnlazada[T]) Encolar(valor T) {
	nuevoNodo := nodoCrear(valor)

	if c.EstaVacia() {
		c.primero, c.ultimo = nuevoNodo, nuevoNodo
		return
	}

	c.ultimo.siguiente = nuevoNodo
	c.ultimo = nuevoNodo
}

// Desencolar saca el primer elemento de la cola. Si la cola tiene elementos, se quita el primero de la misma,
// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}

	valor := c.primero.dato
	c.primero = c.primero.siguiente

	if c.primero == nil {
		c.ultimo = nil
	}

	return valor
}
