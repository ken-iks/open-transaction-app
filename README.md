# Message API

## Usage

In terminal, start backend:
```bash
cd backend
go run .
```

In seperate terminal, start frontend:
```bash
cd frontend
npm run start
```
This will take you to the login page, where you login with your Account Number and Routing Number.

### OAuth:

When Attempting to submit a message to the Database, the `sender_an` and `sender_rtn` are cross referenced with the currently logged in user's account number and routing number as a step to ensure a user is only submitting messages where they are the sender.

### Message Submission:

To submit a message, you can put your message in the textbox and press the submit button. This will only submit the message if your message is in the structure: `seq=A;sender_rtn=B;sender_an=C;receiver_rtn=D;receiver_an=E;amount=F` (in any order), where {A-F} all represent strings that only contain integers, and `sender_an` and `sender_rtn` are aligned with the current logged in user's information

On the backend, this is handled by the `Create_message` function and the `DB_insert` function.

### Message Retrieval:

To retrieve a message, you choose the sequence number of the particular message you want to retrieve, and it will be displayed in table format - using 'ag-grid-react' on the frontend. If you want to retrieve all messages, just choose any negative sequence number and the full table will be displaying all messages in the Database. 'ag-grid-react' automatically handles pagination and sorting so I did not need seperate logic for it.

On the backend, this is handled by the `Db_fetch` function.
