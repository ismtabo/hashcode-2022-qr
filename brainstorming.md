// 3 contributors and 3 projects
3 3

// 1 skill
Anna 1
// 2 skill level for c++
C++ 2

Bob 2
HTML 5
CSS 5
Maria 1
Python 3

// 5 number of days to complete the project
// 10 points to complete the project
// 5 the best before day for the project
// 1 number of roles in the project
Logging 5 10 5 1
C++ 3

WebServer 7 10 7 2
HTML 3
C++ 2
WebChat 10 20 20 2
Python 3
HTML 3


possible approaches:


- if contributor doesn't has the level required (1 lower) and the project has only 1 seat for a contributor
	- assign them another project, when they finish, they'll level up +1
	- then assign them to the initial project since now they'll have the skills


- A contributor doesn't upgrade their skills if they are in a project that requires a lower level for the skill


- If the project ends after the best before but inside the deadline, subtract strictly end day from best before. The result needs to be subtracted from the last day of work -1. That result is the final score 

ig 
Project Logging's last day of work is 11 (so it's completed strictly before day 12), while its "best before" day was 5. It is late by (12−5=) 7 days and receives a score of: (10−7=) 3 points.

- complete first the lower number of best before



------------------------------------------------------------------------------------------------------------

ALGORITMO PARA ASIGNACIÓN DE PROYECTOS:

	- RECORREMOS PROYECTOS
		- COGEMOS EL PROYECTO CON EL MENOR BESTBEFORE, MAYOR SCORE Y MENOR NUMERO DE ROLES NECESARIOS
		- QUITAMOS LOS TRABAJADORES ASIGNADOS DE LA LISTA DE POSIBLES TRABAJADORES Y EL PROYECTO ELEGIDO
		  DE LA LISTA DE POSIBLES PROYECTOS ELEGIDOS.
		- BUSCAMOS MÁS PROYECTOS QUE PODAMOS REALIZAR CON EL MENOR BESTBEFORE POSIBILE Y MAYOR SCORE.
		- CUANDO NO SE PUEDAN ASIGNAR MÁS PROYECTOS AUMENTAMOS EL CONTADOR DE DÍA ACTUAL Y REPETIMOS EL BUCLE.
		- DIA_ACTUAL = DEADLINE_MENOR CALCULAS EL SCORE DEL PROYECTO

------------------------------------------------------------------------------------------------------------

ALGORITMO PARA LA ASIGNACIÓN DE TRABAJADORES A UN PROYECTO

	1.- Los trabajadores con SKILL = necesaria para el proyecto (SIEMPRE QUE SEA < MAXSKILL) Y NSKILL = 1
	2.- Los trabajadores con SKILL > necesaria que no tengan ninguna otra skill (NSKILL = 1) Y NSKILL = 1
	4.- LOS TRABAJADORES CON SKILL = necesaria para el proyecto Y NSKILL > 1
	5.- Los trabajadores con SKILL > necesaria con NSKILL > 1
	6.- Los trabajadores con SKILL menor que 1 para la necesaria en los proyectos que pueda haber MENTORING
		3.1 - MENTORES QUE PUEDAN CUMPLIR ALGUNA OTRA FUNCIÓN EN EL PROYECTO
		3.2 - MENTORES QUE SOLO SIRVAN PARA AYUDAR AL TRABAJADOR ASIGNADO CUYA SKILL SEA = NECESARIA y NSKILL = 1
		3.3 - MENTORES QUE SOLO SIRVAN PARA AYUDAR AL TRABAJADOR ASIGNADO CUYA SKILL SEA > NECESARIA Y NSKILL > 1
		3.4 - MENTORES QUE SOLO SIRVAN PARA AYUDAR AL TRABAJADOR ASIGNADO CUYA SKILL SEA MAXIMA
	7.- Los trabjadores con SKILL > necesaria y que tengan la MAXIMA SKILL.

------------------------------------------------------------------------------------------------------------

ALGORITMO PARA LA ASIGNACIÓN DE SCORE DE PROYECTO

	1.- If the project ends after the best before but inside the deadline, subtract strictly end day from best before.
	2.- The result needs to be subtracted from the last day of work -1. That result is the final score 

------------------------------------------------------------------------------------------------------------