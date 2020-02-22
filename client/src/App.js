import React from "react";
import "./App.css";

// import Container Component de semantic-ui-react
import { Container } from "semantic-ui-react";

// import ToDoList component
import ToDoList from "./components/ToDoList";

function App() {
  return (
    <div>
      <Container>
        <ToDoList />
      </Container>
    </div>
  );
}
export default App;