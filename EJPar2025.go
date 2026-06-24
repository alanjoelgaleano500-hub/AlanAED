ACCION EjTemaA es
AMBIENTE
	Fecha = REGISTRO
		año: N(4)
		mes: 6
		dia: 1..30
	freg;
	Estacion = REGISTRO
		n_est: N(5)
		Tipo_V: ("SD", "SV", "UT")
		n_car: N(5)
		F_car: Fecha
		Dur_car: 1..60
		Costo_car: real
	freg;
	Salida = REGISTRO
		n_est: N(5)
		rec_UT: real
	freg;
	r_est: Estacion
	r_sal: Salida
	est: archivo de Estacion ordenado por n_est, Tipo_V, y n_car
	sal: archivo de Salida
	resg_est: N(5)
	resg_V: ("SD", "SV", "UT")
	tot_est, tot_v, tot_ut, tot_sd, tot_gen: real
	procedimiento inicializar() es
		tot_est:= 0; tot_v:= 0; tot_sd:= 0; tot_ut:= 0;
		resg_est:= r_est.n_est;
		resg_V:= r_est.Tipo_V;
	fprocedimiento;
	procedimiento quincena() es
		si r_est.F_car.dia < 16 entonces
			escribir("Carga en 1er quincena");
			escribir("ESTACION | TIPO V | FECHA CARGA |");
			escribir(r_est.n_est, r_est.Tipo_V, r_est.F_car);
		fsi;
	fprocedimiento;
	procedimiento corte_v() es
		escribir("Nro estacion: ", resg_est);
		escribir("Tipo Vehiculo: " resg_V);
		escribir("Total recaudado: " tot_v);
		tot_est:= tot_est + tot_v;
		si resg_V = "SD" entonces
			tot_sd:= tot_v;
		sino
			si resg_V = "UT" entonces
				tot_ut:= tot_v;
			fsi;
		fsi;
		tot_v:= 0;
		resg_V:= r_est.Tipo_V;
	fprocedimiento;
	procedimiento corte_est() es
		corte_v();
		escribir("Nro estacion: ", resg_est);
		escribir("Total recaudado: ", tot_est);
		si tot_sd > tot_ut entonces;
			escribir("La estacion: ", resg_est, "tuvo mayor recaudacion en Vehiculos de tipo SEDAN");
		fsi;
		r_sal.n_est:= resg_est;
		r_sal.rec_UT:= tot_ut;
		escribir(sal, r_sal);
		tot_est:= 0; tot_sd:= 0; tot_ut:= 0;
		resg_est:= r_est.n_est;
	fprocedimiento;
	procedimiento tratar_corte() es
		si resg_est <> r_est.n_est entonces
			corte_est();
		sino
			si resg_V <> r_est.Tipo_V entonces
				corte_v();
			fsi;
		fsi;
	fprocedimiento;
PROCESO
	ABRIR E/ (est); ABRIR /S (sal);
	Leer(est, r_est);
	inicializar();
	mientras NFDA(est) hacer
		tratar_corte();
		quincena();
		tot_v:= tot_v + r_est.Costo_car;
		Leer(est, r_est);
	fmientras;
	corte_est();
	CERRAR(est); CERRAR(sal);
FinAccion.