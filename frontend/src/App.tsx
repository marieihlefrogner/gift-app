import React from 'react';
import {Routes, Route, BrowserRouter} from 'react-router-dom';
import {DisplayGift} from "./components/gift/gift";
import './style/reset.css';
import './style/App.scss';

function App() {
  return (
    <BrowserRouter>
      <div className="app">
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
