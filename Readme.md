ObjectModel\
|\
|----- Object\
|----- InputObject


ObjectRepository\
|\
|----- CreateObject(Object) Object error\
|----- GetExistingObject(id uint) Object


ObjectService\
|\
|----- CreateObject(InputObject) Object error\
|----- GetExistingObject(id string) Object


ObjectRepository is the database layer and just accepts raw data that gets filtered by the service layer