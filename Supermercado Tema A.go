ACCION superTemaA (ConvertiraNumero(caracter)) ES
AMBIENTE
	sucursales: secuencia de caracter
	ventas: secuencia de entero
	salida: secuencia de entero
	suc: caracter
	vent, i, j, mayor_imp, mayor_tick, n_tick, cant_tick: entero
	procedimiento salida() ES
		si vent > 50000 entonces
			ESC(salida, n_tick);
			ESC(salida, vent);
		fsi;
	fprocedimiento.
	procedimiento digitostick() ES
		para i:= 1 hasta 3 hacer
			cant_tick:= cant_tick*10 + ConvertiraNumero(suc);
			AVZ(sucursales, suc);
		fpara;
	procedimiento importe() ES
		Para i:= 1 hasta cant_tick hacer
			AVZ(ventas, vent);
			n_tick:= vent;
			Para j:= 1 hasta 2 hacer
				AVZ(ventas, vent);
			fpara;
			si vent > mayor_imp entonces
				mayor_imp:= vent;
				mayor_tick:= n_tick;
			fsi;
			salida();
			AVZ(ventas, vent);
		fpara;
	fprocedimiento;
PROCESO
	ARR(sucursales); AVZ(sucursales, suc);
	ARR(ventas); AVZ(ventas, vent);
	CREAR(salida);
	mientras NFDS(sucursales) hacer
		cant_tick:= 0;
		mientras suc <> "&" hacer
			AVZ(sucursales, suc);
		fmientras;
		AVZ(sucursales, suc);
		digitostick();
		mayor_imp:= 0;
		importe();
		Escribir("El ticket de mayor importe fue el: ", mayor_tick, "Y el importe fue de: ", mayor_imp);
		AVZ(sucursales, suc);
	fmientras;
	CERRAR(sucursales); CERRAR(ventas); CERRAR(salida);
FIN ACCION.
