import React from 'react';
import {Routes, Route, BrowserRouter} from 'react-router-dom';
import {DisplayGift} from "./components/gift";
import './App.css';

function App() {

  return (
    <BrowserRouter>
      <div className="App">
        <h1>Gift App</h1>
        <Routes>
          <Route path="gifts">
            <Route path=":id" element={<DisplayGift />}/>
          </Route>
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
