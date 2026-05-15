ACCION 7 es
AMBIENTE
	n1, n2, n3: entero
PROCESO
	escribir("Este algoritmo pide 3 numeros y devuelve el menor, medio y mayor.");
	escribir("Ingrese 3 numeros enteros: ");
	leer(n1, n2, n3);
	si (n1 > n2) y (n1 > n3) entonces
		si (n2 > n3) entonces
			escribir("El mayor es: ", n1);
			escribir("El menor es: ", n3);
			escribir("El medio es: ", n2);
		sino
			escribir("El mayor es: ", n1);
			escribir("El menor es: ", n2);
			escribir("El medio es: ", n3);
		fsi;
	sino
		si (n2 > n3) entonces
			escribir("El mayor es: ", n2);
			escribir("El menor es: ", n1);
			escribir("El medio es: ", n3);
		sino
			escribir("El mayor es: ", n3);
			escribir("El menor es: ", n1);
			escribir("El medio es: ", n2);
		fsi;
	fsi;
FINACCION.

ACCION 8 es
AMBIENTE
	n1, n2, suma: entero
PROCESO
	escribir("Ingrese dos numeros: ");
	leer(n1, n2);
	suma:= n1 + n2;
	segun suma hacer
		<=50: escribir("<=50");
		<=100: escribir("<=100");
		<=200: escribir("<=200");
		otro: escribir(">200");
	fsegun;
FINACCION.

ACCION 9 es
AMBIENTE
	dN, dAc: 1..31
	mN, mAc: 1..12
	aN, aAc: N(4)
	edad: entero
PROCESO
	escribir("Ingrese su fecha de nacimiento en orden DIA, MES AÑO: ");
	leer(dN, mN, aN);
	escribir("Ingrese la fecha actual en mismo orden: ");
	leer(dAc, mAc, aAc);
	si (mN <= mAc) entonces
		si (dN <= dAc) entonces
			edad:= (aAc - aN);
		sino
			edad:= (aAc - aN) - 1;
		fsi;
	sino
		edad:= (aAc - aN) - 1;
	fsi;
	escribir("Su edad es de ", edad, "años.");
FINACCION.

ACCION 14 es
AMBIENTE
	mont: real
	opcion: (1..3)
PROCESO
	escribir("Ingrese monto de compra: ");
	leer(mont);
	si mont > 200000 entonces
		escribir("El monto supera los 200.000 entonces gozara de un 15% de descuento equivalente a: ", mont*0,15);
		mont:= mont - (mont * 0,15);
	fsi;
		escribir("Elija una opcion de entrega: OPCION 1, OPCION 2, OPCION 3");
		leer(opcion);
		segun opcion hacer
			1:mont:= mont;
			2:mont:= mont * 1,05;
			3:mont:= mont * 1,10;
		fsegun;
	escribir("El monto final es igual a: ", mont);
FINACCION.

ACCION 15 es
AMBIENTE
	pisos: entero
	altpiso: real
	altura: real
	edificios: entero
PROCESO
//PUNTO A://
	escribir("Ingrese la cantidad de pisos del edificio: ");
	leer(pisos);
	escribir("Ingrese la altura promedio de cada piso en metros: ");
	leer(altpiso);
	altura:= (pisos * altpiso) * 3,28;
	escribir("La altura del edificio en pies es igual a: ", altura);
//PUNTO B://
	para edificios:= 1 hasta 50 hacer
		escribir("Ingrese la cantidad de pisos del edificio: ");
		leer(pisos);
		escribir("Ingrese la altura promedio de cada piso en metros: ");
		leer(altpiso);
		altura:= (pisos * altpiso) * 3,28;
		escribir("La altura del edificio en pies", edificios, "es igual a: ", altura);
	fpara;
//PUNTO C://
	escribir("Ingrese cantidad de edificios a calcular altura: ");
	leer(edificios);
	mientras (edificios >= 1) hacer
		escribir("Ingrese la cantidad de pisos del edificio: ");
		leer(pisos);
		escribir("Ingrese la altura promedio de cada piso en metros: ");
		leer(altpiso);
		altura:= (pisos * altpiso) * 3,28;
		escribir("La altura del edificio", edificios, "en pies es igual a: ", altura);
		edificios:= edificios - 1;
	fmientras;
FINACCION.

ACCION 16 es
AMBIENTE
	cant_isi, cant_iem, cant_iq, total, isi, iq, iem: entero
	p_isi, p_iem, p_iq: real
	seguir: ("S"/"N")
PROCESO
	//punto B//
	cant_isi:= 0; cant_iem:= 0; cant_iq:= 0;
	seguir:= "S";
	mientras (seguir = "S") hacer
		escribir("Ingrese la cantidad de egresados de cada carrera en la facultad en orden ISI, IQ, IEM: ");
		leer(isi, iq, iem);
		cant_isi:=cant_isi + isi;
		cant_iem:=cant_iem + iem;
		cant_iq:=cant_iq + iq;
		total:= isi + iq + iem;
		p_isi:= (isi * 100) / total;
		p_iem:= (iem * 100) / total;
		p_iq:= (iq * 100) / total;
		escribir("Los porcentajes de egresados en orden (isi, iq, iem), son: ", p_isi, p_iq, p_iem);
		escribir("Desea continuar? S/N");
		leer(seguir);
	fmientras;
	escribir("La cantidad de alumnos por carrera en orden (isi, iq, iem), es: ", cant_isi, cant_iq, cant_iem);
	escribir("La cantidad total general de alumnos de todas las carreras es: ", total);
FINACCION.
