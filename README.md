# SagwaBooks - BACKEND

a project to a publishing company. to sell books.

made with go lang and mongoDB usign monitoring with bugsnag.

Available at :1323

to run, download project and at terminal digit: go run main.go

## ROUTES

### User

    POST /login    ValidateUser
    entrada:
``` sh 
    {
        "email": "email@email.com",
        "password":"senha"
    } 

    return: userId "userId"
```


	POST /cadastro CreateUser
    entrada:
``` sh 
    {
        "name": "name",
        "email":"email@email.com",
        "password": "password",
        "telefone": "telefone",
    } 

    return: ""
```
	POST /user/edit EditUser
    entrada:
``` sh 
    {
        "userId": "number",
        "name": "name",
        "email":"email@email.com",
        "password": "password",
        "telefone": "telefone",
    } 

    return: ""
```   
	POST /user/delete DeleteUser
    entrada:
``` sh 
    {
        "userId": "number",
    } 

    return: ""
``` 
	/getInformation GetInformationByUserId
    entrada:
``` sh 
    {
        "userId": "number",
    } 

    return:    
    {
        "userId": "number",
        "name": "name",
        "email":"email@email.com",
        "password": "password",
        "telefone": "telefone",
    }
``` 

### Book

	/book/create CreateBook
	/book/edit EditBook
	/book/delete DeleteBook
	/book/name GetBookByName
	/book/autor GetBookByAutor
	/book/genero GetBookByGenero

### Card

	/card/create CreateCard
	/card/edit EditCard
	/card/delete DeleteCard
	/card/user GetCardsByUserId
	/card/validate ValidatedCard