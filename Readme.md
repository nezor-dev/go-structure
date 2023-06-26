ObjectModel \n
| \n
|----- Object \n
|----- InputObject \n


ObjectRepository \n
| \n
|----- CreateObject(Object) Object error \n
|----- GetExistingObject(id uint) Object \n


ObjectService \n
| \n
|----- CreateObject(InputObject) Object error \n
|----- GetExistingObject(id string) Object \n


ObjectRepository is the database layer and just accepts raw data that gets filtered by the service layer