ACCION 19 es
AMBIENTE
	n, div: entero
PROCESO
	escribir("ingrese un numero para saber si es primo: ");
	leer(n);
	div:= 2;
	mientras (div < n) y (n mod div <> 0) hacer
		div:= div + 1;
	fmientras;
	si (div = n-1) o (div = n) entonces
		escribir("El numero", n, "es primo.");
	sino
		escribir("El numero", n, "no es primo.");
	fsi;
FINACCION.

ACCION cond6 es
AMBIENTE
	nom: AN(50)
	sexo: ("F","M");
PROCESO
	escribir("Este es un algoritmo que te permite saber a que grupo de alumnos perteneces segun tu nombre y sexo");
	escribir("Ingrese su nombre: "); leer(nom);
	escribir("Ingrese su sexo (F/M): "); leer(sexo);
	si sexo = "M" entonces
		si nombre > "N" entonces
			escribir("Su grupo es el "A".");
		sino
			escribir("Su grupo es el "B".");
	sino
		si nombre < "M" entonces
			escribir("Su grupo es el "A".");
		sino
			escribir("Su grupo es el "B".");
		fsi;
	fsi;
FINACCION.

ACCION cond8 es
AMBIENTE
	punt: real
PROCESO
	escribir("Este algoritmo indica el nivel de un empleado basado en su puntuacion.");
	escribir("Ingrese su puntaje: ");
	leer(punt);
	segun punt hacer
		0.0:
			escribir("Su nivel es INACEPTABLE.");
			escribir("Su cantidad de dinero conseguido es de: ", 2400*punt);
		0.4:
			escribir("Su nivel es ACEPTABLE.");
			escribir("Su cantidad de dinero conseguido es de: ", 2400*punt);
		>=0.6:
			escribir("Su nivel es MERITORIO.");
			escribir("Su cantidad de dinero conseguido es de: ", 2400*punt);
		otro:
			escribir("Puntuacion no valida.");
	fsegun;
FINACCION.

ACCION 32.04 es
AMBIENTE
	clave, n: entero
	opcion: ("S", "N")
PROCESO
	escribir("Este algoritmo devuelve una clave segun un numero ingresado por el usuario.");
	opcion:= "S";
	mientras opcion = "S" hacer
		escribir("Ingrese un numero: ");
		leer(n);
		si n < 0 entonces
			clave:= -1
		sino
			si n >= 0 entonces
				clave:= 0;
				mientras n > 0 hacer
					clave:= clave + (n div 10);
					n:= n div 10;
				fmientras;
				clave:= clave mod 7;
			fsi;
		fsi;
		escribir("Su clave es: ", clave);
		escribir("Seguir? ('S', 'N'): "); leer(opcion);
	fmientras;
FINACCION.

ACCION 32.06 es
AMBIENTE
	unidades: entero
	precio: real
	opcion: ("S", "N")
PROCESO
	opcion:= "S";
	mientras opcion = "S" hacer	
		escribir("Ingrese el precio del producto: ");
		leer(precio);
		escribir("Ingrese las unidades a comprar: ");
		leer(unidades);
		segun unidades hacer
			>6: escribir("Descuento de 4%: " (precio*unidades)*0.04);
			>12: escribir("Descuento de 10%: ", (precio*unidades)*0.10);
			otro: escribir("Descuento de 0%." precio*unidades);
		fsegun;
		escribir("Seguir? (S/N): ");
		leer(opcion);
	fmientras;
FINACCION.
