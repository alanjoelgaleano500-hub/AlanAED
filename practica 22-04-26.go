ACCION 32.04 es
AMBIENTE
	num: entero
	opcion: ("S","N")
	Funcion obtenclave (n: entero) es
	AMBIENTE
		clave: entero
	PROCESO
		si n >= 0 entonces
			clave:= 0;
			mientras n > 0 hacer
				clave:= clave + (n div 10);
				n:= n div 10;
			fmientras;
			clave:= clave mod 7;
		sino
			clave:= -1;
		fsi;
		obtenclave:= clave;
	finfuncion;
PROCESO
	opcion:= "S";
	mientras (opcion = "S") hacer
		escribir("Ingrese un numero para obtener su clave: ");
		leer(num);
		escribir("Su clave es: ", obtenclave(num));
		escribir("Seguir? ('S','N'): ");
		leer(opcion);
	fmientras;
FINACCION.

ACCION 32.11 es
AMBIENTE
	opcion, n1, n2: entero
	salir: ("S","N")
	Procedimiento Menu() es
	PROCESO
		escribir("		 MENU		");
		escribir("1) Suma");
		escribir("2) Resta");
		escribir("3) Multiplicacion");
		escribir("4) Division");
	Finprocedimiento;
	Funcion operacionsuma(num1, num2: entero): entero es
	PROCESO
		operacionsuma:= num1 + num2;	
	finfuncion;
	Funcion operacionresta(num1, num2: entero): entero es
	PROCESO
		operacionresta:= num1 - num2;	
	finfuncion;
	Funcion operacionmult(num1, num2: entero): entero es
	PROCESO
		operacionmult:= num1 * num2;	
	finfuncion;
	Funcion operaciondiv(num1, num2: entero): real es
	PROCESO
		si num2 <> 0 entonces
			operaciondiv:= num1 / num2;	
		sino
			escribir("Error: No es posible dividir por 0.");
			operaciondiv:= 0;
	finfuncion;
PROCESO
	salir:= "N";
	mientras (salir = "N") hacer
		escribir("Ingrese numeros para una operacion: ");
		leer(n1, n2);
		Menu();
		escribir("Seleccione una opcion: ");
		leer(opcion);
		segun opcion hacer
			1:
				escribir("El resultado de la operacion seleccionada es: ", operacionsuma(n1, n2));
			2:
				escribir("El resultado de la operacion seleccionada es: ", operacionresta(n1, n2));
			3:
				escribir("El resultado de la operacion seleccionada es: ", operacionmult(n1, n2));
			4:
				escribir("El resultado de la operacion seleccionada es: ", operaciondiv(n1, n2));
			otro:
				escribir("Error: Opcion ingresada no es valida.");
		fsegun;
		escribir("Salir? ('S','N'): ");
		leer(salir);
	fmientras;
FINACCION.
