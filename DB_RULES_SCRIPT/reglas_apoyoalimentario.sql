INSERT INTO dominio (nombre,descripcion)VALUES('APOYOALIMENTTARIO','Dominio para deterinar puntaje del programa apoyo alimentario de Bienestar');--ID dominio 18
INSERT INTO tipo_predicado (nombre)VALUES ('predicado');--ID tipo_predicado 1
INSERT INTO tipo_predicado (nombre)VALUES ('regla');--ID tipo_predicado 2

INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'estrato(1,10).','Determina un puntaje de 10 si el estudiante posee un estrato 1',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'estrato(2,10).','Determina un puntaje de 10 si el estudiante posee un estrato 2',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'estrato(3,10).','Determina un puntaje de 10 si el estudiante posee un estrato 3',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'estrato(X,Y):- X @> 3, Y is 0 .','Determina un puntaje de 0 si el estudiante posee un estrato superior a 3',2);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'matricula(X,Y):- between(0,200000,X,20,Y); between(200001,400000,X,16,Y); between(400001,600000,X,12,Y); between(600001,800000,X,8,Y); between(800001,900000,X,4,Y); between(900001,X,X,0,Y).','Determina un puntaje dependiendo del valor de la matricula del estudiante',2);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'between(V,W,X,Y,Z):- X @>= V, X @=< W, Z is Y.','Determina un rango de valores y un puntaje a asignar',2);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'betweenresultado(V,W,X,Y,Z):- X @>= V, X @=< W, Z = Y.','Determina un rango de valores y un puntaje a asignar',2);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'sostenimientoPropio(si,5).','Determina un puntaje de 5 si el estudiante posee sostenibilidad economica propia',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'sostenimientoPropio(no,0).','Determina un puntaje de 0 si el estudiante NO posee sostenibilidad economica propia',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'sostieneHogar(si,5).','Determina un puntaje de 5 si el estudiante sostiene economicamente el lugar donde reside',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'sostieneHogar(no,0).','Determina un puntaje de 0 si el estudiante NO sostiene economicamente el lugar donde reside',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'viveFueraNucleo(si,4).','Determina un puntaje de 4 si el estudiante vive fuera de su nucleo familiar',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'viveFueraNucleo(no,0).','Determina un puntaje de 0 si el estudiante NO vive fuera de su nucleo familiar',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'personasAcargo(si,6).','Determina un puntaje de 6 si el estudiante tiene personas a cargo',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'personasAcargo(no,0).','Determina un puntaje de 0 si el estudiante NO tiene personas a cargo',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'empleadorOarriendo(si,5).','Determina un puntaje de 5 si el estudiante convive con el empleador o paga arriendo',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'empleadorOarriendo(no,0).','Determina un puntaje de 0 si el estudiante NO convive con el empleador y  NO paga arriendo',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'fueraDeBogota(si,5).','Determina un puntaje de 5 si el estudiante proviene de municipio o ciudad fuera de Bogota',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'fueraDeBogota(no,0).','Determina un puntaje de 0 si el estudiante  proviene de Bogota',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'poblacionEspecial(X,Y):- X \== n, Y is 5.','Determina un puntaje de 5 si el estudiante pertenece a alguna poblacion especial',2);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'poblacionEspecial(n,0).','Determina un puntaje de 0 si el estudiante NO pertenece a alguna poblacion especial',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'discapacidad(si,5).','Determina un puntaje de 5 si el estudiante posee alguna discapacidad fisica o mental',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'discapacidad(no,0).','Determina un puntaje de 0 si el estudiante NO posee alguna discapacidad fisica o mental',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'patologiaAlimenticia(si,5).','Determina un puntaje de 5 si el estudiante posee alguna patologia alimenticia o de nutricion',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'patologiaAlimenticia(no,0).','Determina un puntaje de 0 si el estudiante NO posee alguna patologia alimenticia o de nutricion',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'suma(A,B,C,D,E,F,G,H,I,J,K,L,X):- X is A+B+C+D+E+F+G+H+I+J+K+L.','Determina la suma total de todos los hechos y reglas',2);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'total(X,Y):- betweenresultado(0,24,X,ss,Y); betweenresultado(25,44,X,b,Y); betweenresultado(45,59,X,a,Y); betweenresultado(60,X,X,t,Y).','Determina el tipo de apoyo alimentario segun la regla de suma',2);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'resultado(AA,A,B,C,C1,D,E,F,G,H,I,J,K,L,Z):- estadoPrograma(AA), estrato(A,Y1), matricula(B,Y2), multiplicacion(C1,2,V1), multiplicacion(C1,3,V2), multiplicacion(C1,4,V3), ingresos(C,C1,V1,V2,V3,Y3), sostenimientoPropio(D,Y4), sostieneHogar(E,Y5), viveFueraNucleo(F,Y6), personasAcargo(G,Y7), empleadorOarriendo(H,Y8), fueraDeBogota(I,Y9), poblacionEspecial(J,Y10), discapacidad(K,Y11), patologiaAlimenticia(L,Y12), suma(Y1,Y2,Y3,Y4,Y5,Y6,Y7,Y8,Y9,Y10,Y11,Y12,X), total(X,Z).','Regla que hace el llamado de los hechos y reglas con los valores ingresados por el estudiante',2);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'ingresos(X,X1,X2,X3,X4,Y):- between(0,X1,X,30,Y); between(X1,X2,X,20,Y); between(X2,X3,X,10,Y); between(X3,X4,X,5,Y); between(X4,X,X,0,Y).','Determina un puntaje dependiendo el valor de ingreso del estudiante',2);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'estadoPrograma(1).','Determina el estado actual de un estudiante en proceso de inscripcion',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'estadoPrograma(4).','Determina el estado actual de un estudiante en proceso de inscripcion',1);
INSERT INTO predicado(dominio, nombre, descripcion, tipo_predicado) VALUES (18,'multiplicacion(X,Y,P):- P is X*Y.','Determina el valor de comparacion de ingresos',2);

