CREATE TABLE books{
    book_id    INT PRIMARY KEY AUTO_INCREMENT,
    name       VARCHAR(50),
    author     VARCHAR(50),
    status     BOOLEAN DEFAULT true
}
CREATE TABLE  loans{
    loan_id    INT PRIMARY KEY AUTO_INCREMENT,
    book_id    INT NOT NULL, 
    user_id    INT NOT NULL,
    loan_date    DATE,
    return_date    DATE,
    FOREIGN KEY (book_id) REFERENCES books(book_id)
}
CREATE TABLE user{
    user_id     INT PRIMARY KEY,

}

/*Sistema de Biblioteca*/
--Ejercicio Nº1 
 SELECT * FROM books WHERE status == true ORDER BY book_id ASC
--Ejercicio Nº2
SELECT * FROM books WHERE name == 'Harry Potter' OR  author == 'Oscar Flores'
--Ejercicio N°3
SET @book_id = 1;
UPDATE books SET status = false WHERE book_id = @book_id;

INSERT INTO loans (book_id, user_id, loan_date, return_date)
VALUES (?,?,CURDATE(), NULL)
--Ejercicio N°4
SELECT * FROM books WHERE status == false ORDER BY book_id ASC
--Ejercicio N°5
Set @book_id = 1;
UPDATE books SET status = true WHERE book_id = @book_id
INSERT INTO loans(user_id, loan_date, return_date) VALUES(?,CURDATE(), NULL)
--Ejercicio Nº6 
SELECT book_id, COUNT(loan_id) FROM loans 
GROUP BY book_id ORDER BY DESC

--------------------------------------

CREATE TABLE pacientes{
    paciente_id   INT PRIMARY KEY AUTO_INCREMENT,
    nombre        VARCHAR(50) NOT NULL,
    estado         INT CHECK (estado IN (1,2,3)),
    num_seg        INT NOT NULL LIMIT(8)
}

CREATE TABLE medicos{
    medico_id     INT NOT NULL PRIMARY KEY, -- 1
    nombre        VARCHAR(50) NOT NULL, -- pepe
    -- id_de_proxima_cita int, -- 10
    -- FOREIGN KEY (id_de_proxima_cita) REFERENCES citas(cita_id)
}

CREATE TABLE citas{
    cita_id         INT PRIMARY KEY AUTO_INCREMENT, 
    nombre_paciente  VARCHAR(50) NOT NULL,
    nombre_medico    VARCHAR(50), 
    fecha_cita      DATE, 
    FOREIGN KEY (nombre_paciente) REFERENCES pacientes(nombre)
    FOREIGN KEY (nombre_medico) REFERENCES medicos(nombre)
}

CREATE TABLE historial_medico{
    historia_id    INT PRIMARY KEY NOT NULL,
    paciente_id    INT NOT NULL,
    estado         INT CHECK(estado in (1,2,3))
    FOREIGN KEY (paciente_id) REFERENCES pacientes(paciente_id)
}

--Ejercicio N°1
SELECT paciente_id FROM pacientes WHERE estado ORDER BY ASC
--Ejercicio N°2
SELECT * FROM pacientes WHERE nombre == ? AND num_seg == ? 
--Ejercicio N°3
INSERT INTO citas(nombre_paciente,nombre_medico, CURDATE()) RETURN id
UPDATE medicos set 
--Ejercicio N°4
SELECT * FROM citas WHERE nombre_medico = ? AND fecha_cita IS NOT NULL ORDER BY  fecha_cita ASC;
--Ejercicio N°5
SET paciente = ?
SET new_diagnostico = ?
UPDATE pacientes SET estado == new_diagnostico WHERE nombre = paciente
--Ejercicio N°6
SET fecha_busqueda = CURDATE()
SELECT * FROM citas WHERE fecha_cita == fecha_busqueda
 
-----------------------------------
CREATE TABLE productos{
    productos_id    INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    nombre          VARCHAR(20) NOT NULL,
    precio          INT,
    proveedor       VARCHAR(20) NOT NULL,
    cantidad        INT NOT NULL,
    FOREIGN KEY(proveedor) REFERENCES proveedor(proveedor_id)
}
CREATE TABLE ventas{
    venta_id        INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    nombre_producto VARCHAR(50) NOT NULL,precio          INT,
    fecha_venta     DATE,
    FOREIGN KEY venta_id REFERENCES productos(productos_id)
}
CREATE TABLE proveedor{
    proveedor_id    INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    productos       VARCHAR(50)
}

--Ejercicio N°1
SELECT * FROM productos WHERE stock == true
--Ejercicio N°2
SET nombre_buscar = ?
SET proveedor_buscar = ?
SELECT * FROM productos WHERE nombre == nombre_buscar AND proveedor == proveedor_buscar
--Ejercicio N°3
INSERT INTO ventas(nombre,fecha_venta) VALUES('aceite',2200, CURDATE())
--Ejercicio N°4 
SELECT * FROM ventas WHERE fecha_venta == CURDATE()
--Ejercicio N°5
INSERT INTO ventas(nombre, fecha_venta) VALUES('Arroz',CURDATE())
UPDATE productos SET cantidad = cantidad-1 WHERE cantidad IS NOT NULL
--Ejercicio N°6
SET fecha_busqueda = CURDATE()
SELECT * FROM ventas WHERE fecha_venta == fecha_busqueda

-----------------------------------------------------------
CREATE TABLE clientes{
    cliente_id      INT PRIMARY KEY AUTO_INCREMENT,
    nombre          VARCHAR(50) NOT NULL,

}
CREATE TABLE reservas{
    reserva_id      INT PRIMARY KEY AUTO_INCREMENT,
    cliente_id      VARCHAR(20) NOT NULL,
    mesa_id         INT,
    fecha_reserva   DATE,
    hora_reserva    TIME,
    cantidad_gente  INT,
    FOREIGN KEY(cliente_id) REFERENCES clientes(cliente_id)      
    FOREIGN KEY(mesa_id) REFERENCES mesas(mesa_id)      
}  

CREATE TABLE mesas{
    mesa_id        INT PRIMARY KEY AUTO_INCREMENT,
    capacidad       INT
}

--Ejercicio N°1
SELECT DAYNAME(fecha_reserva) AS dia_semana, HOUR(hora_reserva) AS hora_dia, cantidad_gente, COUNT(*) AS cantidad_reservas FROM reservas GROUP BY dia_semana, hora_dia,cantidad_gente ORDER BY dia_semana, hora_dia,cantidad_gente