import React, { useState, useRef, useEffect } from "react";
import { AgGridReact } from "ag-grid-react";
import { GridOptions, GridApi, ModuleRegistry, ColDef, PaginationModule, ClientSideRowModelModule } from "ag-grid-community";
import "ag-grid-community/styles/ag-grid.css";
import "ag-grid-community/styles/ag-theme-alpine.css";
import { fetchMessages } from "./api";
import { Message } from "./types";

// Register AG Grid modules
ModuleRegistry.registerModules([PaginationModule, ClientSideRowModelModule]);

const FetchMessage = () => {
  const [messageId, setMessageId] = useState(1);
  const [rowData, setRowData] = useState<Message[]>([]);
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const gridApi = useRef<GridApi | null>(null);

  // Define column headers for AG Grid
  const columnDefs: ColDef[] = [
    { headerName: "Sequence Number", field: "seq", sortable: true, filter: true },
    { headerName: "Sender Routing Number", field: "sender_info.routing_num", sortable: true, filter: true },
    { headerName: "Sender Account Number", field: "sender_info.account_num", sortable: true, filter: true },
    { headerName: "Receiver Routing Number", field: "receiver_info.routing_num", sortable: true, filter: true },
    { headerName: "Receiver Account Number", field: "receiver_info.account_num", sortable: true, filter: true },
    { headerName: "Amount", field: "amount", sortable: true, filter: true },
  ];

  // Define grid options
  const gridOptions: GridOptions = {
    columnDefs,
    pagination: true, // Enables pagination
    paginationPageSize: 10, // Number of rows per page
    domLayout: "autoHeight", // Adjust height automatically
  };

  // Fetch message from backend
  const handleFetchMessage = async () => {
    const data = await fetchMessages(messageId);
    if (data.length === 0) {
      setErrorMessage("No messages found");
      setRowData([]);
    } else {
      setRowData(data);
      setErrorMessage(null);
    }
  };

    // Store the Grid API when the grid is ready
    const onGridReady = (params: { api: GridApi }) => {
        gridApi.current = params.api; // Store API instance
    };

  // Refresh the table with updated data
  const handleReRender = async () => {
    const data = await fetchMessages(messageId);
    setRowData(data);
  };

  return (
    <div style={{ textAlign: "center", marginTop: "50px" }}>
      <h2>Fetch Message by ID</h2>
      <h6>A negative ID will fetch all messages in DB</h6>

      <div style={{ display: "flex", justifyContent: "center", alignItems: "center", gap: "10px" }}>
        <button onClick={() => setMessageId((prev) => prev - 1)} style={{ padding: "5px 10px", fontSize: "16px" }}>-</button>
        <input
          type="number"
          value={messageId}
          onChange={(e) => setMessageId(Number(e.target.value))}
          style={{ padding: "10px", fontSize: "16px", textAlign: "center", width: "80px" }}
        />
        <button onClick={() => setMessageId((prev) => prev + 1)} style={{ padding: "5px 10px", fontSize: "16px" }}>+</button>
      </div>

      <button onClick={handleFetchMessage} style={{ marginTop: "20px", padding: "10px 20px", fontSize: "16px" }}>
        Fetch Messages
      </button>

      {/* Re-fetch and update table */}
      <button onClick={handleReRender} style={{ marginLeft: "10px", padding: "5px 10px", fontSize: "14px", backgroundColor: "#007bff", color: "white", border: "none", borderRadius: "4px", cursor: "pointer" }}>
        🔄 Re-render Table
      </button>

      {/* Display error message */}
      {errorMessage && (
        <p style={{ marginTop: "20px", fontSize: "18px", fontWeight: "bold", color: "red" }}>
          {errorMessage}
        </p>
      )}

      {/* Display AG Grid Table */}
      <div
        className="ag-theme-alpine"
        style={{
          width: "90%",
          margin: "20px auto",
          height: rowData.length > 0 ? `${Math.min(rowData.length * 50, 500)}px` : "150px",
        }}
      >
        <AgGridReact
          onGridReady={onGridReady}
          gridOptions={gridOptions}
          rowData={rowData}
          pagination={true}
          paginationPageSize={10}
          domLayout="autoHeight"
        />
      </div>
    </div>
  );
};

export default FetchMessage;


