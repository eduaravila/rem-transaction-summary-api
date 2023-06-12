
## functional requirements
- REST api to get the user transaction summary
- save user transactions in a database
- save user transactions in a file (csv)
- send transaction summary through email
- register user in the system (name, email) 

## non functional requirements

- error prone
- scalable
- Usability
- maintainable
- testable

## data model

```
 Transaction {
    id int
    Date date
    Transaction float    
 }

  User {
      id int
      name string
      email string
      transactions []Transaction
  }

  TransactionSummary {
      user User
      total float
      averageCredit float
      averageDebit float
      numberOfTransactionsPerMonth []TransactionPerMonth
  }

  TransactionPerMonth {
      month int
      numberOfTransactions int
  }  

```

## CSV storage 

to keep it simple we are going to execute a linear search O(n) to look and see if the user does exist 

index file
    user_id, Email, file_name

user file
    Id, Date, Transaction

## postgres storage

User table
    Id uuid
    Name string
    Email string

Transaction table
    Id uuid
    UserId uuid
    Date date
    Transaction float

## ports | REST api

- GET /users/{id}/transactions/summary - 200 OK
  - res: { summaryId: int,  status: "INPROGRESS", notification-data:{
            notificationId: int,
            recipient: string,
            subject: string
            timestamp: date
        }}
        
- POST /users/{id}/transactions - 202 Accepted
  - body: { date: "2020-01-01", transaction: 100.00 }

- POST /users - 201 Created
  - body: { name: "John Doe", email: "example@example.com" }
  - res: { id: int, name: string, email: string }