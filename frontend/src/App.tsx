import React, {useEffect, useState} from 'react';
import {Gift} from "./types";
import {decreaseGiftAmount, getGift} from "./api";
import './App.css';

function App() {
  const [gift, setGift] = useState<Gift>();

  const loadGift = async () => {
    const gift = await getGift(1);
    setGift(gift);
  }

  useEffect(() => {
    loadGift();
  }, []);

  const onClickDecrease = async () => {
    if (!gift) return;

    console.log('using gift', gift)

    await decreaseGiftAmount(gift);

    loadGift();
  }

  return (
    <div className="App">
      <h1>Gift App</h1>
      {!gift && <div>Loading...</div>}
      {gift &&
        <div>
          <ul>
            <li>{gift.name}</li>
            <li>{gift.displayText}</li>
            <li>{gift.amount}</li>
          </ul>
          <button onClick={onClickDecrease}>Use gift</button>
        </div>
      }
    </div>
  );
}

export default App;
