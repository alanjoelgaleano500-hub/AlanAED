ACCION EJ2 es
AMBIENTE
	Fecha = REGISTRO
		año: N(4)
		mes: 1..12
		dia: 1..31
	freg;
	Super = REGISTRO
		cod_s: N(2)
		rubro: AN(20)
		cod_art: N(5)
		F_UltR: Fecha
		St_seg: entero
		St_act: entero
	freg;
	Salida = REGISTRO
		cod_s: N(2)
		rubro: AN(20)
		c_art: entero
	freg;
	r_sup: Super
	r_sal: Salida
	stock: archivo de Super ordenado por cod_s, rubro y cod_art
	sal: archivo de Salida
	resg_s: N(2)
	resg_rub: AN(20)
	tot_s, tot_rub, tot_gen: entero
	procedimiento inicializar() es
		ABRIR E/ (stock); ABRIR /S (sal);
		Leer(stock, r_sup);
		tot_s:= 0; tot_rub:= 0; tot_gen:= 0;
		resg_s:= r_sup.cod_s;
		resg_rub:= r_sup.rubro;
	fprocedimiento;
	procedimiento tratar_reg() es
		si r_sup.St_act < r_sup.St_seg entonces
			escribir("SUCURSAL | RUBRO | ARTICULO |");
			escribir(r_sup.cod_s, r_sup.rubro, r_sup.cod_art);
			tot_rub:= tot_rub + 1;
		fsi;
	fprocedimiento;
	procedimiento corte_rub() es
		escribir("Sucursal: ", resg_s);
		escribir("Rubro: ", resg_rub);
		escribir("Total de Articulos por Rubro: ", tot_rub);
		tot_s:= tot_s + tot_rub;
		r_sal.c_art:= tot_rub;
		r_sal.rubro:= resg_rub;
		r_sal.cod_s:= resg_s;
		escribir(sal, r_sal);
		tot_rub:= 0;
		resg_rub:= r_sup.rubro;
	fprocedimiento;
	procedimiento corte_suc() es
		corte_rub();
		escribir("Sucursal: ", resg_s);
		escribir("Total de Articulos por Sucursal: ", tot_s);
		tot_gen:= tot_gen + tot_s;
		tot_s:= 0;
		resg_s:= r_sup.cod_s;
	fprocedimiento;
	procedimiento tratar_corte() es
		si r_sup.cod_s <> resg_s entonces
			corte_suc();
		sino
			si r_sup.rubro <> resg_rub entonces
				corte_rub();
			fsi;
		fsi;
	fprocedimiento;
PROCESO
	inicializar();
	mientras NFDA(stock) hacer
		tratar_corte();
		tratar_reg();
		Leer(stock, r_sup);
	fmientras;
	corte_suc();
	escribir("Total General de Articulos a reponer: ", tot_gen);
	CERRAR(stock); CERRAR(sal);
FIN ACCION.