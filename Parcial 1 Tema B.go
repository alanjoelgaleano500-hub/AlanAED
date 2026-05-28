ACCION TEMA_B es
AMBIENTE
	Editorial: secuencia de caracter
	Libros: secuencia de caracter
	Sal: secuencia de caracter
	Ed, lib: caracter
	Pag, Pal, EdAuto, i: entero
	bandEd, bandSal: booleano
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
	Crear(Sal); 
	EdAuto:= 0;
	Mientras NFDS(Editorial) hacer
		Pal:= 0;
		Mientras Ed <> '&' hacer
			AVZ(Editorial, Ed);
		fmientras;
		AVZ(Editorial, Ed);
		bandEd:= Falso; bandSal:= Falso;
		Mientras lib <> '@' hacer
			si lib = 'A' entonces
				bandEd:= Verdadero;
			fsi;
			si lib = 'I' entonces
				bandSal: Verdadero;
			fsi;
			AVZ(Libros, lib);
			Pag:= 0;
			Para i:= 0 hasta 3 hacer
				Pag:= Pag*10 + Conv(lib);
				AVZ(Libros, lib);
			fpara;
			Mientras lib <> '&' hacer
				si (bandSal ^ Pag < 150) entonces
					Esc(Sal, lib);
				fsi;
				AVZ(Libros, lib);
			fmientras;
			AVZ(Libros, lib);
			mientras lib <> '#' hacer
				mientras lib = '' hacer
					AVZ(Libros, lib);
				fmientras;
				Pal:= Pal + 1;
				mientras lib <> '' hacer
					AVZ(Libros, lib);
				fmientras;
			fmientras;
			AVZ(Libros, lib);
		fmientras;
		Para i:=1 hasta 2 hacer
			AVZ(Editorial, Ed);
		fpara;
		Escribir("Cantidad de Palabras de todos los resumenes de la Editorial: ", Pal);
		si bandEd entonces
			EdAuto:= EdAuto + 1;
		fsi;
	fmientras;
	Escribir("Cantidad de Editoriales con almenos un Libro de Autoayuda: ", EdAuto);
	Cerrar(Editorial); Cerrar(Libros); Cerrar(Sal);
FIN ACCION.