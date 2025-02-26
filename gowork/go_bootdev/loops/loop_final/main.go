/*
Ejemplo con 5 personas:

Visualización rápida:

	Persona 1 se conecta con 2, 3, 4 y 5 → 4 conexiones
	Persona 2 ya está conectada con 1, así que se conecta con 3, 4 y 5 → 3 más
	Persona 3 se conecta con 4 y 5 → 2 más
	Persona 4 se conecta solo con 5 → 1 más
	Persona 5 ya está conectada con todos

Total = 4 + 3 + 2 + 1 = 10 conexiones
*/
package main

func countConnections(groupSize int) int {
	conections_between_members := 0

	for i := 1; i < groupSize; i++ {
		conections_between_members += i
	}
	return conections_between_members
}
