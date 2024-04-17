package calculadora

type Calculadora interface {

	//A partir de una operación en notación polaca inversa, Operar devuelve resultado y un booleano que indica la existencia de un error o no.
	Operar(comando string) (int64, bool)
}
