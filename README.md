# Open Transaction API

## Overview

This project is an **open transaction system** that allows users to initiate private transactions between accounts while maintaining a fully visible ledger of all transactions in a centralized database. Think of it as a **private blockchain-inspired system** where every user can send transactions securely, but all transactions are publicly accessible for transparency.

Each transaction is recorded in a structured format, ensuring integrity and accountability while maintaining sender authentication.

## Getting Started

### Backend Setup (Go)
To start the backend server, run:
```bash
cd backend
go run .
```

### Frontend Setup (React)
In a separate terminal, start the frontend:
```bash
cd frontend
npm install  # Install dependencies (if not already installed)
npm start
```
This will launch the frontend, directing you to the login page.

## Authentication & Session Management

### Login
Users log in using their **Account Number** and **Routing Number**. A session is created upon successful login, allowing users to submit transactions securely.

### OAuth & Message Submission
Before a transaction is added to the database, the system performs an **OAuth-like check** to verify that the **sender's account and routing numbers match the currently logged-in user**. This prevents unauthorized transactions from being created on behalf of other users.

## Transactions

### Submitting a Transaction
To send a transaction, users must input their message in the format:

```
seq=A;sender_rtn=B;sender_an=C;receiver_rtn=D;receiver_an=E;amount=F
```

- `{A-F}` are integer values.
- `sender_an` and `sender_rtn` **must match** the logged-in user's credentials.
- Transactions failing this validation will be rejected.

This process is managed on the backend by:
- `Create_message` → Parses and validates transaction input.
- `DB_insert` → Stores valid transactions in the database.

### Retrieving Transactions
Users can fetch transactions using a **sequence number**:
- Entering a **specific sequence number** returns the corresponding transaction.
- Entering a **negative number** retrieves **all transactions** in the database.

Transactions are displayed in a **sortable and paginated table** using `ag-grid-react`.

Backend logic for retrieval:
- `DB_fetch` → Queries the database based on the sequence number.

## Future Improvements
- **Encryption**: Adding cryptographic security for transactions.
- **Distributed Ledger**: Exploring decentralization or blockchain-like implementation.
- **Smart Contracts**: Implementing rules-based automation for transactions.
- **User Profiles**: Enhancing account management and permissions.


