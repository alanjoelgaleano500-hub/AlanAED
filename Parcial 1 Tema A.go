ACCION TEMA_A es
AMBIENTE
	Editorial: secuencia de caracter
	Libros: secuencia de caracter
	Sal: secuencia de caracter
	Ed, lib: caracter
	cant_lib, long_r, Pag, i, mayor, prov, provmayor: entero
	bandSal: booleano
	Funcion Conv(v: caracter): entero es
		segun v hacer
			'0': conv:= 0;
			'1': conv:= 1;
			'2': conv:= 2;
			'3': conv:= 3;
			'4': conv:= 4;
			'5': conv:= 5;
			'6': conv:= 6;
			'7': conv:= 7;
			'8': conv:= 8;
			'9': conv:= 9;
		fsegun;
	ffuncion.
PROCESO
	ARR(Editorial);AVZ(Editorial, Ed);
	ARR(Libros);AVZ(Libros, lib);
	Crear(Sal); bandSal:= Falso;
	Mientras NFDS(Editorial) hacer
		cant_lib:= 0; long_r:= 0;
		Mientras Ed <> '&' hacer
			AVZ(Editorial, Ed);
		fmientras;
		AVZ(Editorial, Ed);
		Para i:=1 hasta 2 hacer
			prov:= prov*10 + Conv(Ed);
			AVZ(Editorial, Ed);
		fpara;
		cant_lib:= cant_lib + 1;
		Mientras lib <> '@' hacer
			si lib = 'L' entonces
				bandSal:= Verdadero;
			fsi;
			AVZ(Libros, lib);
			Para i:= 0 hasta 3 hacer
				Pag:= Pag*10 + Conv(lib);
				AVZ(Libros, lib);
			fpara;
			Mientras lib <> '&' hacer
				si (bandSal ^ (Pag >= 200 ^ Pag <= 300)) entonces
					Esc(Sal, lib);
				fsi;
				AVZ(Libros, lib);
			fmientras;
			AVZ(Libros, lib);
			mientras lib <> '#' hacer
				mientras lib = '' hacer
					AVZ(Libros, lib);
				fmientras;
				long_r:= long_r + 1;
				mientras lib <> '' hacer
					AVZ(Libros, lib);
				fmientras;
			fmientras;
			AVZ(Libros, lib);
			Escribir("Longitud promedio de resumenes: ", long_r/cant_lib);
			si cant_lib > mayor entonces
				mayor:= cant_lib;
				provmayor:= prov;
			fsi;
		fmientras;
	fmientras;
	Escribir("La Provincia con mayor cantidad de libros presentados es: ", provmayor);
	Cerrar(Editorial); Cerrar(Libros); Cerrar(Sal);
FIN ACCION.