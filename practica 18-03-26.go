ACCION 6.01 es
AMBIENTE
	inflacion = 0.04;
	año_ac, año_fut: N(4);
	precio_prod, precio_ac: real;
PROCESO
	escribir ("Ingrese el año futuro a considerar: ");
	leer (año_fut);
	escribir ("Ingrese el precio del producto: ");
	leer (precio_prod);
	escribir ("Ingrese el año actual: ");
	leer (año_ac);
	precio_ac := precio_prod * (1 + inflacion) ** (año_fut - año_act);
	escribir ("El precio actual del producto es: ", precio_ac);
FIN ACCION.

ACCION 6.02 es
AMBIENTE
	valor: N(3);
	centena, decena, unidad, suma_digitos: entero;
PROCESO
	escribir ("ingrese un numero de tres cifras: ");
	leer (valor);
	centena := valor DIV 100
	decena := (valor DIV 10) MOD 10;
	unidad := valor MOD 10;
	suma_digitos := unidad + valor + decena;
	escribir ("la unidad es: ", unidad, " la decena es: ", decena, " la centena es: ", centena);
	escribir ("La suma de los dígitos es: ", suma_digitos);
FIN ACCION.

ACCION 6.03 es
AMBIENTE
	iva = 0.21;
	ganancias_pc = 0.12;
	ganancias_imp = 0.07;
	precio_pc, precio_imp, precio_fpc, precio_fimp: real;
PROCESO
	escribir ("Ingrese el precio de costo del pc: ");
	leer (precio_pc);
	escribir ("Ingrese el precio de costo de la impresora: ");
	leer (precio_imp);
	precio_fpc := precio_pc * (1 + ganancias_pc) * (1 + iva);
	precio_fimp := precio_imp * (1 + ganancias_imp) * (1 + iva);
	escribir ("El precio final del pc es: ", precio_fpc);
	escribir ("El precio final de la impresora es: ", precio_fimp);
FIN ACCION.

ACCION 6.04 es
AMBIENTE
	base_mayor, base_menor, altura: real;
	superficie: real;
PROCESO
	escribir ("Ingrese la base mayor del trapecio: ");
	leer (base_mayor);
	escribir ("Ingrese la base menor del trapecio: ");
	leer (base_menor);
	escribir ("Ingrese la altura del trapecio: ");
	leer (altura);
	superficie := ((base_mayor + base_menor) * altura) / 2;
	escribir ("Las bases del trapecio son: ", base_mayor, " y ", base_menor, " la altura es: ", altura);
	escribir ("La superficie del trapecio es: ", superficie);
FIN ACCION.