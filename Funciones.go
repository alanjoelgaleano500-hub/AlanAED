ACCION Funciones es
AMBIENTE
	a, b, c, alt1, alt2, mult, nomult, cantf, x, y: entero
	seguir: ('V', 'F')
	opcion: 1..2
	procedimiento coeficientes() es
		Escribir("Ingrese los valores de los coeficientes a,b,c en orden:");
		Escribir("Valor a:"); Leer(a);
		Escribir("Valor b:"); Leer(b);
		Escribir("Valor c:"); Leer(c);
	fprocedimiento.
	funcion es_multiplo(f: entero): booleano es
	AMBIENTE
		cont: entero
	PROCESO
		mientras f > 0 hacer
			si cont mod 2 = 0 entonces
				alt1:= alt1 + f MOD 10;
				cont:= cont + 1;
			sino
				alt2:= alt 2 + f MOD 10;
				cont:= cont + 1;
			fsi;
			f:= f DIV 10;
		fmientras;
		si (ABSO(alt1 - alt2) MOD 11 = 0) entonces
			es_multiplo:= Verdadero;
		sino
			es_multiplo:= Falso;
		fsi;
	ffuncion.
	procedimiento resultados() es
		Escribir("Las coordenadas (x, y) son iguales a: ", x, y);
		si es_multiplo(ABSO(y)) entonces
			Escribir("El numero es multiplo de 11");
			mult:= mult + 1;
		sino
			Escribir("El numero NO es multiplo de 11");
			nomult:= nomult + 1;
		fsi;
	fprocedimiento.
	procedimiento multiplos() es
		Escribir("Los y multiplos de 11 son: ", mult);
		Escribir("Los y NO multiplos de 11 son: ", nomult);
	fprocedimiento.
	procedimiento funciones() es
		para x:= 20 hasta -20, -2 hacer
			mult:= 0; nomult:= 0;
			y:= a*(x**2) + b*x + c;
			resultados();
		fpara;
		multiplos();
	fprocedimiento.
	procedimiento continuar() es
		Escribir("Desea seguir? ('V','F'): ");
		Leer(seguir);
	fprocedimiento.
PROCESO
	Escribir("Ingrese una opcion para tratar funciones (1, 2): ");
	Leer(opcion); seguir:= 'V'
	segun opcion hacer
		1:	Escribir("Ingrese cantidad de funciones a tratar: ");
			Leer(cantf);
			para cantf hasta 0, -1 hacer
				coeficientes();
				funciones();
			fpara;
		2:	mientras seguir = 'V' hacer
				coeficientes();
				funciones();
			fmientras;
		otro:	Escribir("Error: Opcion NO valida.");
	fsegun;
FINACCION.