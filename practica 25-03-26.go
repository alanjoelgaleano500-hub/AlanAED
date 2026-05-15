ACCION 6.05 es
AMBIENTE
	seguridad = 0.25;
	Seccion, Corriente, Conductividad: real;
PROCESO
	escribir("Ingrese la Corriente Electrica: ");
	leer(Corriente);
	escribir("Ingrese la Conductividad del material: ");
	leer(Conductividad);
	Seccion:= (Corriente * Conductividad) * (1 + seguridad);
	escribir("La seccion es: ", Seccion);
FIN ACCION.

ACCION 6.06 es
AMBIENTE
	Disp_A, Disp_B, Disp_C, CantV_A, CantV_B, CantV_C: entero;
	Porc_A, Porc_B, Porc_C: N(2.2);
PROCESO
	escribir("Bienvenido al Programa Calculador de Porcentaje de Ventas Concesionarias.");
	escribir("Ingrese la cantidad disponible para cada modelo (En orden A, B, C): ");
	leer(Disp_A, Disp_B, Disp_C);
	escribir("Ingrese la cantidad vendida en el mismo orden especificado anteriormente: ");
	leer(CantV_A, CantV_B, CantV_C);
	Porc_A:= (CantV_A / Disp_A) * 100;
	Porc_B:= (CantV_B / Disp_B) * 100;
	Porc_C:= (CantV_C / Disp_C) * 100;
	escribir("Porcentajes calculados, imprimiendo resultados: ");
	escribir("Porcentaje de Modelos 'A' vendidos: ", Porc_A);
	escribir("Porcentaje de Modelos 'B' vendidos: ", Porc_B);
	escribir("Porcentaje de Modelos 'C' vendidos: ", Porc_C);
FIN ACCION.

ACCION 6.07 es
AMBIENTE
	p_actual, p_nuevo: real;
	pcj_aumento: N(1.2);
PROCESO
	escribir("Este es un programa que calcula el nuevo precio de la nafta wachin.");
	escribir("Ingrese el precio actual del combustible: ");
	leer(p_actual);
	escribir("Ingrese el porcentaje de aumento del combustible (dividido en 100): ");
	leer(pcj_aumento);
	p_nuevo:= p_actual * (1 + pcj_aumento);
	escribir("El precio actual del combustible es: ", p_actual);
	escribir("El porcentaje de aumento fue de: ", pcj_aumento*100);
	escribir("El nuevo precio del combustible es: ", p_nuevo);
FIN ACCION.

ACCION 6.08 es
AMBIENTE
	X, Y, Z: real;
	media: real;
PROCESO
	escribir("Este programa calcula la media geometrica de 3 valores.");
	escribir("Ingrese los valores en orden (X, Y, Z): ");
	leer(X, Y, Z);
	media:= (X*Y*Z) / 3;
	escribir("Imprimiendo resultados...");
	escribir("Valor de X: ", X);
	escribir("Valor de Y: ", Y);
	escribir("Valor de Z: ", Z);
	escribir("Valor de la media: ", media);
FIN ACCION.

ACCION 12 es
AMBIENTE
	num, unidad, decena, centena: entero;
PROCESO
	escribir("Este programa recibe un numero y lo descompone por unidades y comprueba si es multiplo de 3.");
	escribir("Ingrese el numero a descomponer: ");
	leer(num);
	si (num > 100) y (num < 1000) entonces
		centena:= num DIV 100;
		decena:= (num DIV 10) MOD 10;
		unidad:= num MOD 10;
			escribir("El numero se descompone en: ");
			escribir("Centena: ", centena);
			escribir("Decena: ", decena);
			escribir("Unidad: ", unidad);
		si (num MOD 3 = 0) entonces
			escribir("El numero si es multiplo de 3.");
		sino 
			escribir("El numero no es multiplo de 3.");
		fsi;
	sino
		escribir("Ingrese un numero entero que este entre 100 y 1000");
	fsi;
FIN ACCION.