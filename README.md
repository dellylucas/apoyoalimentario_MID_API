# apoyoalimentario_MID_API

Este proyecto fue generado con Framework BeeGo version 1.9.1.

Esta aplicación sirve para realizar evaluacion de reglas de negocio de la inscripción al Apoyo Alimentario de la Universidad Distrital Francisco José de Caldas

IMPORTANTE
Es necesario que se tenga el CLIENTE y CRUD_API

Configuración de puerto
Modifique la propiedad llamada httpport del archivo app.conf que está ubicado en conf/httpport

#Configuración de ruta del CRUD_API

en el archivo models/utility.go
Se encuentran la constante CrudPath la cual devuelve la respuesta al CRUDAPI
tiene por defecto el valor http://localhost:8086/v1/information/

Para desplegar
Debe ejecutar una vez instalado Go y Beego, el comando bee run para ejecutarlo en un servidor local, la aplicación corre por defecto en http://localhost:8090/

Si desea cambiarla, modifique la propiedad httpport del archivo app.conf esta ubicado en la carpeta conf del proyecto.
