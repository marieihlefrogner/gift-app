import React, {useCallback, useEffect, useState} from "react";
import { useParams } from "react-router-dom";
import api from "../api";
import { Gift } from "../types";

export const useGift = (id?: string) => {
  const [gift, setGift] = useState<Gift>();

  const loadGift = useCallback(async () => {
    if (!id) return;

    const gift = await api.getGift(id);
    setGift(gift);
  }, [id]);

  useEffect(() => {
    loadGift();
  }, [loadGift, id]);

  return { gift, loadGift };
}

export const DisplayGift = () => {
  const { id } = useParams<Record<string, string | undefined>>();
  const { gift, loadGift } = useGift(id);

  const onClickUse = async () => {
    if (!gift) return;

    console.log('using gift', gift)

    await api.postUseGift(gift);

    loadGift();
  }

  return (
    <div>
      {!gift && <div>Loading...</div>}
      {gift &&
          <div>
              <ul>
                  <li>{gift.name}</li>
                  <li>{gift.displayText}</li>
                  <li>{gift.amount}</li>
              </ul>
              <button onClick={onClickUse}>Use gift</button>
          </div>
      }
    </div>
  )
}