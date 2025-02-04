package app

type Nodo struct {
	Persona  Persona
	Izquierdo *Nodo
	Derecho   *Nodo
}

type ArbolBinario struct {
	Raiz *Nodo
}

func (a *ArbolBinario) Insertar(persona Persona) {
	a.Raiz = insertarNodo(a.Raiz, persona)
}

func insertarNodo(nodo *Nodo, persona Persona) *Nodo {
	if nodo == nil {
		return &Nodo{Persona: persona}
	}
	if persona.Id < nodo.Persona.Id {
		nodo.Izquierdo = insertarNodo(nodo.Izquierdo, persona)
	} else {
		nodo.Derecho = insertarNodo(nodo.Derecho, persona)
	}
	return nodo
}


func (a *ArbolBinario) Eliminar(id int) {
	a.Raiz = eliminarNodo(a.Raiz, id)
}

func eliminarNodo(nodo *Nodo, id int) *Nodo {
	if nodo == nil {
		return nil
	}
	if id < nodo.Persona.Id {
		nodo.Izquierdo = eliminarNodo(nodo.Izquierdo, id)
	} else if id > nodo.Persona.Id {
		nodo.Derecho = eliminarNodo(nodo.Derecho, id)
	} else {
		if nodo.Izquierdo == nil {
			return nodo.Derecho
		} else if nodo.Derecho == nil {
			return nodo.Izquierdo
		}
	
		sucesor := encontrarMin(nodo.Derecho)
		nodo.Persona = sucesor.Persona
		nodo.Derecho = eliminarNodo(nodo.Derecho, sucesor.Persona.Id)
	}
	return nodo
}


func encontrarMin(nodo *Nodo) *Nodo {
	actual := nodo
	for actual.Izquierdo != nil {
		actual = actual.Izquierdo
	}
	return actual
}

func (a *ArbolBinario) ObtenerTodos() []Persona {
	var personas []Persona
	recorrerInOrder(a.Raiz, &personas)
	return personas
}


func recorrerInOrder(nodo *Nodo, lista *[]Persona) {
	if nodo != nil {
		recorrerInOrder(nodo.Izquierdo, lista)
		*lista = append(*lista, nodo.Persona)
		recorrerInOrder(nodo.Derecho, lista)
	}
}
