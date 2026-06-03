ACCION HotelTemaA (ConvertiraNumero(caracter)) ES
AMBIENTE
	hotel: secuencia de caracter
	reservas: secuencia de entero
	salida: secuencia de entero
	hot: caracter
	resv, i, j, cant_noches, mayor_resv, n_resv: entero
	procedimiento salida() ES
		si resv < 90000 entonces
			ESC(salida, n_resv);
			ESC(salida, resv);
		fsi;
	fprocedimiento.
	procedimiento reserva() ES
		Para i:= 1 hasta ConvertiraNumero(hot) hacer
			n_resv:= resv;
			Para j:= 1 hasta 2 hacer
				AVZ(reservas, resv);
			fpara;
			si resv > cant_noches entonces
				cant_noches:= vent;
				mayor_resv:= n_resv;
			fsi;
			salida();
			AVZ(reservas, resv);
		fpara;
	fprocedimiento;
PROCESO
	ARR(hotel); AVZ(hotel, hot);
	ARR(reservas); AVZ(reservas, resv);
	CREAR(salida);
	mientras NFDS(hotel) hacer
		mientras hot <> "&" hacer
			AVZ(hotel, hot);
		fmientras;
		AVZ(hotel, hot);
		cant_noches:= 0;
		reserva();
		Escribir("La reserva con mayor cantidad de noches fue la: ", mayor_resv, "Y la cantidad de noches fue de: ", cant_noches);
		AVZ(hotel, hot);
	fmientras;
	CERRAR(hotel); CERRAR(reservas); CERRAR(salida);
FIN ACCION.
